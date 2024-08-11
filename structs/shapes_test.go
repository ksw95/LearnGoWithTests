package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{20.0, 15.0}
	got := Perimeter(rectangle)
	want := 70.0

	if got != want {
		t.Errorf("got %.3f want %.3f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Length: 15.0, Width: 12.0}, expectedArea: 180.0},
		{name: "Circle", shape: Circle{Radius: 5.0}, expectedArea: math.Pi * 5 * 5},
		{name: "Triangle", shape: Triangle{10.0, 16.0}, expectedArea: 80},
	}

	for _, testShape := range areaTests {
		t.Run(testShape.name, func(t *testing.T) {
			got := testShape.shape.Area()
			want := testShape.expectedArea

			if got != want {
				t.Errorf("%#v got %g want %g", testShape.shape, got, testShape.expectedArea)
			}
		})
	}

}
