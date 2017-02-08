package main

import (
	"github.com/mikespook/gorbac"
	"fmt"
)
type Data []byte
type DataFile interface {
	// 读取一个数据块。
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块。
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号。
	Rsn() int64
	// 获取最后写入的数据块的序列号。
	Wsn() int64
	// 获取数据块的长度
	DataLen() uint32
}

var rbac = gorbac.New()

type Tes struct {

}

func (x *Tes)init()  {
	fmt.Print("dddddddddddd")
}

func main()  {

	t := new(Tes)
	fmt.Print(t)

	rA := gorbac.NewStdRole("role-a") //角色
	rB := gorbac.NewStdRole("role-b")
	rC := gorbac.NewStdRole("role-c")
	rD := gorbac.NewStdRole("role-d")
	rE := gorbac.NewStdRole("role-e")

	pA := gorbac.NewStdPermission("permission-a")  //权限
	pA2 := gorbac.NewStdPermission("permission-a2")  //权限
	pB := gorbac.NewStdPermission("permission-b")
	pC := gorbac.NewStdPermission("permission-c")
	pD := gorbac.NewStdPermission("permission-d")
	pE := gorbac.NewStdPermission("permission-e")

	rA.Assign(pA) //给角色增加
	rA.Assign(pA2) //给角色增加
	rB.Assign(pB)
	rC.Assign(pC)
	rD.Assign(pD)
	rE.Assign(pE)

	rbac.Add(rA)
	rbac.Add(rB)
	rbac.Add(rC)
	rbac.Add(rD)
	rbac.Add(rE)


	rbac.SetParent("role-a", "role-b")
	rbac.SetParents("role-b", []string{"role-c", "role-d"})
	rbac.SetParent("role-e", "role-d")

	if rbac.IsGranted("role-a", pA, nil) &&
		rbac.IsGranted("role-a", pB, nil) &&
		rbac.IsGranted("role-a", pC, nil) &&
		rbac.IsGranted("role-a", pD, nil) {
		fmt.Println("The role-a has been granted permis-a, b, c and d.")
	}

	r,p,_ := rbac.Get("role-a")
	fmt.Println(r)
	fmt.Println(p)



}
