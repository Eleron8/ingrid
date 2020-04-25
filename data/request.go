package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

//Response - form for take response from project-osrm
type Response struct {
	Routes    []Route    `json:"routes"`
	Waypoints []Waypoint `json:"waypoints"`
	Code      string     `json:"code"`
}

type Leg struct {
	Summary  string  `json:"summary"`
	Weight   float32 `json:"weight"`
	Duration float32 `json:"duration"`
	Steps    []bool  `json:"steps"`
	Distance float32 `json:"distance"`
}

type Route struct {
	Legs       []Leg   `json:"legs"`
	WeightName string  `json:"weight_name"`
	Weight     float32 `json:"weight"`
	Duration   float32 `json:"duration"`
	Distance   float32 `json:"distance"`
}

type Waypoint struct {
	Hint     string    `json:"hint"`
	Distance float64   `json:"distance"`
	Name     string    `json:"name"`
	Location []float64 `json:"location"`
}

type RoutInfo struct {
	Destination string  `json:"destination"`
	Duration    float32 `json:"duration"`
	Distance    float32 `json:"distance"`
}

//List - response from this service
type List struct {
	Source string     `json:"source"`
	Routes []RoutInfo `json:"routes"`
}

func Request(input [][]string) (*List, error) {
	str := urlString(input)
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%s?overview=false", str)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	response := Response{}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	source := fmt.Sprintf("%s, %s", input[0][0], input[0][1])
	list := List{
		Source: source,
	}
	for i, v := range input {
		if i > 0 {
			s := fmt.Sprintf("%s, %s", v[0], v[1])
			route := RoutInfo{
				Destination: s,
			}
			list.Routes = append(list.Routes, route)
		}
	}
	sortList := sortInfo(response, &list)
	return sortList, err
}

func urlString(input [][]string) string {
	var str string
	for i, v := range input {
		if i == len(input)-1 {
			s := fmt.Sprintf("%s,%s", v[0], v[1])
			str = str + s
		} else {
			s := fmt.Sprintf("%s,%s;", v[0], v[1])
			str = str + s
		}
	}
	return str
}

func sortInfo(r Response, list *List) *List {
	for _, v := range r.Routes {
		for i, k := range v.Legs {
			list.Routes[i].Duration = k.Duration
			list.Routes[i].Distance = k.Distance
		}
	}
	sort.Slice(list.Routes, func(i, j int) bool { return list.Routes[i].Duration < list.Routes[j].Duration })
	var duration float32
	for i, v := range list.Routes {
		if i > 0 {
			if v.Duration == duration {
				if v.Distance < list.Routes[i-1].Distance {
					list.Routes[i] = list.Routes[i-1]
					list.Routes[i-1] = v
				}
			}
		}
		duration = v.Duration
	}
	return list
}
