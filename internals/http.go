package internals

import (
	"mediadb/middleware"
	"mediadb/routers"
	"mediadb/utils"
	"net/http"

	"github.com/google/uuid"
)

type HttpConfig struct {
	Addr string
}

type HttpServer struct {
	program *Program
	config  *HttpConfig
}

func NewHttpServer(prog *Program, conf *HttpConfig) HttpServer {
	server := HttpServer{
		program: prog,
		config:  conf,
	}
	return server
}

func (self *HttpServer) LaunchHttpServer() {
	ctx := http.NewServeMux()

	sessionUUID, err := uuid.NewUUID()

	if err != nil {
		self.program.log.Error(err)
		return
	}

	movieRouter := routers.MovieRouter{
		Mongo: self.program.mongoConn,
		Log:   self.program.log,
		Ldap:  self.program.ldapConn,
	}

	authRouter := routers.AuthRouter{
		SessionUUID: sessionUUID.String(),
		Ldap:        self.program.ldapConn,
		Log:         self.program.log,
	}

	authMiddleware := middleware.Authenticator{
		SessionUUID: sessionUUID.String(),
		Log:         self.program.log,
	}

	stack := utils.CreateStack(
		self.program.log.Middleware,
		authMiddleware.Middleware,
	)

	ctx.Handle("/auth", authRouter.GetHandler())
	ctx.Handle("/movies", movieRouter.GetHandler())

	server := http.Server{
		Addr:    self.config.Addr,
		Handler: stack(ctx),
	}

	server.ListenAndServe()
}
