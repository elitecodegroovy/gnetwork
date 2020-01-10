package flyweight

import "time"

type Item struct {
	Id      int32
	Name    string
	Content string
}

type HeavyStone struct {
	Tag   string
	Items []Item
}

func CreateHeavyStore(tag string) HeavyStone {
	return HeavyStone{
		Tag: tag,
		Items: []Item{{Id: 1,
			Name:    "1 value",
			Content: "00000000000000000000000000000000000",
		},
			{Id: 2,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 2,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 3,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 4,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 5,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 6,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 7,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 2,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 8,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
			{Id: 9,
				Name:    "2 value",
				Content: "10000000000000000000000000000000000",
			},
		},
	}
}

type Flyweight struct {
	heavyStoneMap map[string]*HeavyStone
}

func NewFlyweight() Flyweight {
	return Flyweight{heavyStoneMap: make(map[string]*HeavyStone, 0)}
}

func (f *Flyweight) GetHeavyStone(tag string) *HeavyStone {
	time.Sleep(100 * time.Microsecond)
	if f.heavyStoneMap[tag] != nil {
		return f.heavyStoneMap[tag]
	}
	heavyStone := CreateHeavyStore(tag)
	f.heavyStoneMap[tag] = &heavyStone
	return &heavyStone
}
