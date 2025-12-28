package postgresql

import "fmt"

func GetUrl(dbParams DBParams) string {
	url := "postgres://%s:%s@%s:%s/%s?sslmode=%s"
	return fmt.Sprintf(
		url,
		dbParams.Username,
		dbParams.Password,
		dbParams.Host,
		dbParams.Port,
		dbParams.Name,
		dbParams.SSLMode,
	)
}
