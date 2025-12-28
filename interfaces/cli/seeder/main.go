package main

import (
	"flag"
	"fmt"
	"strings"

	"iam-service/constants"
	seeder "iam-service/interfaces/cli/seeder/handlers/main_db/iam"
	"iam-service/interfaces/cli/seeder/helpers"
)

func main() {
	// flags
	dbFlag := flag.String("db", constants.DB_LIST_MAIN, "database to use (main_db, external_db, etc)")
	flag.Parse()

	// module argument (positional)
	var module string
	if flag.NArg() > 0 {
		module = flag.Arg(0)
	}

	if *dbFlag == "" {
		*dbFlag = constants.DB_LIST_MAIN
	}

	type Seeder struct {
		Name string
		Run  helpers.SeederFunc
	}

	seeders := []Seeder{
		{"user", seeder.UserSeeder},
		{"role", seeder.RoleSeeder},
		{"permission", seeder.PermissionSeeder},
		{"role_permission", seeder.RolePermissionSeeder},
		{"user_role", seeder.UserRoleSeeder},
	}

	// run single seeder
	if module != "" {
		module = strings.ToLower(module)

		for _, s := range seeders {
			if s.Name == module {
				fmt.Println("==============")
				fmt.Println("Running seeder:", s.Name)
				fmt.Println("==============")

				s.Run()
				return
			}
		}

	}

	// run all seeders
	fmt.Println("Running ALL seeders")
	for _, seeder := range seeders {
		fmt.Println("==============")
		fmt.Println("Running seeder:", seeder.Name)
		fmt.Println("==============")

		seeder.Run()
	}
}
