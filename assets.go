package main

import "github.com/adinfinit/zombies-on-ice/g"

// Assets assets struct
type Assets struct {
	Textures map[string]*g.Texture
}

// NewAssets returns a new Assets struct
func NewAssets() *Assets {
	return &Assets{
		Textures: make(map[string]*g.Texture),
	}
}

// Reload reload the textures
func (assets *Assets) Reload() {
	for _, tex := range assets.Textures {
		tex.Reload()
	}
}

// Texture get a texture by path
func (assets *Assets) Texture(path string) *g.Texture { return assets.texture(path, false) }

// SpriteFont font
func (assets *Assets) SpriteFont(path string, glyphSize g.V2, glyphs string) *g.Font {
	tex := assets.Texture(path)

	return &g.Font{
		Texture:   tex,
		Glyphs:    glyphs,
		GlyphSize: glyphSize,
	}
}

// TextureRepeat repeat a texture
func (assets *Assets) TextureRepeat(path string) *g.Texture { return assets.texture(path, true) }

func (assets *Assets) texture(path string, repeat bool) *g.Texture {
	npath := path
	if repeat {
		npath = "@" + path
	}

	tex, ok := assets.Textures[npath]
	if !ok {
		tex = &g.Texture{}
		tex.Path = path
		tex.Repeat = repeat
		tex.Reload()

		assets.Textures[npath] = tex
	}

	return tex
}

// Unload unloads all the textures
func (assets *Assets) Unload() {
	for _, tex := range assets.Textures {
		tex.Delete()
	}
	*assets = *NewAssets()
}
