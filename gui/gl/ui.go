package dota2bot

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/gurupras/dota2bot"
	"github.com/gurupras/dota2bot/gui"
	"github.com/llgcode/draw2d/draw2dgl"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/sirupsen/logrus"
)

type MiniMap struct {
	*gui.BaseGUI
	window   *glfw.Window
	mapImage image.Image
}

func NewMiniMap(gameInfo dota2bot.GameInfo) (*MiniMap, error) {
	img, err := getImageFromFilePath("/home/guru/workspace/go/src/github.com/gurupras/dota2bot/server/resources/map-7.27-1024x1024.jpg")
	if err != nil {
		return nil, err
	}
	baseGUI := gui.NewBaseGUI(gameInfo, img.Bounds())
	m := &MiniMap{
		baseGUI,
		nil,
		img,
	}
	logrus.Printf("Game bounds:   %v", gameInfo.WorldBounds)
	logrus.Printf("Image bounds:  %v", img.Bounds())
	return m, nil
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	if err != nil {
		logrus.Errorf("Image has unknown format in getImageFromFilePath")
	}
	return image, err
}

func (m *MiniMap) initGlfw() error {
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("Failed to initialize glfw: %v", err)
	}

	// glfw.WindowHint(glfw.Resizable, glfw.False)
	// glfw.WindowHint(glfw.ContextVersionMajor, 2) // OR 2
	// glfw.WindowHint(glfw.ContextVersionMinor, 1)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(1000, 1000, "Dota2 MiniMap", nil, nil)
	if err != nil {
		return err
	}
	window.MakeContextCurrent()
	m.window = window
	return nil
}

func (m *MiniMap) reshape(window *glfw.Window, w, h int) {
	gl.ClearColor(1, 1, 1, 1)
	/* Establish viewing area to cover entire window. */
	gl.Viewport(0, 0, int32(w), int32(h))
	/* PROJECTION Matrix mode. */
	gl.MatrixMode(gl.PROJECTION)
	/* Reset project matrix. */
	gl.LoadIdentity()
	/* Map abstract coords directly to window coords. */
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	/* Invert Y axis so increasing Y goes down. */
	gl.Scalef(1, -1, 1)
	/* Shift origin up to upper-left corner. */
	gl.Translatef(0, float32(-h), 0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.DEPTH_TEST)
}

func (m *MiniMap) draw() {
	gc := draw2dgl.NewGraphicContext(1000, 1000)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gc.DrawImage(m.mapImage)
	gc.BeginPath()
	for _, entry := range m.LastUpdate {
		draw2dkit.Circle(gc, entry.X, entry.Y, entry.Radius)
		gc.SetFillColor(color.RGBA{uint8(entry.Red), uint8(entry.Green), uint8(entry.Blue), 0xff})
		gc.Fill()
	}
	gl.Flush()
}

func (m *MiniMap) Run() {
	runtime.LockOSThread()
	if err := m.initGlfw(); err != nil {
		logrus.Errorf("%v", err)
		return
	}
	window := m.window
	defer glfw.Terminate()
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		logrus.Errorf("gl failed to init: %v", err)
		return

	}
	m.reshape(window, m.Bounds.Dx(), m.Bounds.Dy())
	for !window.ShouldClose() {
		m.Lock()
		m.draw()
		m.Unlock()
		// TODO
	}
}
