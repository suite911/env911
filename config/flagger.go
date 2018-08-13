package config

// Flagger is the interface for a flag replacement library.
type Flagger interface {
	Bool(string, bool, string)
	BoolP(string, string, bool, string)
	BoolVar(*bool, string, bool, string)
	BoolVarP(*bool, string, string, bool, string)
	Int(string, int, string)
	IntP(string, string, int, string)
	IntVar(int*, string, int, string)
	IntVarP(int*, string, string, int, string)
	String(string, int, string)
	StringP(string, string, int, string)
	StringVar(string, int, string)
	StringVarP(string, string, int, string)
}
