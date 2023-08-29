package model

type SegmentServiceModel struct {
	Segments []string
}

type SegmentDbEntity struct {
	Segments []string
}

func SegEntityToModel(seg SegmentDbEntity) *SegmentServiceModel {
	return &SegmentServiceModel{
		Segments: seg.Segments,
	}
}

type ServChangedSegments struct {
	To_add    []string
	To_delete []string
	User_id   string
}

type DbChangedSegments struct {
	To_add    []string
	To_delete []string
	User_id   string
}

func ChangeSegsModelToEntity(segs ServChangedSegments) *DbChangedSegments {
	return &DbChangedSegments{
		To_add:    segs.To_add,
		To_delete: segs.To_delete,
		User_id:   segs.User_id,
	}
}
