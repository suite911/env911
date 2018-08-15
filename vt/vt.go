package vt

const (
	Reset            = "\x1b[0m\x1b[m"
	NewLine          = Reset + "\n"
	AttrBoldOff      = "\x1b[22m"
	AttrBoldOn       = "\x1b[1m"
	AttrDimOff       = "\x1b[22m"
	AttrDimOn        = "\x1b[2m"
	AttrItalicOff    = "\x1b[23m"
	AttrItalicOn     = "\x1b[3m"
	AttrUnderlineOff = "\x1b[24m"
	AttrUnderlineOn  = "\x1b[4m"
	AttrBlinkOff     = "\x1b[25m"
	AttrBlinkOn      = "\x1b[5m"
	AttrReverseOff   = "\x1b[27m"
	AttrReverseOn    = "\x1b[7m"
	AttrInvisibleOff = "\x1b[28m"
	AttrInvisibleOn  = "\x1b[8m"
	AttrStrikeOff    = "\x1b[29m"
	AttrStrikeOn     = "\x1b[9m"
)

// Wrap in attribute 1 (bold).
func B(text string) string {
	return AttrBoldOn + text + AttrBoldOff
}

// Wrap in attribute 2 (dim).
func Dim(text string) string {
	return AttrDimOn + text + AttrDimOff
}

// Wrap in attribute 3 (italic).
func I(text string) string {
	return AttrItalicOn + text + AttrItalicOff
}

// Wrap in attribute 4 (underline).
func U(text string) string {
	return AttrUnderlineOn + text + AttrUnderlineOff
}

// Wrap in attribute 5 (blink).
func Blink(text string) string {
	return AttrBlinkOn + text + AttrBlinkOff
}

// Wrap in attribute 7 (reverse).
func Reverse(text string) string {
	return AttrReverseOn + text + AttrReverseOff
}

// Wrap in attribute 8 (invisible).
func Invisible(text string) string {
	return AttrInvisibleOn + text + AttrInvisibleOff
}

// Wrap in attribute 9 (strike).
func S(text string) string {
	return AttrStrikeOn + text + AttrStrikeOff
}
