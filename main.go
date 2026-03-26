package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// --- Models ---
type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:nvarchar(100);unique" json:"name"`
}

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	ImageURL  string    `gorm:"type:nvarchar(max)" json:"image_url"`
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

	// --- User Endpoints (เพิ่มใหม่เพื่อให้ยิง Postman ได้) ---
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		if err := DB.Create(&user).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "ID หรือ Name ซ้ำในระบบ"})
		}
		return c.Status(201).JSON(user)
	})

	// --- Post Endpoints ---
	app.Get("/posts", func(c *fiber.Ctx) error {
		var posts []Post
		DB.Preload("User").Preload("Comments.User").Order("created_at desc").Find(&posts)
		return c.JSON(posts)
	})

	app.Post("/posts", func(c *fiber.Ctx) error {
		post := new(Post)
		if err := c.BodyParser(post); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		post.CreatedAt = time.Now()
		DB.Create(&post)
		return c.JSON(post)
	})

	// --- Comment Endpoints ---
	app.Post("/comments", func(c *fiber.Ctx) error {
		comment := new(Comment)
		if err := c.BodyParser(comment); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		DB.Create(&comment)
		return c.JSON(comment)
	})

	// --- Delete Endpoint (สำหรับล้างข้อมูลใน SQL Server) ---
	app.Delete("/clear-database", func(c *fiber.Ctx) error {
		DB.Exec("DELETE FROM comments")
		DB.Exec("DELETE FROM posts")
		return c.JSON(fiber.Map{"message": "ลบข้อมูล Post และ Comment ทั้งหมดแล้ว"})
	})

	log.Fatal(app.Listen(":3000"))
}