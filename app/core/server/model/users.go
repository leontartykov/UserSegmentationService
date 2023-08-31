package model

type UsersWithNeedSegmentDb struct {
	UsersId []int
}

type UsersWithNeedSegmentServ struct {
	UsersId []int
}

func UsersWithNeedSegmentDbToServEntity(dbEntity UsersWithNeedSegmentDb) *UsersWithNeedSegmentServ {
	return &UsersWithNeedSegmentServ{
		UsersId: dbEntity.UsersId,
	}
}
