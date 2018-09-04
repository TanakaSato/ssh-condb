package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Sshconfig is
type Sshconfig struct {
	ID       int    `gorm:"primary_key" yaml:"id"`
	Hostname string `gorm:"hostname" yaml:"hostname"`
	Password string `gorm:"password" yaml:"password"`
	Username string `gorm:"username" yaml:"username"`
	Authkey  string `gorm:"authkey" yaml:"authkey"`
	Proxy    int    `gorm:"proxy" yaml:"proxy"`
	Port     int    `gorm:"port" yaml:"port"`
}

// NewSshconfig is
func NewSshconfig(hostname, password, username, authkey string, proxy, port int) *Sshconfig {

	a := &Sshconfig{}

	a.ID = 0
	a.Hostname = hostname
	a.Password = password
	a.Username = username
	a.Authkey = authkey
	a.Proxy = proxy
	a.Port = port

	return a
}

// CompSshconfig is
func (a *Sshconfig) CompSshconfig(b Sshconfig) bool {

	if a.Hostname == b.Hostname && a.Password == b.Password && a.Username == b.Username &&
		a.Authkey == b.Authkey && a.Proxy == b.Proxy && a.Port == b.Port {
		return true
	}

	// log.Println("------------------------------------- a data is -------------------------------------")
	// log.Println(a)
	// log.Println("------------------------------------- b data is -------------------------------------")
	// log.Println(b)

	return false
}

var dbname = string("mysql")
var dbuser = string("root")
var dbpassword = string("mysql")
var protcol = string("tcp(127.0.0.1:3306)")
var conndbname = string("test_db")

func initDB() *gorm.DB {

	conn := dbuser + ":" + dbpassword + "@" + protcol + "/" + conndbname

	db, err := gorm.Open(dbname, conn)
	if err != nil {
		log.Println(err)
	}

	return db
}

// CreateNewHostData is a
func CreateNewHostData(hostname, password, username, authkey string, proxy, port int) {

	conf := NewSshconfig(hostname, password, username, authkey, proxy, port)

	newconf := InsertDB(*conf)

	if newconf.CompSshconfig(*conf) {
		log.Println("your config data is db insert success!")
	} else {
		log.Println("your config data is db insert FAILED!")
	}
}

// InsertDB is
func InsertDB(conf Sshconfig) Sshconfig {

	db := initDB()
	defer db.Close()

	db.Create(&conf)

	return conf
}

// InsertDBs is
func InsertDBs(confs []Sshconfig) {

	db := initDB()
	defer db.Close()

	for _, c := range confs {
		_ = InsertDB(c)
	}
}

// UpdateDB is
func UpdateDB() {

	db := initDB()
	defer db.Close()

	// TODO
}

// DeleteDB is
func DeleteDB(id int) {

	db := initDB()
	defer db.Close()

	conf := GetID(id)
	if conf.Hostname == "" {
		log.Println("unknown data")
		return
	}

	db.Delete(conf)

	newconf := GetID(id)

	if newconf.Hostname == "" {
		log.Println("remove data success!")
	} else {
		log.Println("remove data FAILED!")
	}

}

// GetHost is get from db any host sshconfig setting
func GetHosts(hostname string) []Sshconfig {

	db := initDB()
	defer db.Close()

	// log.Println("Search hostname is " + hostname)
	confs := []Sshconfig{}

	if hostname == "*" {
		db.Find(&confs)
	} else {
		db.Find(&confs, "hostname=?", hostname)
	}

	return confs
}

// GetID is get from db single host sshconfig setting to id
func GetID(id int) Sshconfig {

	db := initDB()
	defer db.Close()

	conf := Sshconfig{}

	conf.ID = id

	db.Find(&conf)

	return conf
}

// GetProxy is piyo
func GetProxy(id int) []Sshconfig {

	db := initDB()
	defer db.Close()

	confs := []Sshconfig{}

	db.Find(&confs, "proxy=?", id)

	return confs
}
