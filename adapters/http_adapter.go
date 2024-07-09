package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tanasinp/go-inventory-management/core"
	"github.com/tanasinp/go-inventory-management/database"
)

type httpProductHandler struct {
	service core.ProductService
}

func NewHttpProductHandler(service core.ProductService) *httpProductHandler {
	return &httpProductHandler{service: service}
}

func (h *httpProductHandler) CreateSupplierFiber(c *fiber.Ctx) error {
	var supplier database.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.CreateSupplier(&supplier); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(supplier)
}

func (h *httpProductHandler) CreateCategoryFiber(c *fiber.Ctx) error {
	var category database.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.CreateCategory(&category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(category)
}

func (h *httpProductHandler) GetAllSupplierFiber(c *fiber.Ctx) error {
	suppliers, err := h.service.GetAllSupplier()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(suppliers)
}

func (h *httpProductHandler) GetAllCategoryFiber(c *fiber.Ctx) error {
	categories, err := h.service.GetAllCategory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}

func (h *httpProductHandler) CreateProductFiber(c *fiber.Ctx) error {
	var product database.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := h.service.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(product)
}
