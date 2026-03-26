# 🐾 IT 08-1: Pet Feed Project

โปรเจคระบบ Feed รูปภาพสัตว์เลี้ยง พัฒนาด้วย **Go (Fiber + GORM)** และ **SQL Server** พร้อมเชื่อมต่อกับ Frontend (Angular)

## 🛠️ Tech Stack
- **Backend:** Go (Golang)
- **Framework:** [Fiber v2](https://gofiber.io/)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** Microsoft SQL Server
- **Frontend:** Angular (Standalone Component)

## 🚀 การติดตั้งและใช้งาน (Getting Started)

### 1. ตั้งค่า Database
สร้าง Database ใน SQL Server ชื่อว่า `Testcase` และแก้ไข DSN ในไฟล์ `main.go` ให้ตรงกับเครื่องของคุณ:
```go
dsn := "sqlserver://@localhost:1433?database=Testcase&integrated+security=true&encrypt=disable&TrustServerCertificate=true"
