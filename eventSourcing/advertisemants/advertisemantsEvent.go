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
	Type EventType `json:"type"`
	Data []byte    `json:"data"`
}

type EventType int
