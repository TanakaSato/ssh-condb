package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Sshconfig struct {
	ID       int
	Hostname string
	Password string
	Username string
	Authkey  string
	Proxy    int
	Port     int
}

func Insert() {
	// TODO
	// _, err = db.Exec("insert into user values (?, ?, ?) ", 1, "hoge", 30)
}

func Update() {
	// TODO
}

func GetSingleHost(hostname string) Sshconfig {
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/testthird?parseTime=true")
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
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/testthird?parseTime=true")
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
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/testthird?parseTime=true")
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
