package state

import "fmt"

// 上下文对象
type document struct {
	Role  string
	State dState
}

func (d *document) doThis() {
	d.State.doThis()
}

func (d *document) changeState(state dState) {
	d.State = state
}

// 状态接口
type dState interface {
	doThis()
}

// 真实状态 草稿
type draftState struct {
	CurrentContext *document
}

func (d *draftState) SetContext(document *document) {
	d.CurrentContext = document
}

func (d *draftState) doThis() {
	fmt.Println("draft do this,state to moderation")
	newState := &moderationState{CurrentContext: d.CurrentContext}
	d.CurrentContext.changeState(newState)
}

// 真实状态 审核
type moderationState struct {
	CurrentContext *document
}

func (m *moderationState) SetContext(document *document) {
	m.CurrentContext = document
}

func (m *moderationState) doThis() {
	if m.CurrentContext.Role == "admin" {
		fmt.Println("moderation do this,state to published")
		newState := &publishedState{CurrentContext: m.CurrentContext}
		m.CurrentContext.changeState(newState)
	} else {
		fmt.Println("not permissions")
	}
}

// 真实状态 发布
type publishedState struct {
	CurrentContext *document
}

func (p *publishedState) SetContext(document *document) {
	p.CurrentContext = document
}

func (p *publishedState) doThis() {
	fmt.Println("publishedState do this")
}

// 客户端代码

func RunApplication() {

	initState := &draftState{}

	d := &document{
		Role: "admin",
	}
	initState.SetContext(d)
	d.State = initState

	d.doThis()
	d.doThis()
	d.doThis()
}
