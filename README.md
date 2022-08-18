# goFractalGenerator
A fractal Generator, viwer, and image maker created in Go

This is the third version of this program. This one is clean enough where I feel comfortable sharing the code.
The goal of this project was to learn Go and to build something cool in the process.

There are currently two implimented algorithms to generate the image data, one is concurrent the other is not. By default the concurrent algorithm is in use for the fractal viwer. The number of threads to use in the fractal generation can be specified using the -threads flag when launching the program from a terminal.

The controls are specified in the game, but to reiterate:
- Move = Arrow Keys
- Zoom in = Z key
- Zoom out = X key
- Increase/decrease pixel test iteration = W/Q keys respectivly
- Increase/decrease color contrast = P/O keys respectivly
- Take a screenshot = space bar
- Increase/decrease screenshot size (in relation to the game screen size) = T/R keys respectivly

Some examples:

![fractal image 1](https://user-images.githubusercontent.com/90068632/185298504-6c5aa8d7-968e-4319-8035-8a03a7e0524b.png)
![fractal image 2](https://user-images.githubusercontent.com/90068632/185299040-28a2fe6e-bafa-4a0e-b053-eb6c6735d763.png)
