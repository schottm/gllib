package gui

import (
	"github.com/schottm/gllib/logic"
)

//interfaces//

type Component interface {

	SetSize(size *logic.Vector2f)
	GetSize() *logic.Vector2f

	Draw(transform *logic.Matrix4f, timeDelta int64)
}

type Container interface {

	Component

	Add(component Component)
	SetLayout(layout Layout)
}

///////////////////////////Panel//////////////////////////////

type Panel struct {

	size logic.Vector2f
}

func (cp *Panel) GetSize() *logic.Vector2f {

	return &cp.size
}

func (cp *Panel) SetSize(size *logic.Vector2f) {

	cp.size = *size
}

func (cp *Panel) Draw(transform *logic.Matrix4f, timeDelta int64) {

}

/////////////////////////////////ContentPane//////////////////////////////

type ContentPane struct {
	Panel

	components []Component
	layout Layout
}

func (cp *ContentPane) Add(component Component) {

	cp.components = append(cp.components, component)
}

func (cp *ContentPane) SetLayout(layout Layout) {

	cp.layout = layout
}

func (cp *ContentPane) Draw(transform *logic.Matrix4f, timeDelta int64) {

	for _, e := range cp.components {

		current := logic.NewIdentityMatrix4f()
		current.Translate(cp.layout.GetOffset(e).Vector3f(0.0))
		current = current.Mul(logic.NewScaleMatrix4f(e.GetSize().X, e.GetSize().Y, 1.0, 1.0))

		e.Draw(transform.Mul(current), timeDelta)
	}
}
