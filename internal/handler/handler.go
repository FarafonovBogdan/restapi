package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func GetPost(w http.ResponseWriter, r *http.Request) {
	var idParm string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParm)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid"))
		return
	}
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return
	}

	post := posts[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {

	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var idParm string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParm)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return

	}
	var updatePost Post
	json.NewDecoder(r.Body).Decode(&updatePost)

	posts[id] = updatePost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatePost)

}

func PatchPost(w http.ResponseWriter, r *http.Request) {
	var idParm string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParm)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return

	}
	post := &posts[id]
	json.NewDecoder(r.Body).Decode(post)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified ID"))
		return
	}
	posts = append(posts[:id], posts[id+1:]...)
	w.WriteHeader(200)
}
