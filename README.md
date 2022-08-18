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

