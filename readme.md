# Generic Design Patterns

This repository aims to provide a generic and explicit approach to the usage of design patterns in Go. In this package, we provide a set of basic generic implementations of various design patterns. These implementations can be easily utilized to incorporate design patterns into your projects in a clear and explicit manner.

This repository is in an experimental state. It was created to generalize certain aspects of the design patterns I use. The idea is to provide a foundation for implementing design patterns. There are some implementations that I find very interesting and others that are less so. I plan to maintain and update this repository as I gain more experience in Go.

Feel free to explore each pattern in detail and understand how they can be applied in your own solutions.

## What are Design Patterns?

Design patterns are proven solutions to recurring design problems in software engineering. They provide a structured approach to solving common challenges, improving code maintainability, flexibility, and scalability. By following established patterns, developers can leverage best practices and create more robust and efficient software solutions.

## Implemented Generic Design Patterns Example

Examples of usage for each implemented generic design pattern can be found in the `example` folder. Feel free to explore each example to understand how the patterns can be applied in your own projects.

"Not available" patterns are not generalized patterns. There are some examples suffixed by "\_wip.go" as an example, but they are not really generic and well implemented, so they are not recommended to use yet.

### Creational Patterns

## 1. Singleton Example

```go
// DBConnection is a struct that represents a database connection
// It should be a singleton
type DBConnection struct {
	Name string
	// others attributes ...
}

// NewDBConnection creates a new instance of DBConnection
func NewDBConnection() *DBConnection {
	return &DBConnection{Name: "Unique instance"}
}

// creation.NewSingleton creates a new singleton instance
var DBConnectionSingleton = creational.NewSingleton(NewDBConnection)

func main() {

	results := make(chan *DBConnection)

	for i := 0; i < 10; i++ {
		go func() {
			// First call will create the instance
			instance := DBConnectionSingleton.GetInstance()
			results <- instance
		}()
	}

	for i := 0; i < 10; i++ {
		instance := <-results
		// all the pointers are the same
		fmt.Printf("%p \n", instance)
	}
}
```

## 2. Factory Usage Example

```go
type Transport interface {
	deliver()
}

type Truck struct{}

func (t *Truck) deliver() {
	println("Delivering by truck")
}

type Ship struct{}

func (s *Ship) deliver() {
	println("Delivering by ship")
}

func GenerateTransportFactory(transportType string) (*creational.Factory[Transport], error) {

	// make function for the factory
	transportMakeFunc := func() Transport {
		switch transportType {
		case "Truck":
			return &Truck{}
		case "Ship":
			return &Ship{}
		default:
			return nil
		}
	}

	// Generate a factory for the given transport type
	factory := creational.NewFactory(transportMakeFunc)

	if factory != nil {
		return factory, nil
	}

	return nil, fmt.Errorf("UNKNOWN TRANSPORT TYPE")
}

func main() {
	// Generate a factory for trucks
	TruckFactory, _ := GenerateTransportFactory("Truck")

	// Make a truck
	truck1 := TruckFactory.Make()
	truck2 := TruckFactory.Make()

	// Deliver by truck
	truck1.deliver()
	truck2.deliver()
}
```

## 3. Functionnal Builder Usage Example

```go
type Car struct {
	Brand  string
	Model  string
	Option string
}

type CarBuilder struct {
	// A builder using a functionnal approach
	// Object is created only when Build() is called
	// store setters as functions to call successively later
	// Building rules should be updated dynamically
	creational.FunctionalBuilder[Car]
}

func (builder *CarBuilder) SetBrand(brand string) *CarBuilder {
	builder.AddAction(func(c *Car) {
		c.Brand = brand
	})
	return builder
}

func (builder *CarBuilder) SetModel(model string) *CarBuilder {
	builder.AddAction(func(c *Car) {
		c.Model = model
	})
	return builder
}

func (builder *CarBuilder) SetOption(option string) *CarBuilder {
	builder.AddAction(func(c *Car) {
		c.Option = option
	})
	return builder
}

func GenerateCarBuilder(carType string) *CarBuilder {

	if carType == "AudiR8" {
		return (&CarBuilder{}).
			SetBrand("Audi").
			SetModel("R8").
			SetOption("V10")
	}
	if carType == "CitroenC3" {
		return (&CarBuilder{}).
			SetBrand("Citroen").
			SetModel("C3")
	}
	return nil
}

func main() {

	audiR8CarBuilder := GenerateCarBuilder("AudiR8")
	citroenC3CarBuilder := GenerateCarBuilder("CitroenC3")

	c1 := audiR8CarBuilder.Build()
	c2 := citroenC3CarBuilder.Build()

	// air conditioning available, update the builder
	citroenC3CarBuilder.SetOption("air conditioning")

	// Next cars will have air conditioning
	c3 := citroenC3CarBuilder.Build()

	// Reset the builder
	citroenC3CarBuilder.Reset()

	// Next car will have no brand, model and option
	c4 := citroenC3CarBuilder.Build()

	fmt.Println(c1) // Output: &{Audi R8 V10}
	fmt.Println(c2) // Output: &{Citroen C3 ""}
	fmt.Println(c3) // Output: &{Citroen C3 air conditioning}
	fmt.Println(c4) // Output: &{"" "" ""}

}

```

