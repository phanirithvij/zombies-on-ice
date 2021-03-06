package main

import (
	"math/rand"

	"github.com/go-gl/gl/v3.3-compatibility/gl"

	"github.com/adinfinit/zombies-on-ice/g"
)

// Zombie a zombie
type Zombie struct {
	Entity

	Direction g.V2
	Distance  float32
	Frame     int
}

// NewZombie a new instance of zombie
func NewZombie(bounds g.Rect) *Zombie {
	zombie := &Zombie{}

	zombie.respawn(bounds)

	zombie.Elasticity = 0.2
	zombie.Mass = 0.5
	zombie.Dampening = 0.9
	zombie.Radius = 0.5

	zombie.Direction = g.V2{}

	zombie.CollisionLayer = ZombieLayer
	zombie.CollisionMask = HammerLayer

	return zombie
}

// Update updates the zombie
func (zombie *Zombie) Update(game *Game, dt float32) {
	var nearest *Entity
	mindist := float32(1000000.0)
	for _, player := range game.Players {
		if player.Dead {
			continue
		}

		dist := player.Survivor.Position.Sub(zombie.Position).Length()
		if dist < mindist {
			nearest = &player.Survivor
			mindist = dist
		}
	}

	if nearest == nil {
		return
	}

	direction := nearest.Position.Sub(zombie.Position).Normalize()
	zombie.AddForce(direction.Scale(1))

	lateral := direction.Rotate90c().Normalize()
	scale := lateral.Dot(zombie.Velocity)

	zombie.AddForce(lateral.Scale(-scale * 2.0))
	zombie.Direction = direction.Add(zombie.Velocity.Normalize()).Scale(0.5)

	zombie.Distance += zombie.Velocity.Scale(dt).Length()
	if zombie.Distance > 0.1 {
		zombie.Distance -= 0.1
		zombie.Frame++
	}
}

func (zombie *Zombie) respawn(bounds g.Rect) {
	zombie.Position = g.RandomV2(bounds)
	switch rand.Intn(4) {
	case 0:
		zombie.Position.X = bounds.Min.X
	case 1:
		zombie.Position.X = bounds.Max.X
	case 2:
		zombie.Position.Y = bounds.Min.Y
	case 3:
		zombie.Position.Y = bounds.Max.Y
	}
	zombie.Velocity = g.V2{}
}

// DeathStrength retuns the death strength
func (zombie *Zombie) DeathStrength() (float32, bool) {
	if len(zombie.Collision) == 0 {
		return 0, false
	}

	total := float32(0.0)
	for _, collision := range zombie.Collision {
		total += collision.VelocityDelta.Length()
	}
	return total, total > 1.0
}

// Respawn respawns a zombie
func (zombie *Zombie) Respawn(bounds g.Rect) {
	_, dead := zombie.DeathStrength()
	if dead {
		zombie.respawn(bounds)
	}
}

// Render renders the zombie
func (zombie *Zombie) Render(game *Game) {
	gl.PushMatrix()
	{
		gl.Translatef(zombie.Position.X, zombie.Position.Y, 0)

		rotation := -(zombie.Direction.Angle() + g.Tau/4)
		gl.Rotatef(g.RadToDeg(rotation), 0, 0, -1)

		var tex *g.Texture
		if zombie.Velocity.Length() < 0.1 {
			tex = game.Assets.TextureRepeat("assets/zombie-idle.png")
		} else {
			if zombie.Frame&1 == 0 {
				tex = game.Assets.TextureRepeat("assets/zombie-walk-0.png")
			} else {
				tex = game.Assets.TextureRepeat("assets/zombie-walk-1.png")
			}
		}
		tex.Draw(g.NewCircleRect(zombie.Radius))
	}
	gl.PopMatrix()
}
