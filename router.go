package pheromones

// Router 路由接口
type Router interface {
	AddRoute(name string, addr interface{}) error

	// 删除路由
	Delete(name string) error

	// 广播发送信息
	DispatchAll(msg []byte) map[string][]byte

	// 单点发送信息
	Dispatch(name string, msg []byte) ([]byte, error)

	// 获取全部peer
	FetchPeers() map[string]interface{}
}
