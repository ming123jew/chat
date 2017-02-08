package common

import(
	"github.com/mikespook/gorbac"
	"log"
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

func CheckAccess(id string,permission string) (r bool,e error)   {

	p := gorbac.NewStdPermission(permission)
	if Rbac.IsGranted(id, p, nil){
		r = true
	}else{
		r = false
		log.Println("role:",id,"perm:",permission,"result:false")
	}
	return
}
