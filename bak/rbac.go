package middleware

import (

	"log"
	. "common"
	"strings"
	"crypto/md5"
	"encoding/hex"

	"github.com/tango-contrib/session"
)

const(
	USER_AUTH_TYPE 		=   2			// 默认认证类型 1 登录认证 2 实时认证
	USER_AUTH_KEY 		= "USER_AUTH_KEY"
	ADMIN_AUTH_KEY 		= "ADMIN_AUTH_KEY"
	_ACCESS_LIST 		= "_ACCESS_LIST"
	USER_AUTH_ON 		= true
	NOT_AUTH_MODULE         = "Public"	// 默认无需认证模块
	REQUIRE_AUTH_MODULE     =  ""		// 默认需要认证模块
	NOT_AUTH_ACTION         =  ""		// 默认无需认证操作
	REQUIRE_AUTH_ACTION     =  ""		// 默认需要认证操作
)

type Rbac struct {
	Middler
}

//权限认证的过滤器方法
//@parm m moduleName
//@parm a actionName
func (x *Rbac)AccessDecision(m string, a string, s *session.Session)  bool{
	//存在认证识别号，则进行进一步的访问决策
	hash := md5.New()
	hash.Write([]byte(m+a))
	cipherText2 := hash.Sum(nil)
	hexText := make([]byte, 32)
	hex.Encode(hexText, cipherText2)
	accessGuid   := string(hexText)

	s.Set("aa","ddd")

	if x.CheckAccess(m,a)==true {
		accessLit := []string{}
		if USER_AUTH_TYPE==2{
			//加强验证和即时验证模式 更加安全 后台权限修改可以即时生效
			//通过数据库进行访问检查
			authid := s.Get(USER_AUTH_KEY).(int)
			if authid==0{
				authid = 1
			}
			accessLit = x.GetModuleAccessList( authid,m )
			accessLit = append(accessLit)

		}else{
			// 如果是管理员或者当前操作已经认证过，无需再次认证
			if s.Get( accessGuid ).(bool) {
				return true
			}
			//登录验证模式，比较登录后保存的权限访问列表
			accessLit := s.Get(_ACCESS_LIST)
			log.Println(accessLit)
			//accessLit = append(accessLit)
		}


	}

	return true
}


// 取得模块的所属记录访问权限列表 返回有权限的记录ID数组
func (x *Rbac)GetRecordAccessList(p int, m string, s *session.Session)  {
	if p==0 {
		p = s.Get(USER_AUTH_KEY) . (int)
	}
	if m=="" {
		panic("module no value")
	}

	access := x.GetModuleAccessList(p,m)
	log.Println(access)

}

//检查当前操作是否需要认证
func (x *Rbac)CheckAccess( moduleName string, actionName string ) bool{
	//如果项目要求认证，并且当前模块需要认证，则进行权限认证
	if USER_AUTH_ON {
		var module map[string][]string
		var action map[string][]string
		module = make(map[string][]string)
		action = make(map[string][]string)
		if REQUIRE_AUTH_MODULE!="" {
			module["yes"] = strings.Split(REQUIRE_AUTH_MODULE,",")
		}else{
			module["no"] = strings.Split(NOT_AUTH_MODULE,",")
		}

		//检查当前模块是否需要认证
		ok,_ := Contains(moduleName,module["no"])
		ok2,_ := Contains(moduleName,module["yes"])
		if (ok==false && module["no"]!=nil) ||  (ok2 && module["yes"]!=nil) {
			if REQUIRE_AUTH_ACTION!="" {
				//需要认证的操作
				action["yes"] = strings.Split(REQUIRE_AUTH_ACTION,",")
			}else{
				//无需认证的操作
				action["no"] = strings.Split(NOT_AUTH_ACTION,",")
			}

			//检查当前操作是否需要认证
			ok3,_ :=  Contains(actionName,action["no"])
			ok4,_ :=  Contains(actionName,action["yes"])
			if (ok3==false && action["no"]!=nil) ||  (ok4 && action["yes"]!=nil) {
				return true
			}else{
				return false
			}
		}

	}
	return false
}

//用于检测用户权限的方法,并保存到Session中
func (x *Rbac)SaveAccessList(p int,s *session.Session)  {
	if p == 0 {
		p = s.Get(USER_AUTH_KEY) . (int)
	}
	// 如果使用普通权限模式，保存当前用户的访问权限列表
	// 对管理员开发所有权限
	a := s.Get(USER_AUTH_KEY).(int)
	b := s.Get(ADMIN_AUTH_KEY).(int)
	if a != 2 &&  b !=0  {
		s.Set(_ACCESS_LIST,x.GetAccessList(p))
	}
}