## 4. Prototype Usage Example

`Not available`

### Structural Patterns

## 5. Adapter Pattern

`Not available`

## 6. Decorator Pattern

```go
type Shape interface {
	Render() string
	Display() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f, ",
		c.Radius)
}

func (c *Circle) Display() string {
	return fmt.Sprintf("Circle of radius %f, ",
		c.Radius)
}

type ColoredShape struct {
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("has the color %s, ", c.Color)
}

func (c *ColoredShape) Display() string {
	return fmt.Sprintf("has the color %s, ", c.Color)
}

type TransparentShape struct {
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("has %f%% transparency.", t.Transparency*100.0)
}

func (t *TransparentShape) Display() string {
	return fmt.Sprintf("has %f%% transparency.", t.Transparency*100.0)
}

func main() {

    // Decorator execution behavior, will be executed for each wrapped items
	renderHandler := func(s Shape) {
        // You should use any Decorable interface method here
		fmt.Println(s.Render())
	}

	structural.
        // Initialize a new Decorator for Shape and Wrap the Circle struct
		NewDecorator[Shape](&Circle{Radius: 2}).
        // Wrap the ColoredShape struct over the Circle
		Wrap(&ColoredShape{Color: "red"}).
        // Wrap the TransparentShape struct over the ColoredShape
		Wrap(&TransparentShape{Transparency: 0.5}).
        // Setting Shape.Render() implementation to be executed for each wrapped items
		SetExecutionHandler(&renderHandler).
        // Output: Circle of radius 2, has the color red, has 50% transparency.
		Execute()

	err := structural.
        // Initialize a new Decorator for Shape and Wrap the Circle struct
		NewDecorator[Shape](&Circle{Radius: 2}).
        // Trying to execute without setting an execution handler
		Execute()

	if err != nil {
        // Output: You should must set an execution handler before Executing it
		fmt.Println(err)
	}

	// You can use Factory Generator to Generate Custom Decorators Factories and avoid error risks

	displayHandler := func(s Shape) {
		fmt.Println(s.Display())
	}

	coloredTransparentCircleDisplayFactory := structural.GenerateDecoratorFactory[Shape](
        // Decorator execution behavior
		&displayHandler,
        // Initial Decorable struct
		&Circle{Radius: 2},
		&ColoredShape{Color: "red"},
        &TransparentShape{Transparency: 0.5},
        // ... Decorables to be wrapped in order (You can specify order you need in case of state sharing)
	)

	coloredTransparentCircleDisplay := coloredTransparentCircleDisplayFactory.Make()
    // Output: Circle of radius 2, has the color red, has 50% transparency.
	coloredTransparentCircleDisplay.Execute()

}

```

## 7. Bridge Usage Example

