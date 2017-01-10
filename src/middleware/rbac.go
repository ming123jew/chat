package middleware

import (
	. "common"
	"log"

	"strings"
)

type Rb struct {}

func (x Rb) GetAccessList( p int ) {

	var sql = "select node.id,node.name from chat_role as role," +
		"chat_role_user as user,chat_app_access as access ,chat_app_node as node " +
		"where user.userid=? and user.roleid=role.id " +
		"and ( access.roleid=role.id  or (access.roleid=role.pid and role.pid!=0 ) ) " +
		"and role.status=1 and access.nodeid=node.id and node.level=1 and node.status=1"

	result ,err := Orm.Query(sql,p)

	if err != nil {
		log.Println(err)
	}

	access := make(map[string]map[string]interface{})
	for   k ,v := range result{

		// 读取项目的模块权限
		appid := v["id"]
		appname := v["name"]
		log.Println("here k v:",k, "id:",string(appid),"name:",string(appname))

		sql :=   "select node.id,node.name from chat_role as role," +
			"chat_role_user as user,chat_app_access as access ,chat_app_node as node " +
			"where user.userid=? and user.roleid=role.id " +
			"and ( access.roleid=role.id  or (access.roleid=role.pid and role.pid!=0 ) ) " +
			"and role.status=1 and access.nodeid=node.id and node.level=2 and node.pid=? and node.status=1"
		result2 ,err := Orm.Query(sql,p,appid)
		if err != nil {
			log.Println(err)
		}

		// 判断是否存在公共模块的权限
		publicAction  :=  make(map[string]string)
		var result22  []map[string][]byte
		for k2,v2 := range result2{
			moduleid := v2["id"]
			modulename := v2["name"]
			log.Println("here2 k2 v2:",k2, "id:",string(moduleid),"name:",string(modulename))
			if strings.ToUpper(string(modulename))=="PUBLIC" {

				sql := "select node.id,node.name from chat_role as role," +
				"chat_role_user as user,chat_app_access as access ,chat_app_node as node " +
					"where user.userid=? and user.roleid=role.id " +
					"and ( access.roleid=role.id  or (access.roleid=role.pid " +
					"and role.pid!=0 ) ) and role.status=1 and access.nodeid=node.id " +
					"and node.level=3 and node.pid=? and node.status=1"
				result3 ,err := Orm.Query(sql,p,moduleid)
				if err != nil {
					log.Println(err)
				}
				for _,v3 := range result3{
					publicAction[string(v3["name"])] = string(v3["id"])
				}
				log.Println("here22 publicAction:",publicAction)
			}else{
				result22 = append(result22,v2)
			}

			//测试数据
			for k99,v99 := range result22{
				log.Println("here222: no publicAction",k99,string(v99["name"]),string(v99["id"]))
			}

		}


		// 依次读取模块的操作权限
		action  :=  make(map[string]string)
		hp      :=  make(map[string]string)
		for k3,v3 := range result22{
			moduleid := v3["id"]
			modulename := v3["name"]
			sql ="select node.id,node.name from chat_role as role," +
				"chat_role_user as user,chat_app_access as access ," +
				"chat_app_node as node where user.userid=? " +
				"and user.roleid=role.id and ( access.roleid=role.id  or (access.roleid=role.pid and role.pid!=0 ) ) " +
				"and role.status=1 and access.nodeid=node.id and node.level=3 and node.pid=? and node.status=1"
			result3 ,err := Orm.Query(sql,p,moduleid)
			if err != nil {
				log.Println(err)
			}
			log.Println("here4 k3 v3:",k3, "id:",string(moduleid),"name:",string(modulename))
			for  _,v4 := range result3{
				action[string(v4["name"])]=string(v4["id"])
			}
			//测试数据
			for k5,v5 := range action{
				hp[k5] = v5
			}
			for k6,v6 := range publicAction{
				hp[k6] = v6
			}

			// 和公共模块的操作权限合并
			access[strings.ToUpper(string(appname))] = make(map[string]interface{})
			access[strings.ToUpper(string(appname))][string(modulename)] = hp

		}
		log.Println("end",access,strings.ToUpper(string(appname)))

	}


	//return result
}