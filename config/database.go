package config

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

// db 4mysqlClose
var db *sql.DB

// con OverSSHConnection
var con *ssh.Client

type ConfigDB struct {
	SSHHost     string
	SSHPort     int
	SSHUser     string
	SSHPassword string
	DBPort      int
	DBHost      string
	DBUser      string
	DBPassword  string
	Name        string
}

// NewConfigDB 初期化
func NewConfigDB(dbPort, sshPort int, sshHost, sshUser, sshPassword, dbHost, dbUser, dbPassword, dbName string) *ConfigDB {
	return &ConfigDB{
		SSHHost:     sshHost,
		SSHPort:     sshPort,
		SSHUser:     sshUser,
		SSHPassword: sshPassword,
		DBPort:      dbPort,
		DBHost:      dbHost,
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		Name:        dbName,
	}
}

func (d *ConfigDB) Init() (*sql.DB, error) {
	signer, err := ssh.ParsePrivateKey([]byte(d.SSHPassword))
	if err != nil {
		return nil, err
	}
	sshConfig := &ssh.ClientConfig{
		User:            d.SSHUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
	con, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", d.SSHHost, d.SSHPort), sshConfig)
	if err != nil {
		return nil, err
	}
	dialNet := "mysql+tcp"
	dialFunc := func(addr string) (net.Conn, error) {
		return con.Dial("tcp", addr)
	}
	mysql.RegisterDial(dialNet, dialFunc)

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", d.DBUser, d.DBPassword, dialNet, d.DBHost, d.DBPort, d.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		d.CloseSshAndDB()
		return nil, err
	}

	return db, err
}

// CloseSshAndDB SSH及びDB切断
func (d *ConfigDB) CloseSshAndDB() {
	if db != nil {
		db.Close()
	}
	if con != nil {
		con.Close()
	}
}
