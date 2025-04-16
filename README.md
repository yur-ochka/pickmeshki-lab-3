# Pickmeshki Lab 3

This project implements a simple graphics server that accepts commands via HTTP and renders them on a canvas. The server listens on port 17000 for drawing commands.

## Features

- HTTP-based command interface (port 17000)
- Support for basic drawing commands:
  - `white` - Clear canvas with white color
  - `green` - Set color to green
  - `figure x y` - Create a figure at coordinates (x,y)
  - `move x y` - Move all figures to absolute coordinates (x,y)
  - `update` - Apply all pending changes
  - `bgrect x1 y1 x2 y2` - Draw a background rectangle

## Components

- `cmd/painter` - Main server application
- `painter` - Core graphics and command processing logic
- `painter/lang` - Command parsing and validation
- `ui` - User interface components
- `scripts` - Example scripts demonstrating different drawing patterns

## Example Scripts

The `scripts` directory contains several example scripts that demonstrate different drawing patterns:

- `black_square` - Draws a black square with figures at different positions (corners and center)
- `bouncing_figure` - Creates a figure that bounces around the screen boundaries
- `rotating_figure` - Shows a figure rotating in a circular pattern around the center
- `moving_figure` - Demonstrates a figure moving in a square pattern
- `green_frame` - Creates a green frame with figures inside

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/yur-ochka/pickmeshki-lab-3.git
cd pickmeshki-lab-3
```

2. Run the server:

```bash
go run ./.cmd/painter
```

3. Run example scripts:

```bash
go run scripts/black_square/main.go
go run scripts/bouncing_figure/main.go
go run scripts/rotating_figure/main.go
go run scripts/moving_figure/main.go
go run scripts/green_frame/main.go
```

## Dependencies

- Go 1.24 or later
- Standard Go libraries

## Project Structure

```
pickmeshki-lab-3/
├── cmd/
│   └── painter/         # Main server application
├── painter/             # Core graphics logic
│   ├── lang/           # Command parsing
│   └── op.go           # Operation definitions
├── ui/                  # User interface components
├── scripts/            # Example scripts
└── go.mod              # Go module definition
```

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is licensed under the MIT License - see the LICENSE file for details.
