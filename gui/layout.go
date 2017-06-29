package gui

import "github.com/schottm/gllib/logic"

type Layout interface {

	GetOffset(component Component) *logic.Vector2f
	setContentPane(pane Pane)
}

type DefaultLayout struct {

	pane     Pane
	position map[Component]*logic.Vector2f
}

func NewDefaultLayout() *DefaultLayout {

	return &DefaultLayout{nil, map[Component]*logic.Vector2f{}}
}

func (df *DefaultLayout) GetOffset(component Component) *logic.Vector2f {

	if val, ok := df.position[component]; ok {

		return val
	}

	return &logic.Vector2f{}
}

func (df *DefaultLayout) AddComponent(component Component, offset *logic.Vector2f) {

	df.position[component] = offset

	if df.pane != nil {

		df.pane.updateTransform(component)
	}
}

func (df *DefaultLayout) setContentPane(pane Pane) {

	df.pane = pane
}
