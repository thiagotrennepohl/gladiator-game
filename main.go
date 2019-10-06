package main

import (
	"github.com/thiagotrennepohl/gladiator-game/service/warrior"

	"github.com/thiagotrennepohl/gladiator-game/service/opponent"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/thiagotrennepohl/gladiator-game/models"
	"github.com/thiagotrennepohl/gladiator-game/ui/battle"
	"github.com/thiagotrennepohl/gladiator-game/ui/menu"
)

var (
	showMenu = true
	player   models.Character
)

func main() {
	rl.InitWindow(800, 400, "Gladiator Game")
	rl.SetTargetFPS(60)
	rl.SetConfigFlags(rl.FlagVsyncHint)

	warriorSvc := warrior.NewWarriorService(models.BaseAttributes(warriorStatsRules))
	rules := make(map[string]models.BaseAttributes)
	rules["warrior"] = warriorAttributesRules

	classes := make(map[string]models.Character)
	classes["warrior"] = warriorSvc

	opponentSvc := opponent.NewOpponentService(classes, rules)

	menu.InitStuff()
	battle.InitStuff()
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.White)
		rl.BeginDrawing()
		if showMenu {
			showMenu, player = menu.ShowMenu(warriorStatsRules)
		} else {
			battle.Start(player, opponentSvc)
		}
		rl.EndDrawing()
	}
}
