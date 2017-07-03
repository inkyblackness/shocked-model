package model

// ResourceType is an enumeration of resource clusters.
type ResourceType uint16

const (
	// ResourceTypeMfdDataImages refers to the bitmaps used in the MFD data displays, such as logs.
	ResourceTypeMfdDataImages = ResourceType(0x0028)

	// ResourceTypeTrapMessages refers to the texts shown in the MFD based on actions.
	ResourceTypeTrapMessages = ResourceType(0x0867)
	// ResourceTypeWords refers to the texts of WORDS level objects.
	ResourceTypeWords = ResourceType(0x0868)
	// ResourceTypeLogCategories contains the category names of logs.
	ResourceTypeLogCategories = ResourceType(0x0870)
	// ResourceTypeScreenMessages contains the messages shown on screens.
	ResourceTypeScreenMessages = ResourceType(0x0877)
	// ResourceTypeAccessCardNames contains the names of the access cards.
	ResourceTypeAccessCardNames = ResourceType(0x0879)
	// ResourceTypeDataletMessages contains the short text fragments found in cyberspace.
	ResourceTypeDataletMessages = ResourceType(0x087A)

	// ResourceTypePaperTexts refers to the texts written on lose papers.
	ResourceTypePaperTexts = ResourceType(0x003C)
)
