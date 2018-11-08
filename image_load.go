package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	"os"
	"unsafe"
)

// #cgo LDFLAGS: -lm
// #define STB_IMAGE_IMPLEMENTATION
// #include "stb_image.h"
// #include <stdlib.h>
import "C"

func stbLoad(file string) image.Image {
	var x, y, n C.int
	cfn := C.CString(file)
	defer C.free(unsafe.Pointer(cfn))
	data := C.stbi_load(cfn, &x, &y, &n, 4)
	defer C.free(unsafe.Pointer(data))

	godata := C.GoBytes(unsafe.Pointer(data), y*x*4)
	rgba := &image.RGBA{
		Pix:    godata,
		Stride: 4,
		Rect:   image.Rect(0, 0, int(x), int(y)),
	}
	return rgba
}

func imageLoad(file string) image.Image {
	imgFile, err := os.Open(file)
	if err != nil {
		panic(fmt.Sprintf("Can't open: %s", err))
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(fmt.Sprintf("Can't decode: %s", err))
	}

	// convert to RGBA (as that's what we need..)
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.ZP, draw.Src)
	return rgba
}
