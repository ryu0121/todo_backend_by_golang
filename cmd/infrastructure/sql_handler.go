package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"todo_app/cmd/interfaces/database"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewHandler() *SqlHandler {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("PORT")
	dbname := os.Getenv("DB_NAME")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, dbname)
	conn, err := sql.Open(os.Getenv("DB"), connection)
	if err != nil {
		panic(err.Error())
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

// 戻り値を名前付き戻り値にできなかった
// この問題で詰まったのでメモを残す
// 元々は func (handler *SqlHandler) Query(statement string, args ...interface{}) (row database.Row, err error) と名前付きで返していた
// このようにすると最初に変数にdatabase.Rowインターフェース型が割り当てられる
// それによって row.Rows をいうフィールドがないと怒られる...
// → 構造体の状態で戻り値に返し、戻り値の型はインターフェース型にするのが正解
func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Rows, error) {
	rows, err := handler.Conn.Query(statement, args...)
	sqlRows := new(SqlRows)
	if err != nil {
		return sqlRows, err
	}
	sqlRows.Rows = rows
	return sqlRows, err
}

func (handler *SqlHandler) QueryRow(statement string, args ...interface{}) database.Row {
	row := handler.Conn.QueryRow(statement, args...)
	sqlRow := new(SqlRow)
	sqlRow.Row = row
	return sqlRow
}

func (handler *SqlHandler) Exec(statement string, args ...interface{}) (database.Result, error) {
	sqlResult := new(SqlResult)
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return sqlResult, err
	}
	sqlResult.Result = result
	return sqlResult, err
}

type SqlRows struct {
	Rows *sql.Rows
}

func (r *SqlRows) Scan(fields ...interface{}) error {
	return r.Rows.Scan(fields...)
}

func (r *SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r *SqlRows) Close() error {
	return r.Rows.Close()
}

type SqlRow struct {
	Row *sql.Row
}

func (r *SqlRow) Scan(fields ...interface{}) error {
	return r.Row.Scan(fields...)
}

type SqlResult struct {
	Result sql.Result
}

func (r *SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}
