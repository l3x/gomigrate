package gomigrate

import "strings"

type Migratable interface {
	SelectMigrationTableSql() string
	CreateMigrationTableSql() string
	GetMigrationSql() string
	MigrationLogInsertSql() string
	MigrationLogDeleteSql() string
	GetMigrationCommands(string) []string
}

// MYSQL

type Mysql struct{}

func (m Mysql) SelectMigrationTableSql() string {
	return "SELECT table_name FROM information_schema.tables WHERE table_name = ? AND table_schema = (SELECT DATABASE())"
}

func (m Mysql) CreateMigrationTableSql() string {
	return `CREATE TABLE gomigrate (
                  id           INT          NOT NULL AUTO_INCREMENT,
                  migration_id BIGINT       NOT NULL UNIQUE,
                  PRIMARY KEY (id)
                )`
}

func (m Mysql) GetMigrationSql() string {
	return `SELECT migration_id FROM gomigrate WHERE migration_id = ?`
}

func (m Mysql) MigrationLogInsertSql() string {
	return "INSERT INTO gomigrate (migration_id) values (?)"
}

func (m Mysql) MigrationLogDeleteSql() string {
	return "DELETE FROM gomigrate WHERE migration_id = ?"
}

func (m Mysql) GetMigrationCommands(sql string) []string {
	count := strings.Count(sql, ";")
	commands := strings.SplitN(string(sql), ";", count)
	return commands
}

// MARIADB

type Mariadb struct {
	Mysql
}
