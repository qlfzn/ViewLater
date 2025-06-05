package middleware

import (
	"errors"
	"net/url"
	"strings"
)

const (
	OriginTiktok        = "TikTok"
	OriginInstagram     = "Instagram"
	OriginYoutubeShorts = "Youtube Shorts"
)

type Source struct {
	URL    *url.URL
	Origin string
}

func (s *Source) ParseAndValidateUrl(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	s.URL = u

	origin, err := s.checkOrigin(u.Host)
	if err != nil {
		return err
	}
	s.Origin = origin

	return nil
}

func (s *Source) checkOrigin(host string) (string, error) {
	host = strings.ToLower(host)

	switch {
	case strings.Contains(host, "tiktok"):
		return OriginTiktok, nil
	case strings.Contains(host, "instagram"):
		return OriginInstagram, nil
	case strings.Contains(host, "youtube"):
		return OriginYoutubeShorts, nil
	default:
		return "", errors.New("unsupported platform")
	}
}
