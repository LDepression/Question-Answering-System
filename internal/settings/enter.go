package settings

type group struct {
	Log   log
	Page  page
	Dao   mDao
	Sf    sf
	Email email
}

var Group = new(group)

func AllInit() {
	Group.Log.Init()
	Group.Page.Init()
	Group.Dao.Init()
	Group.Sf.Init()
	Group.Email.Init()
}
