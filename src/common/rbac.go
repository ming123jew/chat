package common

import(
	"github.com/mikespook/gorbac"
	"strings"
)

const  (
	USER_AUTH_TYPE 		=   2			// 默认认证类型 0:不认证 1 session登录认证 2 实时认证
	USER_AUTH_KEY 		= "USER_AUTH_KEY"
	ADMIN_AUTH_KEY 		= "ADMIN_AUTH_KEY"
	_ACCESS_LIST 		= "_ACCESS_LIST"
	USER_AUTH_ON 		= true
	NOT_AUTH_MODULE         = "public,user,index"	// 默认无需认证模块
	NOT_AUTH_ACTION         = "login"		// 默认无需认证操作

)

var Rbac = gorbac.New()

func init()  {
	rAdmin := gorbac.NewStdRole("role-admin") 		//角色:超级管理员
	pAdmin := gorbac.NewStdPermission("permission-a")  	//权限
	rAdmin.Assign(pAdmin)						//给角色增加
	Rbac.Add(rAdmin)						//初始化



	rUser := gorbac.NewStdRole("role-user") 		//角色:用户
	pUser := gorbac.NewStdPermission("permission-a")  	//权限
	rUser.Assign(pUser)					//给角色增加
	Rbac.Add(rUser)						//初始化

}
//Determine whether need to verify
func CheckAccess(params []string) bool {
	if len(params) < 2 {
		return false
	}
	//module not need to verify
	for _, nap := range strings.Split(NOT_AUTH_MODULE, ",") {
		//log.Println(nap,":",params[1])
		if params[1] == nap {
			return false
		}
	}
	//action not need to verify
	for _, nap := range strings.Split(NOT_AUTH_ACTION, ",") {
		//log.Println(nap,":",params[2])
		if params[1] == nap {
			return false
		}
	}

	return true
}



/**
	p := gorbac.NewStdPermission(permission)
	if Rbac.IsGranted(id, p, nil){
		r = true
	}else{
		r = false
		log.Println("role:",id,"perm:",permission,"result:false")
	}
 */