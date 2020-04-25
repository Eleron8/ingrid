package data

import (
	"testing"
)

func TestSort(t *testing.T) {
	legs := []Leg{
		Leg{
			Duration: 100,
			Distance: 100,
		},
		Leg{
			Duration: 100,
			Distance: 90,
		},
	}
	resp := Response{
		Routes: []Route{
			Route{
				Legs: legs,
			},
		},
	}
	list := List{
		Source: "test",
		Routes: []RoutInfo{
			RoutInfo{
				Destination: "1",
			},
			RoutInfo{
				Destination: "2",
			},
		},
	}
	sortList := sortInfo(resp, &list)
	for i, v := range sortList.Routes {
		if i == 0 {
			if v.Distance != 90 {
				t.Errorf("sort failed, expected %v, got %v", 90, v.Distance)
			}
		}
	}
}
