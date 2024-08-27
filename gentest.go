package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"testgorm/query"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	//g := gen.NewGenerator(gen.Config{
	//	OutPath: "query",
	//	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	//})
	dsn := "root:48044705@tcp(localhost:3306)/go?charset=utf8&parseTime=True&loc=Local"
	gormdb, _ := gorm.Open(mysql.Open(dsn))
	//g.UseDB(gormdb) // reuse your gorm db
	query.SetDefault(gormdb)

	//// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	//g.ApplyInterface(func(Querier) {}, model.User{})
	//g.GenerateModel("users")
	//g.Execute()

	role, err := query.Q.User.FilterWithNameAndRole("wen", "1")
	if err != nil {
		return
	}
	fmt.Println(role)

}
