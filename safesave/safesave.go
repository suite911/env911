package safesave

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

	save.ok = true
	return save, nil
}

func (save SafeSave) Close() error {
	if !save.ok {
		return errors.New("Open or write failed")
	}

	// The write was successful, but the temporary file is still open

	if err := save.temp.Close(); err != nil {
		return err
	}

	// The temporary file is successfully closed

	os.Remove(save.real)
	if err := os.Rename(save.temp.Name(), save.real); err != nil {
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
