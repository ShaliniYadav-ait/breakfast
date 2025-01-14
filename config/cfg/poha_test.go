package cfg

import (
	"reflect"
	"testing"
)

func Test_getPohaConfig(t *testing.T) {
	tests := []struct {
		name string
		want Poha
	}{
		{
			name: "test ingredients",
			want: Poha{
				Ingredients: "1. Poha , 2. Curry Leaves , 3. Mustard seeds, 4. Jeera, 5. Peanuts, 6. Coriander, 7. Onion",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPohaConfig(); !reflect.DeepEqual(got.Ingredients, tt.want.Ingredients) {
				t.Errorf("getPohaConfig() = %v, want %v", got.Ingredients, tt.want.Ingredients)
			}
		})
	}
}
