package session

import (
	"database/sql"
	"orm/log"
	"strings"
)

type Session struct {
	db      *sql.DB         // 数据库连接实例
	sql     strings.Builder // sql 语句
	sqlVars []interface{}   // 占位符的值
}

// New returns a new session reading from db
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Reset resets the Session to be reading from s
func (s *Session) Reset() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB returns a *sql.DB to be reading from s
func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values)
	return s
}

// Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Reset()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a records from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Reset()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// Query gets a list of records from db
func (s *Session) Query() (rows *sql.Rows, err error) {
	defer s.Reset()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
