package main

import (
	"github.com/go-gl/gl/v3.3-compatibility/gl"

	"github.com/adinfinit/zombies-on-ice/g"
)

// Powerup a powerup
type Powerup struct {
	Entity

	Life     float32
	Rotation float32

	Health       float32
	HammerRadius float32
	HammerMass   float32
}

// NewPowerup a new instance of Powerup
func NewPowerup(bounds g.Rect) *Powerup {
	powerup := &Powerup{}

	powerup.Position = g.RandomV2(bounds.ScaleInv(g.V2{X: 2, Y: 2}))
	powerup.Radius = 0.5
	powerup.Mass = 1
	powerup.CollisionLayer = PowerupLayer
	powerup.CollisionMask = PlayerLayer
	powerup.CollisionTrigger = true

	powerup.Life = 10.0

	powerup.Health = 0.5
	powerup.HammerRadius = 0.15
	powerup.HammerMass = 0.5

	return powerup
}

// Entites entities in this powerup
func (powerup *Powerup) Entites() []*Entity {
	return []*Entity{&powerup.Entity}
}

// Update updates the powerup life and rotation
func (powerup *Powerup) Update(dt float32) {
	powerup.Life -= dt
	powerup.Rotation += dt
}

// Done if powerup needs to be removed
func (powerup *Powerup) Done() bool {
	return powerup.Life < 0.0 || len(powerup.Collision) > 0
}

// Apply applies the powerup to player
func (powerup *Powerup) Apply(player *Player) {
	player.Health = g.Clamp01(player.Health + powerup.Health)
	player.Hammer.Radius += powerup.HammerRadius
	player.Hammer.Mass += powerup.HammerMass
}

// Render renders the powerup
func (powerup *Powerup) Render(game *Game) {
	if powerup.Life < 3.0 {
		if g.Mod(powerup.Life*1.5, 1) < 0.4 {
			return
		}
	}

	gl.PushMatrix()
	{
		gl.Translatef(powerup.Position.X, powerup.Position.Y, 0)
		gl.Rotatef(g.RadToDeg(powerup.Rotation), 0, 0, -1)

		tex := game.Assets.Texture("assets/healthpack.png")
		tex.Draw(g.NewCircleRect(powerup.Radius))
	}
	gl.PopMatrix()
}
