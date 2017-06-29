package gui

import (
	"github.com/schottm/gllib/logic"
)

type Component interface {

	SetSize(size *logic.Vector2f)
	GetSize() *logic.Vector2f

	setTransform(size *logic.Matrix4f)
	setParent(size Pane)
	GetTransform() *logic.Matrix4f

	Draw(timeDelta int64)
}

type Pane interface {

	Component

	updateTransform(component Component)
	Add(component Component)
	SetLayout(layout Layout)
}

///////////////////////////Panel//////////////////////////////

type Panel struct {

	parent Pane

	size logic.Vector2f
	transform *logic.Matrix4f
}

func (cp *Panel) GetTransform() *logic.Matrix4f {

	return cp.transform
}

func (cp *Panel) setTransform(transform *logic.Matrix4f) {

	cp.transform = transform
}

func (cp *Panel) setParent(pane Pane) {

	cp.parent = pane
}

func (cp *Panel) GetSize() *logic.Vector2f {

	return &cp.size
}

func (cp *Panel) SetSize(size *logic.Vector2f) {

	cp.size = *size

	if cp.parent != nil {

		cp.parent.updateTransform(cp)
	} else {

		var dc = logic.NewIdentityMatrix4f()
		dc.Translate(&logic.Vector3f{-1, 1, 0})
		dc = dc.Mul(logic.NewScaleMatrix4f(2.0 * size.X, -2.0 * size.Y, 1.0, 1.0))

		cp.setTransform(dc)
	}
}

func (cp *Panel) Draw(timeDelta int64) {

}

/////////////////////////////////ContentPane//////////////////////////////

type ContentPane struct {
	Panel

	components []Component
	layout Layout
}

func (cp *ContentPane) Add(component Component) {

	cp.components = append(cp.components, component)

	cp.updateTransform(component)
}

func (cp *ContentPane) SetLayout(layout Layout) {

	cp.layout = layout
	layout.setContentPane(cp)

	for _, e := range cp.components {

		cp.updateTransform(e)
	}
}

func (cp *ContentPane) updateTransform(component Component) {

	var current = logic.NewIdentityMatrix4f()
	if cp.layout != nil {
		current.Translate(cp.layout.GetOffset(component).Vector3f(0.0))
	}
	current = current.Mul(logic.NewScaleMatrix4f(component.GetSize().X, component.GetSize().Y, 1.0, 1.0))

	component.setTransform(cp.transform.Mul(current))
}

func (cp *ContentPane) setTransform(transform *logic.Matrix4f) {

	cp.transform = transform

	for _, e := range cp.components {

		cp.updateTransform(e)
	}
}

func (cp *ContentPane) Draw(timeDelta int64) {

	for _, e := range cp.components {

		e.Draw(timeDelta)
	}
}
