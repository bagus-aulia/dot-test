package borrow

import (
	"net/http"

	"github.com/bagus-aulia/dot-test/app/helpers"
	"github.com/bagus-aulia/dot-test/app/pkg/models"
	"github.com/bagus-aulia/dot-test/app/pkg/response"
	borrow_usecase "github.com/bagus-aulia/dot-test/app/pkg/usecase/borrow"
	echo "github.com/labstack/echo/v4"
)

// Borrow represent the httphandler for borrow
type Borrow struct {
	Usecase borrow_usecase.Borrow
}

// NewBorrowHandler will initialize the borrow/ resources endpoint
func NewBorrowHandler(e *echo.Echo, ucase borrow_usecase.Borrow) {
	handler := &Borrow{
		Usecase: ucase,
	}

	e.GET("/borrows", handler.Borrows)
	e.POST("/borrows", handler.CreateBorrow)
	e.GET("/borrows/:uuid", handler.BorrowByUUID)
	e.PATCH("/borrows/:uuid", handler.ReturnBorrow)
}

// Borrows will fetch the borrow list based on given params
func (h *Borrow) Borrows(c echo.Context) error {
	ctx := c.Request().Context()

	borrows, err := h.Usecase.GetBorrowList(ctx)
	if err != nil {
		result := response.BorrowsBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BorrowsBody{
		Message: "SUCCESS",
		Data:    h.transformBorrows(borrows),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// BorrowByUUID will fetch the borrow data based on given params
func (h *Borrow) BorrowByUUID(c echo.Context) error {
	uuid := c.Param("uuid")
	ctx := c.Request().Context()

	borrow, err := h.Usecase.GetBorrowByUUID(ctx, uuid)
	if err != nil {
		result := response.BorrowBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BorrowBody{
		Message: "SUCCESS",
		Data:    h.transformBorrow(&borrow),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// CreateBorrow will store the borrow data based on given params
func (h *Borrow) CreateBorrow(c echo.Context) error {
	userUUID := c.FormValue("user_uuid")
	form, err := c.MultipartForm()
	if err != nil {
		result := response.BorrowsBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	bookUUID := form.Value["book_uuids[]"]
	ctx := c.Request().Context()

	borrow, err := h.Usecase.CreateBorrow(ctx, userUUID, bookUUID)
	if err != nil {
		result := response.BorrowBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BorrowBody{
		Message: "SUCCESS",
		Data:    h.transformBorrow(borrow),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

// ReturnBorrow will finish the borrow season based on given params
func (h *Borrow) ReturnBorrow(c echo.Context) error {
	uuid := c.Param("uuid")
	ctx := c.Request().Context()

	borrow, err := h.Usecase.ReturnBorrow(ctx, uuid)
	if err != nil {
		result := response.BorrowBody{
			Message: "FAILED",
			Data:    nil,
			Error: &response.Error{
				Message: err.Error(),
			},
		}
		return c.JSON(helpers.GetStatusCode(err), result)
	}

	result := response.BorrowBody{
		Message: "SUCCESS",
		Data:    h.transformBorrow(&borrow),
		Error:   nil,
	}

	return c.JSON(http.StatusOK, result)
}

func (h *Borrow) transformBorrow(borrow *models.Borrow) *response.Borrow {
	if borrow == nil {
		return nil
	}

	res := &response.Borrow{
		UUID:         borrow.UUID,
		StartAt:      borrow.StartAt,
		EndAt:        borrow.EndAt,
		User:         h.transformUser(borrow.User),
		BorrowDetail: h.transformBorrowDetails(borrow.BorrowDetail),
	}

	return res
}

func (h *Borrow) transformBorrows(borrows []models.Borrow) []*response.Borrow {
	res := make([]*response.Borrow, 0)

	for _, borrow := range borrows {
		borrowRes := h.transformBorrow(&borrow)
		res = append(res, borrowRes)
	}

	return res
}

func (h *Borrow) transformUser(user *models.User) *response.User {
	if user == nil {
		return nil
	}

	var address string = ""
	if user.Address != nil {
		address = *user.Address
	}

	res := &response.User{
		UUID:    user.UUID,
		Name:    user.Name,
		Address: address,
	}

	return res
}

func (h *Borrow) transformBorrowDetail(detail *models.BorrowDetail) *response.BorrowDetail {
	if detail == nil {
		return nil
	}

	res := &response.BorrowDetail{
		Book: h.transformBook(detail.Book),
	}

	return res
}

func (h *Borrow) transformBorrowDetails(details []models.BorrowDetail) []*response.BorrowDetail {
	res := make([]*response.BorrowDetail, 0)

	for _, detail := range details {
		detailRes := h.transformBorrowDetail(&detail)
		res = append(res, detailRes)
	}

	return res
}

func (h *Borrow) transformBook(book *models.Book) *response.Book {
	if book == nil {
		return nil
	}

	var writer string = ""
	if book.Writer != nil {
		writer = *book.Writer
	}

	res := &response.Book{
		Title:    book.Title,
		Writer:   writer,
		UUID:     book.UUID,
		Borrowed: book.Borrowed,
	}

	return res
}