```go
type Processor interface {
	structural.Bridgeable
	process() string
}

type Intel struct{}

func (i *Intel) process() string {
	return "Intel Processor"
}

type AMD struct{}

func (a *AMD) process() string {
	return "AMD Processor"
}

type OS interface {
	structural.Bridgeable
	boot() string
}

type WindowsOS struct{}

func (w *WindowsOS) boot() string {
	return "Windows OS"
}

type MacOS struct{}

func (m *MacOS) boot() string {
	return "Mac OS"
}

type Computer interface {
	Compute()
}

type Windows struct {
	Proc structural.Bridge[Processor]
	OS   structural.Bridge[OS]
}

func (w *Windows) Compute() {
	fmt.Println("WINDOWS Compute using : ", (*w.Proc.Impl).process(), " and ", (*w.OS.Impl).boot())
}

type Mac struct {
	Proc structural.Bridge[Processor]
	OS   structural.Bridge[OS]
}

func (m *Mac) Compute() {
	fmt.Println("MAC Compute using : ", (*m.Proc.Impl).process(), " and ", (*m.OS.Impl).boot())
}

func main() {
	windows := &Windows{}
	mac := &Mac{}

	intel := &Intel{}
	amd := &AMD{}

	windowsOS := &WindowsOS{}
	macOS := &MacOS{}

	windows.Proc.SetImpl(intel)
	mac.Proc.SetImpl(amd)

	windows.OS.SetImpl(windowsOS)
	mac.OS.SetImpl(macOS)

	windows.Compute()
	mac.Compute()

}
```

## 8. Composite Usage Example

```go
type Searcher interface {
	search(keyword string)
}

type File struct {
	structural.Composable
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

type Folder struct {
	structural.Composable
	structural.Composite[Searcher]
	name string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.Children {
		composite.search(keyword)
	}
}

func NewFolder(name string) *Folder {
	return &Folder{name: name}
}

func NewFile(name string) *File {
	return &File{name: name}
}

func main() {
	folder1 := NewFolder("Folder 1")
	folder2 := NewFolder("Folder 2")
	file1 := NewFile("File 1")
	file2 := NewFile("File 2")

	folder1.Add(file1)
	folder1.Add(file2)
	folder1.Add(folder2)

	folder1.search("rose")
}
```

## 9. Facade Usage Example

`Not available`

## 10. Flyweight Usage Example

```go

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

func NewRed() *Color {
	return &Color{R: 255, G: 0, B: 0}
}

func NewGreen() *Color {
	return &Color{R: 0, G: 255, B: 0}
}

func NewBlue() *Color {
	return &Color{R: 0, G: 0, B: 255}
}

type Color struct {
	structural.IntrinsicState
	R, G, B uint8
}

type OptimizedShape struct {
	Color *Color
	// ...
}

func NewColorFlyweight() *structural.Flyweight[Color] {

	colorFlyweight := structural.NewFlyweight[Color]()

	colorFlyweight.NewCreationHandler(RED, NewRed)
	colorFlyweight.NewCreationHandler(GREEN, NewGreen)
	colorFlyweight.NewCreationHandler(BLUE, NewBlue)

	return colorFlyweight
}

func main() {
	OptimizedCircleFlyweight := NewColorFlyweight()

	OptimizedRedShape := &OptimizedShape{
		Color: OptimizedCircleFlyweight.GetInstance(RED),
	}

	SecondOptimizedRedShape := &OptimizedShape{
		Color: OptimizedCircleFlyweight.GetInstance(RED),
	}

    // Should print the same address for both shapes
	fmt.Printf("first color address: %p, second color address: %p , Color : %v",
		OptimizedRedShape.Color,
		SecondOptimizedRedShape.Color,
		*OptimizedRedShape.Color,
	)
}

```

## 11. Proxy Usage Example

`Not available`

### Behabioral Patterns

## 12. Chain of Responsibility Usage Example

