// Do NOT edit this file manually

package player

import (
	"public/protocol"
)

func init() {
	register_handler(protocol.MSG_CS_PlayerDataRequest, handler_PlayerDataRequest)
	register_handler(protocol.MSG_CS_GMCommandRequest, handler_GMCommandRequest)
	register_handler(protocol.MSG_CS_UseItemRequest, handler_UseItemRequest)
	register_handler(protocol.MSG_CS_MarketBuyRequest, handler_MarketBuyRequest)
	register_handler(protocol.MSG_CS_ChangeNameRequest, handler_ChangeNameRequest)
	register_handler(protocol.MSG_CS_HeroLevelupRequest, handler_HeroLevelupRequest)
	register_handler(protocol.MSG_CS_HeroRefineRequest, handler_HeroRefineRequest)
	register_handler(protocol.MSG_CS_HeroAptitudeRequest, handler_HeroAptitudeRequest)
	register_handler(protocol.MSG_CS_HeroTalentRequest, handler_HeroTalentRequest)
	register_handler(protocol.MSG_CS_QuestListRequest, handler_QuestListRequest)
	register_handler(protocol.MSG_CS_QuestOpRequest, handler_QuestOpRequest)
	register_handler(protocol.MSG_CS_ChapterInfoRequest, handler_ChapterInfoRequest)
	register_handler(protocol.MSG_CS_ChapterFightingRequest, handler_ChapterFightingRequest)
	register_handler(protocol.MSG_CS_ChapterRewardsRequest, handler_ChapterRewardsRequest)
	register_handler(protocol.MSG_CS_ChapterLootRequest, handler_ChapterLootRequest)
}
