package factory

import (
	"errors"

	"design-pattern-golang-example/factory_pattern/concrete"
	"design-pattern-golang-example/factory_pattern/constants"
)

type AnimalFactory struct {
	Cat *concrete.Cat
	Dog *concrete.Dog
	Rat *concrete.Rat
}

func (af *AnimalFactory) GetAnimalFactory(name string) (IAnimalFactory, error) {
	if name == constants.CAT {
		return af.Cat, nil
	} else if name == constants.DOG {
		return af.Dog, nil
	} else if name == constants.RAT {
		return af.Rat, nil
	}

	return nil, errors.New("getting animal factory is failed")
}
