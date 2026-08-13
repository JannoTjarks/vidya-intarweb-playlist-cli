package cmd

type argument struct {
	name      string
	shorthand string
	value     string
	usage     string
	envVar    string
}

var (
	rosterName = argument{
		name:   "roster-name",
		usage:  "The name of the chosen Roster",
		envVar: "VIP_ROSTER_NAME",
		value:  "vip",
	}
	destinationPath = argument{
		name:  "destination-path",
		usage: "The path to the destination, where the files should be downloaded to",
	}
)
