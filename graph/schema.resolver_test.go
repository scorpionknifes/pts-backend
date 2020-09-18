package graph

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/scorpionknifes/pts-backend/graph/model"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var r MockResolver

type MockResolver struct {
	mutationResolver mutationResolver
	queryResolver    queryResolver
	mock             sqlmock.Sqlmock
}

func TestSetupMockDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	gormDB, err := gorm.Open(sqlserver.New(sqlserver.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	resolver := Resolver{
		DB: gormDB,
	}

	r = MockResolver{
		mock:             mock,
		mutationResolver: mutationResolver{Resolver: &resolver},
		queryResolver:    queryResolver{Resolver: &resolver},
	}

	mock.ExpectExec(`CREATE TABLE "stories" \("id" bigint IDENTITY\(1,1\),"name" nvarchar\(MAX\),"count" bigint,"people" bigint,"tags" nvarchar\(MAX\),"created_at" datetimeoffset,"updated_at" datetimeoffset,PRIMARY KEY \("id"\)\)`)
	gormDB.AutoMigrate(&model.Story{}, &model.Turn{}, &model.User{})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateUser(t *testing.T) {
	TestSetupMockDB(t)
	r.mock.ExpectQuery(`INSERT INTO "users" \("name","created_at","updated_at"\) OUTPUT INSERTED\."id" VALUES \(@p1,@p2,@p3\)`)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r.mutationResolver.CreateUser(ctx)

	if err := r.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateStory(t *testing.T) {
	TestSetupMockDB(t)
	r.mock.ExpectQuery(`INSERT INTO "stories" \("name","count","people","tags","created_at","updated_at"\) OUTPUT INSERTED\."id" VALUES \(@p1,@p2,@p3,@p4,@p5,@p6\)`)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r.mutationResolver.CreateStory(ctx, model.StoryInput{})

	if err := r.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateTurn(t *testing.T) {
	TestSetupMockDB(t)
	r.mock.ExpectQuery(`INSERT INTO "turns" \("user_id","story_id","value","created_at","updated_at"\) OUTPUT INSERTED\."id" VALUES \(@p1,@p2,@p3,@p4,@p5\)`)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r.mutationResolver.CreateTurn(ctx, model.TurnInput{})

	if err := r.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStories(t *testing.T) {
	TestSetupMockDB(t)
	r.mock.ExpectQuery(`SELECT \* FROM "stories"`)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r.queryResolver.Stories(ctx)

	if err := r.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestStory(t *testing.T) {
	TestSetupMockDB(t)
	r.mock.ExpectQuery(`SELECT \* FROM "stories" WHERE "stories"\."id" \= @p1 ORDER BY "stories"\."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r.queryResolver.Story(ctx, 0)

	if err := r.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUsers(t *testing.T) {
	TestSetupMockDB(t)
	r.mock.ExpectQuery(`SELECT \* FROM "users"`)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r.queryResolver.Users(ctx)

	if err := r.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser(t *testing.T) {
	TestSetupMockDB(t)
	query := `SELECT \* FROM "users" WHERE "users"\."id" \= @p1 ORDER BY "users"\."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY`
	r.mock.ExpectQuery(query)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r.queryResolver.User(ctx, 0)

	if err := r.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
