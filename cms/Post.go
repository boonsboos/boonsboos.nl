package cms

import (
	"encoding/json"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type PostData struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

type PostInfo struct {
	Info     PostData
	Location string
}

type Post struct {
	PostInfo
	Content template.HTML
}

// sliding window search to prevent the need to use
// split, which requires conversion to a string, then converting
// right back to a byte slice.
func findSeparatorIndices(rawPost []byte) (int, int) {

	if len(rawPost) < 4 {
		return 0, 0
	}

	for i := range rawPost {
		if i > len(rawPost)-4 {
			break
		}
		var current byte
		for _, v := range rawPost[i : i+4] {
			current += v
		}

		if current == 168 {
			return i, i + 4
		}
	}

	return 0, 0
}

func getAllPostInfo() []PostInfo {
	posts := []PostInfo{}

	filepath.WalkDir("./posts", func(path string, d fs.DirEntry, err error) error {

		if !strings.HasSuffix(d.Name(), ".md") {
			return err
		}

		file, er := os.Open(path)
		if er != nil {
			return err
		}

		buf, er := io.ReadAll(file)
		if er != nil {
			return err
		}

		file.Close()

		endOfData, startOfPost := findSeparatorIndices(buf)

		stuff := [][]byte{
			buf[0:endOfData],
			buf[startOfPost:],
		}

		info := PostData{}
		er = json.Unmarshal(stuff[0], &info)
		if er != nil {
			return err
		}

		posts = append(posts, PostInfo{info, strings.TrimSuffix(d.Name(), ".md")})

		return nil
	})

	// TODO: sort posts according to date

	return posts
}

func getAllPosts() []Post {
	posts := []Post{}

	filepath.WalkDir("./posts", func(path string, d fs.DirEntry, err error) error {

		if !strings.HasSuffix(d.Name(), ".md") {
			return err
		}

		file, er := os.Open(path)
		if er != nil {
			return err
		}

		buf, er := io.ReadAll(file)
		if er != nil {
			return err
		}

		endOfData, startOfPost := findSeparatorIndices(buf)

		stuff := [][]byte{
			buf[0:endOfData],
			buf[startOfPost:],
		}

		info := PostData{}
		er = json.Unmarshal(stuff[0], &info)
		if er != nil {
			return err
		}

		posts = append(
			posts,
			Post{
				PostInfo{
					info,
					strings.TrimSuffix(d.Name(), ".md"),
				},
				template.HTML(renderToMarkdown(stuff[1])),
			},
		)
		return nil
	})

	return posts
}

func renderToMarkdown(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	parser := parser.NewWithExtensions(extensions)

	flags := html.CommonFlags | html.HrefTargetBlank
	options := html.RendererOptions{
		Flags:          flags,
		RenderNodeHook: codeRenderHook, // cms/Markdown.go
	}
	renderer := html.NewRenderer(options)

	return markdown.ToHTML(md, parser, renderer)
}
