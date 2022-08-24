package casbindemo

import (
	"fmt"
	"testing"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

func check(e *casbin.Enforcer, sub, domain, obj, act string) {
	ok, _ := e.Enforce(sub, domain, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s in %s\n", sub, act, obj, domain)
	} else {
		fmt.Printf("%s CANNOT %s %s in %s\n", sub, act, obj, domain)
	}
}

func TestRbacTenant(t *testing.T) {
	// e, err := casbin.NewEnforcer("./casbin/model.conf", "./casbin/policy.csv")
	// if err != nil {
	// 	log.Fatalf("NewEnforecer failed:%v\n", err)
	// }

	// a, _ := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/")                                                    //默认
	// a, _ := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin", true)                                        //自定义库
	a, _ := xormadapter.NewAdapterWithTableName("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin", "casbin_tenant_rule", "", true) //自定义库和表

	e, _ := casbin.NewEnforcer("rbac_tenant_model.conf", a)

	check(e, "dajun", "tenant1", "data1", "read")
	check(e, "dajun", "tenant1", "data1", "write")
	check(e, "dajun", "tenant2", "data2", "read")
}
