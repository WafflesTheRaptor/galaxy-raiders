# Represent a Player {#chap:player}

<!-- test:suite=player -->

<!-- test:file=go.mod -->
```
module example.com/developer/space-invaders
```

We will represent the player as a rectangle.  For the moment, we will fill that rectangle in blue.  But, in the future we can draw an image inside the rectangle or even use an animated sprite.

Representing the player as a rectangle allows us to move the player around the screen.  If you've taken high-school math, then you'll have encountered the term `translate`.  In order to move the player around the screen we will _translate_ the player by some distance in a given direction.  Again, high-school math students may recognize that we're translating the rectangle by a _vector_.

## Visual Representation of Translation

## Translation in Code

Represent a vector:

```go
type Vector = image.Point
```

Represent a player as a rectangle:

```go
type Player image.Rectangle

func NewPlayer(min image.Point, max image.Point) Player {
	return Player{
		Min: min,
		Max: max,
	}
}
```

Implement translation

```go
func (p *Player) Translate(v Vector) {
	p.Min = p.Min.Add(v)
	p.Max = p.Max.Add(v)
}
```

<!-- test:file=main.go -->
```go
package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Vector = image.Point
type Player image.Rectangle

func NewPlayer(min image.Point, max image.Point) Player {
	return Player{
		Min: min,
		Max: max,
	}
}

func (p *Player) Translate(v Vector) {
	p.Min = p.Min.Add(v)
	p.Max = p.Max.Add(v)
}

type Game struct {
    player Player
}

func NewGame() *Game {
	g := Game{}
	p := NewPlayer(image.Pt(0, 0), image.Pt(20, 5))
	p.Translate(Vector{screenWidth / 2, screenHeight - 10})
	g.player = p
	return &g
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw player
	ebitenutil.DrawRect(screen,
		float64(g.player.Min.X), float64(g.player.Min.Y),
		float64(g.player.Max.X-g.player.Min.X), float64(g.player.Max.Y-g.player.Min.Y),
		color.RGBA{0, 0, 255, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Space Invaders")
	g := &Game{}
    ebiten.RunGame(g)
}
```

<!-- test:exec -->
```command
go mod tidy
```

<!-- test:exec -->
```
go build -o space-invaders
```

## Update the Player Position

We want the player to go left when the left arrow is pressed, and right when the right arrow is pressed.  Our game engine will call the game `Update` function 60 times per second, so it's important that we move the player only by a small amount each time.

```go
func (g *Game) Update() error {
	// Update the player
	g.player.Update()
}

func (p *Player) Update() {
	// if left key down, move left
	// if right key down move right
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if p.Min.X > 0 {
			p.Translate(Vector{X: -2, Y: 0})
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if p.Max.X < screenWidth {
			p.Translate(Vector{X: 2, Y: 0})
		}
	}
}
```
