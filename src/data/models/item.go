package models

type Item struct {
	Name string
	Cost string
}

func (item *Item) getNameOfItem() string {
	return item.Name
}

func (item *Item) getCostOfItem() string {
	return item.Cost
}
