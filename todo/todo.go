package todo

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"project/filefunc"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func GetWd() string {
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return mydir + "/todo.txt"
}

func Init() {

	var intro = lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#3d4a71")).
		Width(70).
		PaddingBottom(1).
		PaddingTop(1).
		PaddingLeft(1).
		PaddingRight(1).
		Align(lipgloss.Center).MarginLeft(60).
		Render("Todo List:")

	var styleMenu = lipgloss.NewStyle().
		Background(lipgloss.Color("#e70e70")).
		Width(70).Align(lipgloss.Center).MarginLeft(60).PaddingTop(1).PaddingBottom(1).
		Render("1.Create Todo    2.Delete Todo   3.Esc")

	var todoStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#171e3a")).PaddingLeft(20).PaddingBottom(2).
		Width(70).MarginLeft(60)

	for {
		fmt.Println(intro)
		fmt.Println(styleMenu)
		fmt.Println(todoStyle.Render(DisplayTodo()))

		fmt.Print(anotherStyle.Render("\nEnter Your Choice:\n"))
		fmt.Print("\n\t\t\t\t\t\t\t  ->")
		Scanln()

	}
}

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	return line
}
func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var anotherStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#26619c")).
	Background(lipgloss.Color("#9a405e")).
	BorderRight(true).
	BorderBottom(true).
	BorderTop(true).
	BorderLeft(true).MarginLeft(80)

func Scanln() {

	line := scan()
	switch line {
	case "1":
		Clear()
		fmt.Println(anotherStyle.Render("Enter Your Todo:"))
		fmt.Print("\n\t\t\t\t\t\t\t  ->")
		lines := scan()
		filefunc.WriteLine(GetWd(), lines)
		Clear()

	case "2":
		Clear()
		fmt.Println(anotherStyle.Render("Enter number to delete Todo:"))
		fmt.Print("\n\t\t\t\t\t\t\t  ->")
		num := scan()
		todos := filefunc.StoreInSlice(GetWd())
		val, err := strconv.Atoi(num)
		if err != nil {
			panic(err.Error())
		}
		filefunc.DeleteTodo(todos[val], GetWd())
		Clear()

	case "3":
		Clear()
		fmt.Println("Bye Bye!")
		os.Exit(0)

	}

}

func DisplayTodo() string {

	var s string

	todos := filefunc.StoreInSlice("./todo.txt")
	for i, todo := range todos {
		s += fmt.Sprintf("\n[%d] %s", i, todo)
	}
	return s
}
