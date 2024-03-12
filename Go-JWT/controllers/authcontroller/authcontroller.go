package authcontroller

import (
	"encoding/json"
	"github.com/BerkatPS/go-jwt-testing/config"
	"github.com/BerkatPS/go-jwt-testing/helper"
	"github.com/BerkatPS/go-jwt-testing/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// disini kita akan generate token jwt
	var userInput models.User

	// ambil json dengan json decode
	// Mendecode dari parameter http request
	decoder := json.NewDecoder(r.Body)

	/// jika error maka kita akan mengembalikan error.. harus decode dari variabel model user
	if err := decoder.Decode(&userInput); err != nil {
		// Mendapatkan dari helper writeresponse
		response := map[string]string{"message": err.Error()}
		helper.Writejson(w, http.StatusInternalServerError, response)
		// ketika kita mendapatkan error , maka proses akan berhenti sampai disini.
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan username
	var user models.User
	// mengembil data dari username menggunakan first isinya pointer dari user dan mengecek err
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		// mengecek ketika data tidak ditemukan
		case gorm.ErrRecordNotFound:
			response := map[string]string{
				"message": "Tidak Dapat menemukan User",
			}
			helper.Writejson(w, http.StatusUnauthorized, response)
			return

		default:
			response := map[string]string{
				"message": err.Error(),
			}
			helper.Writejson(w, http.StatusInternalServerError, response)
			return
		}
	}

	// memverifikasi password dengan password sesuai database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		response := map[string]string{
			"message": "Password anda salah!!",
		}
		helper.Writejson(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan jwt token
	// membuat waktu expired
	expTime := time.Now().Add(15 * time.Second)
	claims := config.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "cerdasgizi",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// deklarasi algoritma untuk sign in
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	// pada signed kita memasukan key nya
	signedString, err := withClaims.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{
			"message": err.Error(),
		}
		helper.Writejson(w, http.StatusInternalServerError, response)
		return
	}

	// kita akan langsung set ke cookie tidak ke json

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    signedString,
		Path:     "/",
		HttpOnly: true,
	})
	response := map[string]string{
		"message": "Success Sign in!!!",
	}
	helper.Writejson(w, http.StatusOK, response)
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	// proses registerasi user

	// type struct
	var userInput models.User

	// ambil json dengan json decode
	// Mendecode dari parameter http request
	decoder := json.NewDecoder(r.Body)

	/// jika error maka kita akan mengembalikan error.. harus decode dari variabel model user
	if err := decoder.Decode(&userInput); err != nil {
		// Mendapatkan dari helper writeresponse
		response := map[string]string{"message": err.Error()}
		helper.Writejson(w, http.StatusInternalServerError, response)
		// ketika kita mendapatkan error , maka proses akan berhenti sampai disini.
		return
	}

	defer r.Body.Close()

	// Hash password menggunakan bcrypt memakai generatefrompassowrd
	// parameter pertama menggunakan byte dan parameter ke2 defaultCost
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	// Di generatefrompassword menghasilkan byte maka kita akan menggunakan byte
	userInput.Password = string(hashPassword)

	// Insert Database
	if err := models.DB.Create(&userInput).Error; err != nil {
		log.Fatal("Gagal menyimpan data")
	}
	response := map[string]string{
		"message": "success",
	}
	helper.Writejson(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus token dan logout
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // menghapus cookie
	})
	response := map[string]string{
		"message": "Success Logout",
	}
	helper.Writejson(w, http.StatusOK, response)
	return
}
