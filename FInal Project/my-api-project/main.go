package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "time"
)

var (
    db  *gorm.DB
    err error
)

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    Username  string    `json:"username" gorm:"not null"`
    Email     string    `json:"email" gorm:"not null;unique"`
    Password  string    `json:"-" gorm:"not null;size:255"`
    Photos    []Photo   `json:"photos" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Photo struct {
    ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    Title     string    `json:"title"`
    Caption   string    `json:"caption"`
    PhotoURL  string    `json:"photo_url"`
    UserID    uint      `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func main() {
    db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
    db.AutoMigrate(&User{}, &Photo{})

    r := gin.Default()

    r.POST("/users/register", registerUser)

    r.GET("/users/login", loginUser)

    r.PUT("/users/:userId", updateUser)

    r.DELETE("/users/:userId", deleteUser)

    r.Run(":8080")
}

func registerUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func loginUser(c *gin.Context) {

}

func updateUser(c *gin.Context) {
}

func deleteUser(c *gin.Context) {
}
