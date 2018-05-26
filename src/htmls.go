package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Htmls5ac9c22e7576a1ac6808719e0cd81981516f3485 = "<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>Hello World!</title>\n</head>\n<body>\n\t<p>Hello World2!</p>\n\t<img src=\"image/e-1.jpg\"><br>\n</body>\n</html>\n"
var _Htmls2d02122e63c8aaf7ec0d8dc390210672fa100727 = "<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>Hello World!</title>\n</head>\n<body>\n\t<p>Hello World2!</p>\n\t<img src=\"image/e-1.jpg\"><br>\n</body>\n</html>\n"

// Htmls returns go-assets FileSystem
var Htmls = assets.NewFileSystem(map[string][]string{"/": []string{"index.html"}, "/product": []string{"index.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1527299854, 1527299854911388422),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1527271120, 1527271120044247613),
		Data:     []byte(_Htmls5ac9c22e7576a1ac6808719e0cd81981516f3485),
	}, "/product": &assets.File{
		Path:     "/product",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1527299864, 1527299864176617031),
		Data:     nil,
	}, "/product/index.html": &assets.File{
		Path:     "/product/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1527299864, 1527299864176765194),
		Data:     []byte(_Htmls2d02122e63c8aaf7ec0d8dc390210672fa100727),
	}}, "")
