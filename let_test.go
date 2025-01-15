package gosfunc

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  uint
}

func TestLet(t *testing.T) {
	type args struct {
		it    []Person
		block func(it []Person) Person
	}
	tests := []struct {
		name string
		args args
		want Person
	}{
		{
			name: "1",
			args: args{
				it: []Person{
					{
						Name: "233",
						Age:  5,
					},
					{
						Name: "wdf",
						Age:  10,
					},
				},
				block: func(it []Person) Person {
					for _, item := range it {
						if item.Age > 5 {
							return item
						}
					}
					return Person{}
				},
			},
			want: Person{
				Name: "wdf",
				Age:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Let(tt.args.it, tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Let() = %v, want %v", got, tt.want)
			}
		})
	}
}
