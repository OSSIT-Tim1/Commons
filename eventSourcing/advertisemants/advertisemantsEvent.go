package advertisemants

const (
	AdvertisementPosted int = iota
	LikeAdded
	LikeRemoved
	ProfileVisited
	AdvertisementWatched
	ReportGenerated
)

type Event struct {
	Type int    `json:"type"`
	Data []byte `json:"data"`
}

type EventType int
