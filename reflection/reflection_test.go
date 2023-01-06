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
			"struct with one string field",
			struct {
				Name string
			}{"Mf"},
			[]string{"Mf"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Mf", "Ist"},
			[]string{"Mf", "Ist"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Mf", 55},
			[]string{"Mf"},
		},
		{
			"nested fields",
			Person{
				"Mf",
				Profile{55, "Ist"},
			},
			[]string{"Mf", "Ist"},
		},
		{
			"pointers to things",
			&Person{
				"Mf",
				Profile{55, "Ist"},
			},
			[]string{"Mf", "Ist"},
		},
		{
			"slices",
			[]Profile{
				{55, "Ist"},
				{34, "Reykjavík"},
			},
			[]string{"Ist", "Reykjavík"},
		},
		{
			"arrays",
			[2]Profile{
				{55, "Ist"},
				{34, "Reykjavík"},
			},
			[]string{"Ist", "Reykjavík"},
		},
		{
			"maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			assertDeepEqual(t, got, test.ExpectedCalls)
		})
	}

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{55, "Ist"}
			aChannel <- Profile{33, "Berlin"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Ist", "Berlin"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		assertDeepEqual(t, got, want)
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{55, "Ist"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Ist", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})
		assertDeepEqual(t, got, want)
	})
}

func assertDeepEqual(t *testing.T, got, want []string) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but actually got %v", want, got)
	}
}
