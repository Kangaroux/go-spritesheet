package spritesheet_test

import (
	"image"
	"strings"
	"testing"

	ss "github.com/Kangaroux/go-spritesheet"
	"github.com/stretchr/testify/require"
)

func Test_Read_Error(t *testing.T) {
	tests := []struct {
		in string
	}{
		// EOF
		{
			in: ``,
		},
		// Cannot unmarshal string to struct
		{
			in: `foo`,
		},
		// Unknown field foo
		{
			in: `foo: bar`,
		},
		// Rows < 1
		{
			in: `rows: 0`,
		},
		// Cols < 1
		{
			in: `cols: 0`,
		},
		// Size < 1
		{
			in: `size: 0`,
		},
		// Missing sprites field
		{
			in: `
rows: 1
cols: 1
size: 1
image: foo`,
		},
		// Missing image field
		{
			in: `
rows: 1
cols: 1
size: 1
sprites: []`,
		},
		// Sprites field has too many entries
		{
			in: `
rows: 1
cols: 1
size: 1
image: foo
sprites: [a, b]`,
		},
	}

	for _, test := range tests {
		_, err := ss.Read(strings.NewReader(test.in))
		require.Error(t, err)
	}
}

func Test_Read_OK(t *testing.T) {
	tests := []struct {
		in       string
		expected *ss.SpriteSheet
	}{
		{
			in: `
rows: 1
cols: 2
size: 3
image: foo.png
sprites: []`,
			expected: &ss.SpriteSheet{
				Rows:  1,
				Cols:  2,
				Size:  3,
				Image: "foo.png",
				Names: []string{},
			},
		},
		{
			in: `
rows: 2
cols: 2
size: 3
image: foo.png
sprites: [a, b, c, d]`,
			expected: &ss.SpriteSheet{
				Rows:  2,
				Cols:  2,
				Size:  3,
				Image: "foo.png",
				Names: []string{"a", "b", "c", "d"},
			},
		},
	}

	for _, test := range tests {
		sheet, err := ss.Read(strings.NewReader(test.in))

		require.NoError(t, err)
		require.Equal(t, test.expected, sheet)
	}
}

func Test_SpriteSheet_Sprites(t *testing.T) {
	tests := []struct {
		sheet    *ss.SpriteSheet
		expected map[string]*ss.Sprite
	}{
		// Empty
		{
			sheet: &ss.SpriteSheet{
				Rows:  1,
				Cols:  1,
				Names: []string{},
			},
			expected: map[string]*ss.Sprite{},
		},
		// 1x1 sheet, one sprite
		{
			sheet: &ss.SpriteSheet{
				Rows:  1,
				Cols:  1,
				Names: []string{"foo"},
			},
			expected: map[string]*ss.Sprite{
				"foo": {
					Name: "foo",
					Row:  0,
					Col:  0,
				},
			},
		},
		// 2x2 sheet, one sprite
		{
			sheet: &ss.SpriteSheet{
				Rows:  2,
				Cols:  2,
				Names: []string{"foo"},
			},
			expected: map[string]*ss.Sprite{
				"foo": {
					Name: "foo",
					Row:  0,
					Col:  0,
				},
			},
		},
		// 2x2 sheet, four sprites
		{
			sheet: &ss.SpriteSheet{
				Rows:  2,
				Cols:  2,
				Names: []string{"a", "b", "c", "d"},
			},
			expected: map[string]*ss.Sprite{
				"a": {
					Name: "a",
					Row:  0,
					Col:  0,
				},
				"b": {
					Name: "b",
					Row:  0,
					Col:  1,
				},
				"c": {
					Name: "c",
					Row:  1,
					Col:  0,
				},
				"d": {
					Name: "d",
					Row:  1,
					Col:  1,
				},
			},
		},
		// 1x3 sheet, three sprites
		{
			sheet: &ss.SpriteSheet{
				Rows:  1,
				Cols:  3,
				Names: []string{"a", "b", "c"},
			},
			expected: map[string]*ss.Sprite{
				"a": {
					Name: "a",
					Row:  0,
					Col:  0,
				},
				"b": {
					Name: "b",
					Row:  0,
					Col:  1,
				},
				"c": {
					Name: "c",
					Row:  0,
					Col:  2,
				},
			},
		},
		// 1x3 sheet, three sprites, skip
		{
			sheet: &ss.SpriteSheet{
				Rows:  1,
				Cols:  3,
				Names: []string{"a", "_", "c"},
			},
			expected: map[string]*ss.Sprite{
				"a": {
					Name: "a",
					Row:  0,
					Col:  0,
				},
				"c": {
					Name: "c",
					Row:  0,
					Col:  2,
				},
			},
		},
	}

	for _, test := range tests {
		for i, _ := range test.expected {
			test.expected[i].Sheet = test.sheet
		}

		require.Equal(t, test.expected, test.sheet.Sprites())
	}
}

func Test_Sprite_Rect(t *testing.T) {
	tests := []struct {
		size     int
		sprite   *ss.Sprite
		expected image.Rectangle
	}{
		{
			size: 2,
			sprite: &ss.Sprite{
				Row: 0,
				Col: 0,
			},
			expected: image.Rect(0, 0, 2, 2),
		},
		{
			size: 2,
			sprite: &ss.Sprite{
				Row: 0,
				Col: 1,
			},
			expected: image.Rect(2, 0, 4, 2),
		},
		{
			size: 2,
			sprite: &ss.Sprite{
				Row: 1,
				Col: 0,
			},
			expected: image.Rect(0, 2, 2, 4),
		},
		{
			size: 2,
			sprite: &ss.Sprite{
				Row: 1,
				Col: 1,
			},
			expected: image.Rect(2, 2, 4, 4),
		},
	}

	for _, test := range tests {
		test.sprite.Sheet = &ss.SpriteSheet{
			Size: test.size,
		}
		require.Equal(t, test.expected, test.sprite.Rect())
	}
}
