package one

type World []*Entity

type Entity []*Component
func (e *Entity) Add(c *Component)
func (e *Entity) Event(ev *Event)
func (e *Entity) Draw()
func (e *Entity) Has(c *Component) bool
func (e *Entity) Kill()

type Component interface {
	OnAdd(e *Entity)
	onEvent()
	OnRemove()
}

type Event interface {
	Action(e *Entity)
}

type EntityStore []*Entity
