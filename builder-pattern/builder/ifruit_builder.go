package builder

import (
	"design-pattern-golang-example/builder-pattern/concreate_builder"
	"design-pattern-golang-example/builder-pattern/constant"
	"design-pattern-golang-example/builder-pattern/product"
)

type IFruitBuilder interface {
	SetPrice(in *product.InfoInput)
	SetDiscount(in *product.InfoInput)
	SetMoneyPayment(in *product.InfoInput)
	ToOutputMoney() *product.InfoOutput
}

func GetFruitBuilder(fruitName string) IFruitBuilder {
	switch fruitName {
	case constant.Orange:
		return concreate_builder.NewOrangeBuilder()
	case constant.Apple:
		return concreate_builder.NewAppleBuilder()
	}
	return nil
}
