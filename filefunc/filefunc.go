package filefunc

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
)

// var pl = fmt.Println

func Prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Todo: ")
	text, _ := reader.ReadString('\n')

	return (text)
}

func InsertTodo() {
	for {
		todo := Prompt()
		if todo == "0\n" {
			fmt.Println("Bye Bye")
			return
		}

		WriteLine("./test.txt", todo)
	}

}

func FileExist() {
	if _, err := os.Stat("todo.txt"); err != nil {
		os.Create("todo.txt")
	}
}

// check if file is empty or not
func CheckFile(filename string) bool {
	FileExist()
	file, err := os.Stat(filename)

	fsize := file.Size()

	if err != nil {
		panic(err.Error())
		// os.Create("todo.txt")

	}
	if fsize == 0 {
		return true
	} else {
		return false
	}

}

// Delete lines of text stored linearly
func DeleteTodo(todo string, filename string) {

	FileExist()

	lines := StoreInSlice(filename)
	nlines := []string{}

	for _, val := range lines {
		if val != todo {

			nlines = append(nlines, val)
		}
	}

	// fmt.Println(nlines)

	err := os.Truncate(filename, 0)
	if err != nil {
		panic(err)
	}

	for _, val := range nlines {
		val += "\n"
		Append(filename, val)
	}

}

// store text linearly in a slice
func StoreInSlice(filename string) []string {

	FileExist()

	lines := []string{}
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines

}

func CreateFile(fileName string, data string) {

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	_, writeErr := file.WriteString(data)
	if writeErr != nil {
		fmt.Println(err.Error())
		return
	}

}

func WriteBytes(fileName string, data []byte) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()
	size, writeErr := file.Write(data)
	if writeErr != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Wrote %d bytes to file", size)
}

func WriteByLine(filename string, lines []string) {
	file, err := os.Create(filename)
	if err != nil {
		panic("Error")
	}
	defer file.Close()

	for _, val := range lines {
		_, err := fmt.Print(file, "%s\n", val)
		if err != nil {
			panic("Error")
		}
	}

}

// file is empty then writes , else appends
func WriteLine(filename string, line string) {

	if CheckFile(filename) {
		fline := []byte((line))
		err := os.WriteFile(filename, fline, 0644)

		if err != nil {
			panic("Error")
		}
		return
	}

	Append(filename, "\n"+line)

}

// file is appended
func Append(filename, data string) {
	file, err := os.OpenFile(filename,
		os.O_APPEND|os.O_WRONLY,
		fs.ModeAppend)

	if err != nil {
		panic("Error")
	}
	defer file.Close()

	_, fErr := fmt.Fprint(file, data)
	if fErr != nil {
		fmt.Println(err.Error())
		return
	}
}
