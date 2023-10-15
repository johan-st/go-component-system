# go-component-system
exploration of component systems in go




_______________________________________________________

# Ideation
## overview
- Fits well with validation?
- used by unity
- for experimentation, enables change (prototyping)
- system heavy game 
### core strengths
1. Flexibility
2. Reasonability
3. Composability
4. Performance

### notes
- Entries are probably an array.
- find way to minimize allocations?
- component store, entity store, State, dead entities (add in place of old instead of adding)
- create,get,add,remove, has,kill
- next iteration should preferably be next in memory
- archetype version in go. Cool! : https://youtu.be/71RSWVyOMEY?si=qmsnHspe5ZRq35MN
- systems operate on the world ( all entities )
- world holds entities
- entities hold components
- components encapsulates behavior and attributes

### look into
- entity mask - bit mask which tells which component type are on any particular entity.
### validation thought experiment

```go
type Entity struct {
  components []Component
}

func (e *Entity) Add(c Component)
func (e *Entity) Update()
func (e *Entity) Draw()
func (e *Entity) Has(c *Componemt)
func (e *Entity) Kill()

type Component interface {
  OnAdd()
  OnUpdate()
  OnDraw()
  OnKill()
}
```

Alt (hmmm...... Needs work)
```go
type System interface {
  Update(w *World)
}

type World []Entity
type Entity []Component
type Component any{}

// examples

type SystemRender struct {
  renderer *imported.Rederer
}
func (r *SystemRender) Run(w *World){
  // range over entities and draw all entries that has the corresponding Component to the screen.
}

type ComponentSprite struct {
  Texture *imported.Texture
  Pos Vector
  Size Vector 
  Rotation float32
}



```



## Component based e-com
#e-commerce #go #component-based-architecture

Extended with unity like component architecture? Maybe a map of behaviours all with a common API. 

On second thought... dynamically creating new instances might not be optimal? But then again.. in games you do exactly that. They need to be created in a predefined way but this should cover our use-case.

[Reference for architecture](https://youtu.be/5HCdqV4nQkQ?si=JzL7Dlzzgw45TO6j)

```go

// entity holds the value ALL entities plus a slice of components.
type entity struct {
    id uuid.Uuid 
    created time.Time //universally needed?
    components []component
}


type component interface {
    onAction(*Action) error 
}

func (e *entity) onAction(a *Action) error {
    for _, c := range e.components {
       err := c.onAction(a)
       if err != nil {
           return err
           // or maybe run as many actions as possible and collect the errors for later
       }
    }
    return nil
}

func (e *entity) addComponent(newC component) {
    for _, existing := range e.components {
        if reflect.TypeOf(newC) == reflect.TypeOf(existing){
            panic(fmt.Sprintf(
                    "attempted to add components of existing type %v",
                    reflect.TypeOf(newC)
                )
            )
        }
    }
    e.components = append(e.components, newC)
}

func getComponent(c *component) *component {
    for _, existing := range e.components {
        if reflect.TypeOf(c) == reflect.TypeOf(existing){
            return existing
        }
    }
    // you are expected to know what components exist on your entities on compile time. 
    // if you want to handle it differently consider a "getComponentOrNil" function or change the return type of this function to (*component, bool) similar to the standard map implementation.
    panic(fmt.Sprintf(
        "component of type %v does not exist on this entity",
        reflect.TypeOf(c)
        )
}

// component example
		  
type compCartHolder struct {
    cart *Cart
}

func (c *compCartHolder) onAction(a *Action) error {
    switch a.type {
    case ActionType.ItemAddedToCart:
        return c.cart.Add(a.itemId, a.quantity)
    case ActionType.ItemRemovedFromCart:
        return c.cart.Remove(a.itemId, a.quantity)
    case ActionType.CartCleared
        return c.cart.Clear()
    default: 
        return nil
    }
}

// compActionLogger logs all actions.
type compActionLogger struct {
    l *log.Logger
    parentEntity *entity
}

func (c *compActionLogger) onAction(a *Action) error {
    l.Debug("action", 
    "Entity Id", c.parrentEntity.id,
    "ActionType", a.ActionType
    "Action", *a
    )
}
```

Instantiate entity. TODO: Consider how to handle zero values. In component or with constructor?
```go
// instantiate entity

session := newEntity()
if err := session.addComponent(&compCartHolder{}); err {
    // handle error
}
if err = session.addComponent(&compActionLogger{}); err{
    // handle error
}

// on event
session.action(a)
```

Keep a global list of entities in main if you trust your code. Otherwise maybe a protected singleton or similar. Maybe a go routine responding to a channel. 
TODO: Check how this can be accessed by imported packages. 
```go
var (
// in mem sessions
    activeLifetime time.Time // time since last action before moving to stale sessions.
    staleLifetime time.Time // time since last action before forgetting session. After this we should look in the dB.
    
    activeSessions []entity
    staleSessions []entity 
    )


// clean up old seasons with some kind of garbage collector. Js reaper style maybe. Do we need 2 generations? Will lookup be slow if



```