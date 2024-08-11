package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name              string
		shape             PeriShape
		expectedPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Length: 20.0, Width: 15.0}, expectedPerimeter: 70.0},
		{name: "Circle", shape: Circle{Radius: 8.0}, expectedPerimeter: 2 * math.Pi * 8},
	}

	for _, testShape := range perimeterTests {
		t.Run(testShape.name, func(t *testing.T) {
			got := testShape.shape.Perimeter()
			want := testShape.expectedPerimeter

			if got != want {
				t.Errorf("%#v got %g want %g", testShape.shape, got, testShape.expectedPerimeter)
			}
		})
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
