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
	i.Paletted = image.NewPaletted(image.Rect(0, 0, c.Width, c.Height), p)
}