```go

type UserLoginRequestData struct {
	Username string
	Roles    []string
}

type UserLoginResultData struct {
	IsAuth, IsAdmin bool
}

func (u UserLoginResultData) String() string {
	return fmt.Sprintf(" { IsAuth: %v, IsAdmin: %v }", u.IsAuth, u.IsAdmin)
}

type User struct {
	Broker *behavioral.Broker[UserLoginRequestData, UserLoginResultData]
	Name   string
	Roles  []string
}

func NewUser(name string, roles []string, broker *behavioral.Broker[UserLoginRequestData, UserLoginResultData]) *User {
	return &User{Name: name, Roles: roles, Broker: broker}
}

func (u *User) CanAccess() error {
	q := behavioral.Query[UserLoginRequestData, UserLoginResultData]{
		Data:   UserLoginRequestData{Username: u.Name, Roles: u.Roles},
		Result: UserLoginResultData{IsAuth: false, IsAdmin: false},
		Error:  nil,
	}
	u.Broker.Fire(&q)
	return q.Error
}

type isAuthModifier struct {
	behavioral.Handler[UserLoginRequestData, UserLoginResultData]
}

func (authModifier *isAuthModifier) Handle(q *behavioral.Query[UserLoginRequestData, UserLoginResultData]) {
	for _, role := range q.Data.Roles {
		if role == "user" || role == "admin" {
			q.Result.IsAuth = true
			return
		}
	}

	q.Error = fmt.Errorf("UNAUTHENTICATED")
}

type IsAdminModifier struct {
	behavioral.Handler[UserLoginRequestData, UserLoginResultData]
}

func (adminModifier *IsAdminModifier) Handle(q *behavioral.Query[UserLoginRequestData, UserLoginResultData]) {

	for _, role := range q.Data.Roles {
		if role == "admin" {
			q.Result.IsAdmin = true
			return
		}
	}

	q.Error = fmt.Errorf("UNAUTHORIZED")
}

func AuthenticatedRouteCheck(user User) error {
	accessBroker := behavioral.NewBroker[UserLoginRequestData, UserLoginResultData]()
	accessBroker.Subscribe(&isAuthModifier{})
	user.Broker = accessBroker
	return user.CanAccess()

}

func AdminRouteCheck(user User) error {
	accessBroker := behavioral.NewBroker[UserLoginRequestData, UserLoginResultData]()
	accessBroker.Subscribe(&isAuthModifier{})
	accessBroker.Subscribe(&IsAdminModifier{})
	user.Broker = accessBroker
	return user.CanAccess()
}

func Scenario(user User) {
	err := AuthenticatedRouteCheck(user)
	if err != nil {
		fmt.Printf("User %s try to access to Basic route -- Error : %s \n", user.Name, err)
	} else {
		fmt.Printf("User %s try to access to Basic route -- Success \n", user.Name)
	}

	err = AdminRouteCheck(user)
	if err != nil {
		fmt.Printf("User %s try to access to Admin route -- Error : %s \n", user.Name, err)
	} else {
		fmt.Printf("User %s try to access to Admin route -- Success \n", user.Name)
	}

}

func main() {

	// User Created from a valid JWT Token
	jack := NewUser("Jack", []string{}, nil)
	john := NewUser("John", []string{"user"}, nil)
	jane := NewUser("Jane", []string{"admin"}, nil)

	// Try to access to differents routes
	Scenario(*jack)
	Scenario(*john)
	Scenario(*jane)

}
```

## 13. Command Usage Example

`Not available`

## 14. Iterator Usage Example

```go
type Item struct {
	behavioral.Iterable
	Power int
}

// an iterable collection of items in a list form
type ListCollection struct {
	*behavioral.Iterator[Item]
}

func (c *ListCollection) Push(item *Item) {

	wrappedItem := &behavioral.Iterator[Item]{Item: item}

	if c.Iterator == nil {
		c.Iterator = wrappedItem
	} else {
		c.Extend(wrappedItem)
	}
}

func (c *ListCollection) Pop() *Item {

	if c.Iterator == nil {
		return nil
	}

	penultimate := c.Penultimate()
	penultimate.SetNext(nil)

	return penultimate.Item
}

func NewListCollection(items ...*Item) *ListCollection {
	collection := &ListCollection{}

	for _, item := range items {
		collection.Push(item)
	}

	return collection
}

func main() {

	item1 := &Item{Power: 1}
	item2 := &Item{Power: 2}
	item3 := &Item{Power: 3}

	ItemCollection := NewListCollection(item1, item2, item3)

	ItemCollection.Push(&Item{Power: 4})

	iterate := ItemCollection.InitIterator() // Initialize the iterator

	for it := iterate(); it != nil; it = iterate() { // Iterate over the collection
		println(it.Item.Power) // it contains the current item
	}

	ItemCollection.Pop()

	iterate = ItemCollection.InitIterator()
	for it := iterate(); it != nil; it = iterate() {
		println(it.Item.Power)
	}

}
```

## 15. Mediator Usage Example

