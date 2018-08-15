package config

// Flagger is the interface for a flag replacement library.
type Flagger interface {
	Bool(string, bool, string) *bool
	BoolP(string, string, bool, string) *bool
	BoolVar(*bool, string, bool, string)
	BoolVarP(*bool, string, string, bool, string)
	Count(string, int, string) *int
	CountP(string, string, int, string) *int
	CountVar(*int, string, int, string)
	CountVarP(*int, string, string, int, string)
	CountWorks() bool
	Int(string, int, string) *int
	IntP(string, string, int, string) *int
	IntVar(*int, string, int, string)
	IntVarP(*int, string, string, int, string)
	Parse([]string) error
	PVersionsWork() bool
	SetUsageFooter(string)
	SetUsageHeader(string)
	String(string, string, string) *string
	StringP(string, string, string, string) *string
	StringVar(*string, string, string, string)
	StringVarP(*string, string, string, string, string)
}
