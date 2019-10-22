package battle

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/thiagotrennepohl/gladiator-game/models"
)

var (
	playerAttack = false
	startBattle  = false
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
	enemyPosition       rl.Vector2
	frameRec            rl.Rectangle
	bg                  rl.Texture2D
	sky                 rl.Texture2D
	enemy               rl.Texture2D
	enemyPlayer         models.Character
	enemyRectangle      rl.Rectangle
	playerTurn          = true
)

func InitStuff() {
	position = rl.NewVector2(100.0, 325)
	enemyPosition = rl.NewVector2(700, 325)
	characterTexture = rl.LoadTexture("assets/img/idle.png")
	bg = rl.LoadTexture("assets/img/back.png")
	sky = rl.LoadTexture("assets/img/sky.png")
	enemy = rl.LoadTexture("assets/img/enemy2.png")
	frameRec = rl.NewRectangle(0, 0, float32(characterTexture.Width/5), float32(characterTexture.Height))
	enemyRectangle = rl.NewRectangle(0, 0, float32(enemy.Width/7), float32(enemy.Height))
}

func Start(player models.Character, opponentSvc models.OpponentService) {
	showBackground()
	//Checks if the enemy has already been created
	if enemyPlayer == nil {
		//Leveling system not ready yet.
		// enemyPlayer, _ = opponentSvc.CreateNewOpponent(player.GetLevel())
		enemyPlayer, _ = opponentSvc.CreateNewOpponent(1)
	}
	drawPlayer()
	drawEnemy()
	// if match.IsInProgress() {
	// 	 := match.TossCoin()
	// }

	if startBattle {
		rl.DrawText(fmt.Sprintf("PLAYER HP: %.3g", player.GetHealthPoints()), 100, 120, 20, rl.Black)
		rl.DrawText(fmt.Sprintf("ENEMY HP: %.3g", enemyPlayer.GetHealthPoints()), 600, 120, 20, rl.Black)
		if playerTurn {
			playerAttack = raygui.Button(rl.NewRectangle(310, 300, 200, 40), "Attack")
			if playerAttack {
				playerAttackPoints := player.GetAttackPoints() - enemyPlayer.GetDefensePoints()
				if playerAttackPoints > 0 {
					enemyPlayer.DecreaseFromHealthPoints(playerAttackPoints)
				}
				playerTurn = false
			}
		} else {
			enemyAttackPoints := enemyPlayer.GetAttackPoints() - player.GetDefensePoints()
			if enemyAttackPoints > 0 {
				player.DecreaseFromHealthPoints(enemyAttackPoints)
			}
			playerTurn = true

		}

	} else {
		startBattle = raygui.Button(rl.NewRectangle(310, 300, 200, 40), "Start Battle")
	}

}

func showBackground() {
	rl.DrawTextureRec(sky, skyFrame, skyPosition, rl.White)
	rl.DrawTextureRec(bg, groundFrame, groundPosition, rl.White)
}

func drawPlayer() {
	framesCounter++
	if framesCounter >= (60 / framesSpeed) {
		framesCounter = 0
		currentFrame++
		if currentFrame > 4 {
			currentFrame = 0
		}

		frameRec.X = currentFrame * float32(characterTexture.Width) / 5

	}
	rl.DrawTextureRec(characterTexture, frameRec, position, rl.White) // Draw part of the texture
}

func drawEnemy() {
	rl.DrawTextureRec(enemy, enemyRectangle, enemyPosition, rl.White)
}
