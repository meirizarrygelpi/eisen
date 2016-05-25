// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisen

import (
	"fmt"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
)

// A Stein represents an arbitrary-precision Eisenstein integer.
type Stein struct {
	l, r big.Int
}

// Omega returns a pointer to the Eisenstein unit ω.
func Omega() *Stein {
	return &Stein{
		*big.NewInt(0),
		*big.NewInt(1),
	}
}

// Integers returns the pointers to the two integer components of z.
func (z *Stein) Integers() (*big.Int, *big.Int) {
	return &z.l, &z.r
}

// String returns the string version of a Stein value.
//
// If z corresponds to a + bω, then the string is "(a+bω)", similar to
// complex128 values.
func (z *Stein) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%v", &z.l)
	if (&z.r).Sign() == -1 {
		a[2] = fmt.Sprintf("%v", &z.r)
	} else {
		a[2] = fmt.Sprintf("+%v", &z.r)
	}
	a[3] = "ω"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Stein) Equals(y *Stein) bool {
	if (&z.l).Cmp(&y.l) != 0 || (&z.r).Cmp(&y.r) != 0 {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Stein) Copy(y *Stein) *Stein {
	(&z.l).Set(&y.l)
	(&z.r).Set(&y.r)
	return z
}

// New returns a pointer to a Stein value made from two given pointers to
// big.Int values.
func New(a, b *big.Int) *Stein {
	z := new(Stein)
	(&z.l).Set(a)
	(&z.r).Set(b)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Stein) Scal(y *Stein, a *big.Int) *Stein {
	(&z.l).Mul(&y.l, a)
	(&z.r).Mul(&y.r, a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Stein) Neg(y *Stein) *Stein {
	(&z.l).Neg(&y.l)
	(&z.r).Neg(&y.r)
	return z
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Stein) Conj(y *Stein) *Stein {
	(&z.l).Sub(&y.l, &y.r)
	(&z.r).Neg(&y.r)
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Stein) Add(x, y *Stein) *Stein {
	(&z.l).Add(&x.l, &y.l)
	(&z.r).Add(&x.r, &y.r)
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Stein) Sub(x, y *Stein) *Stein {
	(&z.l).Sub(&x.l, &y.l)
	(&z.r).Sub(&x.r, &y.r)
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule is:
// 		Mul(ω, ω) + ω + 1 = 0
// This binary operation is commutative and associative.
func (z *Stein) Mul(x, y *Stein) *Stein {
	a := new(big.Int).Set(&x.l)
	b := new(big.Int).Set(&x.r)
	c := new(big.Int).Set(&y.l)
	d := new(big.Int).Set(&y.r)
	temp := new(big.Int)
	(&z.l).Sub(
		(&z.l).Mul(a, c),
		temp.Mul(d, b),
	)
	(&z.r).Add(
		(&z.r).Mul(d, a),
		temp.Mul(b, c),
	)
	(&z.r).Sub(
		(&z.r),
		temp.Mul(b, d),
	)
	return z
}

// Quad returns the non-negative quadrance of z, a pointer to a big.Int value.
func (z *Stein) Quad() *big.Int {
	quad := new(big.Int)
	temp := new(big.Int)
	quad.Add(
		quad.Mul(&z.l, &z.l),
		temp.Mul(&z.r, &z.r),
	)
	quad.Sub(
		quad,
		temp.Mul(&z.l, &z.r),
	)
	return quad
}

// Quo sets z equal to the quotient of x and y, and returns z. Note that
// truncated division is used.
func (z *Stein) Quo(x, y *Stein) *Stein {
	quad := y.Quad()
	z.Conj(y)
	z.Mul(x, z)
	(&z.l).Quo(&z.l, quad)
	(&z.r).Quo(&z.r, quad)
	return z
}

// Associates returns the six associates of z.
func (z *Stein) Associates() (a, b, c, d, e, f *Stein) {
	a.Copy(z)
	b.Neg(z)
	unit := Omega()
	c.Mul(z, unit)
	unit.Neg(unit)
	d.Mul(z, unit)
	unit.Mul(unit, unit)
	e.Mul(z, unit)
	unit.Neg(unit)
	f.Mul(z, unit)
	return
}

// IsEisensteinPrime returns true if z is an Eisenstein prime.
func (z *Stein) IsEisensteinPrime() bool {
	return false
}

// Generate a random Stein value for quick.Check testing.
func (z *Stein) Generate(rand *rand.Rand, size int) reflect.Value {
	randomStein := &Stein{
		*big.NewInt(rand.Int63()),
		*big.NewInt(rand.Int63()),
	}
	return reflect.ValueOf(randomStein)
}
