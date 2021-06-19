package spritesheet_test

import (
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
		require.Equal(t, sheet, test.expected)
	}
}
