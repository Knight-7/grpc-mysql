package engine

import (
	"fmt"
	"net"
	"rpc-mysql/dao"
	"rpc-mysql/pkg/clientset"
	"rpc-mysql/pkg/config"
	pb "rpc-mysql/pkg/proto"
	"rpc-mysql/rpc"
	"rpc-mysql/rpc-mysql/server/server"
)

type Engine struct {
	daoServer *server.RPCServer
	stopChan  chan struct{}
}

func NewEngine(cfg *config.Config) (*Engine, error) {
	engine := new(Engine)
	engine.stopChan = make(chan struct{})

	gClientset, err := clientset.NewClientset(cfg)
	if err != nil {
		return nil, err
	}

	da := dao.NewDAO(gClientset.MySQL)
	daoRPC := rpc.NewDaoRPC(da)

	engine.daoServer = server.NewRPCServer(cfg.GetServerAddr())
	pb.RegisterDAOServer(engine.daoServer.Server, daoRPC)

	return engine, nil
}

func (e *Engine) Run() {
	_ = e.serverAndListen(e.daoServer.Addr)
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
