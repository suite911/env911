package config

import flag "github.com/ogier/pflag"

type PFlag struct {
	*flag.FlagSet
}

func (pflag PFlag) Count(name string, by int, usage string) *int {
	panic("github.com/ogier/pflag does not support the Count* functions")
}

func (pflag PFlag) CountP(name, shorthand string, by int, usage string) *int {
	panic("github.com/ogier/pflag does not support the Count* functions")
}

func (pflag PFlag) CountVar(p *int, name string, by int, usage string) {
	panic("github.com/ogier/pflag does not support the Count* functions")
}

func (pflag PFlag) CountVarP(p *int, name, shorthand string, by int, usage string) {
	panic("github.com/ogier/pflag does not support the Count* functions")
}

// TODO: make my own CC0-1.0-licensed POSIX-compliant flag replacement library that does

func (pflag PFlag) CountWorks() bool {
	return false
}

func (pflag PFlag) PVersionsWork() bool {
	return true
}

func (pflag PFlag) SetUsageFooter(text string) {
	flag.Usage = usage
	pflag.UsageFooter = text
}

func (pflag PFlag) SetUsageHeader(text string) {
	flag.Usage = usage
	pflag.UsageHeader = text
}

var UsageFooter, UsageHeader string

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:", os.Args[0])
	if len(UsageHeader) > 0 {
		fmt.Fprintf(os.Stderr, UsageHeader)
	}
	flag.PrintDefaults()
	if len(UsageFooter) > 0 {
		fmt.Fprintf(os.Stderr, UsageFooter)
	}
}
