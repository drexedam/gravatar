package gravatar

// DefaultValue is a value suitable for the default parameter
type DefaultValue string

// See https://de.gravatar.com/site/implement/images/#default-image
const (
	NoDefault  DefaultValue = ""
	NotFound   DefaultValue = "404"
	MysteryMan DefaultValue = "mm"
	Identicon  DefaultValue = "identicon"
	Monsterid  DefaultValue = "monsterid"
	Wavatar    DefaultValue = "wavatar"
	Retro      DefaultValue = "retro"
	Blank      DefaultValue = "blank"
)
