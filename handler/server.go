package handler

import (
	"easylib-go/config"
	"easylib-go/manager"
	"easylib-go/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager

	srv            *gin.Engine
	host string
}

func (s *server) Run() {

	store := cookie.NewStore([]byte("secret"))
	s.srv.Use(middleware.LoggerMiddleware())
	s.srv.Use(sessions.Sessions("session", store))

	NewLoginHandler(s.srv, s.usecaseManager.GetLoginUsecase())
	NewUserHandler(s.srv, s.usecaseManager.GetUserUsecase())
	NewBookHandler(s.srv, s.usecaseManager.GetBookListUsecase())
	NewMemberHandler(s.srv, s.usecaseManager.GetMemberUsecase())
	NewBookLoanHandler(s.srv, s.usecaseManager.GetBookLoanUsecase())
	s.srv.Run()

}

func NewServer() Server {
	c := config.NewConfig()

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()
	return &server{
		usecaseManager: usecase,
		srv:            srv,
		host:           c.AppPort,
	}
}
