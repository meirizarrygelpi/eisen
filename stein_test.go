// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package eisen

import (
	"math/big"
	"testing"
	"testing/quick"
)

func TestAddCommutative(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Stein).Add(x, y)
		r := new(Stein).Add(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddAssociative(t *testing.T) {
	f := func(x, y, z *Stein) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Stein), new(Stein)
		l.Add(l.Add(x, y), z)
		r.Add(x, r.Add(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddZero(t *testing.T) {
	zero := new(Stein)
	f := func(x *Stein) bool {
		// t.Logf("x = %v", x)
		l := new(Stein).Add(x, zero)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulCommutative(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l := new(Stein).Mul(x, y)
		r := new(Stein).Mul(y, x)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulAssociative(t *testing.T) {
	f := func(x, y, z *Stein) bool {
		// t.Logf("x = %v, y = %v, z = %v", x, y, z)
		l, r := new(Stein), new(Stein)
		l.Mul(l.Mul(x, y), z)
		r.Mul(x, r.Mul(y, z))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulOne(t *testing.T) {
	one := &Stein{
		*big.NewInt(1),
		*big.NewInt(0),
	}
	f := func(x *Stein) bool {
		// t.Logf("x = %v", x)
		l := new(Stein).Mul(x, one)
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegInvolutive(t *testing.T) {
	f := func(x *Stein) bool {
		// t.Logf("x = %v", x)
		l := new(Stein)
		l.Neg(l.Neg(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestConjInvolutive(t *testing.T) {
	f := func(x *Stein) bool {
		// t.Logf("x = %v", x)
		l := new(Stein)
		l.Conj(l.Conj(x))
		return l.Equals(x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNegConjCommutative(t *testing.T) {
	f := func(x *Stein) bool {
		// t.Logf("x = %v", x)
		l, r := new(Stein), new(Stein)
		l.Neg(l.Conj(x))
		r.Conj(r.Neg(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMulConjAntiDistributive(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Stein), new(Stein)
		l.Conj(l.Mul(x, y))
		r.Mul(r.Conj(y), new(Stein).Conj(x))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScalDouble(t *testing.T) {
	f := func(x *Stein) bool {
		// t.Logf("x = %v", x)
		l, r := new(Stein), new(Stein)
		l.Add(x, x)
		r.Scal(x, big.NewInt(2))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubAntiCommutative(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Stein), new(Stein)
		l.Sub(x, y)
		r.Sub(y, x)
		r.Neg(r)
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddConjDistributive(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Stein), new(Stein)
		l.Add(x, y)
		l.Conj(l)
		r.Add(r.Conj(x), new(Stein).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubConjDistributive(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Stein), new(Stein)
		l.Sub(x, y)
		l.Conj(l)
		r.Sub(r.Conj(x), new(Stein).Conj(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddScalDistributive(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Stein), new(Stein)
		l.Scal(l.Add(x, y), a)
		r.Add(r.Scal(x, a), new(Stein).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSubScalDistributive(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		a := big.NewInt(2)
		l, r := new(Stein), new(Stein)
		l.Scal(l.Sub(x, y), a)
		r.Sub(r.Scal(x, a), new(Stein).Scal(y, a))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestAddNegSub(t *testing.T) {
	f := func(x, y *Stein) bool {
		// t.Logf("x = %v, y = %v", x, y)
		l, r := new(Stein), new(Stein)
		l.Sub(x, y)
		r.Add(x, r.Neg(y))
		return l.Equals(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestQuadPositive(t *testing.T) {
	f := func(x *Stein) bool {
		// t.Logf("x = %v", x)
		return x.Quad().Sign() > 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
