package web

import (
	// "encoding/json"
	// "errors"
	// "fmt"
	// "log"
	"math/rand"
	// "strconv"
	// "strings"
	"time"
	// "unicode"
	// "golang.org/x/crypto/bcrypt"
	// "sw-sys/api-service/cache"
	// "sw-sys/api-service/database"
	// "sw-sys/api-service/logger"
	// "sw-sys/api-service/messaging"
	// "sw-sys/api-service/storage"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	sizeString  = 8
)
