package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"crud/models"
)

// CreateOrder godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.CreateOrder true "CreateOrderRequestBody"
// @Success 201 {object} models.Order "GetOrderBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) CreateOrder(c *gin.Context) {
	var order models.CreateOrder

	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Order().Create(context.Background(), &order)
	if err != nil {
		log.Printf("error whiling Create: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling Create").Error())
		return
	}

	resp, err := h.storage.Order().GetByPKey(
		context.Background(),
		&models.OrderPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByIdOrder godoc
// @ID get_by_id_order
// @Router /order/{id} [GET]
// @Summary Get By Id Order
// @Description Get By Id Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Order "GetOrderBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetOrderById(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storage.Order().GetByPKey(
		context.Background(),
		&models.OrderPrimarKey{Id: id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetListOrder godoc
// @ID get_list_order
// @Router /order [GET]
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.GetListOrderResponse "GetOrderBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) GetOrderList(c *gin.Context) {
	var (
		limit  int
		offset int
		err    error
	)

	limitStr := c.Query("limit")
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	offsetStr := c.Query("offset")
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Printf("error whiling limit: %v\n", err)
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	resp, err := h.storage.Order().GetList(
		context.Background(),
		&models.GetListOrderRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		log.Printf("error whiling get list: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling get list").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateOrder godoc
// @ID update_order
// @Router /order/{id} [PUT]
// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param order body models.UpdateOrder true "CreateOrderRequestBody"
// @Success 200 {object} models.Order "GetOrdersBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) UpdateOrder(c *gin.Context) {

	var (
		order models.UpdateOrder
	)

	order.Id = c.Param("id")

	if order.Id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required order id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required order id").Error())
		return
	}

	err := c.ShouldBindJSON(&order)
	if err != nil {
		log.Printf("error whiling update: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.storage.Order().Update(
		context.Background(),
		&order,
	)

	if err != nil {
		log.Printf("error whiling update: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update").Error())
		return
	}

	if rowsAffected == 0 {
		log.Printf("error whiling update rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling update rows affected").Error())
		return
	}

	resp, err := h.storage.Order().GetByPKey(
		context.Background(),
		&models.OrderPrimarKey{Id: order.Id},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteByIdOrder godoc
// @ID delete_by_id_order
// @Router /order/{id} [DELETE]
// @Summary Delete By Id Order
// @Description Delete By Id Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Order "GetOrderBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *HandlerV1) DeleteOrder(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		log.Printf("error whiling update: %v\n", errors.New("required order id").Error())
		c.JSON(http.StatusBadRequest, errors.New("required order id").Error())
		return
	}

	err := h.storage.Order().Delete(
		context.Background(),
		&models.OrderPrimarKey{
			Id: id,
		},
	)

	if err != nil {
		log.Printf("error whiling delete: %v", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling delete").Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
