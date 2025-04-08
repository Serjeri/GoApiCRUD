package handlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"goApiTask/database/query"
	"goApiTask/models"
	"strconv"
)

type Client struct {
	repository *query.Repository
}

func NewClient(repository *query.Repository) *Client {
	return &Client{repository: repository}
}
func (client *Client) GetPage(fib *fiber.Ctx) error {
	return fib.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tasks retrieved successfully",
	})
}

func (client *Client) Create(fib *fiber.Ctx) error {
	body := new(models.Create)

	err := fib.BodyParser(body)
	if err != nil {
		fib.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}

	result, err := client.repository.CreateTask(context.TODO(), body.Title, body.Description)
	if err != nil {
		return fib.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Failed to create task: %v", err),
			"data":    nil,
		})
	}

	return fib.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Задача была добавлена под № %d", result),
		"data": fiber.Map{
			"id":          result,
			"title":       body.Title,
			"description": body.Description,
			"status":      "new",
		},
	})
}

func (client *Client) GetTasks(fib *fiber.Ctx) error {
	tasks, err := client.repository.GetAllTasks(context.TODO())
	if err != nil {
		fmt.Printf("Error getting tasks: %v\n", err)

		return fib.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve tasks",
			"error":   err.Error(),
			"data":    nil,
		})
	}

	return fib.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Tasks retrieved successfully",
		"data":    tasks,
	})
}

func (client *Client) Update(fib *fiber.Ctx) error {
	idStr := fib.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fib.Status(400).SendString("Некорректный ID")
	}

	product := new(models.Tasks)

	if err := fib.BodyParser(product); err != nil {
		return fib.Status(400).SendString("Неверный формат запроса")
	}

	_, err = client.repository.UpdateTask(context.TODO(), id, product.Title, product.Description)
	if err != nil {
		return fib.Status(500).SendString("Ошибка обновления таски")
	}
	return fib.SendString("Таска успешно обновлена")
}

func (client *Client) Delete(fib *fiber.Ctx) error {
	idStr := fib.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fib.Status(400).SendString("Некорректный ID")
	}

	result, _ := client.repository.DeleteTask(context.TODO(), id)
	if result == false {
		return fib.Status(500).SendString("Ошибка удаления таски: ID не указан")
	}

	return fib.SendString("Такска успешно удолена")
}
