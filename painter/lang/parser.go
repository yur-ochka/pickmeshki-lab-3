package lang

import (
	"bufio"
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"

	"github.com/yur-ochka/pickmeshki-lab-3/painter"
)

type Parser struct {
	lastBgColor painter.Operation
	lastBgRect  *painter.BgRectangle
	figures     []*painter.Figure
	moveOps     []painter.Operation
	updateOp    painter.Operation
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		err := p.parseLine(line)
		if err != nil {
			return nil, err
		}
	}

	var res []painter.Operation
	if p.lastBgColor != nil {
		res = append(res, p.lastBgColor)
	}
	if p.lastBgRect != nil {
		res = append(res, p.lastBgRect)
	}
	if len(p.moveOps) > 0 {
		res = append(res, p.moveOps...)
	}
	for _, f := range p.figures {
		res = append(res, f)
	}
	if p.updateOp != nil {
		res = append(res, p.updateOp)
	}
	return res, nil
}

func (p *Parser) parseLine(line string) error {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return nil
	}
	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "white":
		p.lastBgColor = painter.OperationFunc(painter.WhiteFill)
	case "green":
		p.lastBgColor = painter.OperationFunc(painter.GreenFill)
	case "reset":
		p.lastBgColor = painter.OperationFunc(painter.WhiteFill)
		p.lastBgRect = nil
		p.updateOp = nil
		p.figures = nil
		p.moveOps = nil
	case "update":
		p.updateOp = painter.UpdateOp
	case "bgrect":
		if len(args) != 4 {
			return fmt.Errorf("bgrect expects 4 arguments")
		}
		vals, err := parseInts(args)
		if err != nil {
			return err
		}
		p.lastBgRect = &painter.BgRectangle{
			X1: vals[0], Y1: vals[1], X2: vals[2], Y2: vals[3],
		}
	case "figure":
		if len(args) != 2 {
			return fmt.Errorf("figure expects 2 arguments")
		}
		vals, err := parseInts(args)
		if err != nil {
			return err
		}
		p.figures = append(p.figures, &painter.Figure{
			X: vals[0], Y: vals[1], C: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		})
	case "move":
		if len(args) != 2 {
			return fmt.Errorf("move expects 2 arguments")
		}
		vals, err := parseInts(args)
		if err != nil {
			return err
		}
		move := &painter.Move{X: vals[0], Y: vals[1], Figures: p.figures}
		p.moveOps = append(p.moveOps, move)
	default:
		return fmt.Errorf("unknown command: %s", cmd)
	}
	return nil
}

func parseInts(args []string) ([]float64, error) {
	vals := make([]float64, len(args))
	for i, a := range args {
		v, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float argument: %v", a)
		}
		vals[i] = v
	}
	return vals, nil
}
