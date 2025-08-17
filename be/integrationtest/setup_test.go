package integrationtest

import (
	"context"
	"github.com/filtkac/consult-scheduler/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"

	tcPostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	container := startPostgresContainer(ctx)
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	testDB = initializeTestDB(container)
	config.AutoMigrate(testDB)
	loadTestData(testDB)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func startPostgresContainer(ctx context.Context) *tcPostgres.PostgresContainer {
	container, err := tcPostgres.Run(
		ctx,
		"postgres:17-alpine",
		tcPostgres.WithDatabase("consult_scheduler_test"),
		tcPostgres.WithUsername("cs_app_test"),
		tcPostgres.WithPassword("cs_app_test_password"),
		tcPostgres.BasicWaitStrategies(),
	)

	if err != nil {
		log.Fatalf("could not start postgres container: %s", err)
	}

	return container
}

func initializeTestDB(container *tcPostgres.PostgresContainer) *gorm.DB {
	ctx := context.Background()

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("failed to get connection string: %s", err)
	}

	testDBResult, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test database: %s", err)
	}
	return testDBResult
}

func loadTestData(testDB *gorm.DB) {
	data, err := os.ReadFile("../../test_data.sql")
	if err != nil {
		log.Fatalf("failed to read test data: %s", err)
	}

	result := testDB.Exec(string(data))
	if result.Error != nil {
		log.Fatalf("failed to load test data: %s", err)
	}

	log.Println("Test data loaded successfully")
}
