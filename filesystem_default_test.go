package filesystem

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultCreateFile(t *testing.T) {
	t.Run("it should create new file at provided location", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			// WHEN
			err = CreateFile(fs, fp)

			// THEN
			require.NoError(t, err)

			f, err := os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			require.NoError(t, err)
			defer f.Close()

			content, err := io.ReadAll(f)
			require.NoError(t, err)
			assert.Empty(t, content)
		})
	})
	t.Run("it should forbid creating directory structure if not permitted by arguments", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, path.Join("path", "to", "test-file.txt"))

			// WHEN
			err = CreateFile(fs, fp)

			// THEN
			require.ErrorIs(t, err, ErrUnresolvableDirectoryStructure)

			_, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			assert.True(t, os.IsNotExist(err))
		})
	})
	t.Run("it should create directory structure if allowed by arguments", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, path.Join("path", "to", "test-file.txt"))

			// WHEN
			err = CreateFile(fs, fp, WithAllowCreationOfDirectoryStructure(true))

			// THEN
			require.NoError(t, err)

			f, err := os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			require.NoError(t, err)
			defer f.Close()

			content, err := io.ReadAll(f)
			require.NoError(t, err)
			assert.Empty(t, content)
		})
	})
	t.Run("it should forbid overwriting a file if not permitted by arguments", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			f, err := os.Create(fp)
			require.NoError(t, err)
			_, _ = f.WriteString("TEST")
			_ = f.Close()

			// WHEN
			err = CreateFile(fs, fp)

			// THEN
			require.ErrorIs(t, err, ErrFileFound)

			f, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			require.NoError(t, err)
			defer f.Close()

			content, err := io.ReadAll(f)
			require.NoError(t, err)
			assert.Equal(t, "TEST", string(content))
		})
	})
	t.Run("it should overwrite a file if allowed by arguments", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			f, err := os.Create(fp)
			require.NoError(t, err)
			_, _ = f.WriteString("TEST")
			_ = f.Close()

			// WHEN
			err = CreateFile(fs, fp, WithAllowOverwrite(true))

			// THEN
			require.NoError(t, err)

			f, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			require.NoError(t, err)
			defer f.Close()

			content, err := io.ReadAll(f)
			require.NoError(t, err)
			assert.Empty(t, content)
		})
	})
}

func TestDefaultWriteContentTo(t *testing.T) {
	t.Run("it should append content to the file", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			f, err := os.Create(fp)
			require.NoError(t, err)
			_, _ = f.WriteString("TEST")
			_ = f.Close()

			// WHEN
			err = WriteContentTo[string](fs, fp, "-MORE")

			// THEN
			require.NoError(t, err)

			f, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			require.NoError(t, err)
			defer f.Close()

			content, err := io.ReadAll(f)
			require.NoError(t, err)
			assert.Equal(t, "TEST-MORE", string(content))
		})
	})
	t.Run("it should replace content to the file", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			f, err := os.Create(fp)
			require.NoError(t, err)
			_, _ = f.WriteString("TEST")
			_ = f.Close()

			// WHEN
			err = WriteContentTo[string](fs, fp, "OVERWRITE", WithContentOperation(ContentOperationOverwrite))

			// THEN
			require.NoError(t, err)

			f, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			require.NoError(t, err)
			defer f.Close()

			content, err := io.ReadAll(f)
			require.NoError(t, err)
			assert.Equal(t, "OVERWRITE", string(content))
		})
	})
	t.Run("it should report if file does not exist", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			// WHEN
			err = WriteContentTo[string](fs, fp, "Test")

			// THEN
			require.ErrorIs(t, err, ErrFileNotFound)

			_, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			assert.True(t, os.IsNotExist(err))
		})
	})
}

func TestDefaultStreamContentTo(t *testing.T) {
	t.Run("it should append content to the file", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			f, err := os.Create(fp)
			require.NoError(t, err)
			_, _ = f.WriteString("TEST")
			_ = f.Close()

			// WHEN
			err = StreamContentTo(fs, fp, bytes.NewBufferString("-MORE"))

			// THEN
			require.NoError(t, err)

			f, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			require.NoError(t, err)
			defer f.Close()

			content, err := io.ReadAll(f)
			require.NoError(t, err)
			assert.Equal(t, "TEST-MORE", string(content))
		})
	})
	t.Run("it should replace content to the file", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			f, err := os.Create(fp)
			require.NoError(t, err)
			_, _ = f.WriteString("TEST")
			_ = f.Close()

			// WHEN
			err = StreamContentTo(fs, fp, bytes.NewBufferString("OVERWRITE"), WithContentOperation(ContentOperationOverwrite))

			// THEN
			require.NoError(t, err)

			f, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			require.NoError(t, err)
			defer f.Close()

			content, err := io.ReadAll(f)
			require.NoError(t, err)
			assert.Equal(t, "OVERWRITE", string(content))
		})
	})
	t.Run("it should report if file does not exist", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-file.txt")

			// WHEN
			err = StreamContentTo(fs, fp, bytes.NewBufferString("Test"))

			// THEN
			require.ErrorIs(t, err, ErrFileNotFound)

			_, err = os.OpenFile(fp, os.O_RDONLY, os.FileMode(0644))
			assert.True(t, os.IsNotExist(err))
		})
	})
}

func TestDefaultCreateDirectory(t *testing.T) {
	t.Run("it should create new directory at provided location", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, "test-directory")

			// WHEN
			err = CreateDirectory(fs, fp)

			// THEN
			require.NoError(t, err)

			di, err := os.Stat(fp)
			require.NoError(t, err)
			assert.True(t, di.IsDir())
		})
	})
	t.Run("it should forbid creating directory structure if not permitted by arguments", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, path.Join("path", "to", "test-directory"))

			// WHEN
			err = CreateDirectory(fs, fp)

			// THEN
			require.ErrorIs(t, err, ErrUnresolvableDirectoryStructure)

			_, err = os.Stat(fp)
			assert.True(t, os.IsNotExist(err))
		})
	})
	t.Run("it should create directory structure if allowed by arguments", func(t *testing.T) {
		t.Parallel()

		withinRandomDirectoryScope(func(workdir string) {
			// GIVEN
			fs, err := New()
			require.NoError(t, err)

			fp := path.Join(workdir, path.Join("path", "to", "test-directory"))

			// WHEN
			err = CreateDirectory(fs, fp, WithAllowCreationOfDirectoryStructure(true))

			// THEN
			require.NoError(t, err)

			di, err := os.Stat(fp)
			require.NoError(t, err)
			assert.True(t, di.IsDir())
		})
	})
}

func withinRandomDirectoryScope(fn func(workdir string)) {
	tmpDir, err := os.MkdirTemp("", fmt.Sprintf("create-file-%d", rand.Int()))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()

	fn(tmpDir)
}
