package controllers

import (
	"TestWallet/config"
	"TestWallet/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var users models.User
	if err := config.DB.Unscoped().Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Verify WalletTop-up
func VerifyTopUp(c *gin.Context) {
	var request struct {
		UserID        int    `json:"user_id"`
		Amount        int    `json:"amount"`
		PaymentMethod string `json:"payment_method"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("request.Amount", request.UserID)
	a := float64(0.1)
	b := float64(0.2)
	fmt.Println("%v\n", a+b)

	var user models.User
	if err := config.DB.Where("user_id = ?", request.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if request.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be positive"})
		return
	}

	var existingTransaction models.Transaction
	if err := config.DB.Where("user_id = ?", request.UserID).Order("transaction_id DESC").First(&existingTransaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No existing transaction found"})
		return
	}
	transaction := models.Transaction{
		TransactionID: existingTransaction.TransactionID,
		UserID:        request.UserID,
		Amount:        request.Amount,
		PaymentMethod: request.PaymentMethod,
		Status:        "verified",
		ExpiresAt:     time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}

	config.DB.Create(&transaction)

	c.JSON(http.StatusOK, transaction)
}

// Confirm Top-up API
func ConfirmTopUp(c *gin.Context) {
	var request struct {
		TransactionID string `json:"transaction_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var transaction models.Transaction
	if err := config.DB.Where("transaction_id = ?", request.TransactionID).First(&transaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if transaction.Status != "verified" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction is not verified"})
		return
	}

	expiredTime, _ := time.Parse(time.RFC3339, transaction.ExpiresAt)
	if time.Now().After(expiredTime) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction expired"})
		return
	}
	//check_ update
	var Checktransaction models.Transaction
	if err := config.DB.Where("transaction_id = ?", request.TransactionID).First(&Checktransaction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found for this user_id"})
		return
	}

	Checktransaction.Balance = 50050
	Checktransaction.Status = "completed"

	config.DB.Model(&Checktransaction).Where("transaction_id = ?", request.TransactionID).Updates(&Checktransaction)

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": transaction.TransactionID,
		"user_id":        transaction.UserID,
		"amount":         transaction.Amount,
		"status":         Checktransaction.Status,
		"balance":        Checktransaction.Balance,
	})
}
