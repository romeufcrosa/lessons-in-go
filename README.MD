# Lessons in Go(lang)

## Intro

Lessons in Go(lang) was a project I started to convince my friend and colleague Bruno on the virtues of the Go programming
language.

I wanted to convince him to have a look and I remembered how I had enjoyed the *Ruby Koans* when I was first introduced to
it. So I decided to challenge myself to create a similar approach using Go.

All created tests are failing, it is your mission to discover what causes it, change it, run the test again and see if
you've learned your lesson.

You will find many lessons in this repository, I will go more in detail into each one and how to run then below but first
thank you for checking this out, hopefully you find this at least entertaining if not useful.

All the best on your Go voyage ;)

## How to use them

In each lesson I will include the command that needs to be executed in the terminal to run that specific lesson.
Of course you can do it directly via VS Code (or another editor).

The output will be something like this:
![screenshot](sample_screenshot.png)

I'm considering if I add hints or not, I guess it might depend on each test.

I'm thinking in the future of using a Makefile or having main as an entrypoint so the lessons can be run from the binary,
but for now please follow the individual instructions.

## Lessons

### Lesson One: Basics

**Status**: In Progress

```
go test -run ^TestLesson1 lessons-in-go/lessons
```

This lessons focus on teaching you the basic building blocks of the Go language:
- Declaring functions and variables
- Type Assertion
- Arrays and Slices
- For vs Range
- 

To pass the lesson you will have to now how to declare funcs and variables, primitive types, what is different between
arrays and slices as well as what is different between for and range.

### Lesson Two: TBD