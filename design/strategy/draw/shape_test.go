package draw

import (
	"container/list"
	"itknown.io/design/strategy"
	"testing"
)

func TestDraw(t *testing.T) {
	l := list.New()
	l.PushBack(&Circle{})
	l.PushBack(&Square{})

	for e := l.Front(); e != nil; e = e.Next() {
		if draw, ok := e.Value.(strategy.Draw); ok {
			if draw.Support(strategy.Square) {
				draw.Shape(nil)
			}
		}
	}
}
