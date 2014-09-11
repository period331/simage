package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hanguofeng/freetype-go-mirror/freetype/truetype"
)

type Fonts struct {
	files      []string
	objects    map[string]*truetype.Font
	randObject *rand.Rand
}

func NewFonts(fontPath string) *Fonts {
	f := new(Fonts)
	err := f.AddFonts(fontPath, false)
	if err != nil {
		panic(err)
	}
	f.objects = make(map[string]*truetype.Font)
	f.randObject = rand.New(rand.NewSource(time.Now().UnixNano()))

	return f
}

func (f *Fonts) AddFonts(path string, blockErr bool) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".ttf") {
			addError := f.AddFont(path)
			if !blockErr && addError != nil {
				return addError
			}
		}
		return nil
	})

	return nil
}

func (f *Fonts) AddFont(filePath string) error {
	fontBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		return err
	}
	f.files = append(f.files, filePath)
	f.objects[filePath] = font

	return nil
}
