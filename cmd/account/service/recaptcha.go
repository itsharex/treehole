package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/Jazee6/treehole/cmd/account/rpc"
	"github.com/Jazee6/treehole/pkg/rpcs"
	"github.com/Jazee6/treehole/pkg/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const reUrl = "https://www.recaptcha.net/recaptcha/api/siteverify"

type recaptchaResponse struct {
	Score       float64   `json:"score"`
	Success     bool      `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	Action      string    `json:"action"`
	ErrorCodes  []string  `json:"error-codes"`
}

func Recaptcha(token string) error {
	formValue := url.Values{}
	formValue.Set("secret", recaptchaSec)
	formValue.Set("response", token)
	formStr := formValue.Encode()
	formBytes := []byte(formStr)
	response, err := http.Post(reUrl, "application/x-www-form-urlencoded", bytes.NewReader(formBytes))
	if err != nil {
		return err
	}
	r, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var resp recaptchaResponse
	err = json.Unmarshal(r, &resp)
	if err != nil {
		return err
	}
	if !resp.Success {
		return errors.New("recaptcha failed")
	}
	return nil
}

func (c *AccountService) SendCaptcha(ctx context.Context, request *rpc.SendCaptchaRequest) (*rpc.SendCaptchaResponse, error) {
	err := Recaptcha(request.Token)
	if err != nil {
		if err.Error() == "recaptcha failed" {
			return &rpc.SendCaptchaResponse{
				Code: rpcs.Code_ErrRecaptchaErr,
			}, nil
		}
		return nil, err
	}

	// 验证邮箱是否已经注册
	q := dao.Q.User
	user, err := q.Where(q.Email.Eq(request.Email)).Take()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if user != nil {
		return &rpc.SendCaptchaResponse{
			Code: rpcs.Code_ErrUserExist,
		}, nil
	}

	// 验证码限流
	result, err := r.Get(ctx, "email").Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if err == redis.Nil {
		now := time.Now()
		year, month, day := now.Date()
		err := r.Set(ctx, "email", 0, time.Date(year, month, day, 24, 0, 0, 0, time.Local).Sub(time.Now())).Err()
		if err != nil {
			return nil, err
		}
		result = "0"
	}
	re, err := strconv.Atoi(result)
	if err != nil {
		return nil, err
	}
	if emailMax < re {
		return &rpc.SendCaptchaResponse{
			Code: rpcs.Code_ErrEmailLimit,
		}, nil
	}
	err = r.Incr(ctx, "email").Err()
	if err != nil {
		return nil, err
	}
	code := fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999999))
	err = r.Set(ctx, request.Email, code, time.Minute*time.Duration(captchaExpire)).Err()
	if err != nil {
		return nil, err
	}
	err = utils.SendMail(request.Email, "【"+name+"】"+"验证码", os.Expand(captchaContent, func(s string) string {
		switch s {
		case "code":
			return code
		}
		return "{" + s + "}"
	}))
	if err != nil {
		return nil, err
	}

	return &rpc.SendCaptchaResponse{
		Code: rpcs.Code_Success,
	}, nil
}
