package core

// Create  the interface that each new action should implement  每一个新增动作都应该实现的接口
type Create interface {
	NewConfig()
}
type Core struct {
	AllWorkingGroups []string
}