/**
 +----------------------------------------------------------
 * 取得当前认证号的所有权限列表
 +----------------------------------------------------------
 * @param integer $authId 用户ID
 +----------------------------------------------------------
 * @access public
 +----------------------------------------------------------
 */
func (x *Rbac) GetAccessList( p int ) map[string]map[string]interface{}{

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
	for   _ ,v := range result{

		// 读取项目的模块权限
		appid := v["id"]
		appname := v["name"]
		//log.Println("here k v:",k, "id:",string(appid),"name:",string(appname))

		sql :=   "select node.id,node.name from chat_role as role," +
			"chat_role_user as user,chat_app_access as access ,chat_app_node as node " +
			"where user.userid=? and user.roleid=role.id " +
			"and ( access.roleid=role.id  or (access.roleid=role.pid and role.pid!=0 ) ) " +
			"and role.status=1 and access.nodeid=node.id and node.level=2 and node.pid=? and node.status=1"
		result2 ,err := Orm.Query(sql,p,string(appid))
		if err != nil {
			log.Println(err)
		}

		// 判断是否存在公共模块的权限
		publicAction  :=  make(map[string]string)
		var result22  []map[string][]byte
		for _,v2 := range result2{
			moduleid := v2["id"]
			modulename := v2["name"]
			//log.Println("here2 k2 v2:",k2, "id:",string(moduleid),"name:",string(modulename))
			if strings.ToUpper(string(modulename))=="PUBLIC" {

				sql := "select node.id,node.name from chat_role as role," +
				"chat_role_user as user,chat_app_access as access ,chat_app_node as node " +
					"where user.userid=? and user.roleid=role.id " +
					"and ( access.roleid=role.id  or (access.roleid=role.pid " +
					"and role.pid!=0 ) ) and role.status=1 and access.nodeid=node.id " +
					"and node.level=3 and node.pid=? and node.status=1"
				result3 ,err := Orm.Query(sql,p,string(moduleid))
				if err != nil {
					log.Println(err)
				}
				for _,v3 := range result3{
					publicAction[strings.ToUpper(string(v3["name"]))] = string(v3["id"])
				}
				//log.Println("here22 publicAction:",publicAction)
			}else{
				result22 = append(result22,v2)
			}

			//测试数据
			//for k99,v99 := range result22{
			//	log.Println("here222: no publicAction",k99,string(v99["name"]),string(v99["id"]))
			//}

		}

		log.Println("print:::result22",result22)

		// 依次读取模块的操作权限
		access[strings.ToUpper(string(appname))] = make(map[string]interface{})
		for _,v3 := range result22{
			action  :=  make(map[string]string)
			hp      :=  make(map[string]string)
			moduleid := v3["id"]
			modulename := v3["name"]
			sql ="select node.id,node.name from chat_role as role," +
				"chat_role_user as user,chat_app_access as access ," +
				"chat_app_node as node where user.userid=? " +
				"and user.roleid=role.id and ( access.roleid=role.id  or (access.roleid=role.pid and role.pid!=0 ) ) " +
				"and role.status=1 and access.nodeid=node.id and node.level=3 and node.pid=? and node.status=1"
			result3 ,err := Orm.Query(sql,p,string(moduleid))
			if err != nil {
				log.Println(err)
			}
			//log.Println("here4 k3 v3:",k3, "id:",string(moduleid),"name:",string(modulename))
			for  _,v4 := range result3{
				action[strings.ToUpper(string(v4["name"]))]=string(v4["id"])
			}
			//测试数据
			for k5,v5 := range action{
				hp[k5] = v5
			}
			for k6,v6 := range publicAction{
				hp[k6] = v6
			}
			//log.Println("here4 hp:",hp)
			// 和公共模块的操作权限合并
			access[strings.ToUpper(string(appname))][strings.ToUpper(string(modulename))] = hp

		}
		log.Println("end",access,strings.ToUpper(string(appname)))

	}

	return access
}


// 读取模块所属的记录访问权限
func (x *Rbac)GetModuleAccessList(p int,m string) []string  {
	sql := "select access.nodeid from chat_role as role,chat_role_user as user,cha_app_access as access "+
		"where user.userid=? and user.roleid=role.id " +
		"and ( access.roleid=role.id  or (access.roleid=role.pid and role.pid!=0 ) ) " +
		"and role.status=1 and  access.module=?"

	result,_ := Orm.Query(sql,p,m)
	access := []string{}
	for k,v := range result{
		access[k] = string(v["nodeid"])
	}
	log.Println("GetModuleAccessList:",result,access)
	return access
}