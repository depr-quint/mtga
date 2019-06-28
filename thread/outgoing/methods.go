package outgoing

import "github.com/di-wu/mtga/thread"

const (
	AuthenticateMethod      thread.LogMethod = "Authenticate"
	LogErrorMethod          thread.LogMethod = "Log.Error"
	GetProductCatalogMethod thread.LogMethod = "PlayerInventory.GetProductCatalog"
	TrackDetailMethod       thread.LogMethod = "Quest.GetTrackDetail"

	GetPlayerCourseMethod thread.LogMethod = "Event.GetPlayerCourseV2"
	JoinQueueMethod       thread.LogMethod = "Event.JoinQueue"
)
