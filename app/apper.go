package app

// Apper is the interface for an app.
type Apper interface {
	Cache()       string
	Config()      string
	Data()        string
	Desktop()     string
	Documents()   string
	Downloads()   string
	Home()        string
	Name()        string
	Path()        string
	Pictures()    string
	Screenshots() string
	Vendor()      string
}
