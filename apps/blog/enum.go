package blog

type Status int

const (
	//Article status draft
	STATUS_DRAFT Status = iota
	// Already published
	STATUS_PUBLISHED
)

type UpdateMode int

const (
	// PUT （Full update）
	UPDATE_MODE_PUT UpdateMode = iota
	// PATCH （incremental update）
	UPDATE_MODE_PATCH
)