```go

type StationManager struct {
	behavioral.Mediator
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type PassengerTrain struct {
	behavioral.Component[StationManager]
}

func (g *PassengerTrain) arrive() {
	if !g.Mediator.canArrive(g) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (g *PassengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.notifyAboutDeparture()
}

func (g *PassengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	g.arrive()
}

type FreightTrain struct {
	behavioral.Component[StationManager]
}

func (g *FreightTrain) arrive() {
	if !g.Mediator.canArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (g *FreightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	g.Mediator.notifyAboutDeparture()
}

func (g *FreightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.arrive()
}

func main() {

	stationManager := newStationManger()

	passengerTrain := &PassengerTrain{}
	passengerTrain.Register(stationManager)

	freightTrain := &FreightTrain{}
	freightTrain.Register(stationManager)

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()

}

```

## 16. Memento Usage Example

```go

type ConcreteState struct {
	Value string
}

func main() {

	caretaker := &behavioral.Caretaker[ConcreteState]{
		MementoArray: make([]*behavioral.Memento[ConcreteState], 0),
	}

	originator := &behavioral.Originator[ConcreteState]{
		State: ConcreteState{"A"},
	}

	fmt.Printf("Originator Current State: %s\n", originator.GetState().Value)
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState(ConcreteState{"B"})
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState(ConcreteState{"C"})
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

}

```

## 17. Observer Usage Example

```go

type DoctorService struct {
	behavioral.Observer
	Name string
}

func (d DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor of %s has been called for %s \n", d.Name, data.(string))
}

type Person struct {
	behavioral.Observable[DoctorService]
	Name string
}

func NewPerson(name string) *Person {

	return &Person{
		Observable: behavioral.NewObservable[DoctorService](),
		Name:       name,
	}
}

func (p *Person) CatchACold() {
	p.Notify(p.Name)
}

func main() {

	p := NewPerson("John")

	ds1 := &DoctorService{Name: "Hospital 1"}
	ds2 := &DoctorService{Name: "Hospital 2"}
	p.Subscribe(ds1)
	p.Subscribe(ds2)

	p.CatchACold()

}

```

## 18. State Usage Example

```go

type Document struct {
	Title   string
	Content string
	// ...
	State *behavioral.StateMachine
}

func (d *Document) Moderate(approuved bool) error {
	if approuved {
		return d.State.GoTo(Approved)
	} else {
		return d.State.GoTo(Rejected)
	}
}

// Define the different possible states of the document
const (
	Draft behavioral.State = iota
	Moderation
	Approved
	Rejected
	Published
	MAX_BUILD_STATUS
)

// Define the mapping between state and their string representation
var DocumentStateToString = map[behavioral.State]string{
	Moderation: "Moderation",
	Approved:   "Approved",
	Rejected:   "Rejected",
	Published:  "Published",
}

// Define the transition rules between states
var DocumentStateTransitionRules = map[behavioral.State][]behavioral.State{
	Draft: {
		Moderation,
	},
	Moderation: {
		Approved, Rejected,
	},
	Rejected: {
		Draft,
	},
	Approved: {
		Published,
	},
}

func NewDocument(title, content string, state behavioral.State) (*Document, error) {

	documentStateMachine := (&behavioral.StateMachineBuilder{}).
		SetCurrentState(state).                            // Set the current state
		SetTransitionRules(&DocumentStateTransitionRules). // Set the transition rules between states
		SetStateToString(&DocumentStateToString).          // Set the mapping between state and string representation
		SetMaxUnreachableState(MAX_BUILD_STATUS).          // Set the maximum unreachable state (prevent invalid state value)
		Build()                                            // We can use dedicated builder Functionnal pattern to create a new StateMachine

	return &Document{
		Title:   title,
		Content: content,
		State:   &documentStateMachine,
	}, nil

}

func main() {

	draftDocument, _ := NewDocument("Draft Document", "This is a draft document", Draft) // You can start from any state

	draftDocument.State.GoTo(Moderation) // Transition from Draft to Moderation

	draftDocument.Moderate(true) // Transition from Moderation to Approved

	fmt.Println(draftDocument.State) // StateMachine implement Stringer Interface String() return currentState as string specified from mapping

	err := draftDocument.State.GoTo(Draft) // Trying to do invalid transition from Approved to Draft

	if err != nil {
		fmt.Println(err) // Output: Transition from Moderation to Published is not allowed
	}

}

```

## 19. Strategy Usage Example

`Not available`

## 20. Template Method Usage Example

`Not available`

## 21. Visitor Usage Example

`Not available`
