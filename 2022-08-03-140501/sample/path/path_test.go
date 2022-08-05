package path

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	samplePathStr := "sample.txt"
	samplePath := Path{Filepath: samplePathStr}

	t.Run("Abs", func(t *testing.T) {
		absPathStr, err := filepath.Abs(samplePathStr)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, samplePath.Abs().Filepath, absPathStr)
	})
}
