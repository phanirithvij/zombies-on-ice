package main

import "github.com/adinfinit/zombies-on-ice/g"

// Room a room
type Room struct {
	Bounds g.Rect

	TextureScale float32
}

// NewRoom a new instance of a Room
func NewRoom() *Room {
	room := &Room{}

	//1600, 1000
	room.Bounds.Min = g.V2{X: -16, Y: -10}
	room.Bounds.Max = g.V2{X: 16, Y: 10}
	room.TextureScale = 0.5

	return room
}

// Render renders the room
func (room *Room) Render(game *Game) {
	if false {
		ground := game.Assets.TextureRepeat("assets/ground.png")

		ground.DrawSub(
			room.Bounds,
			g.Rect{
				Min: g.V2{X: 0, Y: 0},
				Max: room.Bounds.Size().Scale(room.TextureScale),
			},
		)
	} else {
		ground := game.Assets.Texture("assets/room.png")
		ground.Draw(room.Bounds)
	}
}
