package outgoing

import "github.com/di-wu/mtga/thread"

const (
	AuthenticateMethod      thread.LogMethod = "Authenticate"
	GetProductCatalogMethod thread.LogMethod = "PlayerInventory.GetProductCatalog"
	TrackDetailMethod       thread.LogMethod = "Quest.GetTrackDetail"

	AIPracticeMethod      thread.LogMethod = "Event.AIPractice"
	DeckSubmitMethod      thread.LogMethod = "Event.DeckSubmitV3"
	DropMethod            thread.LogMethod = "Event.Drop"
	GetPlayerCourseMethod thread.LogMethod = "Event.GetPlayerCourseV2"
	JoinMethod            thread.LogMethod = "Event.Join"
	JoinQueueMethod       thread.LogMethod = "Event.JoinQueue"
)
