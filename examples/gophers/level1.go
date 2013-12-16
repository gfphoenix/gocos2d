package main

import "gocos2d.org/sdk"

type Level struct {
	gocos2d.Scene
}

func NewLevel() *Level {
	l := &Level{gocos2d.NewScene("level")}
	l.AddLayer(gocos2d.NewLayer("global", 1))
	return l
}

func (l *Level) AddChild(layer string, n gocos2d.Node) {
	l.GetChild(layer).AddChild(n.Tag(), n)
}

func (l *Level) Update() error {
	return nil
}
func (l *Level) Draw() error {
	return nil
}
func (l *Level) OnEnter() error {
	return nil
}
func (l *Level) OnExit() error {
	return nil
}
