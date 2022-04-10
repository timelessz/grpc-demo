package service

import (
	"context"
	"github.com/sirupsen/logrus"
	auth "go-kit-demo/v8-gokit-all/auth"
	v8 "go-kit-demo/v8-gokit-all/book"
	"time"
)

// 服务中间件
// service/middleware.go

type BookServiceLogMiddleware struct {
	next   BookService
	logger *logrus.Logger
}

// 创建 BookServiceLogMiddleware
func NewBookServiceLogMiddleware(logger *logrus.Logger) NewMiddlewareService {
	return func(next BookService) BookService {
		return &BookServiceLogMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (mw BookServiceLogMiddleware) GetBookList(ctx context.Context, in *v8.BookQueryParams) (*v8.BookList, error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{
			"method": "GetBook",
			//"id":     id,
			"took": time.Since(begin),
		}).Info("GetBookList")
	}(time.Now())
	return mw.next.GetBookList(ctx, in)
}

func (mw BookServiceLogMiddleware) GetOneBook(ctx context.Context, obp *v8.OneBookQueryParams) (*v8.Book, error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{
			"method": "GetOneBook",
			//"id":     id,
			"took": time.Since(begin),
		}).Info("GetOneBook")
	}(time.Now())
	return mw.next.GetOneBook(ctx, obp)
}

func (mw BookServiceLogMiddleware) CreateBook(ctx context.Context, bi *v8.Book) (*v8.Book, error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{
			"method": "CreateBook",
			//"id":     id,
			"took": time.Since(begin),
		}).Info("CreateBook")
	}(time.Now())
	println("service middleware CreateBook")
	bk, err := mw.next.CreateBook(ctx, bi)
	if err != nil {
		println("Create Book middleware", err.Error())
	}
	return bk, nil
}

// 用户登录service 中间件
type LoginServiceLogMiddleware struct {
	next   AuthService
	logger *logrus.Logger
}

// 创建登录 service 中间件
func NewAuthServiceLogMiddleware(logger *logrus.Logger) NewLoginMiddlewareService {
	return func(next AuthService) AuthService {
		return &LoginServiceLogMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (mw LoginServiceLogMiddleware) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	//TODO implement me
	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{
			"method": "Login",
			//"id":     id,
			"took": time.Since(begin),
		}).Info("login")
	}(time.Now())
	println("service middleware Login")
	return mw.next.Login(ctx, request)
}

func (mw LoginServiceLogMiddleware) GetUserInfo(ctx context.Context, request *auth.UserIdRequest) (*auth.User, error) {
	//TODO implement me
	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{
			"method": "GetUserInfo",
			//"id":     id,
			"took": time.Since(begin),
		}).Info("GetUserInfo")
	}(time.Now())
	return mw.next.GetUserInfo(ctx, request)
}
