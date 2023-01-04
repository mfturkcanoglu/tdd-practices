package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectange := Rectangle{
		Width:  10.0,
		Heigth: 10.0,
	}
	got := Perimeter(rectange)
	want := 40.0

	if got != want {
		t.Errorf("Expected %.2f but found %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v Expected %g but got %g", shape, want, got)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10.0, Heigth: 6.0}, want: 60.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, test := range areaTests {
		checkArea(t, test.shape, test.want)
	}

	// t.Run("rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{
	// 		Width:  10.0,
	// 		Heigth: 6.0,
	// 	}
	// 	want := 60.0

	// 	checkArea(t, rectangle, want)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle := Circle{Radius: 10.0}
	// 	want := 314.1592653589793

	// 	checkArea(t, circle, want)
	// })
}
