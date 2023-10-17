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
		Add(ok.EndsWith("jst.dev")) // case?

	if err := companyEmail.Ok("me@jst.dev"); err != nil {
		t.Error(err)
	}
	if err := companyEmail.Ok("me@jst.dev.com"); err == nil {
		t.Error("should fail. not a jst.dev email")
	}

	if err := companyEmail.Ok("me@google.dev"); err == nil {
		t.Error("extenal email should fail")
	}

	emailWorkButNotSamOrMom := companyEmail.
		New("is not Sam").
		Add(ok.Not("sam@jst.dev")).
		New("is not mom").
		Add(ok.Not("mother@jst.dev"))

	if err := emailWorkButNotSamOrMom.Ok("sam@jst.dev"); err == nil {
		t.Error("should fail. sam is not allowed")
	}
	
	if err := emailWorkButNotSamOrMom.OkAll("sam@jst.dev"); err == nil {
		t.Error("should fail. sam is not allowed")
	}

	if err := emailWorkButNotSamOrMom.Ok("mother@jst.dev"); err == nil {
		t.Error()
	}

	companyEmailShort := companyEmail.
		New("is short").
		Add(ok.Max(12))
	if err := companyEmailShort.Ok("jst@jst.dev"); err != nil {
		t.Error(err)
	}
	if err := companyEmailShort.Ok("jststststststststststststststststststststststststststst@jst.dev"); err == nil {
		t.Error("should fail. not a short email")
	}

	companyEmailLong := companyEmail.
		New("is long").
		Add(ok.Min(20))
	if err := companyEmailLong.Ok("jst@jst.dev"); err == nil {
		t.Error("should fail. not a long email")
	}
	if err := companyEmailLong.Ok("jststststststststststststststststststststststststststst@jst.dev"); err != nil {
		t.Error(err)
	}
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

func TestMinMaxBytes(t *testing.T) {
	tests := []struct {
		name    string
		min     int
		max     int
		lenData int
		want    bool
	}{
		{"all zero", 0, 0, 0, true},

		{"min-1", 128, 256, 127, false},
		{"min", 128, 256, 128, true},
		{"min+1", 128, 256, 129, true},

		{"max-1", 0, 128, 127, true},
		{"max", 0, 128, 128, true},
		{"max+1", 0, 128, 129, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := varOfLen(tt.lenData)
			if got := ok.MinMax(tt.min, tt.max)(data); got != tt.want {
				t.Errorf("MinMaxBytes(%d, %d)(<%d bytes>) = %v, want %v", tt.min, tt.max, tt.lenData, got, tt.want)
			}
		})
	}
}

func varOfLen(n int) []byte {
	v := make([]byte, n)
	for i := 0; i < n; i++ {
		v[i] = byte(i)
	}
	return v

}
