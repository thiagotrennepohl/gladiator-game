package menu

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/thiagotrennepohl/gladiator-game/models"
	warriorSvc "github.com/thiagotrennepohl/gladiator-game/service/warrior"
)

var (
	//Character that has been built
	remainingPoints     = 5
	character           models.Character
	finish              = false
	warrior             = false
	mage                = false
	archer              = false
	startGameClicked    = false
	showMenu            = true
	strSpinner          = 0
	staSpinner          = 0
	dexSpinner          = 0
	intSpinner          = 0
	luckSpinner         = 0
	framesCounter       = 0
	groundFramesCounter = 0
	framesSpeed         = 8
	gFramesSpeed        = 30
	currentFrame        = float32(0)
	currentgroundFrame  = float32(0)
	groundPosition      = rl.NewVector2(0, 255)
	groundFrame         = rl.NewRectangle(20, 20, 1200.0, 1200.0)
	skyPosition         = rl.NewVector2(0, 0)
	skyFrame            = rl.NewRectangle(20, 20, 12000.0, 1020.0)
	characterTexture    = rl.Texture2D{}
	position            = rl.Vector2{}
	frameRec            rl.Rectangle
	bg                  rl.Texture2D
	sky                 rl.Texture2D
)

func InitStuff() {
	position = rl.NewVector2(100.0, 325)
	characterTexture = rl.LoadTexture("assets/img/run.png")
	bg = rl.LoadTexture("assets/img/back.png")
	sky = rl.LoadTexture("assets/img/sky.png")
	frameRec = rl.NewRectangle(0, 0, float32(characterTexture.Width/7), float32(characterTexture.Height))
}

func ShowMenu(rules models.WarriorRules) (bool, models.Character) {
	rl.DrawText("MENU", 380, 20, 20, rl.Black)
	if showMenu {
		startGameClicked = raygui.Button(rl.NewRectangle(310, 70, 200, 40), "Start game")
	}
	if startGameClicked && !finish {
		showMenu = false
		showBackground()
		if !(warrior || mage || archer) {
			showClassSelection()
		} else {
			showCharacterCreationMenu()
			showUpdatedStats()
			if remainingPoints == 0 {
				finish = raygui.Button(rl.NewRectangle(310, 350, 200, 40), "FIGHT!")
			}
		}
	}
	if finish {
		return !finish, createCharacter(rules)
	}
	return true, nil
}

func createCharacter(rules models.WarriorRules) models.Character {
	char := warriorSvc.NewWarriorService(models.BaseAttributes(rules))
	char.SetBaseAttributes(
		models.BaseAttributes{
			Strength:     float32(strSpinner) + 5,
			Stamina:      float32(staSpinner) + 5,
			Dextery:      float32(dexSpinner) + 5,
			Intelligence: float32(intSpinner) + 5,
			HP:           float32(staSpinner)*10 + 100,
		})
	return char
}

func showClassSelection() {
	rl.DrawText("PICK YOUR TOOL", 200, 20, 20, rl.Black)

	mage = raygui.Button(rl.NewRectangle(310, 300, 200, 40), "Staff")
	archer = raygui.Button(rl.NewRectangle(310, 200, 200, 40), "Bow")
	//Under 100 value for Y it sets the var to true automagically
	warrior = raygui.Button(rl.NewRectangle(310, 100, 200, 40), "Sword")

}

func showCharacterCreationMenu() {
	rl.DrawText("PICK YOUR ATTRIBUTES", 200, 20, 20, rl.Black)
	local := 5 - strSpinner - staSpinner - dexSpinner - intSpinner - luckSpinner
	remainingPoints = local
	rl.DrawText(fmt.Sprintf("Remaining points : %d", local), 400, 320, 20, rl.Black)
	rl.DrawText("Strenght", 470, 100, 15, rl.Black)
	strSpinner = raygui.Spinner(rl.NewRectangle(400, 120, 200, 20), strSpinner, 0, local+strSpinner)
	rl.DrawText("Stamina", 470, 150, 15, rl.Black)
	staSpinner = raygui.Spinner(rl.NewRectangle(400, 170, 200, 20), staSpinner, 0, local+staSpinner)
	rl.DrawText("Dextery", 470, 200, 15, rl.Black)
	dexSpinner = raygui.Spinner(rl.NewRectangle(400, 220, 200, 20), dexSpinner, 0, local+dexSpinner)
	rl.DrawText("Intelligence", 470, 250, 15, rl.Black)
	intSpinner = raygui.Spinner(rl.NewRectangle(400, 270, 200, 20), intSpinner, 0, local+intSpinner)
	//rl.DrawText("Stamina", 470, 250, 15, rl.Black)
	//luckSpinner = raygui.Spinner(rl.NewRectangle(500, 370, 200, 20), luckSpinner, 0, local+luckSpinner)

}

func showUpdatedStats() {
	if warrior {
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Attack", float32(strSpinner)*0.5+10), 200, 100, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Defense", float32(staSpinner)*0.4+10), 200, 115, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Critical Chance", float32(dexSpinner)*0.4+10), 200, 130, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "HP", float32(staSpinner)*10+100), 200, 145, 15, rl.Black)
	} else if mage {
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Attack", float32(intSpinner)*0.5+10), 200, 100, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Defense", float32(staSpinner)*0.4+10), 200, 115, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Critical Chance", float32(dexSpinner)*0.4+10), 200, 130, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "HP", float32(staSpinner)*10+100), 200, 145, 15, rl.Black)
	} else {
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Attack", float32(dexSpinner)*0.5+10), 200, 100, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Defense", float32(staSpinner)*0.4+10), 200, 115, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "Critical Chance", float32(dexSpinner)*0.4+10), 200, 130, 15, rl.Black)
		rl.DrawText(fmt.Sprintf("%s: %.3g", "HP", float32(staSpinner)*10+100), 200, 145, 15, rl.Black)
	}

}

func showBackground() {
	framesCounter++
	if framesCounter >= (60 / framesSpeed) {
		framesCounter = 0
		currentFrame++
		if currentFrame > 6 {
			currentFrame = 0
		}

		frameRec.X = currentFrame * float32(characterTexture.Width) / 7

	}

	groundFramesCounter++
	if groundFramesCounter >= (60 / gFramesSpeed) {
		groundFramesCounter = 0
		currentgroundFrame++
		if currentgroundFrame > 512 {
			currentgroundFrame = 0
		}
		groundFrame.X = currentgroundFrame * float32(bg.Width) / 512
	}

	rl.DrawTextureRec(sky, skyFrame, skyPosition, rl.White)
	rl.DrawTextureRec(bg, groundFrame, groundPosition, rl.White)
	rl.DrawTextureRec(characterTexture, frameRec, position, rl.White) // Draw part of the texture

}
