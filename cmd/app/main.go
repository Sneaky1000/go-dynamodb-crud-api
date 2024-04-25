package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sneaky1000/go-dynamodb-crud-api/config"
	"github.com/Sneaky1000/go-dynamodb-crud-api/internal/repository/adapter"
	"github.com/Sneaky1000/go-dynamodb-crud-api/internal/repository/instance"
	"github.com/Sneaky1000/go-dynamodb-crud-api/internal/routes"
	"github.com/Sneaky1000/go-dynamodb-crud-api/internal/rules"
	RulesProduct "github.com/Sneaky1000/go-dynamodb-crud-api/internal/rules/product"
	"github.com/Sneaky1000/go-dynamodb-crud-api/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	configs := config.GetConfig()

	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("Waiting service starting.... ", nil)

	tables, err := checkTables(connection)

	if err != nil {
		logger.PANIC("Error listing tables: ", err)
	}

	for _, tableName := range tables {
		errors := Migrate(connection, tableName, tables)
		if len(errors) > 0 {
			for _, err := range errors {
				logger.PANIC(fmt.Sprintf("Error on migrate for table %s: ", tableName), err)
			}
		}
	}

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.INFO("Service running on port ", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB, currentTable string, allTables []string) []error {
	var errors []error

	for _, existingTableName := range allTables {
		if existingTableName == currentTable {
			logger.INFO(fmt.Sprintf("Table '%s' found. Skipping migration.", currentTable), nil)
			return nil
		}
	}

	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})

	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) ([]string, error) {
	var tableNames []string

	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return nil, err
	}
	for _, tableName := range response.TableNames {
		tableNames = append(tableNames, *tableName)
	}
	return tableNames, nil
}
