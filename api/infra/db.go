package infra

import (
	"api/gen/entv1"
	"errors"
	"fmt"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *entv1.Client {
	mysqlClient := mysqlInit()
	return mysqlClient
}

func mysqlInit() *entv1.Client {
	myENV, err := newENV().
		setEnvMyUser(os.Getenv("MYSQL_USER")).
		setEnvMyPassword(os.Getenv("MYSQL_PASSWORD")).
		setEnvMyHost(os.Getenv("MYSQL_HOST")).
		setEnvMyPort(os.Getenv("MYSQL_PORT")).
		setEnvMyDatabase(os.Getenv("MYSQL_DATABASE")).
		checkError()
	if err != nil {
		fmt.Println(err)
	}

	dsn, err := myENV.combineENV().transDSN()
	if err != nil {
		fmt.Println(err)
	}

	client, err := entv1.Open("mysql", dsn.toString())
	if err != nil {
		panic(fmt.Errorf("[FAILED] opening connection to mysql: %v", err))
	}

	return client
}

type dsn string
type combinedENV string

func (s dsn) toString() string {
	return string(s)
}

func (s combinedENV) toString() string {
	return string(s)
}

func (s combinedENV) transDSN() (dsn, error) {
	r, err := regexp.Compile(`.+:.+@tcp\(.+:\d+\)/.+`)
	if err != nil {
		return "", err
	}
	if !r.MatchString(s.toString()) {
		return "", errors.New("invalid DSN pattern")
	}
	return dsn(s), nil
}

func (me *myENV) combineENV() combinedENV {
	return combinedENV(me.User + ":" + me.Password + "@tcp(" + me.Host + ":" + me.Port + ")/" + me.Database)
}

type myENV struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Err      error
}

func newENV() myENV {
	return myENV{}
}

func (me myENV) setEnvMyUser(s string) myENV {
	if me.Err != nil {
		return me
	}
	if me.User = s; me.User == "" {
		me.Err = errors.New("you can't read mysql_user")
		return me
	}
	return me
}

func (me myENV) setEnvMyPassword(s string) myENV {
	if me.Err != nil {
		return me
	}
	if me.Password = s; me.Password == "" {
		me.Err = errors.New("you can't read mysql_password")
		return me
	}
	return me
}

func (me myENV) setEnvMyHost(s string) myENV {
	if me.Err != nil {
		return me
	}
	if me.Host = s; me.Host == "" {
		me.Err = errors.New("you can't read mysql_host")
		return me
	}
	return me
}

func (me myENV) setEnvMyPort(s string) myENV {
	if me.Err != nil {
		return me
	}
	if me.Port = s; me.Port == "" {
		me.Err = errors.New("you can't read mysql_port")
		return me
	}
	return me
}

func (me myENV) setEnvMyDatabase(s string) myENV {
	if me.Err != nil {
		return me
	}
	if me.Database = s; me.Database == "" {
		me.Err = errors.New("you can't read mysql_database")
		return me
	}
	return me
}

func (me myENV) checkError() (myENV, error) {
	return me, me.Err
}
