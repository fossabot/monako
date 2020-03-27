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
const themeName = "monako-docsy-master"

// CreateHugoPage extracts the Monako theme and copies the hugoconfig and menuconfig to the needed files
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
	configContent := fmt.Sprintf(`# Autogenerated by Monako, do not edit
baseURL = '%s'
title = '%s'
theme = '%s'
monakomenu = '%s'

# Will give values to .Lastmod etc.
enableGitInfo = true

# Language settings
contentDir = "content/"
defaultContentLanguage = "en"
defaultContentLanguageInSubdir = true
# Useful when translating.
enableMissingTranslationPlaceholders = true

disableKinds = ["taxonomy", "taxonomyTerm"]

# Highlighting config
pygmentsCodeFences = true
pygmentsUseClasses = false
# Use the new Chroma Go highlighter in Hugo.
pygmentsUseClassic = false
#pygmentsOptions = "linenos=table"
# See https://help.farbox.com/pygments.html
pygmentsStyle = "tango"

 # First one is picked as the Twitter card image if not set on page.
 #images = ["images/project-illustration.png"]

# Configure how URLs look like per section.
[permalinks]
blog = "/:section/:year/:month/:day/:slug/"

[outputs]
	home = [ "HTML", "JSON" ]
	page = [ "HTML" ]

## Configuration for BlackFriday markdown parser: https://github.com/russross/blackfriday
[blackfriday]
plainIDAnchors = true
hrefTargetBlank = true
angledQuotes = false
latexDashes = true

# Image processing configuration.
[imaging]
resampleFilter = "CatmullRom"
quality = 75
anchor = "smart"

# Additional menu items

[[menu.main]]
    name = "Example Site"
    weight = 40
    url = "https://example.docsy.dev"
[[menu.main]]
    name = "GitHub"
    weight = 50
    url = "https://github.com/google/docsy/"

[services]
[services.googleAnalytics]
# Comment out the next line to disable GA tracking. Also disables the feature described in [params.ui.feedback].
# id = "UA-00000000-0"

# Language configuration

[languages]
[languages.en]
title = "Docsy"
description = "Docsy does docs"
languageName ="English"
# Weight used for sorting.
weight = 1

[markup]
  [markup.goldmark]
    [markup.goldmark.renderer]
      unsafe = true

# Everything below this are Site Params

[params]
copyright = "The Docsy Authors"
privacy_policy = "https://policies.google.com/privacy"

# Menu title if your navbar has a versions selector to access old versions of your site.
# This menu appears only if you have at least one [params.versions] set.
version_menu = "Releases"

# Flag used in the "version-banner" partial to decide whether to display a 
# banner on every page indicating that this is an archived version of the docs.
# Set this flag to "true" if you want to display the banner.
archived_version = false

# The version number for the version of the docs represented in this doc set.
# Used in the "version-banner" partial to display a version number for the 
# current doc set.
version = "0.0"

# A link to latest version of the docs. Used in the "version-banner" partial to
# point people to the main doc site.
url_latest_version = "https://example.com"

# Repository configuration (URLs for in-page links to opening issues and suggesting changes)
github_repo = "https://github.com/google/docsy"
# An optional link to a related project repo. For example, the sibling repository where your product code lives.
github_project_repo = "https://github.com/google/docsy"

# Specify a value here if your content directory is not in your repo's root directory
github_subdir = "userguide"

time_format_blog = "Monday, January 02, 2006"
time_format_default = "January 2, 2006"
# Sections to publish in the main RSS feed.
rss_sections = ["blog"]

# Google Custom Search Engine ID. Remove or comment out to disable search.
gcs_engine_id = "011217106833237091527:la2vtv2emlw"

# Enable Algolia DocSearch
algolia_docsearch = false

#Enable offline search with Lunr.js
offlineSearch = false

# User interface configuration
[params.ui]
# Enable to show the side bar menu in its compact state.
sidebar_menu_compact = false
#  Set to true to disable breadcrumb navigation.
breadcrumb_disable = false
#  Set to true to hide the sidebar search box (the top nav search box will still be displayed if search is enabled)
sidebar_search_disable = false
#  Set to false if you don't want to display a logo (/assets/icons/logo.svg) in the top nav bar
navbar_logo = true

# Adds a H2 section titled "Feedback" to the bottom of each doc. The responses are sent to Google Analytics as events.
# This feature depends on [services.googleAnalytics] and will be disabled if "services.googleAnalytics.id" is not set.
# If you want this feature, but occasionally need to remove the "Feedback" section from a single page,
# add "hide_feedback: true" to the page's front matter.
[params.ui.feedback]
enable = true
# The responses that the user sees after clicking "yes" (the page was helpful) or "no" (the page was not helpful).
yes = 'Glad to hear it! Please <a href="https://github.com/USERNAME/REPOSITORY/issues/new">tell us how we can improve</a>.'
no = 'Sorry to hear that. Please <a href="https://github.com/USERNAME/REPOSITORY/issues/new">tell us how we can improve</a>.'

[params.links]
# End user relevant links. These will show up on left side of footer and in the community page if you have one.
 [[params.links.user]]
	name = "User mailing list"
	url = "https://groups.google.com/forum/#!forum/docsy-users"
	icon = "fa fa-envelope"
        desc = "Discussion and help from your fellow users"
[[params.links.user]]
	name ="Twitter"
	url = "https://twitter.com/docsydocs"
	icon = "fab fa-twitter"
        desc = "Follow us on Twitter to get the latest news!"
# [[params.links.user]]
	# name = "Stack Overflow"
	# url = "https://example.org/stack"
	# icon = "fab fa-stack-overflow"
        # desc = "Practical questions and curated answers"
# Developer relevant links. These will show up on right side of footer and in the community page if you have one.
[[params.links.developer]]
	name = "GitHub"
	url = "https://github.com/google/docsy"
	icon = "fab fa-github"
        desc = "Development takes place here!"
# [[params.links.developer]]
	# name = "Slack"
	# url = "https://example.org/slack"
	# icon = "fab fa-slack"
        # desc = "Chat with other project developers"
# [[params.links.developer]]
	# name = "Developer mailing list"
	# url = "https://example.org/mail"
	# icon = "fa fa-envelope"
        # desc = "Discuss development issues around the project"

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
