package main

import (
	"bufio"
	"fmt"
	"os"
)

var lastestID int = 0

var tasks = map[int]Task{}

func showMainMenu() {
	fmt.Println("===== CHUONG TRINH QUAN LY CONG VIEC =====")
	fmt.Println("Hay chon cac chuc nang sau:")
	fmt.Println("1. Them cong viec")
	fmt.Println("2. Xem danh sach cong viec")
	fmt.Println("3. Check cong viec")
	fmt.Println("4. Xoa cong viec")
	fmt.Println("5. Thoat chuong trinh")
}

func addTask() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("===== THEM CONG VIEC =====")
		fmt.Print("Nhap ten cong viec: ")

		scanner.Scan()
		title := scanner.Text()

		lastestID++
		task := Task{
			id:        lastestID,
			title:     title,
			completed: false,
		}

		tasks[task.id] = task

		fmt.Println("Da them cong viec thanh cong.")
		fmt.Print("Ban co muon them cong viec khac khong (y/n): ")

		scanner.Scan()
		choice := scanner.Text()
		if choice != "y" {
			break
		}
	}
}

func showTasks() {
	fmt.Println("===== DANH SACH CONG VIEC =====")
	if len(tasks) == 0 {
		fmt.Println("Khong co cong viec nao.")
	} else {
		for _, task := range tasks {
			var completedStatus string
			if task.completed {
				completedStatus = "[X]"
			} else {
				completedStatus = "[ ]"
			}

			fmt.Printf("%d. %s %s\n", task.id, completedStatus, task.title)
		}
	}
	for {
		fmt.Println("====================")
		fmt.Println("1. Home")
		var choice int
		fmt.Println("Nhap lua chon cua ban: ")
		fmt.Scan(&choice)
		if choice == 1 {
			return
		}
	}
}

func checkTask() {
	fmt.Println("===== CHECK CONG VIEC =====")
	if len(tasks) == 0 {
		fmt.Println("Khong co cong viec nao.")
		return
	}
	fmt.Println("Nhap ID cong viec can check: ")
	var id int
	fmt.Scan(&id)
	for i, task := range tasks {
		if i == id {
			task.completed = !task.completed
			tasks[i] = task
			fmt.Println("Da check cong viec thanh cong.")
			return
		}
	}
	fmt.Println("Khong tim thay cong viec voi ID da nhap.")
}

func deleteTask() {
	fmt.Println("Nhap ID cong viec can xoa: ")
	var id int
	fmt.Scan(&id)
	delete(tasks, id)
	fmt.Println("Xoa cong viec thanh cong.")
}

func main() {
	for {
		showMainMenu()
		var choice int
		fmt.Println("Nhap lua chon cua ban: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addTask()
		case 2:
			showTasks()
		case 3:
			checkTask()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("Thoat chuong trinh")
			return
		default:
			fmt.Println("Lua chon khong hop le")
		}
	}
}
