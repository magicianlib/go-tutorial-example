package draw

import (
	"fmt"
	"itknown.io/design/strategy"
)

type Square struct {
}

func (d *Square) Support(s strategy.ShapeType) bool {
	return strategy.Square == s
}

func (d *Square) Shape(b *strategy.Blueprint) {
	fmt.Println("draw Square...")
}
