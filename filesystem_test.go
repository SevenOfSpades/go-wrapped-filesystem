package filesystem

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeStreamer struct {
	buff *bytes.Buffer
}

func newFakeStreamer[T string | []byte](data T) *fakeStreamer {
	return &fakeStreamer{buff: bytes.NewBuffer([]byte(data))}
}

func (s *fakeStreamer) Read(p []byte) (n int, err error) {
	return s.buff.Read(p)
}

func (s *fakeStreamer) Close() error {
	s.buff.Reset()
	return nil
}

func TestReadContentOf(t *testing.T) {
	t.Run("it should pass execution to provided handler", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionReadContentOfHandler(func(path string) (Content, error) {
			assert.Equal(t, expectedPath, path)
			return []byte("TEST"), nil
		}))

		// WHEN
		result, err := ReadContentOf(fs, expectedPath)

		// THEN
		require.NoError(t, err)
		assert.Equal(t, "TEST", result.String())
	})
	t.Run("it should return up-stream error", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionReadContentOfHandler(func(path string) (Content, error) {
			assert.Equal(t, expectedPath, path)
			return nil, errors.New("something went wrong")
		}))

		// WHEN
		result, err := ReadContentOf(fs, expectedPath)

		// THEN
		require.EqualError(t, err, "failed to read content of path/to/file: something went wrong")
		assert.Nil(t, result)
	})
}

func TestStreamContentOf(t *testing.T) {
	t.Run("it should pass execution to provided handler", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionStreamContentOfHandler(func(path string) (Streamer, error) {
			assert.Equal(t, expectedPath, path)
			return newFakeStreamer([]byte("TEST")), nil
		}))

		// WHEN
		result, err := StreamContentOf(fs, expectedPath)

		// THEN
		require.NoError(t, err)

		val, _ := io.ReadAll(result)
		assert.Equal(t, "TEST", string(val))
	})
	t.Run("it should return up-stream error", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionStreamContentOfHandler(func(path string) (Streamer, error) {
			assert.Equal(t, expectedPath, path)
			return nil, errors.New("something went wrong")
		}))

		// WHEN
		result, err := StreamContentOf(fs, expectedPath)

		// THEN
		require.EqualError(t, err, "failed to attach reader to path/to/file: something went wrong")
		assert.Nil(t, result)
	})
}
