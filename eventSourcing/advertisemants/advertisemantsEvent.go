package advertisemants

const (
	AdvertisementPosted EventType = iota
	LikeAdded
	LikeRemoved
	ProfileVisited
	AdvertisementWatched
	ReportGenerated
)

type Event struct {
	Type EventType
	Data []byte
}

type EventType int8
