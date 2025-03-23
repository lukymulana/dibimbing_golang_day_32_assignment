package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"strconv"
)

// Item represents an inventory item
type Item struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Stock int    `json:"stock"`
}

var items = []Item{
    {ID: 1, Name: "Laptop", Stock: 10},
    {ID: 2, Name: "Mouse", Stock: 20},
}

// GetItems returns all items
func GetItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

// AddItem adds a new item
func AddItem(c *gin.Context) {
    var newItem Item
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newItem.ID = len(items) + 1
    items = append(items, newItem)
    c.JSON(http.StatusCreated, newItem)
}

// UpdateItem updates an existing item
func UpdateItem(c *gin.Context) {
    id := c.Param("id")
    var updatedItem Item
    if err := c.ShouldBindJSON(&updatedItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    for i, item := range items {
        if item.ID == atoi(id) {
            items[i] = updatedItem
            c.JSON(http.StatusOK, updatedItem)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// DeleteItem deletes an item
func DeleteItem(c *gin.Context) {
    id := c.Param("id")
    for i, item := range items {
        if item.ID == atoi(id) {
            items = append(items[:i], items[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// atoi converts string to int
func atoi(s string) int {
    n, _ := strconv.Atoi(s)
    return n
}

func main() {
    r := gin.Default()

    // Load HTML templates
    r.LoadHTMLGlob("templates/*")

    // Serve static files
    r.Static("/static", "./static")

    // Routes for API
    r.GET("/items", GetItems)
    r.POST("/items", AddItem)
    r.PUT("/items/:id", UpdateItem)
    r.DELETE("/items/:id", DeleteItem)

    // Route for web interface
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "items": items,
        })
    })

    // Start server
    r.Run(":8080")
}
