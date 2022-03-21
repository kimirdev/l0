package server

import (
	"errors"
	"l0/cache"
	"l0/db"
	"l0/models"
	"net/http"

	_ "l0/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	repo       *db.Repository
	localCache *cache.LocalCache
}

func NewHandler(repo *db.Repository, localCache *cache.LocalCache) *Handler {
	return &Handler{
		repo:       repo,
		localCache: localCache,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.GET("/", h.getAllOrders)
			orders.GET("/:id", h.getOrderById)
		}
	}
	return router
}

// @Summary GetAllOrders
// @Tags orders
// @Description Get all orders
// @ID get-all-orders
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Router /api/orders [get]
func (h *Handler) getAllOrders(c *gin.Context) {
	orders := h.localCache.GetAll()

	c.JSON(http.StatusOK, orders)
}

// @Summary GetOrderById
// @Tags orders
// @Description Get order by id
// @ID get-order-by-id
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} models.OrderDTO
// @Router /api/orders/{id} [get]
func (h *Handler) getOrderById(c *gin.Context) {
	order, exist := h.localCache.Get(c.Param("id"))

	if !exist {
		c.AbortWithStatusJSON(400, "error")
		return
	}
	//val, _ := json.Marshal(order)
	c.JSON(http.StatusOK, order)
}

func (h *Handler) Create(data models.OrderDTO) error {
	err := h.repo.Order.Create(data)

	if err != nil {
		return err
	}

	succ := h.localCache.Set(data)

	if !succ {
		return errors.New("failed to add model in cache")
	}

	return nil
}
