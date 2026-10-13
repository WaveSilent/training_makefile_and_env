package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"

	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	user := os.Getenv("NEW_USER")

	switch user {
	case "YES":
		scanner := bufio.NewScanner(os.Stdin)

		ok := scanner.Scan()

		if !ok {
			return
		}

		inputString := scanner.Text()

		sliceString := strings.Fields(inputString)

		fullName, phoneNumber := sliceString[0], sliceString[1]

		if err := simple_sql.InsertRow(ctx, conn, fullName, phoneNumber); err != nil {
			panic(err)
		}

	case "NO":
		users, err := simple_sql.SelectRows(ctx, conn)

		if err != nil {
			panic(err)
		}

		pp.Print("users:", users)

	default:
		fmt.Println("Переменная окружения NEW_USER имеет неизвестное значение!")
		return
	}

	fmt.Println("succeed!")

}
