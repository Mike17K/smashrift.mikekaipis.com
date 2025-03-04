package benchmarking

import (
	"log"
	"sync"
)

type Metric struct {
	ID    string      `json:"id"`
	Value interface{} `json:"value"`
}

var metrics sync.Map

func (m Metric) Validate() bool {
	_, ok := metrics.Load(m.ID)
	return !ok
}

func (m Metric) Show() {
	log.Println(m.ID, m.Value)
}

func (m Metric) Save() {
	metrics.Store(m.ID, m)
}
