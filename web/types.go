package web

import (
	"reflect"
)

// type Store struct {
// 	sessions []*Session

// }

// func NewEntityManager() *Store {
// 	return &Store{
// 		live:      []*Entity{},
// 		dead:      []*uint,
// 	}
// }

// func (em *Store) NewEntity() *Entity {
// 	var e *Entity

// 	// reuse entity if available else create new
// 	if len(em.pool) > 0 {
// 		e, em.pool = em.pool[len(em.pool)-1], em.pool[:len(em.pool)-1]
// 	} else {
// 		e = &Entity{id: em.currentID}
// 		em.currentID++
// 	}
// 	return e
// }

// func (em *Store) ReleaseEntity(e *Entity) {
// 	em.pool = append(em.pool, e)
// }
// func (s *Store) Add(e *Session) uint32 {
// 	if len(s.Dead) > 0 {
// 		id := s.Dead[0]
// 		s.Dead = s.Dead[1:]
// 		s.Live[id] = e
// 		return id
// 	}
// }

// func (s *Store) Remove(id uint32) bool
// func (s *Store) Get(id uint32) *Session

type Session struct {
	Live       bool
	components []*Component
}

func (s *Session) Add(c *Component) { 
	(*s).components = append((*s).components, c)
}
func (s Session) Has(c Component) bool {
	for _, sc := range s {
		if reflect.TypeOf(sc) == reflect.TypeOf(c) {
			return true
		}
	}
	return false
}
func (s *Session) Remove(c Component) bool {
	for _, sc := range *s {
		if reflect.TypeOf(sc) == reflect.TypeOf(c) {
			(*sc).Active = false
			return true
		}
	}
	return false
}
func (s *Session) Event(ev Event) {
	for _, sc := range *s {
		ev(sc)
	}

}

type Component struct {
	OnAdd    func(e *Session)
	OnRemove func()
}

type Event func(c *Component)
