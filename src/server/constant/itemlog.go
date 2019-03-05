package constant

// ============================================================================
// item log id

const (

	// ------------------------------------------------------------------------
	// 不要随便改动顺序以及从中间插入
	// ------------------------------------------------------------------------

	ItemLog_ChapterGot       = iota + 1 // 章节完成奖励
	ItemLog_ChapterLoot                 // 使用道具
	ItemLog_ChapterBreakPass            // 关卡通过奖励
	ItemLog_QuestCommit                 // 任务提交
	ItemLog_QuestFinish                 // 任务提交
	ItemLog_MarketBuy                   // 集市购买
	ItemLog_HeroLvUp                    // 英雄升级
	ItemLog_HeroRefine                  // 英雄精炼
	ItemLog_UseItem                     // 使用道具

	ItemLog_GM = 9999 // GM发放
)
