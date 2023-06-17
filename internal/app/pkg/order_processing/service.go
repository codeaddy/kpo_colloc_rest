package order_processing

import (
	"colloc_rest/internal/app/pkg/order"
	"colloc_rest/internal/app/pkg/product"
	"colloc_rest/internal/app/pkg/product_order"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderProcessingService struct {
	ProductService      *product.Service
	ProductOrderService *product_order.Service
	OrderService        *order.Service
}

func NewService(productService *product.Service, productOrderService *product_order.Service, orderService *order.Service) *OrderProcessingService {
	return &OrderProcessingService{ProductService: productService, ProductOrderService: productOrderService, OrderService: orderService}
}

type createOrderInput struct {
	UserID   int               `json:"user_id"`
	Products []product.Product `json:"products"`
	Status   string            `json:"status"`
}

type createOrderResponse struct {
	Success string `json:"success"`
}

func (s *OrderProcessingService) CreateOrder(c *gin.Context) {
	var input createOrderInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error when unmarshalling input"})
		return
	}

	if input.Status != "pending" && input.Status != "processing" && input.Status != "finished" && input.Status != "cancelled" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid status of order (available: pending, processing, finished, cancelled)"})
		return
	}

	ok := true
	for _, product := range input.Products {
		stockProduct, err := s.ProductService.GetById(c, product.ID)
		if err != nil {
			ok = false
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "there is no such dish named: " + product.Name})
			break
		}
		if product.Quantity > stockProduct.Quantity {
			ok = false
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "there is no such amount of product: " + product.Name + " at the stock"})
			break
		}
	}
	if !ok {
		return
	}

	orderID, err := s.OrderService.Create(c, order.OrderRow{
		UserID: input.UserID,
		Status: input.Status,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok = true
	for _, product := range input.Products {
		_, err := s.ProductOrderService.Create(c, product_order.ProductOrder{
			ProductID: product.ID,
			OrderID:   orderID,
			Quantity:  product.Quantity,
		})
		if err != nil {
			ok = false
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			break
		}
	}
	if !ok {
		return
	}

	c.IndentedJSON(http.StatusOK, createOrderResponse{Success: "order successfully created, id = " + strconv.Itoa(orderID)})
}

type getCartByUserIdInput struct {
	UserID int `json:"user_id"`
}

func (s *OrderProcessingService) GetCartByUserId(c *gin.Context) {
	var input getCartByUserIdInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error when unmarshalling input"})
		return
	}

	orders, err := s.OrderService.GetAllByUserId(c, input.UserID)
	if err != nil || len(orders) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "there are no cart for such user_id"})
	}

	cart := orders[len(orders)-1]
	c.IndentedJSON(http.StatusOK, cart)
}

func (s *OrderProcessingService) GetProducts(c *gin.Context) {
	products, err := s.ProductService.GetAll(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.IndentedJSON(http.StatusOK, products)
}

type addProductInput struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int    `db:"price"`
	Quantity    int    `db:"quantity"`
}

func (s *OrderProcessingService) AddProduct(c *gin.Context) {
	var input addProductInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error when unmarshalling input"})
		return
	}

	id, err := s.ProductService.Create(c, product.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.IndentedJSON(http.StatusOK, id)
}