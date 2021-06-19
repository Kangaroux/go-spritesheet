package spritesheet

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"gopkg.in/yaml.v3"
)

type Sprite struct {
	Name string
	Row  int
	Col  int
}

type SpriteSheet struct {
	Rows, Cols int
	Size       int
	Image      string
	Names      []string `yaml:"sprites"`
}

func (ss *SpriteSheet) Sprites() []Sprite {
	sprites := []Sprite{}

	for i, name := range ss.Names {
		sprites = append(sprites, Sprite{
			Name: name,
			Row:  int(math.Floor(float64(i) / float64(ss.Cols))),
			Col:  i % ss.Cols,
		})
	}

	return sprites
}

func OpenAndReadSpriteSheet(path string) (*SpriteSheet, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return ReadSpriteSheet(data)
}

func ReadSpriteSheet(data []byte) (*SpriteSheet, error) {
	sheet := &SpriteSheet{}
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	if err := decoder.Decode(sheet); err != nil {
		return nil, err
	}

	if sheet.Rows < 1 {
		return nil, errors.New("rows must be at least 1")
	} else if sheet.Cols < 1 {
		return nil, errors.New("cols must be at least 1")
	} else if sheet.Size < 1 {
		return nil, errors.New("size must be at least 1")
	} else if sheet.Image == "" {
		return nil, errors.New("missing image field")
	} else if sheet.Names == nil {
		return nil, errors.New("missing sprites field")
	} else if len(sheet.Names) > sheet.Cols*sheet.Rows {
		return nil, fmt.Errorf(
			"sprites field has too many entries (%d entries, max is %d)",
			len(sheet.Names),
			sheet.Cols*sheet.Rows,
		)
	}

	return sheet, nil
}
