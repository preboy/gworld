@echo off


del /S /Q json\*.*
del /S /Q lua\*.*


rem: record export

exporter jl excel/levelup_�ȼ���.xlsx               Levelup
exporter jl excel/hero_Ӣ�۱�.xlsx                  Hero
exporter jl excel/creature_�����.xlsx              Creature
exporter jl excel/creature_team_��������.xlsx     CreatureTeam
exporter jl excel/skill_���ܱ�.xlsx                 Skill
exporter jl excel/aura_�⻷��.xlsx                  Aura
exporter jl excel/item_���߱�.xlsx                  Item
exporter jl excel/market_���б�.xlsx                Market
exporter jl excel/activity_���.xlsx              Activity

exporter jl excel/refine_������.xlsx                RefineNormal
exporter jl excel/refine_������.xlsx                RefineSuper

exporter jl excel/achv_�ɾͱ�.xlsx                  Growth
exporter jl excel/achv_�ɾͱ�.xlsx                  Achv

exporter jl excel/world_���������.xlsx             Scene
exporter jl excel/world_���������.xlsx             Object


rem: map export
exporter_map jl excel/global_ȫ�ֱ�.xlsx            Global


pause
