package user

type Role int

// Use the iota enumeration type to represent the role
const (
	// Creator
	ROLE_AUTHOR Role = iota
	// Reviewer
	ROLE_AUDITOR
	// System Administrator
	ROLE_ADMIN
)

type DescribeBy int

const (
	DESCRIBE_BY_ID DescribeBy = iota
	DESCRIBE_BY_USERNAME
)
