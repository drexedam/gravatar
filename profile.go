package gravatar

// Profiles is a wrapper for the response
type Profiles struct {
	Entry []Profile
}

// Profile represents a single profile
type Profile struct {
	ID                string   `json:"id"`
	Hash              string   `json:"hash"`
	RequestHash       string   `json:"requestHash"`
	ProfileURL        string   `json:"profileUrl"`
	PreferredUsername string   `json:"prefferedUsername"`
	ThumbnailURL      string   `json:"thumbnailUrl"`
	Photos            []Photo  `json:"photos"`
	Name              []string `json:"name"`
	DisplayName       string   `json:"displayName"`
	Urls              []string `json:"urls"`
}

// Photo represents a single photo entry
type Photo struct {
	Value string
	Type  string
}
