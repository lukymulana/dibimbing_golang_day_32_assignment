package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "strconv"
)

// Item represents an inventory item
type Item struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Stock       int    `json:"stock"`
    Description string `json:"description"`
}

var items = []Item{
    {ID: 1, Name: "Laptop", Stock: 10, Description: "High-end laptop"},
    {ID: 2, Name: "Mouse", Stock: 20, Description: "Wireless mouse"},
}

// Handler function for the API
func handler(c *gin.Context) {
    switch c.Request.Method {
    case http.MethodGet:
        if c.Param("action") == "delete" {
            deleteItem(c)
        } else {
            getItems(c)
        }
    case http.MethodPost:
        if c.Param("action") == "update" {
            updateItem(c)
        } else {
            addItem(c)
        }
    default:
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
    }
}

// GetItems returns all items
func getItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

// AddItem adds a new item
func addItem(c *gin.Context) {
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
func updateItem(c *gin.Context) {
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
func deleteItem(c *gin.Context) {
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

// Exported function for Vercel
func VercelHandler(w http.ResponseWriter, r *http.Request) {
    router := gin.Default()
    router.GET("/api/items", getItems) // Get all items
    router.POST("/api/items", addItem) // Add new item
    router.POST("/api/items/update/:id", updateItem) // Update item
    router.GET("/api/items/delete/:id", deleteItem) // Delete item
    router.ServeHTTP(w, r)
}
