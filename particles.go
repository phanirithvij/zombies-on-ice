package main

import (
	"github.com/go-gl/gl/v3.3-compatibility/gl"

	"github.com/adinfinit/zombies-on-ice/g"
)

const (
	// DecalFadeTime fade time
	DecalFadeTime = 15.0
	// DecalMax max
	DecalMax = 20.0
)

// Particles a list of particles
type Particles struct {
	List []*Particle

	DecalHead int
	Decals    [2048]Particle
}

// Particle a particle
type Particle struct {
	Position        g.V2
	Velocity        g.V2
	Rotation        float32
	AngularVelocity float32
	Radius          float32
	Life            float32
	Fade            float32
}

// NewParticles a new instance of particles
func NewParticles() *Particles { return &Particles{} }

// Spawn spawns the particles
func (ps *Particles) Spawn(amount int, position g.V2, velocity g.V2, radius float32, spread float32) {
	for i := 0; i < amount; i++ {
		rotate := g.RandomBetween(-spread/2, spread/2)
		speed := g.RandomBetween(-spread, spread)
		ps.List = append(ps.List, &Particle{
			Position:        position,
			Velocity:        velocity.Rotate(rotate).Scale(1 + speed),
			Rotation:        g.RandomBetween(0, 7),
			AngularVelocity: g.RandomBetween(-spread, spread),
			Radius:          g.RandomBetween(radius, radius*2),
			Life:            g.RandomBetween(0.1, 0.5),
		})
	}
}

// Update updates the particles
func (ps *Particles) Update(dt float32) {
	for i := range ps.Decals {
		ps.Decals[i].Fade -= dt
	}

	for _, p := range ps.List {
		p.Position = p.Position.AddScale(p.Velocity, dt)
		p.Velocity = p.Velocity.Scale(g.Pow(0.9, dt))
		p.Life -= dt
		p.Radius -= dt * 0.2
		p.Rotation += p.AngularVelocity
	}
}

func (ps *Particles) decalize(p *Particle) {
	ps.DecalHead = ps.DecalHead + 1
	if ps.DecalHead >= len(ps.Decals) {
		ps.DecalHead = 0
	}

	p.Fade = DecalFadeTime
	ps.Decals[ps.DecalHead] = *p
}

// Kill kills particles
func (ps *Particles) Kill(bounds g.Rect) {
	list := ps.List[:0:cap(ps.List)]
	for _, p := range ps.List {
		if p.Radius < 0.0 {
			continue
		}
		if p.Life < 0.0 || !bounds.Contains(p.Position) {
			g.EnforceInside(&p.Position, &p.Velocity, bounds, 0.0)
			ps.decalize(p)
			continue
		}
		list = append(list, p)
	}
	ps.List = list
}

// RenderDecals renders decals
func (ps *Particles) RenderDecals(game *Game) {
	tex := game.Assets.TextureRepeat("assets/blood.png")
	for i := range ps.Decals {
		p := &ps.Decals[i]
		if p.Fade <= 0 {
			continue
		}

		gl.PushMatrix()
		gl.Translatef(p.Position.X, p.Position.Y, 0)
		gl.Rotatef(g.RadToDeg(p.Rotation), 0, 0, -1)

		sat := g.Sat8(p.Fade / DecalMax)
		color := g.Color{R: sat, G: sat, B: sat, A: sat}
		tex.DrawColored(g.NewCircleRect(p.Radius), color)
		gl.PopMatrix()
	}
}

// Render renders the particles
func (ps *Particles) Render(game *Game) {
	tex := game.Assets.TextureRepeat("assets/blood.png")
	for _, p := range ps.List {
		gl.PushMatrix()
		gl.Translatef(p.Position.X, p.Position.Y, 0)
		gl.Rotatef(g.RadToDeg(p.Rotation), 0, 0, -1)
		tex.Draw(g.NewCircleRect(p.Radius))
		gl.PopMatrix()
	}
}
