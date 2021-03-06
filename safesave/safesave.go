package safesave

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

type SafeSave struct {
	real string
	temp *os.File
	ok   bool
}

func New(real, temp string) (*SafeSave, error) {
	return new(SafeSave).Init(real, temp)
}

func (save *SafeSave) Init(real, temp string) (*SafeSave, error) {
	save.real = real

	if len(temp) < 1 {
		temp = filepath.Join(os.TempDir(), "." + strconv.Itoa(os.Getpid()) + ".new")
		// TODO: mimic mkstemp(3)
	}
	os.Remove(temp)
	var err error
	save.temp, err = os.Create(temp)
	if err != nil {
		return nil, err
	}

	info, err := os.Stat(real)
	if err != nil {
		return nil, err
	}

	if err := os.Chmod(save.temp.Name(), info.Mode()); err != nil {
		return nil, err
	}

	save.ok = true
	return save, nil
}

func (save SafeSave) Close() error {
	if !save.ok {
		save.temp.Close()
		os.Remove(save.temp.Name()) // Do not defer this!*
		return errors.New("Open or write failed")
	}

	// The write was successful, but the temporary file is still open

	if err := save.temp.Close(); err != nil {
		os.Remove(save.temp.Name()) // Do not defer this!*
		return err
	}

	// The temporary file is successfully closed

	os.Remove(save.real)
	if err := os.Rename(save.temp.Name(), save.real); err != nil {
		// *Do not remove the temporary file here, it is all the user has left!
		return err
	}

	// The temporary file has been renamed over the real file

	os.Remove(save.temp.Name())
	return nil
}

func (save *SafeSave) Write(p []byte) (n int, err error) {
	if n, err = save.temp.Write(p); err != nil {
		save.ok = false
	}
	return
}
