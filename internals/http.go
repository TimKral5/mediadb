package internals

import (
	"mediadb/routers"
	"mediadb/utils"
	"net/http"
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

	movieRouter := routers.MovieRouter{
		BaseRoute: "/movies",
		Mongo: self.program.mongoConn,
		Log: self.program.log,
	}

	stack := utils.CreateStack(
		self.program.log.Middleware,
		movieRouter.Middleware,
	)

	server := http.Server{
		Addr:    self.config.Addr,
		Handler: stack(ctx),
	}

	server.ListenAndServe()
}
