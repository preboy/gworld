@echo off


del /S /Q json\*.*
del /S /Q lua\*.*


rem: record export
exporter jl excel/hero_Ӣ�۱�.xlsx  HeroProto
exporter jl excel/creature_�����.xlsx  CreatureProto
exporter jl excel/skill_���ܱ�.xlsx  SkillProto 
exporter jl excel/aura_�⻷��.xlsx  AuraProto 
exporter jl excel/item_���߱�.xlsx  ItemProto
exporter jl excel/creature_team_��������.xlsx  CreatureTeam



rem: map export
exporter_map jl excel/global_ȫ�ֱ�.xlsx  global


pause