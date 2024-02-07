package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

var (
	MYSQL_HOST     = os.Getenv("MYSQL_HOST")
	MYSQL_USER     = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	MYSQL_DB       = os.Getenv("MYSQL_DB")
	MYSQL_PORT     = os.Getenv("MYSQL_PORT")
	JWT_SECRET     = os.Getenv("JWT_SECRET")
)

type keyServerContextDB struct{}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func useMethod(method string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	})
}

func createJWT(username string, secret string, authz bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().UTC().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().UTC().Unix(),
		"admin":    authz,
	})
	return token.SignedString([]byte(secret))
}

// HANDLERS
func login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	ctx := r.Context()
	var db *sql.DB = ctx.Value(keyServerContextDB{}).(*sql.DB)

	var user User

	row := db.QueryRow("SELECT email, password FROM user WHERE email = ?", username)

	if err := row.Scan(&user.Email, &user.Password); err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if username != user.Email || password != user.Password {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	jwt, err := createJWT(username, JWT_SECRET, true)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(jwt))
}

func validate(w http.ResponseWriter, r *http.Request) {
	reqToken := r.Header.Get("Authorization")

	if reqToken == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	reqToken = strings.Split(reqToken, " ")[1]

	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(claims)
		return
	}

	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DB))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/login", useMethod(http.MethodPost, login))
	mux.HandleFunc("/validate", useMethod(http.MethodPost, validate))

	ctx, cancelCtx := context.WithCancel(context.Background())

	server := &http.Server{
		Addr:         ":5000",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.WithValue(ctx, keyServerContextDB{}, db)
			return ctx
		},
	}

	err = server.ListenAndServe()

	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}

	cancelCtx()
	os.Exit(1)
}
