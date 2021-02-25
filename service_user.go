package service_user

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
	"strconv"
)
var db, _ = sql.Open("mysql", "root:jerrytran97@tcp(127.0.0.1:3306)/goblog")

type student struct {
	Id int `json:"id"`
	Fullname string `json:"fullname"`
	Age int `json:"age"`
	Location string `json:"location"`
}

func GetUser(ctx *fasthttp.RequestCtx)  {
	id_request, _ := ctx.UserValue("id").(string)
	id, _ := strconv.Atoi(id_request)
	getDb, err := db.Query("SELECT * FROM student WHERE id=?", id)
	if err!= nil {
		panic(err.Error())
	}
	var s student
	for getDb.Next() {
		err = getDb.Scan(&s.Id, &s.Fullname, &s.Age, &s.Location)
		if err!=nil {
			print(err.Error())
		}
	}
	json_student, err := json.Marshal(s)
	if err != nil {
		panic(err.Error())
	}
	ctx.Write(json_student)
}

func GetUsers(ctx *fasthttp.RequestCtx)  {
	var students []student
	getDbs, err := db.Query("SELECT * FROM student")
	if err!=nil {
		panic(err.Error())
	}
	for getDbs.Next() {
		var s student
		err = getDbs.Scan(&s.Id, &s.Fullname, &s.Age, &s.Location)
		students = append(students, s)
	}
	json_students, err := json.Marshal(students)
	if err!=nil {
		panic(err.Error())
	}
	ctx.Write(json_students)
}

func CreateStudent(ctx *fasthttp.RequestCtx)  {
	json_student := ctx.PostBody();
	var s student
	err := json.Unmarshal([]byte(json_student), &s)
	if err!=nil {
		panic(err.Error())
		ctx.SetStatusCode(fasthttp.StatusPreconditionFailed)
	}
	insertDB, err := db.Prepare("INSERT INTO student(id, fullname, age, location) values (?,?,?,?)");
	if err!=nil {
		panic(err.Error())
		ctx.SetStatusCode(fasthttp.StatusPreconditionFailed)
	}
	insertDB.Exec(s.Id, s.Fullname, s.Age, s.Location)
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func DeleteStudent(ctx *fasthttp.RequestCtx)  {
	id_request, _ := ctx.UserValue("id").(string)
	id, _ := strconv.Atoi(id_request)
	deleteDB, err := db.Prepare("DELETE FROM student WHERE id=?")
	if err != nil {
		panic(err)
	}
	deleteDB.Exec(id)
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func UpdateStudent(ctx *fasthttp.RequestCtx) {
	s := new(student)
	json_student := ctx.PostBody();
	err := json.Unmarshal([]byte(json_student), &s)
	id_request, _ := ctx.UserValue("id").(string)
	id, _ := strconv.Atoi(id_request)
	updateDB, err := db.Prepare("UPDATE student SET fullname=?, age=?, location=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	updateDB.Exec(s.Fullname, s.Age, s.Location, id)

}