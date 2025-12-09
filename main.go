package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:81@localhost:5432/todo?sslmode=disable"
	}

	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("openDB error: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)
	ctx := context.Background()

	fmt.Println("Начинаем проверку всех функций...")
	fmt.Println("=====================================")

	// 0. Выведем все задачи
	fmt.Println("\n0. Текущие задачи:")
	allTasks, err := repo.ListTasks(ctx)
	if err != nil {
		log.Printf("ListTasks error: %v", err)
	} else {
		if len(allTasks) == 0 {
			fmt.Println("   ℹЗадач нет")
		} else {
			fmt.Printf("   Всего задач: %d\n", len(allTasks))
			for _, t := range allTasks {
				status := "-"
				if t.Done {
					status = "+"
				}
				fmt.Printf("   %s #%d: %s\n", status, t.ID, t.Title)
			}
		}
	}
	// 1. Создаем несколько начальных задач
	fmt.Println("\n1. Создаем начальные задачи:")
	initialTitles := []string{"Изучить Go", "Сделать домашку", "Прочитать документацию"}
	for _, title := range initialTitles {
		id, err := repo.CreateTask(ctx, title)
		if err != nil {
			log.Printf("CreateTask error: %v", err)
		} else {
			fmt.Printf("   Создана задача #%d: %s\n", id, title)
		}
	}

	// 2. Проверяем ListDone - невыполненные задачи
	fmt.Println("\n2. Список невыполненных задач:")
	undoneTasks, err := repo.ListDone(ctx, false)
	if err != nil {
		log.Printf("ListDone error: %v", err)
	} else {
		if len(undoneTasks) == 0 {
			fmt.Println("   ℹ Невыполненных задач нет")
		} else {
			for _, t := range undoneTasks {
				fmt.Printf("   #%d: %s (создана: %s)\n",
					t.ID, t.Title, t.CreatedAt.Format("01.02.2003 04:05"))
			}
		}
	}

	// 3. Проверяем FindByID - поиск по ID
	fmt.Println("\n3. Поиск задачи по ID:")
	taskID := 1 // ищем первую задачу
	task, err := repo.FindByID(ctx, taskID)
	if err != nil {
		log.Printf("FindByID error: %v", err)
	} else if task != nil {
		fmt.Printf("   Найдена задача #%d:\n", taskID)
		fmt.Printf("      Заголовок: %s\n", task.Title)
		fmt.Printf("      Выполнена: %v\n", task.Done)
		fmt.Printf("      Создана: %s\n", task.CreatedAt.Format("01.02.2003 04:05"))
	} else {
		fmt.Printf("   Задача #%d не найдена\n", taskID)
	}

	// 4. Проверяем CreateMany - массовая вставка
	fmt.Println("\n4.  Массовое создание задач:")
	batchTitles := []string{"Задача из группы 1", "Задача из группы 2", "Задача из группы 3"}
	err = repo.CreateMany(ctx, batchTitles)
	if err != nil {
		log.Printf("CreateMany error: %v", err)
		fmt.Println("    Ошибка при массовой вставке")
	} else {
		fmt.Printf("    Успешно добавлено %d задач:\n", len(batchTitles))
		for i, title := range batchTitles {
			fmt.Printf("      %d. %s\n", i+1, title)
		}
	}

	fmt.Println("\n=====================================")
	fmt.Println("Проверка завершена!")
}
