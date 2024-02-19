package main

import "regexp"

// SpecialCharRegex is a regular expression matching all characters
// that need to be escaped in LaTeX source code
var SpecialCharRegex = regexp.MustCompile(`([\\\^\%~\#\$%&_\{\}])`)

// ReplacementMap maps each LaTeX special character to its escaped version
var ReplacementMap = map[string]string{
	"#":  `\#`,               // pound/hashtag
	"$":  `\$`,               // dollar sign
	"%":  `\%`,               // percent
	"&":  `\&`,               // ampersand
	"~":  `\~{}`,             // tilde
	"_":  `\_`,               // underscore
	"^":  `\^{}`,             // caret
	"\\": `\textbackslash{}`, // backslash
	"{":  `\{`,               // left curly brace
	"}":  `\}`,               // right curly brace
}

func EscapeLatex(src string) string {
	return SpecialCharRegex.ReplaceAllStringFunc(src, func(value string) string {
		return ReplacementMap[value]
	})
}
