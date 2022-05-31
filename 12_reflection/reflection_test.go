package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Jeff"},
			[]string{"Jeff"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Jeff", "Denver"},
			[]string{"Jeff", "Denver"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Jeff", 37},
			[]string{"Jeff"},
		},
		{
			"nested fields",
			Person{
				"Jeff",
				Profile{37, "Denver"},
			},
			[]string{"Jeff", "Denver"},
		},
		{
			"pointers to things",
			&Person{
				"Jeff",
				Profile{37, "Denver"},
			},
			[]string{"Jeff", "Denver"},
		},
		{
			"slices",
			[]Profile{
				{37, "Denver"},
				{23, "Anchorage"},
			},
			[]string{"Denver", "Anchorage"},
		},
		{
			"arrays",
			[2]Profile{
				{37, "Denver"},
				{23, "Anchorage"},
			},
			[]string{"Denver", "Anchorage"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}
type Profile struct {
	Age  int
	City string
}
