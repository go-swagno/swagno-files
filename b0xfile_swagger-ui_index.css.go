// Code generaTed by fileb0x at "2022-09-01 14:22:46.84184 +0300 +03 m=+0.119103106" from config file "b0x.yaml" DO NOT EDIT.
// modified(2022-09-01 13:54:08.571317329 +0300 +03)
// original path: swagger-ui/index.css

package static

import (
	"os"
)

// FileSwaggerUIIndexCSS is "swagger-ui/index.css"
var FileSwaggerUIIndexCSS = []byte("\x68\x74\x6d\x6c\x20\x7b\x0a\x20\x20\x20\x20\x62\x6f\x78\x2d\x73\x69\x7a\x69\x6e\x67\x3a\x20\x62\x6f\x72\x64\x65\x72\x2d\x62\x6f\x78\x3b\x0a\x20\x20\x20\x20\x6f\x76\x65\x72\x66\x6c\x6f\x77\x3a\x20\x2d\x6d\x6f\x7a\x2d\x73\x63\x72\x6f\x6c\x6c\x62\x61\x72\x73\x2d\x76\x65\x72\x74\x69\x63\x61\x6c\x3b\x0a\x20\x20\x20\x20\x6f\x76\x65\x72\x66\x6c\x6f\x77\x2d\x79\x3a\x20\x73\x63\x72\x6f\x6c\x6c\x3b\x0a\x7d\x0a\x0a\x2a\x2c\x0a\x2a\x3a\x62\x65\x66\x6f\x72\x65\x2c\x0a\x2a\x3a\x61\x66\x74\x65\x72\x20\x7b\x0a\x20\x20\x20\x20\x62\x6f\x78\x2d\x73\x69\x7a\x69\x6e\x67\x3a\x20\x69\x6e\x68\x65\x72\x69\x74\x3b\x0a\x7d\x0a\x0a\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x6d\x61\x72\x67\x69\x6e\x3a\x20\x30\x3b\x0a\x20\x20\x20\x20\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x3a\x20\x23\x66\x61\x66\x61\x66\x61\x3b\x0a\x7d\x0a")

func init() {

	f, err := FS.OpenFile(CTX, "swagger-ui/index.css", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileSwaggerUIIndexCSS)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
