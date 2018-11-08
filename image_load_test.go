package main

import (
	"testing"
)

func BenchmarkLoadImageOneSTB(b *testing.B) {
	file := "Free_Spring_Blossoms_on_Blue_Creative_Commons_(3457362713).jpg"
	for n := 0; n < b.N; n++ {
		stbLoad(file)
	}
}
func BenchmarkLoadImageOneGo(b *testing.B) {
	file := "Free_Spring_Blossoms_on_Blue_Creative_Commons_(3457362713).jpg"
	for n := 0; n < b.N; n++ {
		imageLoad(file)
	}
}

func BenchmarkLoadImageTwoSTB(b *testing.B) {
	file := "Free_Unedited_Happy_Little_Yellow_Stars_in_Pink_Flower_Creative_Commons_(2898759838).jpg"
	for n := 0; n < b.N; n++ {
		stbLoad(file)
	}
}
func BenchmarkLoadImageTwoGo(b *testing.B) {
	file := "Free_Unedited_Happy_Little_Yellow_Stars_in_Pink_Flower_Creative_Commons_(2898759838).jpg"
	for n := 0; n < b.N; n++ {
		imageLoad(file)
	}
}
