package main

import (
	"fmt"

	"design-pattern-golang-example/factory_pattern/concrete"
	"design-pattern-golang-example/factory_pattern/constants"
	"design-pattern-golang-example/factory_pattern/factory"
)

func main() {
	// init
	cat := concrete.NewCat()
	dog := concrete.NewDog()
	rat := concrete.NewRat()
	animalFactory := &factory.AnimalFactory{
		Cat: cat,
		Dog: dog,
		Rat: rat,
	}

	// factory
	processor, err := animalFactory.GetAnimalFactory(constants.CAT)
	if err != nil {
		fmt.Println(err)
	}
	sound := processor.GetSound()
	fmt.Println(sound)
}
