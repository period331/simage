package images

import (
	"image"
)

type CaptchaConfig struct {
	Width     int
	Height    int
	FontSize  int
	FontFiles []string
}

type Captcha struct {
	*image.Paletted
	config *CaptchaConfig
}

func NewCaptcha(c *CaptchaConfig) *Captcha {
	i := new(Captcha)
	i.config = c
	i.Paletted = image.NewPaletted(image.Rect(0, 0, c.Width, c.Height), randomPalette())
}

func randomPalette() color.Palette {

	p := make([]color.Color, colorCount+1)
	// Transparent color.
	p[0] = color.RGBA{0xFF, 0xFF, 0xFF, 0x00}
	// Primary color.
	prim := color.RGBA{
		uint8(rnd(0, 255)),
		uint8(rnd(0, 255)),
		uint8(rnd(0, 255)),
		0xFF,
	}
	p[1] = prim
	// Circle colors.
	for i := 2; i <= colorCount; i++ {
		p[i] = randomBrightness(prim, 255)
	}

	return p
}

func randomBrightness(c color.RGBA, max uint8) color.RGBA {
	minc := min3(c.R, c.G, c.B)
	maxc := max3(c.R, c.G, c.B)
	if maxc > max {
		return c
	}
	n := rnd(0, int(max-maxc)) - int(minc)
	return color.RGBA{
		uint8(int(c.R) + n),
		uint8(int(c.G) + n),
		uint8(int(c.B) + n),
		uint8(c.A),
	}
}

func min3(x, y, z uint8) (m uint8) {
	m = x
	if y < m {
		m = y
	}
	if z < m {
		m = z
	}
	return
}

func max3(x, y, z uint8) (m uint8) {
	m = x
	if y > m {
		m = y
	}
	if z > m {
		m = z
	}
	return
}
