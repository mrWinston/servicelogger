package templates

import (
	"fmt"
	"github.com/geowa4/servicelogger/pkg/config"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	managedNotificationsDirName = "managed-notifications"
	opsSOPDirName               = "ops-sop"
)

var (
	templateVarRegexp = regexp.MustCompile("\\$\\{[A-Z0-9_]+}")
)

func GetTemplateVarRegexp() *regexp.Regexp {
	return templateVarRegexp
}

type Template struct {
	Severity      string   `json:"severity"`
	ServiceName   string   `json:"service_name"`
	Summary       string   `json:"summary"`
	Description   string   `json:"description"`
	InternalOnly  bool     `json:"internal_only"`
	EventStreamId string   `json:"event_stream_id,omitempty"`
	Tags          []string `json:"_tags,omitempty"`
	SourcePath    string   `json:"-"`
}

func (t *Template) String() string {
	md := fmt.Sprintf(
		"# %s\n\n%s",
		t.Summary,
		t.Description,
	)
	if len(t.Tags) > 0 {
		md += fmt.Sprintf("\n\n_Tags_: %s", strings.Join(t.Tags, ", "))
	}
	return templateVarRegexp.ReplaceAllStringFunc(md, func(match string) string {
		return fmt.Sprintf("*%s*", match)
	})
}

func GetRelativePathForManagedNotifications(filePath string) string {
	split := strings.Split(filePath, string(os.PathSeparator)+managedNotificationsDirName+string(os.PathSeparator))
	re := regexp.MustCompile("^.*master" + string(os.PathSeparator))
	return re.ReplaceAllString(split[1], "")
}

// GetServiceLogTemplatesDir returns the directory to use to find all templates
// and returns empty string if there was an unlikely error
func GetServiceLogTemplatesDir() string {
	cloneTarget, err := config.GetCacheDir(managedNotificationsDirName)
	if err != nil {
		return ""
	}
	return cloneTarget
}

// GetOsdServiceLogTemplatesDir returns the directory to use to find templates for OSD
// and returns empty string if there was an unlikely error
func GetOsdServiceLogTemplatesDir() string {
	return filepath.Join(GetServiceLogTemplatesDir(), "osd")
}

// GetOpsSOPDir returns the directory to use to find Ops SOPs
// and returns empty string if there was an unlikely error
func GetOpsSOPDir() string {
	cloneTarget, err := config.GetCacheDir(opsSOPDirName)
	if err != nil {
		return ""
	}
	return cloneTarget
}
