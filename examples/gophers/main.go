//http://mortdeus.mit-license.org/
package main

import "gocos2d.org/sdk"

var dir = gocos2d.Director

func main() {
	dir.Title = "Gophers"
	Init()
	for dir.Running {
		dir.Update()
		dir.Draw()
	}
	dir.Cleanup()

}
func Init() {
	dir.Init()
	hog := NewGroundhog()
	lvl1 := NewLevel()
	lvl1.AddChild("global", hog)
	dir.Push(lvl1)
}
