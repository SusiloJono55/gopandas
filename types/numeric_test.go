package types

import (
	"testing"
)

func TestNumericType(t *testing.T) {
	var s Numeric

	switch s.Type() {
	case NUMERIC:
		return
	default:
		t.Error("Wrong type")
	}
}

func TestNumericAdd(t *testing.T) {
	n1 := Numeric(1)
	n2 := Numeric(2)
	n3 := Numeric(3)
	if n1.Add(n2.Add(n3)) != n2.Add(n1.Add(n3)) {
		t.Errorf("Error association")
	}
	if n1.Add(Numeric(0)) != n1 {
		t.Errorf("Error neutral element")
	}
}

func TestNumericMul(t *testing.T) {
	n1 := Numeric(1)
	n2 := Numeric(2)
	n3 := Numeric(3)
	if n1.Mul(n2.Mul(n3)) != n2.Mul(n1.Mul(n3)) {
		t.Errorf("Error association")
	}
	if n1.Mul(Numeric(1)) != n1 {
		t.Errorf("Error neutral element")
	}
}

func TestNumericDiv(t *testing.T) {
	n1 := Numeric(1)
	n2 := Numeric(2)
	if n1.Div(n2) != Numeric(1).Div(n2.Div(n1)) {
		t.Errorf("Error test x/y != 1/(y/x)")
	}
	if n1.Div(Numeric(1)) != n1 {
		t.Errorf("Error neutral element")
	}
	if n1.Div(Numeric(0)) != NewNan() {
		t.Errorf("Error div by zero")
	}
}

func TestNumericGreat(t *testing.T) {
	if Numeric(3).Great(Numeric(2)) == false {
		t.Errorf("Error Great")
	}
	if Numeric(0).Great(Numeric(0)) {
		t.Errorf("Error Great")
	}
}

func TestNumericLess(t *testing.T) {
	if Numeric(3).Less(Numeric(2)) {
		t.Errorf("Error Less")
	}
	if Numeric(0).Less(Numeric(0)) {
		t.Errorf("Error Less")
	}
}

func TestNumericEqual(t *testing.T) {
	if Numeric(0).Equal(Numeric(0)) == false {
		t.Errorf("Error Equal")
	}
	if Numeric(-1).Equal(Numeric(-1)) == false {
		t.Errorf("Error Equal")
	}
	if Numeric(-2).Equal(Numeric(-1)) {
		t.Errorf("Error Equal")
	}
}

func TestNumericNotEqual(t *testing.T) {
	if Numeric(0).NotEqual(Numeric(0)) == true {
		t.Errorf("Error Equal")
	}
	if Numeric(-1).NotEqual(Numeric(-1)) == true {
		t.Errorf("Error Equal")
	}
	if Numeric(-2).NotEqual(Numeric(-1)) == false {
		t.Errorf("Error Equal")
	}
}

func TestNumeric_String(t *testing.T) {
	tests := []struct {
		name string
		n    Numeric
		want string
	}{
		{name: "test float to string", n: 22222.2222, want: "22222.2222"},
		{name: "test int to string", n: 23, want: "23"},
		{name: "test zero to string", n: 0, want: "0"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("Numeric.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
