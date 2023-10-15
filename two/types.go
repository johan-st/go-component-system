package two

type System interface {
	Update(w *World)
  }
  
  type World []*Entity
  type Entity []*Component
  type Component *any

  
  // examples
  
//   type SystemRender struct {
// 	renderer *imported.Rederer
//   }
//   func (r *SystemRender) Run(w *World){
// 	// range over entities and draw all entries that has the corresponding Component to the screen.
//   }
  
//   type ComponentSprite struct {
// 	Texture *imported.Texture
// 	Pos Vector
// 	Size Vector 
// 	Rotation float32
//   }
  
  