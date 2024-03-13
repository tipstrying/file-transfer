package web

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets4548c4de5a42a759758868d96443b68ee4bacf04 = "<!DOCTYPE html>\r\n<html lang=\"zh\">\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n    <title>Document</title>\r\n</head>\r\n<body>\r\n    <h4>\r\n        欢迎使用文件传输助手\r\n    </h4>\r\n    <div>\r\n        <form id=\"uploadbanner\" enctype=\"multipart/form-data\" method=\"post\" action=\"/\">\r\n            <input id=\"fileupload\" name=\"file\" type=\"file\" />\r\n            <input type=\"submit\" value=\"submit\" id=\"submit\" />\r\n         </form>\r\n    </div>\r\n    <div>\r\n        <ul>\r\n            {{ range .files }}\r\n            <li >\r\n                <a href=\"/{{ . }}\">{{ . }}</a>\r\n            </li>\r\n            {{ end }}\r\n        </ul>\r\n        \r\n    </div>\r\n</body>\r\n</html>\r\n<style>\r\n    li{\r\n        cursor: pointer;\r\n        user-select: none;\r\n    }\r\n</style>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1710305409, 1710305409888512200),
		Data:     nil,
	}, "/index.tmpl": &assets.File{
		Path:     "/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1710308900, 1710308900981277100),
		Data:     []byte(_Assets4548c4de5a42a759758868d96443b68ee4bacf04),
	}}, "")
