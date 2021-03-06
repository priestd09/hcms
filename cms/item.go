package cms

import (
	"time"
)

type Item struct {
	// The project or workspace, an item belongs to
	ProjectID string

	// ID is an abstract id of an item, which may be autogenerated by the underlaying store.
	ID string

	// The collection, the Item belongs to
	Collection string

	// TechName is a unique technical name which can be used for technical referencing.
	// For products, this could be the SKU or it could be e.g. the filename for images.
	// Only limited characters are allowd: [a-zA-Z0-9\-_]
	TechName string

	Data

	// The revision id of this storable.
	// This is a number, which increases on each change of the
	// storable at least by one.
	Revision int

	// Time of creation.
	Created time.Time

	// Time of change.
	Changed time.Time

	// The username of the creator
	CreatedBy string

	// The username of the last editor
	ChangedBy string

	// A list of tags
	Tags []string

	// Flag, if the item is a build in item, which can only be changed by deployment and not
	// during runtime over the api
	BuildIn bool
}

type Meta struct {
	// The Mime Type, as in the HTTP-Content Type.
	MimeType string

	// An optional sub type to further describe the expected schema
	// for the Data Content.
	ItemType string

	// Flag if the part is published
	Publish bool

	// Time to start showing the item.
	PublishFrom time.Time

	// Time to end showing the item.
	PublishUntil time.Time
}
type Data struct {
	Meta

	// The actual content, depending on the type
	Content []byte

	Parts []Data
}

type TypeDescription struct {
	// The project or workspace, an item belongs to
	ProjectID string

	// ID is an abstract id
	ID string

	// The Name
	Name string

	// The json Schema for the core data of the type
	JsonSchema string

	// The UI Schema for the core data of the type
	UISchema string
}
