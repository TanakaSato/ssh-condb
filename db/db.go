package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Sshconfig struct {
	ID       int    `yaml:"id"`
	Hostname string `yaml:"hostname"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	Authkey  string `yaml:"authkey"`
	Proxy    int    `yaml:"proxy"`
	Port     int    `yaml:"port"`
}

func InsertDB(confs []Sshconfig) {

	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(0)

	rows, err := db.Query("SELECT * FROM sshconfig")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	n := 0
	for rows.Next() {
		n = n + 1
	}

	stmt, err := db.Prepare(`INSERT INTO inserttest.sshconfig(id, hostname, password, username, authkey, proxy, port) VALUES (?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, m := range confs {
		n = n + 1
		_, err := stmt.Exec(
			n,
			&m.Hostname,
			&m.Password,
			&m.Username,
			&m.Authkey,
			&m.Proxy,
			&m.Port)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func UpdateDB() {
	// TODO
}

func GetSingleHost(hostname string) Sshconfig {

	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(0)

	user := Sshconfig{}
	if err := db.QueryRow("SELECT * FROM sshconfig WHERE hostname = ?", hostname).Scan(
		&user.ID,
		&user.Hostname,
		&user.Password,
		&user.Username,
		&user.Authkey,
		&user.Proxy,
		&user.Port); err != nil {
		log.Fatal(err)
	}

	return user
}

func GetAnyHost() []Sshconfig {

	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(0)

	rows, err := db.Query("SELECT * FROM sshconfig")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []Sshconfig{}

	for rows.Next() {
		user := Sshconfig{}

		err := rows.Scan(
			&user.ID,
			&user.Hostname,
			&user.Password,
			&user.Username,
			&user.Authkey,
			&user.Proxy,
			&user.Port)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func GetID(id int) Sshconfig {

	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(0)

	user := Sshconfig{}
	if err := db.QueryRow("SELECT * FROM sshconfig WHERE id = ?", id).Scan(
		&user.ID,
		&user.Hostname,
		&user.Password,
		&user.Username,
		&user.Authkey,
		&user.Proxy,
		&user.Port); err != nil {
		log.Fatal(err)
	}

	return user
}
