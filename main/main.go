package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"os"
	"fmt"
	"bytes"
	"strings"
	"runtime"
	"github.com/schottm/gllib/logic"
	"math"
	"github.com/schottm/gllib/gui"
	"github.com/schottm/gllib"
	"sync"
)

var (
	triangle = []float32{
		0, 0,
		1, 0,
		0.5, 1}
	translU int32
	colourU int32
	prog uint32
	vao uint32
)


type TestPanel struct {
	gui.Panel

	tick int64
}

func (tp *TestPanel) Draw(transform *logic.Matrix4f, timeDelta int64) {

	gl.UseProgram(prog)

	tp.tick += timeDelta

	x := float64(tp.tick) / 500000000

	current := 0.5 + float32(math.Sin(x)) / 2.0
	mat := logic.NewIdentityMatrix4f()
	mat.Translate(&logic.Vector3f{(1 - current) /  2, (1 - current) / 2, 0})
	scale := logic.NewScaleMatrix4f(current, current, 1, 1)
	mat = mat.Mul(scale)

	gl.BindVertexArray(vao)
	gl.UniformMatrix4fv(translU, 1, false, &transform.Mul(mat).M00)
	gl.Uniform3f(colourU,
		0.5 + float32(math.Sin(x) / 2),
		0.5 + float32(math.Sin(x + math.Pi * 2 / 3) / 2),
		0.5 + float32(math.Sin(x + math.Pi * 4 / 3) / 2))
	//gl.Uniform3f(colourU, 1, 1, 1)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle))/2)
}

func main() {

	runtime.LockOSThread()

	defer glfw.Terminate()

	var display = gllib.NewDisplay(500, 500, "gllib", true, gllib.USE_VSYNC)
	defer display.Destroy()

	prog = createProgram()
	translU = gl.GetUniformLocation(prog,  gl.Str("transform" + "\x00"))
	colourU = gl.GetUniformLocation(prog,  gl.Str("colour" + "\x00"))
	defer gl.DeleteProgram(prog)

	vao, _ = createVAO(triangle)

	var pane gui.Container = &gui.ContentPane{}
	pane.SetSize(&logic.Vector2f{1, 1})

	var layout = gui.NewDefaultLayout()
	pane.SetLayout(layout)

	var a gui.Component = &TestPanel{}
	a.SetSize(&logic.Vector2f{0.5, 0.5})
	pane.Add(a)
	layout.AddComponent(a, &logic.Vector2f{0,0})

	var b gui.Component = &TestPanel{}
	b.SetSize(&logic.Vector2f{0.5, 0.5})
	pane.Add(b)
	layout.AddComponent(b, &logic.Vector2f{0.5, 0})

	var c gui.Component = &TestPanel{}
	c.SetSize(&logic.Vector2f{0.5, 0.5})
	pane.Add(c)
	layout.AddComponent(c, &logic.Vector2f{0.25, 0.5})

	var context gllib.Context = gllib.NewUIOverlay(pane, gllib.OPENGL_ALIGNMENT)

	display.AddContext(context)

	var mutex = sync.Mutex{}

	for !display.ShouldClose() {

		//add regular update thread management with mutex

		mutex.Lock()
		glfw.PollEvents()

		display.Update()
		mutex.Unlock()
	}
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