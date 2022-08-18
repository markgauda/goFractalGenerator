/*
This program was created by Mark Gauda in the Summer of 2022
This is the program that will test everything else

*/

package main

import (
	"fmt"
	"math/big"
	"testing"
)

func ExampleGame() {
	main()

	println("test")
	//Output:
	//test
}

func TestArbitraryPrecision(t *testing.T) {
	var left, right, answer arbPrecComplex
	left.real.Set(big.NewFloat(1))
	left.imaginary.Set(big.NewFloat(1))
	right.real.Set(big.NewFloat(1))
	right.imaginary.Set(big.NewFloat(1))
	answer = left.add(right)
	if answer.real.Cmp(big.NewFloat(2)) != 0 || answer.imaginary.Cmp(big.NewFloat(2)) != 0 {
		t.Error("Adding Failed, (1+1i) + (1+1i) != " + fmt.Sprintf("%s+%si", answer.real.Text('f', 10), answer.imaginary.Text('f', 10)))
	}
	answer = left.multiply(right)
	if answer.real.Cmp(big.NewFloat(0)) != 0 || answer.imaginary.Cmp(big.NewFloat(2)) != 0 {
		t.Error("Adding Failed, (1+1i) * (1+1i) != " + fmt.Sprintf("%s+%si", answer.real.Text('f', 10), answer.imaginary.Text('f', 10)))
	}
	//println(answer.real.Text('f', 10), "+", answer.imaginary.Text('f', 10), "i")

	left.real.Set(big.NewFloat(2))
	left.imaginary.Set(big.NewFloat(2))
	right.real.Set(big.NewFloat(2))
	right.imaginary.Set(big.NewFloat(2))
	answer = left.add(right)
	if answer.real.Cmp(big.NewFloat(4)) != 0 || answer.imaginary.Cmp(big.NewFloat(4)) != 0 {
		t.Error("Adding Failed, (2+2i) + (2+2i) != " + fmt.Sprintf("%s+%si", answer.real.Text('f', 10), answer.imaginary.Text('f', 10)))
	}
	//println(answer.real.Text('f', 10), "+", answer.imaginary.Text('f', 10), "i")
	answer = left.multiply(right)
	if answer.real.Cmp(big.NewFloat(0)) != 0 || answer.imaginary.Cmp(big.NewFloat(8)) != 0 {
		t.Error("Adding Failed, (2+2i) * (2+2i) != " + fmt.Sprintf("%s+%si", answer.real.Text('f', 10), answer.imaginary.Text('f', 10)))
	}
	//println(answer.real.Text('f', 10), "+", answer.imaginary.Text('f', 10), "i")

	left.real.Set(big.NewFloat(2))
	left.imaginary.Set(big.NewFloat(-0))
	right.real.Set(big.NewFloat(2))
	right.imaginary.Set(big.NewFloat(-0))
	answer = left.add(right)
	answer.print()
	if answer.real.Cmp(big.NewFloat(4)) != 0 || answer.imaginary.Cmp(big.NewFloat(0)) != 0 {
		t.Error("Adding Failed, (2+0i) + (2+0i) != " + fmt.Sprintf("%s+%si", answer.real.Text('f', 10), answer.imaginary.Text('f', 10)))
	}
	//println(answer.real.Text('f', 10), "+", answer.imaginary.Text('f', 10), "i")
	answer = left.multiply(right)
	answer.print()
	if answer.real.Cmp(big.NewFloat(4)) != 0 || answer.imaginary.Cmp(big.NewFloat(0)) != 0 {
		t.Error("Adding Failed, (2+0i) * (2+0i) != " + fmt.Sprintf("%s+%si", answer.real.Text('f', 10), answer.imaginary.Text('f', 10)))
	}

	//Output:
	//2.0000000000 + 2.0000000000 i
	//0.0000000000 + 2.0000000000 i
	//4.0000000000 + 4.0000000000 i
	//0.0000000000 + 8.0000000000 i
}
