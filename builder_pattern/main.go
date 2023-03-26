package main

import (
	"fmt"

	"design-pattern-golang-example/builder_pattern/builder"
	"design-pattern-golang-example/builder_pattern/constant"
	"design-pattern-golang-example/builder_pattern/director"
	"design-pattern-golang-example/builder_pattern/product"
)

func main() {
	in := &product.InfoInput{
		Price:    100,
		Discount: 10,
		Amount:   15,
	}

	builderApple := builder.GetFruitBuilder(constant.Apple)
	directorApple := director.NewDirector(builderApple)
	outApple := directorApple.BuildOutput(in)

	builderOrange := builder.GetFruitBuilder(constant.Orange)
	directorOrange := director.NewDirector(builderOrange)
	outOrange := directorOrange.BuildOutput(in)

	fmt.Printf("Pay for orange : $%v\n", outOrange.MoneyPayment)
	fmt.Printf("Pay for apple : $%v\n", outApple.MoneyPayment)
}
