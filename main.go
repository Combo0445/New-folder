package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:nvarchar(100);unique" json:"name"`
}

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	ImageURL  string    `gorm:"type:nvarchar(max)" json:"image_url"` // เก็บ URL รูปจากเน็ต
	CreatedAt time.Time `json:"created_at"`
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Message   string    `gorm:"type:nvarchar(max)" json:"message"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

var DB *gorm.DB

func initDatabase() {
	dsn := "sqlserver://@localhost:1433?database=Testcase&integrated+security=true&encrypt=disable&TrustServerCertificate=true"
	var err error
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection failed: ", err)
	}
	DB.AutoMigrate(&User{}, &Post{}, &Comment{})
}

func main() {
	initDatabase()
	app := fiber.New()
	app.Use(cors.New())

	// เพิ่ม User (ต้องมี ID: 1 ก่อนถึงจะคอมเมนต์ได้)
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)
		c.BodyParser(user)
		DB.Create(&user)
		return c.JSON(user)
	})

	// ดึงข้อมูล Feed ทั้งหมด
	app.Get("/posts", func(c *fiber.Ctx) error {
		var posts []Post
		DB.Preload("User").Preload("Comments.User").Order("created_at desc").Find(&posts)
		return c.JSON(posts)
	})

	// บันทึกโพสต์ใหม่ (ส่ง image_url เป็นลิงก์รูปจากเน็ต)
	app.Post("/posts", func(c *fiber.Ctx) error {
		post := new(Post)
		c.BodyParser(post)
		post.CreatedAt = time.Now()
		DB.Create(&post)
		return c.JSON(post)
	})

	// ส่งคอมเมนต์
	app.Post("/comments", func(c *fiber.Ctx) error {
		comment := new(Comment)
		c.BodyParser(comment)
		DB.Create(&comment)
		return c.JSON(comment)
	})

	log.Fatal(app.Listen(":3000"))
}