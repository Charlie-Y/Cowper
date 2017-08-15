package bibleweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Chapter struct {
	Auditid   string         `json:"auditid"`
	Label     string         `json:"label"`
	Text      string         `json:"text"`
	Chapter   string         `json:"chapter"`
	ID        string         `json:"id"`
	OsisEnd   string         `json:"osis_end"`
	Parent    BookRefWrap    `json:"parent"`
	Next      ChapterRefWrap `json:"next"`
	Previous  ChapterRefWrap `json:"previous"`
	Copyright string         `json:"copyright"`
}

type ChapterRef struct {
	Path string `json:"path"`
	Name string `json:"name"`
	ID   string `json:"id"`
}

type ChapterRefWrap struct {
	Chapter ChapterRef `json:"chapter"`
}

// ChapterResponse is the response for grabbing a chapter
type ChapterResponse struct {
	Response struct {
		Chapters []Chapter    `json:"chapters"`
		Meta     ResponseMeta `json:"meta"`
	} `json:"response"`
}

func (c *Chapter) GetNakedText() (string, error) {
	doc, err := html.Parse(strings.NewReader(c.Text))
	if err != nil {
		return "", err
	}

	str := ""
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {

			str += "\t" + n.Data + "\n"
		}

		// if n.Type == html.ElementNode && n.Data == "span" {
		// 	for _, a := range n.Attr {
		// 		if a.Key == "href" {
		// 			fmt.Println(a.Val)
		// 			str += a.Val
		// 			break
		// 		}
		// 	}
		// }
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	fmt.Println("GOOGO")

	return str, nil
}

// GetPsalm is a temporary fn,
// TODO make this a wrapper around GetChapter
func (a *API) GetPsalm(version Version, chapter int) (*Chapter, error) {
	// psalmURL := "https://bibles.org/v2/chapters/eng-ESV:Ps.1.js"
	psalmURL := fmt.Sprintf("https://bibles.org/v2/chapters/%s:Ps.%d.js", version, chapter)

	client := &http.Client{}
	req, err := http.NewRequest("GET", psalmURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(a.Key, "X")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	chapterResponse := &ChapterResponse{}
	err = json.NewDecoder(resp.Body).Decode(chapterResponse)
	if err != nil {
		return nil, err
	}
	// fmt.Println(chapterResponse)

	// chapterJSON, err := json.MarshalIndent(chapterResponse, "", "    ")
	// fmt.Println(string(chapterJSON))
	// if err != nil {
	// 	return nil, err
	// }

	if len(chapterResponse.Response.Chapters) != 1 {
		return nil, fmt.Errorf("chapterResponse has too many chapters")
	}
	return &chapterResponse.Response.Chapters[0], nil
}
