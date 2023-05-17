# Breaking the Problem Down {#chap:problem}

<!-- test:suite=problem -->

A _value_ is something like a number or a sequence of written characters.  The number $42$ is a value, so si the sequence of characters "The quick brown fox".  In programming values are drawn from sets definde by the programming languges.  For example, the numbe $42$ comes from the set of integers.  As we use integers very often, we will shorten it to `int`.  Therefore the value $42$ is an `int`.  Strangly, we name the set of all sequence of characters to be the set of `string`s.  So, "The quick brown fox" is a `string`.  We call the sets from which `value`s are drawn `type`s.  Programmers will talk about $42$ being of `type` `int`.

A _variable_ is a way to name a value.  We declare a vaiable using the `:=` syntax.  As an example, `x := 42` creates a variable named `x` which holds the `int` value `42`.  The Go system is clever enough to know that when we write `x := 42` that `42` is an `int`.  A variable is called a variable because it can be re-assigned.  The following code declares `x` to be a variable holding the value `42`.  The next line re-assigns the variable `x` to hold the value `-2`.  Note that we declare a variable using the `:=` syntax, but re-assign it using the `=` syntax.

```go
x := 42
x = -1
```

The Go programming language provides a cool way to combine `types`.  Often we want to to represent a player's name and their score.  Their name is of type `string` as it is a squence of characters.  The score is of type `int`.  We would like to combine these into a new type that we can use.  The cool way to combine `types` is called a `struct`:

TODO: more explanation here

```go
type Player struct {
	Score int
	Name string
}
```

We can create a value of `type` `Player` cunningly called `playerOne`:

```go
playerOne := Player {
	Score 0
	Name "Caitlín"
}
```

A _function_ is a piece of code that performs a specific task.

For example, suppose we want to increase the player's score.  The solution is to add $10$ points to the `Score` componant of `playerOne`.

```go
func increaseScore(p Player) {
	return Player {
		Score: p.Score + 10,
		Name: p.Name
	}
}
```

TODO More explanation....

```go
func (p Player) increaseScoreReciever() {
	return Player {
		Score: p.Score + 10,
		Name: p.Name
	}
}
```

<!--
For example, suppose we want to translate a point $(x, y)$ along hte $x$-axis by 10 units.  The solution to this is to add $10$ to the $x$ value resulting in $(x + 10, y)$.  We can define a reusalbe function that translates a point along the $x$-axis by any number.  In this case we might call the function `translateX`.  Our `translateX` might take any whole number, which we call an integter.  The the function `translateX` will return the translated point.

```go
import "point"

func translateX(p Point, tx int) {
	return Point{
		p.X + tx,
		p.Y
	}
}
```

Now we have a reusable piece of code that will translate any point to any distance along the $x$-axis.  Translating the point `$(6, 5)$ by $5$ units along the $x$-axis becomes

```go
original := Point{X: 6, Y:5}
translated := translateX(original, 5)
```
--> 

In general a function is declared using the `func` keyword.  A function has a name, and it can take multiple _parameters_.  The parameters are a comma separate list inside the `(` and `)` paranthesis.

A function is a _reciever_ funtion if it is associated with a struct.  Programmers say that `increaseScoreReciever` is a reciever of `Player` [^method].

[^method]: Note, other programming lanugages such as C#, Java and Python call _reciever_ functions "methods".  The technical term _reciever_ is mostly used by Go programmers.

<!-- test:file=go.mod -->
```
module example.com/developer/galaxy-raiders
```

We need a place for our program to start. The main function gives a place for the program to start.

<!-- test:file=main.go -->
```go
package main

func main() {
}
```

These commands run in the terminal. The `go mod tidy` installs any modules the program needs. The `go build -o galaxy-raiders` turns the source code into a executable command which can be run in the terminal.

<!-- test:exec -->
```command
go mod tidy
go build -o galaxy-raiders
```	

The `import` statement imports the game engine called `ebiten v2`. The `screenWidth` and `screenHeight` set the size of the window we want to appear.  The `Game struct` command defines a game by saying that a game is anything that has the Update, Draw and Layout functions. The `main` command sets the desired window size and title and then runs the game. Setting the Update command to run every 1/60 of a second and the the draw command 1/60 of a second after it.

<!-- test:file=main.go -->
```go
package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Space Invaders")
	g := &Game{}
    ebiten.RunGame(g)
}
```

`do mod tidy` makes sure the game engine is ready to use by installing any modules the program needs. 

<!-- test:exec -->
```command
go mod tidy
```

The `go build -o space invaders` turns the source code into an executable command which can be run in the terminal.

<!-- test:exec -->
```
go build -o galaxy-raiders
```

Finally, we can now run our game using:

```console
./galaxy-raiders
```
