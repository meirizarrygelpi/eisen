// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisen

import (
	"fmt"
	"math/big"
	"strings"
)

// A Stein represents an arbitrary-precision Eisenstein integer.
type Stein struct {
	l, r big.Int
}

// L returns the left component of z.
func (z *Stein) L() *big.Int {
	return &z.l
}

// R returns the right component of z.
func (z *Stein) R() *big.Int {
	return &z.r
}

// SetL sets the left component of z equal to a.
func (z *Stein) SetL(a *big.Int) {
	z.l = *a
}

// SetR sets the right component of z equal to b.
func (z *Stein) SetR(b *big.Int) {
	z.r = *b
}

// Components returns the two components of z.
func (z *Stein) Components() (a, b *big.Int) {
	a = z.L()
	b = z.R()
	return
}

// String returns the string version of a Stein value.
//
// If z corresponds to a + bω, then the string is "(a+bω)", similar to
// complex128 values.
func (z *Stein) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", z.L())
	if z.R().Sign() == -1 {
		a[2] = fmt.Sprintf("%v", z.R())
	} else {
		a[2] = fmt.Sprintf("+%v", z.R())
	}
	a[3] = "ω"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Stein) Equals(y *Stein) bool {
	if z.L().Cmp(y.L()) != 0 || z.R().Cmp(y.R()) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Stein) Copy(y *Stein) *Stein {
	z.SetL(y.L())
	z.SetR(y.R())
	return z
}

// New returns a pointer to a Stein value made from two given pointers to
// big.Int values.
func New(a, b *big.Int) *Stein {
	z := new(Stein)
	z.SetL(a)
	z.SetR(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Stein) Scal(y *Stein, a *big.Int) *Stein {
	z.SetL(new(big.Int).Mul(y.L(), a))
	z.SetR(new(big.Int).Mul(y.R(), a))
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Stein) Neg(y *Stein) *Stein {
	z.SetL(new(big.Int).Neg(y.L()))
	z.SetR(new(big.Int).Neg(y.R()))
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Stein) Conj(y *Stein) *Stein {
	z.SetL(new(big.Int).Sub(y.L(), y.R()))
	z.SetR(new(big.Int).Neg(y.R()))
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Stein) Add(x, y *Stein) *Stein {
	z.SetL(new(big.Int).Add(x.L(), y.L()))
	z.SetR(new(big.Int).Add(x.R(), y.R()))
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Stein) Sub(x, y *Stein) *Stein {
	z.SetL(new(big.Int).Sub(x.L(), y.L()))
	z.SetR(new(big.Int).Sub(x.R(), y.R()))
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(ω, ω) + ω + 1 = 0
// This binary operation is commutative and associative.
func (z *Stein) Mul(x, y *Stein) *Stein {
	a, b := x.L(), x.R()
	c, d := y.L(), y.R()
	s, t, u, v := new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	z.SetL(s.Sub(
		s.Mul(a, c),
		u.Mul(d, b),
	))
	z.SetR(t.Add(
		t.Mul(d, a),
		u.Mul(b, c),
	))
	z.SetR(v.Sub(
		z.R(),
		v.Mul(b, d),
	))
	return z
}

// Quad returns the quadrance of z, a pointer to a big.Int value.
func (z *Stein) Quad() *big.Int {
	quad := new(big.Int)
	quad.Add(
		quad.Mul(z.L(), z.L()),
		new(big.Int).Mul(z.R(), z.R()),
	)
	quad.Sub(
		quad,
		new(big.Int).Mul(z.L(), z.R()),
	)
	return quad
}

// Quo sets z equal to the quotient of x and y, and returns z. Note that
// truncated division is used.
func (z *Stein) Quo(x, y *Stein) *Stein {
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	z.SetL(new(big.Int).Quo(z.L(), quad))
	z.SetR(new(big.Int).Quo(z.R(), quad))
	return z
}

// Associates returns the six associates of z.
func (z *Stein) Associates() (a, b, c, d, e, f *Stein) {
	ω := New(
		big.NewInt(0),
		big.NewInt(1),
	)
	a = new(Stein).Copy(z)
	b = new(Stein).Neg(z)
	c = new(Stein).Mul(z, ω)
	ω.Neg(ω)
	d = new(Stein).Mul(z, ω)
	ω.Mul(ω, ω)
	e = new(Stein).Mul(z, ω)
	ω.Neg(ω)
	f = new(Stein).Mul(z, ω)
	return
}
