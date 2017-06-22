package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
	"log"
	"os"
	"fmt"
	"bytes"
	"strings"
)

func main() {

	runtime.LockOSThread()

	window := createWindow(500, 500)
	defer glfw.Terminate()

	initOpenGL()

	prog := createProgram()
	defer gl.DeleteProgram(prog)

	for !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(prog)

		//draw some shit

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
}

func createProgram() uint32 {

	assets := getAssetsLocation()

	fragShader, err := compileShader(getFileContent(assets + "/default.frag") + "\x00", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	vertShader, err := compileShader(getFileContent(assets + "/default.frag") + "\x00", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertShader)
	gl.AttachShader(prog, fragShader)
	gl.LinkProgram(prog)

	return prog
}

func createWindow(width, height int) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "opengl_app", nil, nil)

	//set the window centered
	var x, y = window.GetSize()
	x = glfw.GetPrimaryMonitor().GetVideoMode().Width - x
	y = glfw.GetPrimaryMonitor().GetVideoMode().Height - y

	window.SetPos(x / 2, y / 2)

	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func getAssetsLocation() string {

	assets, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return assets + "/src/github.com/schottm/opengl_app/assets"
}

func getFileContent(fName string) string {

	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	return buf.String()
}