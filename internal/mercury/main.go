package mercury

import (
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

const source = "https://www.ismercuryinretrograde.com/"

var answerFinder = regexp.MustCompile(`https://www.ismercuryinretrograde.com/(\w+).html`)

func parse(node *html.Node) string {
	for n := node; n != nil; n = n.NextSibling {
		if n.Data == "meta" && len(n.Attr) >= 2 {
			if n.Attr[0].Key == "property" && n.Attr[0].Val == "og:url" {
				match := answerFinder.FindStringSubmatch(n.Attr[1].Val)
				return match[1]
			}
		}

		if n.FirstChild != nil {
			answer := parse(n.FirstChild)

			if len(answer) > 0 {
				return answer
			}
		}
	}

	return ""
}

type mercury struct{}

func (m *mercury) Retrograde() (bool, error) {
	resp, err := http.Get(source)

	if err != nil {
		return false, fmt.Errorf("error getting %v: %v", source, err)
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return false, fmt.Errorf("error parsing html: %v", err)
	}

	answer := parse(doc)

	return answer == "yes", err
}

func New() *mercury {
	return &mercury{}
}
