package znet

import "Game_Zinx/src/zinx/ziface"

type Request struct {
	conn ziface.Iconnection
	data []byte
}

func (r *Request) GetConnection() Iconnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
