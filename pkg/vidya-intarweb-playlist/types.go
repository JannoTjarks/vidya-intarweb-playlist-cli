package vidyaintarwebplaylist

type Roster struct {
	Changelog string `json:"changelog"`
	URL       string `json:"url"`
	Ext       string `json:"ext"`
	NewID     int    `json:"new_id"`
	Tracks    []track
}

type track struct {
	ID          int    `json:"id"`
	Game        string `json:"game"`
	Title       string `json:"title"`
	Comp        string `json:"comp"`
	Arr         string `json:"arr"`
	File        string `json:"file"`
	SourceID    int    `json:"s_id,omitempty"`
	SourceTitle string `json:"s_title,omitempty"`
	SourceFile  string `json:"s_file,omitempty"`
}
