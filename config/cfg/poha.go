package cfg

import "sync"

var p Poha
var once sync.Once

type Poha struct{
	Image string `json:"image"`
	Ingredients string `json:"ingredients"`
	Recipie string `json:"recipie"`
}

func getPohaConfig() Poha{
	once.Do(func() {
		LoadConfig("poha.json", &p)
	})
	return p
}