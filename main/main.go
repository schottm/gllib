package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"os"
	"fmt"
	"bytes"
	"strings"
	"runtime"
	"github.com/schottm/gllib/logic"
	"math"
	"time"
	"github.com/schottm/gllib/gui"
)

const (

	fps = 120
	use_vsync = true
)

var (
	has_vsync_extension = false

	triangle = []float32{
		0.5, 1,
		0, 0,
		1, 0}
	translU int32
	colourU int32
	vao uint32
)


type TestPanel struct {
	gui.Panel

	tick int64
}

func (tp *TestPanel) Draw(timeDelta int64) {

	tp.tick += timeDelta

	x := float64(tp.tick) / 500000000

	current := 0.5 + float32(math.Sin(x)) / 2.0
	mat := logic.NewIdentityMatrix4f()
	mat.Translate(&logic.Vector3f{(1 - current) /  2, (1 - current) / 2, 0})
	scale := logic.NewScaleMatrix4f(current, current, 1, 1)
	mat = mat.Mul(scale)


	gl.BindVertexArray(vao)
	gl.UniformMatrix4fv(translU, 1, false, &tp.GetTransform().Mul(mat).M00)
	gl.Uniform3f(colourU,
		0.5 + float32(math.Sin(x) / 2),
		0.5 + float32(math.Sin(x + math.Pi * 2 / 3) / 2),
		0.5 + float32(math.Sin(x + math.Pi * 4 / 3) / 2))

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle))/2)
}

func main() {

	runtime.LockOSThread()

	window := createWindow(500, 500)
	defer glfw.Terminate()
	defer window.Destroy()

	initOpenGL()

	prog := createProgram()
	translU = gl.GetUniformLocation(prog,  gl.Str("transform" + "\x00"))
	colourU = gl.GetUniformLocation(prog,  gl.Str("colour" + "\x00"))
	defer gl.DeleteProgram(prog)

	vao, _ = createVAO(triangle)

	width, height := window.GetSize()

	//Initialize Component
	var context = &gui.ContentPane{}
	context.SetSize(&logic.Vector2f{1, 1})

	var layout = gui.NewDefaultLayout()
	context.SetLayout(layout)

	var a gui.Component = &TestPanel{}
	a.SetSize(&logic.Vector2f{0.5, 0.5})
	context.Add(a)
	layout.AddComponent(a, &logic.Vector2f{0,0})

	var b gui.Component = &TestPanel{}
	b.SetSize(&logic.Vector2f{0.5, 0.5})
	context.Add(b)
	layout.AddComponent(b, &logic.Vector2f{0.25, 0.5})

	var c gui.Component = &TestPanel{}
	c.SetSize(&logic.Vector2f{0.5, 0.5})
	context.Add(c)
	layout.AddComponent(c, &logic.Vector2f{0.5, 0})

	var lastTime = time.Now()
	var sleptTime time.Duration = 0

	for !window.ShouldClose() {

		//calculate delta time
		currentTime := time.Now()
		deltaTime := time.Since(lastTime)

		fmt.Println(deltaTime)

		if !use_vsync || !has_vsync_extension {

			drawingTime := deltaTime - sleptTime
			time.Sleep((time.Second / fps) - drawingTime)
			sleptTime = (time.Second / fps) - drawingTime
		}

		lastTime = currentTime

		//check window resize
		cwidth, cheight := window.GetSize()
		if cwidth != width || cheight != height {
			width = cwidth
			height = cheight
			gl.Viewport(0, 0, int32(width), int32(height))
		}

		//clear frame buffer
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		//use default shader
		gl.UseProgram(prog)

		context.Draw(int64(deltaTime))

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

	fragShader, err := compileShader(getFileContent(assets + "/default.frag"), gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	vertShader, err := compileShader(getFileContent(assets + "/default.vert"), gl.VERTEX_SHADER)
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
	//glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "gllib", nil, nil)


	window.SetSizeLimits(width, height, glfw.GetPrimaryMonitor().GetVideoMode().Width, glfw.GetPrimaryMonitor().GetVideoMode().Height)
	window.SetAspectRatio(1, 1)
	//set the window centered
	var x, y = window.GetSize()
	x = glfw.GetPrimaryMonitor().GetVideoMode().Width - x
	y = glfw.GetPrimaryMonitor().GetVideoMode().Height - y

	window.SetPos(x / 2, y / 2)

	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	has_vsync_extension = glfw.ExtensionSupported("WGL_EXT_swap_control_tear") || glfw.ExtensionSupported("GLX_EXT_swap_control_tear")

	if use_vsync && has_vsync_extension {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}

	return window
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logC := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logC))

		return 0, fmt.Errorf("failed to compile %v: %v", source, logC)
	}

	return shader, nil
}

func createVAO(source []float32) (uint32, uint32) {

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(source) * 4, gl.Ptr(source), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)

	return vao, vbo
}

func getAssetsLocation() string {

	assets, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return assets + "/src/github.com/schottm/gllib/assets"
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