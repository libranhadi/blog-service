package controller

import (
	"blog-service/helper"
	"blog-service/model"
	"blog-service/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PostController struct {
	service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{
		service: service,
	}
}

func (c *PostController) GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	post, err := c.service.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := helper.WebResponse{
		Code:   200,
		Status: "OK!",
		Data:   post,
	}

	json.NewEncoder(w).Encode(response)
}

func (c *PostController) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	posts, err := c.service.GetAllPosts(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalItems := len(posts)

	totalPages := 0

	if totalItems > 0 {
		totalPages = totalItems / limit
	}

	response := helper.PaginationResponse{
		TotalItems: totalItems,
		TotalPages: totalPages,
		Limit:      limit,
		Offset:     offset,
		Posts:      posts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := c.service.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Post created successfully")
}

func (c *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var post model.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	post.Id = id
	if err := c.service.UpdatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Post updated successfully")
}

func (c *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	post, err := c.service.GetPostByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := c.service.DeletePost(post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Post deleted successfully")
}

func (c *PostController) GetPostByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	post, err := c.service.GetPostByStatus(vars["status"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := helper.WebResponse{
		Code:   200,
		Status: "OK!",
		Data:   post,
	}

	json.NewEncoder(w).Encode(response)
}
