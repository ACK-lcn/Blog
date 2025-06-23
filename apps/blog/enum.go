package blog

type Status int

const (
	STATUS_DRAFT Status = iota
	STATUS_PUBLISHED
)

type UpdateMode int

const (
	// PUT （Full update）
	UPDATE_MODE_PUT UpdateMode = iota
	// PATCH （incremental update）
	UPDATE_MODE_PATCH
)
