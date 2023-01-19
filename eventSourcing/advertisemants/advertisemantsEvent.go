package advertisemants

const (
	AdvertisementPosted EventType = iota
	LikeAdded
	LikeRemoved
	ProfileVisited
	AdvertisementWatched
)

type Event struct {
	Type EventType
	Data []byte
}

type EventType int8
