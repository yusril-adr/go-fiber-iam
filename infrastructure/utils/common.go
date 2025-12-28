package utils

import (
	"encoding/json"
	"iam-service/infrastructure/errors"
	"math/rand"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Find is a generic helper function to find the first element in a slice
// that satisfies a given predicate function.
func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, item := range slice {
		if predicate(item) {
			return item, true // Return the element and true (found)
		}
	}
	var zero T         // Zero value of type T
	return zero, false // Return zero value and false (not found)
}

// Map transforms a slice of type A into a slice of type B
// using the provided transformation function f.
func Map[A, B any](arr []A, f func(A) B) []B {
	newArr := make([]B, 0, len(arr))
	for _, a := range arr {
		newArr = append(newArr, f(a))
	}
	return newArr
}

func Filter[T any](arr []T, predicate func(T) bool) []T {
	return slices.DeleteFunc(arr, func(item T) bool {
		return !predicate(item)
	})
}

func StringToSnakeCase(s string) string {
	var result strings.Builder
	result.Grow(len(s) + 5) // Approximate additional space for underscores

	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}

// If returns the value of trueVal if cond is true, otherwise falseVal.
func If[T any](cond bool, trueVal, falseVal T) T {
	if cond {
		return trueVal
	}
	return falseVal
}

func UUIDChecker(uuidString string, fieldKey string) uuid.UUID {
	id, err := uuid.Parse(uuidString)
	if err != nil {
		panic(errors.ValidationError(fieldKey, err.Error()))
	}

	return id
}

func LogJson(data interface{}, logLevel logrus.Level) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.WithField("data", string(jsonData)).Log(logLevel, "")
}

// GenerateSlug converts a string to a slug (lowercase with underscores, alphanumeric only)
// Example: "Super Admin" → "super_admin", "Test Role!" → "test_role"
func GenerateSlug(s string) string {
	// Trim whitespace
	s = strings.TrimSpace(s)

	// Convert to lowercase
	s = strings.ToLower(s)

	// Replace spaces with underscores
	s = strings.ReplaceAll(s, " ", "_")

	// Remove special characters (keep only alphanumeric and underscores)
	var result strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// GenerateRandomSuffix generates a random alphanumeric string of specified length
// Example: GenerateRandomSuffix(4) → "a7k2", "m9p4", etc.
func GenerateRandomSuffix(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
