package mdhtml

import (
	"io/ioutil"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/shurcooL/go/github_flavored_markdown"
)

// GenerateHTML takes in a Markdown file and generates an HTML file
func GenerateHTML(sourceFile string) string {
	htmlHeader := `
<!DOCTYPE html>
<html>
  <head>

    <link href="css/reset.css" rel="stylesheet" />

    <meta charset="utf-8" />
    <meta name="viewport" content="width=1024" />
    <meta name="apple-mobile-web-app-capable" content="yes" />
    <link rel="shortcut icon" href="css/favicon.png" />
    <link rel="apple-touch-icon" href="css/apple-touch-icon.png" />
    <!-- Code Prettifier: -->
    <link href="css/highlight.css" type="text/css" rel="stylesheet" />
    <script type="text/javascript" src="js/highlight.pack.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
    <link href="css/style.css" rel="stylesheet" />
    <link href="http://fonts.googleapis.com/css?family=Lato:300,900" rel="stylesheet" />
  </head>
  <body>
    <div class="fallback-message">
      <p>Your browser <b>doesn't support the features required</b> by impress.js, so you are presented with a simplified version of this presentation.</p>
      <p>For the best experience please use the latest <b>Chrome</b>, <b>Safari</b> or <b>Firefox</b> browser.</p>
    </div>
  `

	htmlCSSstyle := `
    <style>
    .slide {
      color: #00786e;
    }
    h1 {
      color: orange;
    }
    </style>
    <div style="background-color: white; height: 100%;">
      <div>
        <img style="position: absolute; bottom: 0; width: 100%" src="http://i.imgur.com/QtxV5NQ.jpg" />
      </div>
    </div>
    <div id="impress">
      <div class='step slide' >
  `

	htmlFooter := `
      </div>
    </div>
    <script src="js/impress.js"></script>
    <script>impress().init();</script>
  </body>
</html>
  `

	sourceFileRead, _ := ioutil.ReadFile(sourceFile)
	markdownToHTML := blackfriday.MarkdownBasic(sourceFileRead)
	htmlFromMarkdown := string(bluemonday.UGCPolicy().SanitizeBytes(github_flavored_markdown.Markdown(markdownToHTML))[:])
	return htmlHeader + htmlCSSstyle + htmlFromMarkdown + htmlFooter
}