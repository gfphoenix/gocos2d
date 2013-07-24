package gocos2d

import (
	glfw "github.com/go-gl/glfw3"
	gl "github.com/mortdeus/egles/es2"
	"log"
)

type window struct {
	win           *glfw.Window
	Width, Height int
	AppTitle      string
}

func errorCallback(err glfw.ErrorCode, desc string) {
	log.Printf("%v: %v\n", err, desc)
}
func (w *window) Init() {
	glfw.SetErrorCallback(errorCallback)
	if !glfw.Init() {
		println("glfw init failure")
	}
	if w.Width < 150 || w.Height == 150 {
		w.Width = 680
		w.Height = 480
	}
	if w.AppTitle == "" {
		w.AppTitle = "gocos2d"
	}
	glfw.WindowHint(glfw.ClientApi, glfw.OpenglEsApi)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
	var err error
	w.win, err = glfw.CreateWindow(w.Width, w.Height, w.AppTitle, nil, nil)
	if err != nil {
		panic(err)
	}
	w.win.MakeContextCurrent()
	glfw.SwapInterval(1)
}

func (w *window) reshape(width, height int) {
	w.Width, w.Height = width, height
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}
func (w *window) cleanup() {
	w.win.Destroy()
	glfw.Terminate()
}