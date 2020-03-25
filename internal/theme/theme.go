package theme

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/artdarek/go-unzip"
	"github.com/snipem/monako/internal/config"
)

const monakoMenuDirectory = "monako_menu_directory"
const themeName = "monako-book-master"

// CreateHugoPage fetches the Monako theme and copies the hugoconfig and menuconfig to the needed files
func CreateHugoPage(composeConfig config.ComposeConfig, menuconfig string) {

	dir := "compose/content/" + monakoMenuDirectory
	dst := dir + "/index.md"

	extractTheme()
	err := createHugoConfig(composeConfig)
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir(dir, os.FileMode(0744))

	data, err := ioutil.ReadFile(menuconfig)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func createHugoConfig(c config.ComposeConfig) error {
	configContent := fmt.Sprintf(`
baseURL = '%s'
title = '%s'
theme = '%s'

# Book configuration
disablePathToLower = true
enableGitInfo = true

# Needed for mermaid/katex shortcodes
[markup]
[markup.goldmark.renderer]
unsafe = true

[markup.tableOfContents]
startLevel = 1

[params]
# (Optional, default true) Controls table of contents visibility on right side of pages.
# Start and end levels can be controlled with markup.tableOfContents setting.
# You can also specify this parameter per page in front matter.
BookToC = true

# (Optional, default none) Set the path to a logo for the book. If the logo is
# /static/logo.png then the path would be logo.png
# BookLogo = 'logo.png'

# (Optional, default none) Set leaf bundle to render as side menu
# When not specified file structure and weights will be used
BookMenuBundle = '/%s'

# (Optional, default docs) Specify section of content to render as menu
# You can also set value to '*' to render all sections to menu
BookSection = 'docs'

# Set source repository location.
# Used for 'Last Modified' and 'Edit this page' links.
#BookRepo = 'https://github.com/alex-shpak/hugo-book'

# Enable "Edit this page" links for 'doc' page type.
# Disabled by default. Uncomment to enable. Requires 'BookRepo' param.
# Path must point to 'content' directory of repo.
#BookEditPath = 'edit/master/exampleSite/content'

# Configure the date format used on the pages
# - In git information
# - In blog posts
BookDateFormat = 'Jan 2, 2006'

# (Optional, default true) Enables search function with flexsearch,
# Index is built on fly, therefore it might slowdown your website.
# Configuration for indexing can be adjusted in i18n folder per language.
BookSearch = true

# (Optional, default true) Enables comments template on pages
# By default partals/docs/comments.html includes Disqus template
# See https://gohugo.io/content-management/comments/#configure-disqus
# Can be overwritten by same param in page frontmatter
BookComments = true
	`, c.BaseURL, c.Title, themeName, monakoMenuDirectory)
	return ioutil.WriteFile("compose/config.toml", []byte(configContent), os.FileMode(0700))
}

func extractTheme() {
	themezip, err := Asset("tmp/theme.zip")
	if err != nil {
		log.Fatalf("Error loading theme %s", err)
	}

	// TODO Don't use local filesystem, keep it in memory
	tmpFile, err := ioutil.TempFile(os.TempDir(), "monako-theme-")
	if err != nil {
		log.Fatalf("Cannot create temporary file %s", err)
	}
	tmpFile.Write(themezip)
	tempfilename := tmpFile.Name()

	if err != nil {
		log.Fatalf("Error writing temp theme %s", err)
	}

	// TODO Don't use a library that depends on local files
	uz := unzip.New(tempfilename, "compose/themes")
	err = uz.Extract()
	if err != nil {
		log.Fatalf("Error extracting theme: %s ", err)
	}
	os.RemoveAll(tempfilename)
}