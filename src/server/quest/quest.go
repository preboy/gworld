package quest

type Quest struct {

	// 接任务条件 (等级/性别，时间段)

	Id    int32
	QType int32 // 主线，日常， 帮会

	Content string // 介绍文本

	// 提交Npc
	// {
	// 杀怪、对话、送信 (侦听事件)
	// }

	//  *Rewards    []*Items {Id, Cnt},
}
