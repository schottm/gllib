package gllib

import (
	"github.com/schottm/gllib/logic"
	"github.com/schottm/gllib/gui"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v3.3-core/gl"
	"time"
	"log"
)


type Alignment int
type FrameRate int

const (
	openglMajorVersion = 3
	openglMinorVersion = 3

	OPENGL_ALIGNMENT Alignment = 0
	SIZE_ALIGNMENT Alignment = 1
	NORMAL_ALIGNMENT Alignment = 2

	USE_VSYNC = -1
	FPS_NO_LIMIT = 0
)

//display//

type Context interface {

	draw(timeDelta int64, display *Display)
}

type Display struct {

	window *glfw.Window
	context []Context

	frameRate FrameRate

	lastUpdate time.Time
	sleptTime time.Duration

	width, height int
}

func NewDisplay(width, height int, title string,  resizeable bool, frameRate FrameRate) *Display {

	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, openglMajorVersion)
	glfw.WindowHint(glfw.ContextVersionMinor, openglMinorVersion)
	if !resizeable {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	//window.SetSizeLimits(width, height, glfw.GetPrimaryMonitor().GetVideoMode().Width, glfw.GetPrimaryMonitor().GetVideoMode().Height)
	//window.SetAspectRatio(1, 1)
	//set the window centered
	var x, y = window.GetSize()
	x = glfw.GetPrimaryMonitor().GetVideoMode().Width - x
	y = glfw.GetPrimaryMonitor().GetVideoMode().Height - y

	window.SetPos(x / 2, y / 2)

	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	extension := glfw.ExtensionSupported("WGL_EXT_swap_control_tear") || glfw.ExtensionSupported("GLX_EXT_swap_control_tear")

	if frameRate == USE_VSYNC && extension {
		glfw.SwapInterval(1)
	} else if frameRate == FPS_NO_LIMIT {
		glfw.SwapInterval(0)
	} else {
		frameRate = FrameRate(60)
	}

	initOpenGL()

	return &Display{window, []Context{}, frameRate, time.Now(), 0, x, y}
}

func (display *Display) AddContext(context Context) {

	display.context = append(display.context, context)
}

func (display *Display) Update() {

	display.window.MakeContextCurrent()

	currentTime := time.Now()
	deltaTime := time.Since(display.lastUpdate)

	x, y := display.window.GetSize()
	if display.width != x || display.height != y {
		display.width = x
		display.height = y
		gl.Viewport(0, 0, int32(x), int32(y))
	}

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, e := range display.context {

		e.draw(int64(deltaTime), display)
	}

	if display.frameRate != USE_VSYNC && display.frameRate != FPS_NO_LIMIT {

		drawingTime := deltaTime - display.sleptTime
		time.Sleep((time.Second / time.Duration(display.frameRate)) - drawingTime)
		display.sleptTime = (time.Second / time.Duration(display.frameRate)) - drawingTime
	}

	display.lastUpdate = currentTime
	display.window.SwapBuffers()
}

func (display *Display) Destroy() {

	display.window.Destroy()
}

func (display *Display) ShouldClose() bool {

	return display.window.ShouldClose()
}

func initOpenGL() {

	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
}


//the GUI//
type UIOverlay struct {

	component gui.Component
	alignment Alignment
}

func NewUIOverlay(component gui.Component, alignment Alignment) *UIOverlay {

	return &UIOverlay{component, alignment}
}

func (uio *UIOverlay) draw(timeDelta int64, display *Display){

	if uio.alignment == OPENGL_ALIGNMENT {

		uio.component.Draw(logic.NewIdentityMatrix4f(), timeDelta)
	} else if uio.alignment == SIZE_ALIGNMENT {
		//TODO : set size alignment
		uio.component.Draw(logic.NewIdentityMatrix4f(), timeDelta)
	} else {
		matrix := logic.NewIdentityMatrix4f()
		matrix.Translate(&logic.Vector3f{-1, 1, 0})
		matrix = matrix.Mul(logic.NewScaleMatrix4f(2.0 * uio.component.GetSize().X ,
			-2.0 * uio.component.GetSize().Y, 1.0, 1.0))

		uio.component.Draw(matrix, timeDelta)
	}
}

