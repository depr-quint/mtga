package outgoing

import "github.com/di-wu/mtga/thread"

const (
	AuthenticateMethod      thread.LogMethod = "Authenticate"
	LogErrorMethod          thread.LogMethod = "Log.Error"
	GetProductCatalogMethod thread.LogMethod = "PlayerInventory.GetProductCatalog"
	TrackDetailMethod       thread.LogMethod = "Quest.GetTrackDetail"

	DeckSubmitMethod      thread.LogMethod = "Event.DeckSubmitV3"
	GetPlayerCourseMethod thread.LogMethod = "Event.GetPlayerCourseV2"
	JoinMethod            thread.LogMethod = "Event.Join"
	JoinQueueMethod       thread.LogMethod = "Event.JoinQueue"
)
