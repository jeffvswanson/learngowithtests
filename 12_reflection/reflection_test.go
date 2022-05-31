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
	t.Run("maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{37, "Denver"}
			aChannel <- Profile{23, "Anchorage"}
			close(aChannel)
		}()
		var got []string
		want := []string{"Denver", "Anchorage"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{37, "Denver"}, Profile{23, "Anchorage"}
		}

		var got []string
		want := []string{"Denver", "Anchorage"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, got []string, want string) {
	t.Helper()
	contains := false
	for _, x := range got {
		if x == want {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q, but value not present", got, want)
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
