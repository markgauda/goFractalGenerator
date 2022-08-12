/*
This program was created by Mark Gauda in the Summer of 2022
This file will implement arbitrary precision complex numbers

*/

package main

import "math/big"

//an arbitrary-precision complex number
type arbPrecComplex struct {
	real      big.Float
	imaginary big.Float
}

//perfomes the operation left * right aka (a+bi) * (c+di)
func (left arbPrecComplex) multiply(rigth arbPrecComplex) arbPrecComplex {
	var return_value arbPrecComplex
	// expands to ac - bd + adi + bci
	var ac big.Float
	var bd big.Float
	var adi big.Float
	var bci big.Float
	ac.Mul(&left.real, &rigth.real)
	bd.Mul(&left.imaginary, &rigth.imaginary)
	adi.Mul(&left.real, &rigth.imaginary)
	bci.Mul(&left.imaginary, &rigth.real)
	return_value.real.Sub(&ac, &bd)
	return_value.imaginary.Add(&adi, &bci)
	return return_value
}

//performs the operation left + right aka (a+bi) + (c+di)
func (left arbPrecComplex) add(right arbPrecComplex) arbPrecComplex {
	var return_value arbPrecComplex
	//factor to (a+c) + (b+d)i
	return_value.real.Add(&left.real, &right.real)
	return_value.imaginary.Add(&left.imaginary, &right.imaginary)
	return return_value
}

//performs the operation |(a+bi)|
func (c arbPrecComplex) abs() big.Float {
	var return_value big.Float
	//expands to √(a*a + b*b)
	var aSqr big.Float
	var bSqr big.Float
	var aSqrPlusbSqr big.Float
	aSqr.Mul(&c.real, &c.real)
	bSqr.Mul(&c.imaginary, &c.imaginary)
	aSqrPlusbSqr.Add(&aSqr, &bSqr)
	return_value.Sqrt(&aSqrPlusbSqr)
	return return_value
}
