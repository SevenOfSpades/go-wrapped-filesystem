package filesystem

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeReadCloser struct {
	buff *bytes.Buffer
}

func newFakeReadCloser[T string | []byte](data T) *fakeReadCloser {
	return &fakeReadCloser{buff: bytes.NewBuffer([]byte(data))}
}

func (s *fakeReadCloser) Read(p []byte) (n int, err error) {
	return s.buff.Read(p)
}

func (s *fakeReadCloser) Close() error {
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
		require.NoError(t, err)

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
		require.NoError(t, err)

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
		fs, err := New(OptionStreamContentOfHandler(func(path string) (io.ReadCloser, error) {
			assert.Equal(t, expectedPath, path)
			return newFakeReadCloser([]byte("TEST")), nil
		}))
		require.NoError(t, err)

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
		fs, err := New(OptionStreamContentOfHandler(func(path string) (io.ReadCloser, error) {
			assert.Equal(t, expectedPath, path)
			return nil, errors.New("something went wrong")
		}))
		require.NoError(t, err)

		// WHEN
		result, err := StreamContentOf(fs, expectedPath)

		// THEN
		require.EqualError(t, err, "failed to attach reader to path/to/file: something went wrong")
		assert.Nil(t, result)
	})
}

func TestCheckIfExists(t *testing.T) {
	t.Run("it should pass execution to provided handler", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionCheckIfExistsHandler(func(path string) (bool, error) {
			assert.Equal(t, expectedPath, path)
			return true, nil
		}))
		require.NoError(t, err)

		// WHEN
		result, err := CheckIfExists(fs, expectedPath)

		// THEN
		require.NoError(t, err)
		assert.True(t, result)
	})
}

func TestCreateFile(t *testing.T) {
	t.Run("it should pass execution to provided handler with default arguments values", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionCreateFileHandler(func(path string, arg Arguments) error {
			assert.Equal(t, expectedPath, path)
			assert.Equal(t, ModeAllReadWrite, arg.Mode)
			assert.False(t, arg.AllowCreationOfDirectoryStructure)
			assert.False(t, arg.AllowOverwrite)
			return nil
		}))

		// WHEN
		err = CreateFile(fs, expectedPath)

		// THEN
		require.NoError(t, err)
	})
	t.Run("it should pass execution to provided handler with defined arguments values", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionCreateFileHandler(func(path string, arg Arguments) error {
			assert.Equal(t, expectedPath, path)
			assert.Equal(t, ModeUserRead, arg.Mode)
			assert.True(t, arg.AllowCreationOfDirectoryStructure)
			assert.True(t, arg.AllowOverwrite)
			return nil
		}))

		// WHEN
		err = CreateFile(
			fs,
			expectedPath,
			WithMode(ModeUserRead),
			WithAllowOverwrite(true),
			WithAllowCreationOfDirectoryStructure(true),
		)

		// THEN
		require.NoError(t, err)
	})
}

func TestWriteContentTo(t *testing.T) {
	t.Run("it should pass execution to provided handler with string value", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionWriteContentToHandler(func(path string, content []byte, arg Arguments) error {
			assert.Equal(t, "Test", string(content))
			assert.Equal(t, expectedPath, path)
			assert.Equal(t, ContentOperationAppend, arg.ContentOperation)
			return nil
		}))
		require.NoError(t, err)

		// WHEN
		err = WriteContentTo[string](fs, expectedPath, "Test")

		// THEN
		require.NoError(t, err)
	})
	t.Run("it should pass execution to provided handler with []byte value", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionWriteContentToHandler(func(path string, content []byte, arg Arguments) error {
			assert.Equal(t, "Test", string(content))
			assert.Equal(t, expectedPath, path)
			assert.Equal(t, ContentOperationAppend, arg.ContentOperation)
			return nil
		}))
		require.NoError(t, err)

		// WHEN
		err = WriteContentTo[[]byte](fs, expectedPath, []byte("Test"))

		// THEN
		require.NoError(t, err)
	})
}

func TestStreamContentTo(t *testing.T) {
	t.Run("it should pass execution to provided handler with string value", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionStreamContentToHandler(func(path string, reader io.Reader, arg Arguments) error {
			content, _ := io.ReadAll(reader)
			assert.Equal(t, "Test", string(content))

			assert.Equal(t, expectedPath, path)
			assert.Equal(t, ContentOperationAppend, arg.ContentOperation)
			return nil
		}))
		require.NoError(t, err)

		// WHEN
		err = StreamContentTo(fs, expectedPath, bytes.NewBuffer([]byte("Test")))

		// THEN
		require.NoError(t, err)
	})
	t.Run("it should pass execution to provided handler with []byte value", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionStreamContentToHandler(func(path string, reader io.Reader, arg Arguments) error {
			content, _ := io.ReadAll(reader)
			assert.Equal(t, "Test", string(content))

			assert.Equal(t, expectedPath, path)
			assert.Equal(t, ContentOperationAppend, arg.ContentOperation)
			return nil
		}))
		require.NoError(t, err)

		// WHEN
		err = StreamContentTo(fs, expectedPath, bytes.NewBuffer([]byte("Test")))

		// THEN
		require.NoError(t, err)
	})
}

func TestCreateDirectory(t *testing.T) {
	t.Run("it should pass execution to provided handler", func(t *testing.T) {
		t.Parallel()

		// GIVEN
		expectedPath := "path/to/file"
		fs, err := New(OptionCreateDirectory(func(path string, _ Arguments) error {
			assert.Equal(t, expectedPath, path)
			return nil
		}))
		require.NoError(t, err)

		// WHEN
		err = CreateDirectory(fs, expectedPath)

		// THEN
		require.NoError(t, err)
	})
}
