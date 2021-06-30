package rtm

type RTMRegressiveConnectStrategy struct {
	connectFailedMaxIntervalMilliseconds int // 从连接成功到断开多少毫秒内算闪断，记一次连接失败
	startConnectFailedCount              int // 连接失败多少次后，开始退行性处理
	firstIntervalSeconds                 int // 第一次退行性间隔基数
	maxIntervalSeconds                   int // 退行性重连最大时间间隔
	linearRegressiveCount                int // 从第一次退行性连接开始，到最大连接时间，允许尝试几次连接，每次时间间隔都会增大
}

func NewRTMRegressiveConnectStrategy(connectFailedMaxIntervalMilliseconds int, startConnectFailedCount int,
	firstIntervalSeconds int, maxIntervalSeconds int, linearRegressiveCount int) *RTMRegressiveConnectStrategy {

	strategy := &RTMRegressiveConnectStrategy{
		connectFailedMaxIntervalMilliseconds,
		startConnectFailedCount,
		firstIntervalSeconds,
		maxIntervalSeconds,
		linearRegressiveCount}
	return strategy
}

func (strategy *RTMRegressiveConnectStrategy) SetConnectFailedInterval(milliseconds int) {
	strategy.connectFailedMaxIntervalMilliseconds = milliseconds
}

func (strategy *RTMRegressiveConnectStrategy) SetStartConnectFailedCount(count int) {
	strategy.startConnectFailedCount = count
}

func (strategy *RTMRegressiveConnectStrategy) SetFirstIntervalTime(second int) {
	strategy.firstIntervalSeconds = second
}

func (strategy *RTMRegressiveConnectStrategy) SetMaxInterval(second int) {
	strategy.maxIntervalSeconds = second
}

func (strategy *RTMRegressiveConnectStrategy) SetLinearRegressiveCount(count int) {
	strategy.linearRegressiveCount = count
}

var DefaultRegressiveStrategy = &RTMRegressiveConnectStrategy{
	1500,
	5,
	2,
	120,
	5,
}
