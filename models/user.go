package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户
type User struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(100)"`
	Nickname string `orm:"size(100)"`
	Pwd      string `orm:"size(100)"`
	Email    string `orm:"size(100)"`
	Sex      string `orm:"size(2)"`
	Roleid   string `orm:"size(100)"`
	Status   int64
	Phone    string `orm:"size(16)"`
}

//新增用户
func Create(uid int64, name string, nickname string, pwd string, email string,
	sex string, roleId string, status int64, phone string) (user User) {

	//查询用户是否已存在
	user, err := QueryById(uid)
	if err == true {
		return user
	} else {
		o := orm.NewOrm()
		o.Using("user")
		newuser := new(User)
		//赋值给模型
		newuser.Id = uid
		newuser.Name = name
		newuser.Nickname = nickname
		newuser.Pwd = pwd
		newuser.Email = email
		newuser.Sex = sex
		newuser.Roleid = roleId
		newuser.Status = status
		newuser.Phone = phone
		//新增数据
		o.Insert(newuser)

		return *newuser
	}
}

//删除用户
func DeleteById(id int64) bool {

	o := orm.NewOrm()
	o.Using("default")
	//根据ID得到用户模型
	if num, err := o.Delete(&User{Id: id}); err == nil {
		fmt.Println("删除影响的行数:")
		fmt.Println(num)
		return true
	} else {
		return false
	}
}

//更新用户
func UpdateById(id int, table string, filed map[string]interface{}) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable(
		table).Filter(
		"Id", id).Update(
		filed)
	if err == nil {
		return true
	}
	return false
}

//根据用户ID查询用户
func QueryById(uid int64) (User, bool) {

	o := orm.NewOrm()
	u := User{Id: uid}

	err := o.Read(&u)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return u, false
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return u, false
	} else {
		fmt.Println(u.Id, u.Name)
		return u, true
	}
}

//根据用户名称查询用户
func QueryByName(name string) (User, error) {
	var user User

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	err := qs.Filter("Name", name).One(&user)
	fmt.Println(err)
	if err == nil {
		fmt.Println(user.Name)
		return user, nil
	}
	return user, err
}

//根据用户数据列表
func DataList() (users []User) {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var us []User
	cnt, err := qs.Filter("id__gt", 0).OrderBy("-id").Limit(10, 0).All(&us)
	if err == nil {
		fmt.Printf("count", cnt)
	}
	return us
}

//查询语句，sql语句的执行
//格式类似于:o.Raw("UPDATE user SET name = ? WHERE name = ?", "testing", "slene")
//
func QueryBySql(sql string, qarms []string) bool {

	o := orm.NewOrm()

	//执行sql语句
	o.Raw(sql, qarms)

	return true
}

//根据用户分页数据列表
func LimitList(pagesize int, pageno int) (users []User) {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var us []User
	cnt, err := qs.Limit(pagesize, (pageno-1)*pagesize).All(&us)
	if err == nil {
		fmt.Printf("count", cnt)
	}
	return us
}

//根据用户数据总个数
func GetDataNum() int64 {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var us []User
	num, err := qs.Filter("id__gt", 0).All(&us)
	if err == nil {
		return num
	} else {
		return 0
	}
}

//初始化模型
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}
