package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	worldWidth   = 2000
	worldHeight  = 2000
)

type Player struct {
	x int
	y int
}

type Game struct {
	player Player
}

func (g *Game) Update() error {
	// Mettre à jour la position du joueur en fonction de l'entrée du joueur
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.x -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.x += 5
	}

	// Limiter la position du joueur à la taille du monde
	if g.player.x < 0 {
		g.player.x = 0
	}
	if g.player.x > worldWidth {
		g.player.x = worldWidth
	}
	if g.player.y < 0 {
		g.player.y = 0
	}
	if g.player.y > worldHeight {
		g.player.y = worldHeight
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Calculer les coordonnées de la caméra en fonction de la position du joueur
	cameraX := g.player.x - screenWidth/2
	cameraY := g.player.y - screenHeight/2

	// Afficher le joueur sur l'écran en prenant en compte la position de la caméra
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.player.x-cameraX), float64(g.player.y-cameraY))
	screen.DrawImage(playerImage, op)

	// Afficher le texte avec les coordonnées du joueur en haut à gauche de l'écran
	text := "Position: (" + fmt.Sprint(g.player.x) + ", " + fmt.Sprint(g.player.y) + ")"
	text2 := "Mouse: (" + fmt.Sprint(cameraX) + ", " + fmt.Sprint(cameraY) + ")"
	ebitenutil.DebugPrint(screen, text)
	ebitenutil.DebugPrint(screen, text2)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var playerImage *ebiten.Image

func main() {
	// Charger l'image du joueur
	img, _, err := ebitenutil.NewImageFromFile("image.png")
	if err != nil {
		panic(err)
	}
	playerImage = img

	// Initialiser le jeu avec la position du joueur au centre du monde
	game := &Game{
		player: Player{x: worldWidth / 2, y: worldHeight / 2},
	}

	// Lancer le jeu avec la résolution spécifiée
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Camera system with Ebiten")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
