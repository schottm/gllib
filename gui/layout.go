package gui

import "github.com/schottm/gllib/logic"

type Layout interface {

	GetOffset(component Component) *logic.Vector2f
}

type DefaultLayout struct {

	position map[Component]*logic.Vector2f
}

func NewDefaultLayout() *DefaultLayout {

	return &DefaultLayout{map[Component]*logic.Vector2f{}}
}

func (df *DefaultLayout) GetOffset(component Component) *logic.Vector2f {

	return df.position[component]
}

func (df *DefaultLayout) AddComponent(component Component, offset *logic.Vector2f) {

	df.position[component] = offset
}