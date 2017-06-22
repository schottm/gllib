package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
	"log"
	"os"
	"bytes"
	"fmt"
)

func main() {

	runtime.LockOSThread()

	window := createWindow(500, 500)
	defer glfw.Terminate()

	initOpenGL()

	//createProgram()

	fmt.Println(os.Args[0])


	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.ClearColor(0.5, 0.5, 0.5, 1.0)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}

func initOpenGL() {

	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	/*
	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
	*/
}

func createProgram() uint32 {

	file, err := os.Open("assets/shader/default.frag")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	log.Println(buf.String())


	return 0
}

func createWindow(width, height int) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	//glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "opengl_app", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}
