package spritesheet_test

import (
	"testing"

	ss "github.com/Kangaroux/go-spritesheet"
	"github.com/stretchr/testify/require"
)

func Test_ReadSpriteSheet(t *testing.T) {
	tests := []struct {
		in       string
		expected *ss.SpriteSheet
	}{
		{
			in: `
rows: 1
cols: 2
size: 3`,
			expected: &ss.SpriteSheet{
				Rows: 1,
				Cols: 2,
				Size: 3,
			},
		},
	}

	for _, test := range tests {
		sheet, err := ss.ReadSpriteSheet([]byte(test.in))

		require.NoError(t, err)
		require.Equal(t, sheet, test.expected)
	}
}
