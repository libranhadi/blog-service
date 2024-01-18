package helper

import (
	"blog-service/model"
	"net/http"
	"strconv"
)

func ParsePaginationQueryParams(r *http.Request) (int, int) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return 0, 0
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		return 0, 0
	}

	return page, pageSize
}

type PaginationResponse struct {
	TotalItems int          `json:"total_items"`
	Limit      int          `json:"limit"`
	TotalPages int          `json:"total_pages"`
	Offset     int          `json:"offset"`
	Posts      []model.Post `json:"posts"`
}
