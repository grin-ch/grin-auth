package grpc

import "google.golang.org/grpc"

// 初始化客户端
func initClients(resolve func(string) (*grpc.ClientConn, error), connectors ...ClientConnector) {
	for _, c := range connectors {
		name, fn := c()
		initClient(resolve, name, fn)
	}
}

func initClient(resolve func(string) (*grpc.ClientConn, error),
	serverName string, fn func(*grpc.ClientConn)) {
	conn, err := resolve(serverName)
	if err != nil {
		panic(err)
	}
	fn(conn)
}

type ClientConnector func() (string, func(*grpc.ClientConn))
