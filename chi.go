package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func chimain() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/articles", func(r chi.Router) {
		r.With(paginate).Get("/", listArticles)                           // GET /articles
		r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017

		r.Post("/", createArticle)       // POST /articles
		r.Get("/search", searchArticles) // GET /articles/search

		// Subrouters:
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleCtx)
			r.Get("/", getArticle)       // GET /articles/123
			r.Put("/", updateArticle)    // PUT /articles/123
			r.Delete("/", deleteArticle) // DELETE /articles/123
		})
	})

	// Mount the admin sub-router
	r.Mount("/admin", adminRouter())

	http.ListenAndServe(":3333", r)
}

func ArticleCtx(handler http.Handler) http.Handler {
	panic("implement me")
}

func searchArticles(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

type Article string

func getArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(*Article)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(*Article)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(*Article)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))
}

func listArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%s", "article")))
}
func listArticlesByDate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%s", "article")))
}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	article, ok := ctx.Value("article").(*Article)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))
}

// A completely separate router for administrator routes
func adminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(AdminOnly)
	r.Get("/", adminIndex)
	r.Get("/accounts", adminListAccounts)
	return r
}

func adminIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:1")))
}
func adminListAccounts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:1")))
}

type YourPermissionType struct {
	Permission string `json:"permission"`
}

func (t YourPermissionType) IsAdmin() bool {
	return false
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		perm, ok := ctx.Value("acl.permission").(YourPermissionType)
		if !ok || !perm.IsAdmin() {
			http.Error(w, http.StatusText(403), 403)
			return
		}
		next.ServeHTTP(w, r)
	})
}
