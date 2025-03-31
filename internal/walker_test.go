package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNewPath(t *testing.T) {
	new_path := GetNewPath(
		"/home/l/.local/share/nvim/lee/pack/plug/opt/auto-session/nvim/test.lua",
		"/home/l/.local/share/nvim/lee/pack/plug/opt/auto-session/",
		"/home/l/temp/packnvim/",
	)

	var expected = "/home/l/temp/packnvim/nvim/test.lua"
	assert.Equal(t, expected, new_path, "Should be equal")
}
