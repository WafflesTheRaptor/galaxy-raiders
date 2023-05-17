// Copyright 2022 Aidan Delaney <aidan.delaney@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240

	MaxRows         = 3
	MaxAliensPerRow = 6
)

type Vector = image.Point

type Alien struct {
	image.Rectangle
	direction Vector
}

func NewAlien(min image.Point, max image.Point) Alien {
	return Alien{
		Rectangle: image.Rectangle{
			Min: min,
			Max: max,
		},
		direction: Vector{X: 2, Y: 0},
	}
}

func (a *Alien) Update() {
	// move in the direction
	// Right, Down, Left
	// if next update would take us too far right, move down, flip direction
	// if next update would take us too far left, move down, flip direction
	if a.Min.X < 0 || a.Max.X > screenWidth {
		a.Translate(Vector{X: 0, Y: 20})
		a.direction.X = a.direction.X * -1
	}

	a.Translate(a.direction)
}

func (a *Alien) Translate(v Vector) {
	a.Rectangle.Min = a.Rectangle.Min.Add(v)
	a.Rectangle.Max = a.Rectangle.Max.Add(v)
}

type Player image.Rectangle

func NewPlayer(min image.Point, max image.Point) Player {
	return Player{
		Min: min,
		Max: max,
	}
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

func (p *Player) Translate(v Vector) {
	p.Min = p.Min.Add(v)
	p.Max = p.Max.Add(v)
}

type Bullet image.Rectangle

func NewBullet(loc image.Point) Bullet {
	b := Bullet{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: 2,
			Y: 2,
		},
	}
	b.Translate(loc)
	return b
}

func (b *Bullet) Update() {
	b.Translate(Vector{X: 0, Y: -2})
}

func (b *Bullet) Translate(v Vector) {
	b.Min = b.Min.Add(v)
	b.Max = b.Max.Add(v)
}

type Game struct {
	aliens  []Alien
	player  Player
	bullets []Bullet
	armed   sync.Mutex
}

func NewGame() *Game {
	g := Game{}
	g.aliens = make([]Alien, MaxRows*MaxAliensPerRow)
	for i := 0; i < MaxRows; i++ {
		for j := 0; j < MaxAliensPerRow; j++ {
			a := NewAlien(image.Pt(0, 0), image.Pt(10, 10))
			a.Translate(Vector{X: 50 * j, Y: 30 * i})
			g.aliens[(i*MaxAliensPerRow)+j] = a
		}
	}

	p := NewPlayer(image.Pt(0, 0), image.Pt(20, 5))
	p.Translate(Vector{screenWidth / 2, screenHeight - 10})
	g.player = p

	ticker := time.NewTicker(250 * time.Millisecond)
	go func() {
		for range ticker.C {
			g.armed.TryLock()
			g.armed.Unlock()
		}
	}()
	return &g
}

func (g *Game) Update() error {
	// Update the player
	g.player.Update()

	// Cull offscreen bullets
	for i, b := range g.bullets {
		if b.Min.Y < 0 {
			g.bullets = append(g.bullets[:i], g.bullets[i+1:]...)
		}
	}

	// Detect intersection of bullets & aliens
	for k, b := range g.bullets {
		for i := range g.aliens {
			r := image.Rectangle{
				Min: b.Min,
				Max: b.Max,
			}
			if g.aliens[i].Overlaps(r) {
				g.bullets = append(g.bullets[:k], g.bullets[k+1:]...)
				g.aliens = append(g.aliens[:i], g.aliens[i+1:]...)
				break
			}
		}
	}

	// Update all aliens
	for i := range g.aliens {
		g.aliens[i].Update()
	}

	// Update all bullets
	for i := range g.bullets {
		g.bullets[i].Update()
	}

	// Create new bullet if spacebar is pressed and we're armed
	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.armed.TryLock() {
		b := NewBullet(g.player.Min.Add(g.player.Max).Div(2))
		g.bullets = append(g.bullets, b)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, a := range g.aliens {
		ebitenutil.DrawRect(screen,
			float64(a.Min.X), float64(a.Min.Y),
			float64(a.Max.X-a.Min.X), float64(a.Max.Y-a.Min.Y),
			color.RGBA{255, 0, 0, 255})
	}

	for _, b := range g.bullets {
		ebitenutil.DrawRect(screen,
			float64(b.Min.X), float64(b.Min.Y),
			float64(b.Max.X-b.Min.X), float64(b.Max.Y-b.Min.Y),
			color.RGBA{255, 255, 255, 255})
	}

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
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
