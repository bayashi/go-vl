package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"

	flag "github.com/spf13/pflag"
)

var (
	cmd = "vl"

	version     = ""
	installFrom = "Source"

	exitOK  int = 0
	exitErr int = 1
)

type options struct {
	noPager bool
	grep    []string
	grepRe  []*regexp.Regexp
	label   string
	labels  []string
}

func parseArgs() *options {
	o := &options{}

	var (
		flagHelp    bool
		flagVersion bool
	)

	flag.BoolVarP(&flagHelp, "help", "h", false, "Show help (This message) and exit")
	flag.BoolVarP(&flagVersion, "version", "v", false, "Show version and build info and exit")
	flag.BoolVarP(&o.noPager, "no-pager", "", false, "Output without pager")
	flag.StringArrayVarP(&o.grep, "grep", "g", []string{}, "Grep condition to filter lines")
	flag.StringVarP(&o.label, "label", "l", "", "Show only matching items of labels")

	flag.Parse()

	if flagHelp {
		putHelp(fmt.Sprintf("Version %s", getVersion()))
	}

	if flagVersion {
		putErr(versionDetails())
		os.Exit(exitOK)
	}

	for _, r := range o.grep {
		o.grepRe = append(o.grepRe, regexp.MustCompile(regexp.QuoteMeta(r)))
	}

	if strings.Contains(o.label, ",") {
		o.labels = strings.Split(o.label, ",")
	} else if o.label != "" {
		o.labels = append(o.labels, o.label)
	}

	return o
}

func versionDetails() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	compiler := runtime.Version()

	return fmt.Sprintf(
		"Version %s - %s.%s (compiled:%s, %s)",
		getVersion(),
		goos,
		goarch,
		compiler,
		installFrom,
	)
}

func getVersion() string {
	if version != "" {
		return version
	}
	i, ok := debug.ReadBuildInfo()
	if !ok {
		return "Unknown"
	}

	return i.Main.Version
}

func putErr(message ...interface{}) {
	fmt.Fprintln(os.Stderr, message...)
}

func putUsage() {
	putErr(fmt.Sprintf("Usage: some-command | %s OPTIONS", cmd))
}

func putHelp(message string) {
	putErr(message)
	putUsage()
	putErr("Options:")
	flag.PrintDefaults()
	os.Exit(exitOK)
}
