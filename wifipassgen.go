package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	lowercase  = "abcdefghijklmnopqrstuvwxyz"
	uppercase  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits     = "0123456789"
	symbols    = "!@#$%^&*()-_=+[]{}|;:,.<>?"
	allChars   = lowercase + uppercase + digits + symbols
	minLength  = 12
	maxLength  = 32
	defaultLen = 16
)

func generatePassword(length int) (string, error) {
	if length < minLength {
		return "", fmt.Errorf("la longitud mínima es %d", minLength)
	}
	if length > maxLength {
		return "", fmt.Errorf("la longitud máxima es %d", maxLength)
	}

	var password strings.Builder
	charTypes := []string{lowercase, uppercase, digits, symbols}

	// Asegurar que al menos un carácter de cada tipo esté presente
	for _, charType := range charTypes {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charType))))
		if err != nil {
			return "", err
		}
		password.WriteByte(charType[n.Int64()])
	}

	// Llenar el resto de la contraseña
	for i := password.Len(); i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(allChars))))
		if err != nil {
			return "", err
		}
		password.WriteByte(allChars[n.Int64()])
	}

	// Mezclar los caracteres
	passwordRunes := []rune(password.String())
	for i := len(passwordRunes) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}
		passwordRunes[i], passwordRunes[j.Int64()] = passwordRunes[j.Int64()], passwordRunes[i]
	}

	return string(passwordRunes), nil
}

func main() {
	length := defaultLen
	if len(os.Args) > 1 {
		userLength, err := strconv.Atoi(os.Args[1])
		if err == nil && userLength >= minLength && userLength <= maxLength {
			length = userLength
		}
	}

	password, err := generatePassword(length)
	if err != nil {
		fmt.Printf("Error al generar la contraseña: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Contraseña WiFi generada (longitud %d): %s\n", length, password)
}