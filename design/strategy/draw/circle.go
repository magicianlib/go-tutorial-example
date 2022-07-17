package draw

import (
	"fmt"
	"itknown.io/design/strategy"
)

type Circle struct {
}

func (c *Circle) Support(s strategy.ShapeType) bool {
	return strategy.Circle == s
}

func (c *Circle) Shape(b *strategy.Blueprint) {
	fmt.Println("draw circle...")
}
