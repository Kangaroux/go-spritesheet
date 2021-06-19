package spritesheet

import (
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

	if err := yaml.Unmarshal(data, sheet); err != nil {
		return nil, err
	}

	return sheet, nil
}
