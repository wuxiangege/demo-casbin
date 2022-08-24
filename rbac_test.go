package casbindemo

import (
	"fmt"
	"testing"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"

	xormadapter "github.com/casbin/xorm-adapter/v2"
)

func TestRbac(t *testing.T) {
	a, _ := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/") //默认
	// a, _ := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin", true)                                 //自定义库
	// a, _ := xormadapter.NewAdapterWithTableName("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin", "casbin_rule", "", true) //自定义库和表

	e, _ := casbin.NewEnforcer("rbac_model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	enable, _ := e.Enforce("person1", "data1", "read")
	fmt.Println(enable)

	// sub(从jwt中解析) + obj + act
	enable2, _ := e.Enforce("chenjun", "data1", "read")
	fmt.Println(enable2)

	// Modify the policy.
	// e.AddPolicy("wuxian", "data1", "read")
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	err := e.SavePolicy()
	fmt.Println(err)
}
