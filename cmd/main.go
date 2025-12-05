package main

import (
	"github.com/drhythm/ekit/concurrent_list"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 1. Define DB Model
type ListItem struct {
	gorm.Model
	Value string
}

func main() {
	// Connect to Database
	db, err := gorm.Open(sqlite.Open("list.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate.
	db.AutoMigrate(&ListItem{})

	// Initialize Memory List.
	list := concurrent_list.NewConcurrentArrayList[string](10)

	// 5. [Warm-up] Load data from DB to Memory on startup.
	var items []ListItem
	db.Find(&items)
	for _, item := range items {
		list.Append(item.Value)
	}
	log.Printf("🔥 System Started: Loaded %d items from database!", len(items))

	// Setup Web Server.
	r := gin.Default()

	// API 1: Add item (Persist to DB + Cache in Memory)
	r.GET("/add", func(c *gin.Context) {
		val := c.Query("val")
		if val == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "value is required"})
			return
		}

		// Step A: Save to Database (Hard Drive)
		db.Create(&ListItem{Value: val})

		// Step B: Append to Memory List (Fast Access)
		list.Append(val)

		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"added":  val,
			"total":  list.Len(), // Now this count matches DB count
		})
	})

	// API 2: Get item (Read from Memory directly - Super Fast!)
	r.GET("/get", func(c *gin.Context) {
		indexStr := c.Query("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index"})
			return
		}

		// Read from ConcurrentArrayList
		val, err := list.Get(index)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"index": index,
			"value": val,
			"source": "memory_cache", // Tell user this came from RAM
		})
	})

	r.Run(":8080")
}