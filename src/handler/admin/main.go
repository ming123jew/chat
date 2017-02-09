package handler

type AdminMain struct {
	AdminHandler
}

func (x *AdminMain)Get()  {
	x.Ctx.ResponseWriter.Write( []byte("ok"))
}

func (x *AdminMain)Post() ()  {

}