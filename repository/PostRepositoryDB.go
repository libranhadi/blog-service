package repository

import (
	"blog-service/model"
	"database/sql"
	"errors"
)

type PostRepositoryDB struct {
	db *sql.DB
}

func NewPostRepositoryDB(db *sql.DB) *PostRepositoryDB {
	return &PostRepositoryDB{
		db: db,
	}
}

func (r *PostRepositoryDB) FindByID(id int) (*model.Post, error) {
	var post model.Post
	row := r.db.QueryRow(`SELECT * FROM posts WHERE id = ?`, id)
	if err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Category, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepositoryDB) FindByStatus(status string) (*model.Post, error) {
	var post model.Post
	row := r.db.QueryRow(`SELECT * FROM posts WHERE status = ?`, status)
	if err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Category, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepositoryDB) FindAll(page, pageSize int) ([]model.Post, error) {
	// Calculate the offset
	offset := (page - 1) * pageSize

	// Get the posts on the current page
	rows, err := r.db.Query(`SELECT * FROM posts ORDER BY created_date DESC LIMIT ? OFFSET ?`, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Category, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepositoryDB) Save(post *model.Post) error {
	stmt, err := r.db.Prepare(`INSERT INTO posts (title, content, category, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.Category, post.CreatedAt, post.UpdatedAt)
	return err
}

func (r *PostRepositoryDB) Update(post *model.Post) error {
	stmt, err := r.db.Prepare(`UPDATE posts SET title = ?, content = ?, category = ?, updated_date = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.Category, post.UpdatedAt, post.Id)
	return err
}

func (r *PostRepositoryDB) Delete(post *model.Post) error {
	_, err := r.db.Exec(`DELETE FROM posts WHERE id = ?`, post.Id)
	return err
}
