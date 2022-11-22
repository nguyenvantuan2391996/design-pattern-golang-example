package main

import (
	"fmt"

	"design-pattern-golang-example/builder-pattern/builder"
	"design-pattern-golang-example/builder-pattern/constant"
	"design-pattern-golang-example/builder-pattern/director"
	"design-pattern-golang-example/builder-pattern/product"
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

	fmt.Println(outApple.MoneyPayment)

	builderOrange := builder.GetFruitBuilder(constant.Orange)
	directorOrange := director.NewDirector(builderOrange)
	outOrange := directorOrange.BuildOutput(in)

	fmt.Println(outOrange.MoneyPayment)
}
