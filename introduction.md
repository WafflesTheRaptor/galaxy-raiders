# Introduction {#chap:introduction}

## Space Invaders

Developing software is a creative process where you have to see a problem from multiple perspectives.

The key question is how do I _model_ a problem?  And the technical word "_model_" is hard to define; unless you have had a lot of experience writing software.  So, let's take an example to figure out why something simple might have different _models_.

Take the example of a Rectangle.

![Sketch of rectangle on squared paper]()

We all know what a rectangle is.  It has four sides, each side is connected to another in a $90^\circ$ angle.

The description of a rectangle as having four sides connected in $90^\circ$ angles is a good description for mathematicians, but not a great description for a computer.  The problem with our description is that it does not tell us how to _create_ a rectangle.  And computers need explanations that allow them to create things.  There are two parts to the description of creation, one part is the `model`.

A _model_ of a rectangle must provide all the information needed so that we can _construct_ the rectangle.  We can _model_ a rectangle in many ways, all of which using a rectangular coordinate system:

* as a sequence of line segments
  - line segment from `(x, y)` to `(x, y)` ...
* as a list of `(x, y)` points inside the rectangle
  - `(), (), (), (),\ldots` this description gets quite long, so I've provided only the first few points in the list
* as the centre point of the rectangle, the rectangle's width and the rectangles height
  - `(x, y)`, width of `w` and height of `h`
* the top-left corner of the rectangle and the bottom-right corner of the rectangle
  - `(x, y)` and `(x, y)`

I'm going to choose to represent a rectangle by modeling it as the top-left and bottom-right corner of a rectangle.  I'm choosing this _model_ because I know that this is the same model used in the **Go** programming language representation of a rectangle.  But we will see later that our game engine represents rectangles using a different model where it uses the top-left corner point in addition to the width and height of the rectangle.

This brings me back to my original point about seeing a problem from multiple perspectives.  We need to _model_ a rectangle from the perspective of the Go programming language.  We will need to convert that into the perspective of a rectangle used by our game engine.  We will deal with these different _models_ in @chapter:problem.

In our discussion of _model_ we ...

## Go Programming Language

I've chosen to use the **Go** programming language in this book.  Go is developed by Google and is available for download from [go.dev](https://go.dev).  It was designed to be a very efficient programming language that is easy to learn.  In particular Go excels when writing large-scale networked software systems.  Our Space Invaders project is not a large-scale networked system but we will find that Go provides tools that makes programming just a little bit easier.

```yaml
---
# It is worth mentioning that programming languages are tools; Go is not a "better" programming language than Python or Java, however there are situations where each of these programming languages excel.  Software development professionals are usually confident in a variety of programming languages.  Therefore, it's worth learning Go 
```

We will use some features of the Go programming language.

* selection (`if` statements)
* repetition (`for` loops)
* data types (records, data structures and arrays)
* boolean logic
* functions

## Ebiten v2 Game Engine

## Chapter Outline

1. Breaking the Problem Down
   * Create a 640x480 screen
   * Put a box on the screen
     * Create:
     * Draw: (screen layout)
     * Update: ()
2. Represent a Player
   * Create: Represent the Player as a Rectangle (type)
   * Draw: Translate the Player to the correct (Vector + translate)
   * Update: The Player on each frame
3. Represent an Alien
   * Crate:
   * Draw:
   * Update:
4. Represent our Bullets
   * Crate:
   * Draw:
   * Update:
5. Detect Collisions between Aliens and Bullets
   * Update:
6. Scoreboards, Sprites & End Game
   * Adding a Scoreboards
   * Draw Images instead of Rectangles
   * Add an End Game screen
