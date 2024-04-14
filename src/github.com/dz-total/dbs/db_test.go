package dbs

import (
	"FinalTaskAppGoBasic/configs"
	"FinalTaskAppGoBasic/logs"
	"FinalTaskAppGoBasic/models"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestUsersTransactionsIntegration(t *testing.T) {
	fmt.Printf("initiation logrus... ")
	logs.InitializeLogging("../configs")
	fmt.Println("logrus Init Done.")

	configApp, err := configs.LoadConfig(".././configs/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	Connect(configApp.DataBase)
	Migrate()
	defer Close() //DB.Migrator().DropTable(&models.Users{}, &models.Transactions{})

	// Create a user
	r := NewRand()
	userName, userEmail := GenUserData(r) //Гененируем рандомные данные пользователя
	fmt.Println(userName, userEmail)
	user := models.Users{
		Username: userName,
		Email:    userEmail,
		Password: "password",
	}
	if err := DB.Create(&user).Error; err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Create a transaction
	transaction := models.Transactions{
		UserID:          user.ID,
		Amount:          randomdFloats(3, r),
		Currency:        randomString(3, r),
		TransactionType: "income",
		Category:        "Test-" + randomString(1, r),
		TransactionDate: time.Now(),
		Description:     "Test_Transaction",
		Commission:      0.0,
	}
	if err := DB.Create(&transaction).Error; err != nil {
		t.Fatalf("Failed to create transaction: %v", err)
	}

	// Fetch and verify the transaction
	var fetchedTx models.Transactions

	DB.First(&fetchedTx, transaction.ID)

	fmt.Println(*&transaction)
	fmt.Println(*&fetchedTx)

	if err := DB.First(&fetchedTx, transaction.ID).Error; err != nil {
		t.Fatalf("Failed to fetch transaction: %v", err)
	}

	if !models.IsTransactionType(fetchedTx.TransactionType) {
		t.Errorf("Invalid transaction type: %v", fetchedTx.TransactionType)
	}
}

// NewRand создаёт и возвращает новый экземпляр rand.Rand, инициализированный текущим временем
func NewRand() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}

func randomString(n int, r *rand.Rand) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[r.Intn(len(letters))]
	}
	return string(s)
}
func randomdFloats(n int, r *rand.Rand) float64 {
	var digits = []rune("0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = digits[r.Intn(len(digits))]
	}
	floatValue, err := strconv.ParseFloat(string(s), 64)

	if err != nil {
		fmt.Println("Ошибка при конвертации:", err)
		return 0.
	}
	return floatValue
}

func GenUserData(r *rand.Rand) (string, string) {
	user := randomString(4, r)   // Генерируем случайное имя пользователя длиной 10 символов
	domain := randomString(4, r) // Генерируем случайное доменное имя длиной 5 символов
	email := fmt.Sprintf("%s@%s.org", user, domain)
	return user, email
}
