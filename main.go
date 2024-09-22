package main

import (
    "log"
    "fmt"
    "github.com/hajimehoshi/ebiten/v2"
    // "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "image"
   // "image/color"
    "image/png"
    "os"
)

type Game struct{
  player Player
}

type Player struct {
  x, y float64
  width, height int
}

func (g *Game) Update() error {
  // Handle Keyboard Input
//  currentKeyPress := ebiten.IsKeyPressed(ebiten.KeyA)
  controller := map[string]bool {
    "up_arrow": ebiten.IsKeyPressed(ebiten.KeyArrowUp),
    "down_arrow": ebiten.IsKeyPressed(ebiten.KeyArrowDown),
    "right_arrow": ebiten.IsKeyPressed(ebiten.KeyArrowRight),
    "left_arrow": ebiten.IsKeyPressed(ebiten.KeyArrowLeft),
  }

  switch {
    case controller["up_arrow"]:
      MoveUp(&g.player)
    case controller["down_arrow"]:
      MoveDown(&g.player)
    case controller["right_arrow"]:
      MoveRight(&g.player)
    case controller["left_arrow"]:
      MoveLeft(&g.player)
  }

  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

  // Create Wizard image
  wizard_img_file, err := os.Open("resources/wizard.png")

  if err != nil {
    fmt.Println(err) 
  }
  img_png, err := png.Decode(wizard_img_file)
  if err != nil {
    fmt.Println(err)
  }
  newImg := ebiten.NewImageFromImage(img_png)

  subImgRect := image.Rect(0, 0, 16, 16)
  myImg := newImg.SubImage(subImgRect).(*ebiten.Image)

 //  myImg := ebiten.NewImage(g.player.width, g.player.height)
  //myImg.Fill(color.RGBA{200, 0, 0, 200})
  op := &ebiten.DrawImageOptions{}
  op.GeoM.Translate(g.player.x, g.player.y)
  screen.DrawImage(myImg, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
  return 320, 240
}

func main() {
  g := &Game{
    player: Player {
      x: 50,
      y: 50,
      width: 50,
      height: 50,
    },
  }
  if err := ebiten.RunGame(g); err != nil {
    log.Fatal(err)
  }
}
