package rpc

import "rpc-mysql/dao"

type DaoRPC struct {
	d *dao.DAO
}

func NewDaoRPC(d *dao.DAO) *DaoRPC {
	return &DaoRPC{
		d: d,
	}
}
