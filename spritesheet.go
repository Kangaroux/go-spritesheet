package spritesheet

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Sprite struct {
	Path string
	Name string
}

type SpriteSheet struct {
	Rows, Cols int
	Size       int
	Sprites    []Sprite
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
	}

	return sheet, nil
}
