package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ricardoraposo/blogs-api-go/internal/database"
	"github.com/ricardoraposo/blogs-api-go/internal/entity"
	"github.com/ricardoraposo/blogs-api-go/internal/utils"
)

type BlogPostHandler struct {
	blogPostDB     database.BlogPostDBInterface
	postCategoryDB database.PostCategoryDBInterface
}

type CreateBlogPostDTO struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	UserID     string   `json:"user_id"`
	Categories []string `json:"categories"`
}

func NewBlogPostHandler(blogPostDB database.BlogPostDBInterface, postCategoryDB database.PostCategoryDBInterface) *BlogPostHandler {
	return &BlogPostHandler{blogPostDB, postCategoryDB}
}

func (h *BlogPostHandler) CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	var p CreateBlogPostDTO
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.WriteToJson(w, map[string]string{"error": "invalid request"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	blogPost := entity.NewBlogPost(p.Title, p.Content, p.UserID)

	post, err := h.blogPostDB.CreateBlogPost(blogPost)
	if err != nil {
		utils.WriteToJson(w, map[string]string{"error": "could not create post"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, categoryID := range p.Categories {
		err := h.postCategoryDB.CreatePostCategory(post.ID, categoryID)
		if err != nil {
			log.Println(err)
			utils.WriteToJson(w, map[string]string{"error": "could not create post category"})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	utils.WriteToJson(w, post)
	w.WriteHeader(http.StatusCreated)
}

func (h *BlogPostHandler) GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.blogPostDB.GetBlogPosts()
	if err != nil {
		utils.WriteToJson(w, map[string]string{"error": "could not get posts"})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.WriteToJson(w, posts)
	w.WriteHeader(http.StatusOK)
}

func (h *BlogPostHandler) GetBlogPostByID(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	post, err := h.blogPostDB.GetBlogPostByID(postID)
	if err != nil {
		utils.WriteToJson(w, map[string]string{"error": "could not fin post"})
		w.WriteHeader(http.StatusNotFound)
		return
	}

	utils.WriteToJson(w, post)
	w.WriteHeader(http.StatusOK)
}
