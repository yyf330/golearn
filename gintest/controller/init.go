package controller

import (
	"github.com/unrolled/render"
	"fmt"
)

var r *render.Render

func init() {
	fmt.Println("this test!!")

	r = render.New(render.Options{
		Directory:     "views",
		IndentJSON:    true,
		Layout:        "container",
		IsDevelopment: true,
		Extensions:    []string{".html"},
	})
}
