package ok_test

import (
	"fmt"
	"testing"

	"github.com/johan-st/go-component-system/ok"
)

func TestAPI(t *testing.T) {

	// with composable rules. Make UI with checkboxes for rules and an input field.
	companyEmail := ok.Validation{}.
		New("is internal email").
		Add(ok.String).
		Add(ok.Email()).
		Add(ok.Contains("jst.dev")) // case?

	if err := companyEmail.Ok("me@jst.dev"); err != nil {
		t.Log(err)
		t.Fail()
	}

	if err := companyEmail.Ok("me@google.dev"); err == nil {
		t.Fail()
	}

	emailWorkButNotSamOrMom := companyEmail.
		New("is not Sam").
		Add(ok.Not("sam@jst.dev")).
		New("is not mom").
		Add(ok.Not("mother@jst.dev"))

	if err := emailWorkButNotSamOrMom.Ok("sam@jst.dev"); err != nil {
		fmt.Print(err)
	} else {
		t.Fail()
	}
	if err := emailWorkButNotSamOrMom.OkAll("sam@jst.dev"); err != nil {
		fmt.Print(err)
	} else {
		t.Fail()
	}

	if err := emailWorkButNotSamOrMom.Ok("mother@jst.dev"); err != nil {
		fmt.Print(err)
	} else {
		t.Fail()
	}
	// Validation error:
	//   GOT: "sam@jst.dev"
	//   FAILED: is not Sam
	//   RULES APPLIED:
	//    ✓ is internal email
	//    × is not Sam
	//      is not mom

}

// TestFuncs

func TestString(t *testing.T) {
	tests := []struct {
		name string
		arg  any
		want bool
	}{
		{"nil - NOT OK", nil, false},
		{"empty string - OK", "", true},
		{"string - OK", "bobby tables was here", true},
		{"[]byte(\"hello\") - OK", []byte("hello"), true},
		{"[]byte{0xff} - NOT OK", []byte{104, 101, 108, 108, 111, 200}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ok.String(tt.arg); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmail(t *testing.T) {
	tests := []struct {
		arg  any
		want bool
	}{
		{"test@test.com", true},
		{"test@test.co.uk", true},
		{"test@test.", false},
		{"test.com", false},
		{"@test.com", false},
		{"test_at_test.com", false},
		{"test @test.com", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s__%v", tt.arg, tt.want), func(t *testing.T) {
			if got := ok.Email()(tt.arg); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
