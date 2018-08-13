package config

// Flagger is the interface for a flag replacement library.
type Flagger interface {
	Bool(string, bool, string)
	BoolP(string, string, bool, string)
	BoolVar(*bool, string, bool, string)
	BoolVarP(*bool, string, string, bool, string)
	Count(string, int, string)
	CountP(string, string, int, string)
	CountVar(*int, string, int, string)
	CountVarP(*int, string, string, int, string)
	Int(string, int, string)
	IntP(string, string, int, string)
	IntVar(*int, string, int, string)
	IntVarP(*int, string, string, int, string)
	Parse([]string) error
	String(string, string, string)
	StringP(string, string, string, string)
	StringVar(*string, string, string, string)
	StringVarP(*string, string, string, string, string)
}
