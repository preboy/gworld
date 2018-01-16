@echo off


del /S /Q json\*.*
del /S /Q lua\*.*


rem: record export
exporter jl excel/hero_英雄表.xlsx  HeroProto
exporter jl excel/creature_怪物表.xlsx  CreatureProto
exporter jl excel/skill_技能表.xlsx  SkillProto 
exporter jl excel/aura_光环表.xlsx  AuraProto 
exporter jl excel/item_道具表.xlsx  ItemProto
exporter jl excel/creature_team_怪物队伍表.xlsx  CreatureTeam



rem: map export
exporter_map jl excel/global_全局表.xlsx  global


pause