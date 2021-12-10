package repository

import (
	"github.com/netoudi/codepay-process-transaction/adapter/repository/fixture"
	"github.com/netoudi/codepay-process-transaction/domain/entity"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir)
	defer fixture.Down(db, migrationsDir)
	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 16.9, entity.APPROVED, "")
	assert.Nil(t, err)
}
