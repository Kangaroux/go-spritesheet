package spritesheet

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Sprite represents a single sprite within a sprite sheet.
type Sprite struct {
	Name  string
	Row   int
	Col   int
	Sheet *SpriteSheet
}

// Rect returns the area where this sprite is in the sprite sheet.
func (s *Sprite) Rect() image.Rectangle {
	p0 := image.Pt(s.Col*s.Sheet.Size, s.Row*s.Sheet.Size)
	p1 := image.Pt(p0.X+s.Sheet.Size, p0.Y+s.Sheet.Size)
	return image.Rectangle{Min: p0, Max: p1}
}

// SpriteSheet represents a sprite sheet config file loaded from YAML.
type SpriteSheet struct {
	Rows, Cols int
	Size       int
	Image      string
	Names      []string `yaml:"sprites"`
}

// Sprites returns a map of all the sprites declared in the sprite sheet.
// The map keys are the sprite names.
func (ss *SpriteSheet) Sprites() map[string]*Sprite {
	m := make(map[string]*Sprite)

	for i, name := range ss.Names {
		if name == "_" {
			continue
		}

		m[name] = &Sprite{
			Name:  name,
			Row:   int(math.Floor(float64(i) / float64(ss.Cols))),
			Col:   i % ss.Cols,
			Sheet: ss,
		}
	}

	return m
}

// OpenAndRead reads and returns the sprite sheet config file at the given path.
func OpenAndRead(path string) (*SpriteSheet, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()
	data, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return Read(bytes.NewReader(data))
}

// Read reads a sprite sheet config file, parses it, and returns it.
func Read(r io.Reader) (*SpriteSheet, error) {
	sheet := &SpriteSheet{}
	decoder := yaml.NewDecoder(r)
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

	// Check that all of the sprite names are unique
	if len(sheet.Names) > 0 {
		dupes := []string{}
		names := make(map[string]struct{})

		for _, name := range sheet.Names {
			if name == "_" {
				continue
			}

			if _, exists := names[name]; exists {
				dupes = append(dupes, name)
			} else {
				names[name] = struct{}{}
			}
		}

		if len(dupes) > 0 {
			return nil, fmt.Errorf(
				"sprite names must be unique (duplicated: %s)",
				strings.Join(dupes, ", "),
			)
		}
	}

	return sheet, nil
}
