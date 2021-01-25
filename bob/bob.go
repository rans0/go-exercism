// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	remark string
}

func (r Remark) isSilence() bool {
	return r.remark == ""
}

func (r Remark) isYelling() bool {
	isEmot := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	isUppercase := strings.ToUpper(r.remark) == r.remark
	return isUppercase && isEmot
}

func (r Remark) isQuestion() bool {
	hasQuestion := strings.HasSuffix(r.remark, "?")
	return hasQuestion
}

func (r Remark) isYellandAsk() bool {
	return r.isYelling() && r.isQuestion()
}

func newRemark(remark string) Remark {
	return Remark{strings.TrimSpace(remark)}
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := newRemark(remark)
	switch {
	case r.isSilence():
		return "Fine. Be that way!"
	case r.isYellandAsk():
		return "Calm down, I know what I'm doing!"
	case r.isYelling():
		return "Whoa, chill out!"
	case r.isQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
