package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// бесконечно крутимся пока не завершимся
	for {
		fmt.Print(">>> ")
		// читаем юзерский ввод
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// подчищаем лишние пробелы
		input = strings.TrimSpace(input)

		// сеанс поддерживается до тех пор, пока не будет введена команда выхода (\quit)
		if input == `\quit` {
			os.Exit(0)
		}

		// делим строку на аргументы
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}
		// проверка на пустые аргументы
		switch args[0] {
		// cd <args> - смена директории
		case "cd":
			if len(args) > 1 {
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			} else {
				home, err := os.UserHomeDir()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				err = os.Chdir(home)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		// pwd - показать путь до текущего каталога
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Println(dir)
			}
		// echo <args> - вывод аргумента в STDOUT
		case "echo":
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " "))
			}
		// kill <args> - "убить" процесс, переданный в качестве аргумента
		case "kill":
			if len(args) > 1 {
				pid := args[1]
				cmd := exec.Command("kill", pid)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		// ps - выводит общую информацию по запущенным процессам в формате
		case "ps":
			cmd := exec.Command("ps", "aux")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		// ловим непрописанные команды (ошибки), пробрасываем в поток ошибок и выводим на консоль
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
