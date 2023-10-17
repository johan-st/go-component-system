package ok

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

type TestFunc func(any) bool

type Rule struct {
	tests  []TestFunc
	msg    string
	passed bool
	ran    bool
}

type Validation []*Rule

func (v Validation) New(msg string) *Validation {
	// fmt.Println("adding rule " + msg)
	v = append(v, &Rule{msg: msg})
	// fmt.Println("vlen:",len(v))
	return &v
}

func (v *Validation) Add(t TestFunc) *Validation {
	last := (*v)[len(*v)-1]
	// fmt.Println("add before:", len(last.tests))
	last.tests = append(last.tests, t)
	// TODO: Consider checling for conflicting duplicates when adding tests

	// fmt.Println("adding test to", last.msg)
	// fmt.Println("add after:", len(last.tests))
	return v
}

func (v *Validation) Ok(a any) error {
	return v.ok(a, false)
}

func (v *Validation) OkAll(a any) error {
	return v.ok(a, true)
}

func (v *Validation) ok(a any, full bool) error {
	var failed []int
	// fmt.Println("running...")

validation:
	for ruleIndex, rule := range *v {
		// fmt.Println("rule: " + rule.msg)
		// test rule
		rule.ran = true
		rule.passed = true // no tests on rule should fails validation?
		for _, t := range rule.tests {
			if !t(a) {
				// fmt.Println("test failed")
				rule.passed = false
				failed = append(failed, ruleIndex)
				if full {
					break
				}
				break validation
			}
			// fmt.Println("test ok")
		}
		// fmt.Println("end: " + rule.msg)
	}

	// if any failed
	if len(failed) > 0 {
		firstFailedMsg := (*v)[failed[0]].msg
		strB := strings.Builder{}
		strB.WriteString("Validation Error:\n")
		strB.WriteString(fmt.Sprintf("  GOT: %#v\n", a))
		strB.WriteString(fmt.Sprintf("  FAILED: %s\n", firstFailedMsg))
		for _, r := range *v {
			if r.ran && r.passed {
				strB.WriteString(fmt.Sprintf("   ✓ %s\n", r.msg))
			} else if r.ran && !r.passed {
				strB.WriteString(fmt.Sprintf("   × %s\n", r.msg))
			} else {
				strB.WriteString(fmt.Sprintf("     %s\n", r.msg))
			}
		}
		// fmt.Println(strB.String())
		return errors.New(strB.String())
	}
	return nil
}

// buildt-in testFuncs

// String checks for type string or valid utf8 []byte
func String(a any) bool {
	switch a := a.(type) {
	case string:
		return true
	case []byte:
		// fmt.Println(a)
		// fmt.Println(string(a))
		return utf8.Valid(a)
	default:
		return false
	}
}

func Regex(re *regexp.Regexp) TestFunc {
	return func(a any) bool {
		switch a := a.(type) {
		case []byte:
			return re.Match(a)
		case string:
			return re.MatchString(a)
		default:
			return false
		}
	}
}

// Email checks if we have a valid email address
func Email() TestFunc {
	return Regex(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`))
}

// Contains checks if contains he given string (case insensitive)
func Contains(str string) TestFunc {
	return func(a any) bool {
		switch a := a.(type) {
		case string:
			return strings.Contains(a, str)
		case []byte:
			return bytes.Contains(a, []byte(str))
		default:
			return false
		}
	}
}

func EndsWith(str string) TestFunc {
	return func(a any) bool {
		switch a := a.(type) {
		case string:
			return strings.HasSuffix(a, str)
		case []byte:
			return bytes.HasSuffix(a, []byte(str))
		default:
			return false
		}
	}
}

// Not makes sure the given string does not exist in value
func Not(str string) TestFunc {
	return func(a any) bool {
		return !Contains(str)(a)
	}
}

func Min(min int) TestFunc {
	return func(a any) bool {
		switch a := a.(type) {
		case string:
			return (min <= len(a))
		case []byte:
			return (min <= len(a))
		case int:
			return (min <= a)
		default:
			return false
		}
	}
}

func Max(max int) TestFunc {
	return func(a any) bool {
		switch a := a.(type) {
		case string:
			return (len(a) <= max)
		case []byte:
			return (len(a) <= max)
		case int:
			return (a <= max)
		default:
			return false
		}
	}
}

func MinMaxBytes(min, max int) TestFunc {
	return func(a any) bool {
		switch a := a.(type) {
		case []byte:
			return (min <= len(a) && len(a) <= max)
		default:
			return false

		}
	}
}
