package model

import "time"

type Relative struct {
	UserId   string
	Type     RelativeType
	Families []Family
}

type Family struct {
	Surname   string
	Relatives []Relative
	Creator   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateFamilyRequestDto struct {
	Surname   string
	CreatorId string
}

type AddRelativesRequestDto struct {
	FamilyId  string
	Relatives []Relative
}

type RemoveRelativeRequestDto struct {
	FamilyId string
	Relative Relative
}

type GetRelativesRequestDto struct {
	FamilyId string
	Page     int64
	PageSize int64
}

type GetRelativesResponseDto struct {
	Relatives []Relative
}

type RelativeType int

const (
	UNKNOWN RelativeType = iota
	MOTHER
	FATHER
	BROTHER
	SISTER
	GRANDMOTHER
	GRANDFATHER
	AUNT
	UNCLE
	COUSIN_BROTHER
	COUSIN_SISTER
	GREAT_GRANDMOTHER
	GREAT_GRANDFATHER
	SECOND_COUSIN_BROTHER
	SECOND_COUSIN_SISTER
)

func (rt RelativeType) String() string {
	return [...]string{
		"UNKNOWN",
		"MOTHER",
		"FATHER",
		"BROTHER",
		"SISTER",
		"GRANDMOTHER",
		"GRANDFATHER",
		"AUNT",
		"UNCLE",
		"COUSIN_BROTHER",
		"COUSIN_SISTER",
		"GREAT_GRANDMOTHER",
		"GREAT_GRANDFATHER",
		"SECOND_COUSIN_BROTHER",
		"SECOND_COUSIN_SISTER",
	}[rt]
}
