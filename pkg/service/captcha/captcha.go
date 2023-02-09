package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/grin-ch/grin-api/api/grpc/captcha"
	"github.com/grin-ch/grin-auth/pkg/auth"
	"github.com/grin-ch/grin-auth/pkg/service/internal"
	"github.com/grin-ch/grin-auth/pkg/util"
	"github.com/grin-ch/grin-utils/log"
	"github.com/grin-ch/grin-utils/safe"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type captchaEmail struct {
	port       int
	host       string
	from       string
	secret     string
	subject    string
	recipients []string
	body       string
}

func (e *captchaEmail) Host() string         { return e.host }
func (e *captchaEmail) Port() int            { return e.port }
func (e *captchaEmail) Secret() string       { return e.secret }
func (e *captchaEmail) From() string         { return e.from }
func (e *captchaEmail) Recipients() []string { return e.recipients }
func (e *captchaEmail) Subject() string      { return e.subject }
func (e *captchaEmail) Body() []byte         { return []byte(e.body) }

type captchaService struct {
	auth.Unauther
	redisClient *redis.Client

	expires int
	port    int
	subject string
	format  string
	host    string
	from    string
	secret  string
}

// NewCaptchaServer impl captcha service
func NewCaptchaServer(client *redis.Client, expires, port int,
	subject, emailFormat, host, from, secret string) captcha.CaptchaServiceServer {
	return &captchaService{
		redisClient: client,
		expires:     expires,

		port:    port,
		subject: subject,
		format:  emailFormat,
		host:    host,
		from:    from,
		secret:  secret,
	}
}

// 异步验证码 用于短信/邮箱等
func (s *captchaService) AsyncCode(ctx context.Context, req *captcha.AsyncCodeReq) (*captcha.AsyncCodeRsp, error) {
	var (
		code   = string(util.GenFormSet(6, util.EasyRead()))
		sender func()
	)

	if util.ValidateEmail(req.Contact) {
		mail := s.makeEmail([]string{req.Contact}, s.subject, fmt.Sprintf(s.format, code))
		sender = func() {
			util.SendEmail(mail)
		}
	} else if util.ValidatePhoneNumber(req.Contact) {
		// TODO
		return nil, status.Errorf(codes.Unimplemented, "the sms unimplemented")
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "contact invalid")
	}
	c := internal.Captcha{
		Key:     req.Contact,
		Content: code,
		Purpose: int32(req.Purpose),
	}
	if err := s.setCaptcha(c); err != nil {
		log.Logger.Errorf("set captcha err:%s", err.Error())
		return nil, status.Errorf(codes.Internal, "set captcha err")
	}
	safe.Go(sender, safe.OptLog(
		func(err error) {
			log.Logger.Errorf("send async code err:%v", err)
		},
	))

	return &captcha.AsyncCodeRsp{}, nil
}

func (s *captchaService) makeEmail(recipients []string, subject string, body string) *captchaEmail {
	return &captchaEmail{
		port:       s.port,
		host:       s.host,
		from:       s.from,
		secret:     s.secret,
		recipients: recipients,
		subject:    subject,
		body:       body,
	}
}

// 图形验证码
func (s *captchaService) GraphCaptcha(ctx context.Context, req *captcha.GraphCaptchaReq) (*captcha.GraphCaptchaRsp, error) {
	text := util.GenFormSet(4, util.EasyRead())
	// 设置key,value
	key := util.MD5(text)
	val := string(text)
	c := internal.Captcha{
		Key:     key,
		Content: val,
		Purpose: int32(req.Purpose),
	}
	if err := s.setCaptcha(c); err != nil {
		log.Logger.Errorf("set captcha err:%s", err.Error())
		return nil, status.Errorf(codes.Internal, "set captcha err")
	}

	imgByte, err := util.NewImg(val)
	if err != nil {
		log.Logger.Errorf("new captcha image err:%s", err.Error())
		return nil, status.Errorf(codes.Internal, "new captcha image err")
	}
	base64str := base64.StdEncoding.EncodeToString(imgByte)
	return &captcha.GraphCaptchaRsp{
		Captcha: &captcha.Captcha{
			Key:     key,
			Purpose: req.Purpose,
			Content: base64str,
		},
	}, nil
}

// 验证码校验
func (s *captchaService) Verify(ctx context.Context, req *captcha.VerifyReq) (*captcha.VerifyRsp, error) {
	ok, err := s.verifyCaptcha(internal.Captcha{
		Key:     req.Key,
		Content: req.Value,
		Purpose: int32(req.Purpose),
	})
	if err != nil {
		log.Logger.Errorf("verify captcha err:%s", err.Error())
		return nil, status.Errorf(codes.Internal, "verify captcha err")
	}

	return &captcha.VerifyRsp{
		Success: ok,
	}, nil
}

// 设置验证码
func (s *captchaService) setCaptcha(c internal.Captcha) error {
	return s.redisClient.Set(c.Key, &c, time.Duration(s.expires)*time.Second).Err()
}

// 验证验证码
func (s *captchaService) verifyCaptcha(c internal.Captcha) (bool, error) {
	cmd := s.redisClient.Get(c.Key)
	if err := cmd.Err(); err != nil {
		return false, err
	}

	var val internal.Captcha
	if err := cmd.Scan(&val); err != nil && err != redis.Nil {
		return false, err
	}
	if c.Purpose != val.Purpose ||
		!strings.EqualFold(c.Content, val.Content) {
		return false, nil
	}

	s.redisClient.Del(c.Key)
	return true, nil
}
