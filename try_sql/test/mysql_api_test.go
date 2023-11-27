package test

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_sql/conn"
	"fmt"
	"testing"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		_ = fmt.Errorf("mysql error: %s\n", err)
	}
	err = conn.InitMysqlConfig()
	if err != nil {
		_ = fmt.Errorf("mysql error: %s\n", err)
	}
}

type User struct {
	Id       int    `db:"id, primary"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

func TestSelectSql(t *testing.T) {
	var u []User
	err := conn.DB.Select(&u, "SELECT * FROM `user`")
	if err != nil {
		_ = fmt.Errorf("mysql error: %s\n", err)
	}
	for i := range u {
		fmt.Println("id:", u[i].Id, "username:", u[i].Username, "email:", u[i].Email)
	}
}

// 查询单条数据
func TestQueryRow(t *testing.T) {
	u := User{}
	row := conn.DB.QueryRow("SELECT * FROM `user`")
	err := row.Scan(&u.Id, &u.Username, &u.Email)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u.Id, u.Username, u.Email)
}

// 查多单条数据
func TestQuerySql(t *testing.T) {
	rows, err := conn.DB.Query("SELECT *\nFROM\n`user`\nWHERE \nusername=?", "ledger")
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		var id int
		var username string
		var email string
		err = rows.Scan(&id, &username, &email)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(id, username, email)
	}
}

func TestExec(t *testing.T) {
	result, err := conn.DB.Exec("INSERT INTO `user` VALUES (?,?,?)", "H0132132HH", "yb", "123@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("affected: %d\n", affected)
}
