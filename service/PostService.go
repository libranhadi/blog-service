// service.go
package service

import (
	"blog-service/model"
	"blog-service/repository"
)

type PostService struct {
	repo *repository.PostRepositoryDB
}

func NewPostService(repo *repository.PostRepositoryDB) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) GetPostByID(id int) (*model.Post, error) {
	return s.repo.FindByID(id)
}

func (s *PostService) GetPostByStatus(status string) (*model.Post, error) {
	return s.repo.FindByStatus(status)
}

func (s *PostService) GetAllPosts(page, pageSize int) ([]model.Post, error) {
	return s.repo.FindAll(page, pageSize)
}

func (s *PostService) CreatePost(post *model.Post) error {
	return s.repo.Save(post)
}

func (s *PostService) UpdatePost(post *model.Post) error {
	return s.repo.Update(post)
}

func (s *PostService) DeletePost(post *model.Post) error {
	return s.repo.Delete(post)
}
