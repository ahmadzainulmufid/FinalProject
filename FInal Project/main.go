package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
    db  *gorm.DB
    err error
)

type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"-"`
    Profile  string `json:"profile"`
}

func main() {
    // Inisialisasi database SQLite (bisa diganti sesuai kebutuhan)
    db, err = gorm.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    db.AutoMigrate(&User{})

    r := gin.Default()

    // Endpoint untuk registrasi pengguna
    r.POST("/users/register", registerUser)

    // Middleware untuk otentikasi
    r.Use(authMiddleware)

    // Endpoint untuk mengupload foto profil
    r.POST("/users/upload-profile", uploadProfile)

    // Endpoint untuk menghapus foto profil
    r.DELETE("/users/delete-profile/:id", deleteProfile)

    r.Run(":8080") // Ganti port sesuai kebutuhan Anda
}

// Handler untuk registrasi pengguna
func registerUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Simpan data pengguna ke database
    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// Middleware otentikasi
func authMiddleware(c *gin.Context) {
    // Dapatkan token atau sesi otentikasi dari header atau cookie
    // Lakukan verifikasi sesuai dengan mekanisme autentikasi yang digunakan (misalnya JWT)

    // Jika otentikasi berhasil, set data pengguna di context
    // Jika otentikasi gagal, berikan respons dengan status Unauthorized (401)
}

// Handler untuk mengupload foto profil
func uploadProfile(c *gin.Context) {
    user := getUserFromContext(c)
    if user == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Lakukan logika upload gambar profil
    // Simpan gambar ke direktori yang sesuai
    // Perbarui kolom "Profile" dalam data pengguna di database

    c.JSON(http.StatusOK, gin.H{"message": "Profile picture uploaded successfully"})
}

// Handler untuk menghapus foto profil
func deleteProfile(c *gin.Context) {
    user := getUserFromContext(c)
    if user == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Dapatkan ID gambar yang akan dihapus dari parameter URL
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
        return
    }

    // Lakukan verifikasi bahwa gambar dengan ID tersebut dimiliki oleh pengguna yang sedang login
    // Hapus gambar dari direktori dan kolom "Profile" dalam data pengguna di database

    c.JSON(http.StatusOK, gin.H{"message": "Profile picture deleted successfully"})
}

// Fungsi utilitas untuk mendapatkan data pengguna dari context
func getUserFromContext(c *gin.Context) *User {
    // Dapatkan data pengguna dari otentikasi yang telah dilakukan
    // Anda dapat mengimplementasikan cara yang sesuai dengan mekanisme autentikasi yang digunakan (misalnya JWT)

    // Contoh sederhana:
    // Ambil ID pengguna dari sesi atau token
    userID := c.Get("userID")
    if userID != nil {
        var user User
        db.First(&user, userID)
        return &user
    }

    return nil
}
