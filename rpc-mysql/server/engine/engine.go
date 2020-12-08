package engine

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"rpc-mysql/dao"
	"rpc-mysql/interceptor"
	"rpc-mysql/pkg/clientset"
	"rpc-mysql/pkg/config"
	pb "rpc-mysql/pkg/proto"
	"rpc-mysql/rpc"
	"rpc-mysql/rpc-mysql/server/server"
)

type Engine struct {
	daoServer *server.RPCServer
	log       *logrus.Logger
	stopChan  chan struct{}
}

func NewEngine(cfg *config.Config) (*Engine, error) {
	engine := new(Engine)
	engine.stopChan = make(chan struct{})

	gClientset, err := clientset.NewClientset(cfg)
	if err != nil {
		return nil, err
	}

	engine.log = gClientset.GetLogger()

	// 初始化dao和daoRPC服务
	da := dao.NewDAO(gClientset.GetMySQL())
	daoRPC := rpc.NewDaoRPC(da)

	// 注册拦截器
	var options []grpc.ServerOption
	interceptor.SetupAuthAndLog(cfg)
	options = append(options, grpc.UnaryInterceptor(interceptor.NewServerUnaryInterceptor()))
	options = append(options, grpc.StreamInterceptor(interceptor.NewServerStreamInterceptor()))

	// 注册DAO服务
	engine.daoServer = server.NewRPCServer(cfg.GetServerAddr(), options...)
	pb.RegisterDAOServer(engine.daoServer.Server, daoRPC)

	return engine, nil
}

func (e *Engine) Run() {
	fmt.Println("start grpc server")
	_ = e.serverAndListen(e.daoServer.Addr)
}

func (e *Engine) Stop() {

}

func (e *Engine) serverAndListen(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		//TODO: add log
		fmt.Println(err)
		return err
	}

	if err = e.daoServer.Server.Serve(listen); err != nil {
		//TODO:add log
		fmt.Println(err)
		return err
	}
	return nil
}
