package gravatar

// Rating is a value suitable for the rating parameter
type Rating string

// See https://de.gravatar.com/site/implement/images/#rating
const (
	NoRating Rating = ""
	G        Rating = "g"
	Pg       Rating = "pg"
	R        Rating = "r"
	X        Rating = "x"
)
