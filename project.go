package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Project is a unique build project that is bundled with a Nintendo Labo Kit.
type Project struct {
	Icon        *Image
	Image       *Image
	Name        string
	Screenshots []*Image
}

// NewProject is a constructor function that instantiates and returns a new Project struct pointer.
func NewProject(s *goquery.Selection) (*Project, error) {
	var (
		name        string
		screenshots []*Image
	)
	var (
		image, _ = NewImage(s.Find(".toycon-image picture img"))
		icon, _  = NewImage(s.Find(".toycon-icon img"))
	)
	nameSelection := s.Find(".toycon-icon p")
	name = strings.TrimSpace(nameSelection.Text())
	name = strings.ToUpper(name)
	screenshotsSelection := s.Find(".screenshot")
	screenshotsSelection.Each(func(i int, s *goquery.Selection) {
		screenshot, err := NewImage(s.Find("picture img"))
		if err != nil {
			return
		}
		screenshots = append(screenshots, screenshot)
	})
	project := Project{
		Icon:        icon,
		Image:       image,
		Name:        name,
		Screenshots: screenshots}
	return &project, nil
}
