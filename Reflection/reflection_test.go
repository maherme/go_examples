package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"maherme"},
			[]string{"maherme"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"maherme", "Barcelona"},
			[]string{"maherme", "Barcelona"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"maherme", 33},
			[]string{"maherme"},
		},
		{
			"Nested fields",
			Person{
				"maherme",
				Profile{33, "Barcelona"},
			},
			[]string{"maherme", "Barcelona"},
		},
		{
			"Pointers to things",
			&Person{
				"maherme",
				Profile{33, "Barcelona"},
			},
			[]string{"maherme", "Barcelona"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Barcelona"},
				{34, "Sevilla"},
			},
			[]string{"Barcelona", "Sevilla"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Barcelona"},
				{34, "Sevilla"},
			},
			[]string{"Barcelona", "Sevilla"},
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

	t.Run("with maps", func(t *testing.T) {
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

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Madrid"}
			aChannel <- Profile{34, "Valencia"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Madrid", "Valencia"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Barcelona"}, Profile{34, "Sevilla"}
		}

		var got []string
		want := []string{"Barcelona", "Sevilla"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
