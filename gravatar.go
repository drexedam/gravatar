package gravatar

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Gravatar is a struct for configuring and generating a
// Gravatar-URL
type Gravatar struct {
	hash         string
	defaultURL   string
	defaultValue string
	size         int
	forceDefault bool
	rating       string
}

// New creates a new Gravatar instance based
// on the given E-Mail
func New(email string) *Gravatar {
	hash := md5.Sum([]byte(email))
	return &Gravatar{hash: fmt.Sprintf("%x", hash)}
}

// URL generates the URL to the Gravatar profile of the given email
func (g *Gravatar) URL() string {
	return "https://www.gravatar.com/" + g.hash
}

// AvatarURL generates the URL to get the avatar of the user
// based on the given configuration
func (g *Gravatar) AvatarURL() string {
	url := "https://www.gravatar.com/avatar/" + g.hash

	if g.forceDefault {
		url = g.addParameter(url, "f", "y")
	}

	if g.defaultURL != "" {
		url = g.addParameter(url, "d", g.defaultURL)
	} else if g.defaultValue != "" {
		url = g.addParameter(url, "d", g.defaultValue)
	}

	if g.rating != "" {
		url = g.addParameter(url, "r", g.rating)
	}

	if g.size > 0 {
		url = g.addParameter(url, "s", strconv.Itoa(g.size))
	}

	return url
}

// JSONURL generates the URL to fetch profile data as json
func (g *Gravatar) JSONURL() string {
	return "https://www.gravatar.com/" + g.hash + ".json"
}

// JSONURLCallback returns the URL to fetch profile data as json
// and sets the callback parameter (See https://de.gravatar.com/site/implement/profiles/json/#request-options)
func (g *Gravatar) JSONURLCallback(callback string) string {
	return g.addParameter(g.JSONURL(), "callback", callback)
}

// Size sets the size of the requested image
// If size is zero the parameter is not used
// Valid sizes are from 1px up to 2048px
func (g *Gravatar) Size(size int) *Gravatar {
	if size > 0 && size < 2049 {
		g.size = size
	} else {
		g.size = 0
	}
	return g
}

// DefaultURL sets a URL to use as default
// image (See https://de.gravatar.com/site/implement/images/#default-image)
// An invalid URL will be ignored
func (g *Gravatar) DefaultURL(urlString string) *Gravatar {
	u, err := url.Parse(urlString)
	if err == nil { // Invalid urls will be ignored
		g.defaultURL = u.String()
	}

	return g
}

// Default sets a default value to be used if no
// image is available
func (g *Gravatar) Default(value DefaultValue) *Gravatar {
	g.defaultValue = string(value)
	return g
}

// Rating sets the rating appropriate for your audience
func (g *Gravatar) Rating(rating Rating) *Gravatar {
	g.rating = string(rating)
	return g
}

// ForceDefault sets if the default avatar should be forced o be returned
func (g *Gravatar) ForceDefault(force bool) *Gravatar {
	g.forceDefault = force
	return g
}

// Profiles fetches and parses the profile data
func (g *Gravatar) Profiles() (*Profiles, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(g.JSONURL())
	if err != nil {
		return nil, err
	}

	defer close(r.Body)

	profiles := &Profiles{}
	err = json.NewDecoder(r.Body).Decode(profiles)
	return profiles, err
}

func (g *Gravatar) addParameter(url, key, value string) string {
	if strings.HasSuffix(url, g.hash) || strings.HasSuffix(url, ".json") {
		url = url + "?"
	} else {
		url = url + "&"
	}

	return url + key + "=" + value
}

func close(c io.Closer) {
	_ = c.Close()
}
