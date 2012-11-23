package gocos2d

import ()

type (
	Node_ interface {
		Tag() string
		ZOrder() int
		SetParent(Node_)
		OnEnter()
		OnExit()
		Cleanup()
		Draw()
		Visit()
		Transform()
		Schedule(string, func(), int)
		Unschedule(string)
		RunAction(Action_)
		StopAction(Action_)
		AddChild(Node_)
		RemoveChild(uint)
		ConvertTo()
	}
	Node struct {
		tag      string
		Position Position
		Z        int
		Camera   Camera
		Anchor   Anchor
		Parent   Node_
		Children *ChildList
	}
)

func (this *Node) Tag() string {
	return this.tag
}
func (this *Node) ZOrder() int {
	return this.Z
}
func (this *Node) SetParent(parent Node_) {
	this.Parent = parent
}

func (this *Node) Init(tag string) *Node {
	this = new(Node)
	this.tag = tag
	list := new(ChildList).Init()
	this.Children = list

	return this
}
func (this *Node) Cleanup() {
	for e := this.Children.Front(); e != nil; e = e.Next() {
		e.Value.Cleanup()
	}
	this = nil
}
func (this *Node) AddChild(child Node_) {
	child.SetParent(this)
	_ = this.Children.PushFront(child)
}
func (this *Node) RemoveChild(id uint) {

}
func (this *Node) Draw() {

}
func (this *Node) Visit() {

}
func (this *Node) Transform() {

}
func (this *Node) Schedule(name string, cb func(), interval int) {

}
func (this *Node) Unschedule(name string) {

}
func (this *Node) RunAction(act Action_) {

}
func (this *Node) StopAction(act Action_) {

}
func (this *Node) ConvertTo() {

}
func (this *Node) OnEnter() {

}
func (this *Node) OnExit() {

}