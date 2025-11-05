package infrastructure

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/hw2-labyrinths/internal/domain"
)

func PrintHelp() {
	fmt.Println(`Usage: maze-app [-hV] [COMMAND]
Maze generator and solver CLI application.
  -h, --help      Show this help message and exit.
  -V, --version   Print version information and exit.
Commands:
  generate  Generate a maze with specified algorithm and dimensions.
  solve     Solve a maze with specified algorithm and points.`)
}

func HandleGenerate(args []string) {
	fs := flag.NewFlagSet("generate", flag.ExitOnError)
	algorithm := fs.StringP("algorithm", "a", "", "Algorithm for generation (dfs, prim, kruskal)")
	width := fs.IntP("width", "w", 0, "Width of the maze")
	height := fs.IntP("height", "h", 0, "Height of the maze")
	output := fs.StringP("output", "o", "", "Output file name [optional]")
	unicodeFlag := fs.BoolP("unicode", "u", false, "Image of the maze with unicode")

	err := fs.Parse(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *width == 0 || *height == 0 {
		fs.PrintDefaults()
		return
	}

	var gen application.Generator
	switch *algorithm {
	case "dfs":
		gen = application.NewDFSGen()
	case "prim":
		gen = application.NewPrimGen()
	case "kruskal":
		gen = application.NewKruskalGen()
	default:
		fs.PrintDefaults()
		return
	}

	maze, err := gen.Generate(*width, *height)

	var mazeDraw string
	if *unicodeFlag {
		mazeDraw = maze.DrawMazeUnicode()
	} else {
		mazeDraw = maze.DrawMazeASCII()
	}

	if *output != "" {
		err := os.WriteFile(*output, []byte(mazeDraw), 0644)
		if err != nil {
			fmt.Println("file write error:", err)
			return
		}
		fmt.Println("maze was saved:", *output)
	} else {
		fmt.Println(mazeDraw)
	}
}

func HandleSolve(args []string) {
	fs := flag.NewFlagSet("solve", flag.ExitOnError)
	algorithm := fs.String("algorithm", "", "Algorithm for solving (astar, dijkstra, bellman-ford")
	file := fs.String("file", "", "Input maze file")
	start := fs.String("start", "", "Start point (x,y)")
	end := fs.String("end", "", "End point (x,y)")
	output := fs.String("output", "", "Output file name [optional]")
	unicodeFlag := fs.Bool("unicode", false, "Image of the maze with unicode")

	err := fs.Parse(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *file == "" || *start == "" || *end == "" {
		fs.PrintDefaults()
		return
	}

	var solver application.Solver
	switch *algorithm {
	case "astar":
		solver = application.NewAStarSolver()
	case "dijkstra":
		solver = application.NewDijkstraSolver()
	case "bellman-ford":
		solver = application.NewBellmanFordSolver()
	default:
		fs.PrintDefaults()
		return
	}

	startPoint, err := parsePoint(*start)
	if err != nil {
		fmt.Println(err)
		return
	}

	endPoint, err := parsePoint(*end)
	if err != nil {
		fmt.Println(err)
		return
	}

	srcMaze, err := loadMazeFromFile(*file)
	if err != nil {
		fmt.Println(err)
		return
	}

	maze, err := solver.Solve(startPoint, endPoint, srcMaze)
	if err != nil {
		fmt.Println(err)
		return
	}

	var mazeDraw string
	if *unicodeFlag {
		mazeDraw = maze.DrawMazeUnicode()
	} else {
		mazeDraw = maze.DrawMazeASCII()
	}

	if *output != "" {
		err := os.WriteFile(*output, []byte(mazeDraw), 0644)
		if err != nil {
			fmt.Println("file write error:", err)
			return
		}
		fmt.Println("solution was saved:", *output)
	} else {
		fmt.Println(mazeDraw)
	}
}

func parsePoint(s string) (domain.Point, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return domain.Point{}, fmt.Errorf("invalid point format: %s, expected format: x,y", s)
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return domain.Point{}, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return domain.Point{}, err
	}

	return domain.Point{X: x, Y: y}, nil
}

func loadMazeFromFile(filename string) (*domain.Maze, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) == 0 {
		return nil, errors.New("empty file")
	}

	height := len(lines)
	width := len([]rune(lines[0]))
	maze, err := domain.NewMaze(width-2, height-2)
	if err != nil {
		return nil, err
	}

	for y, line := range lines {
		runes := []rune(line)
		for x, char := range runes {
			switch char {
			case ' ', '\u2B1B': // ⬛
				maze.Cells[x][y] = domain.Space
			case '#', '\U0001F9F1': // 🧱
				maze.Cells[x][y] = domain.Wall
			case '$', '\U0001FA99': // 🪙
				maze.Cells[x][y] = domain.Coin
			case '+', '\U0001F335': // 🌵
				maze.Cells[x][y] = domain.Tree
			case '@', '\U0001FAA8': // 🪨
				maze.Cells[x][y] = domain.Rock
			default:
				return nil, errors.New("invalid character in maze from file")
			}
		}
	}
	return maze, nil
}
