package conposite

import "testing"

func TestComposite(t *testing.T) {
	circle := &Circle{
		Center: Point{X: 100, Y: 100},
		Radius: 50,
	}

	square := &Square{
		Location: Point{X: 50, Y: 50},
		Side:     20,
	}

	layer := &Layer{
		Elements: []VisualElement{
			circle,
			square,
		},
	}

	layer.Draw(&Drawer{})
}
