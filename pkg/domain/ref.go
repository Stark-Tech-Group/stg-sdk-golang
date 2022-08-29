package domain

import (
	"regexp"
	"strings"
)

const (
	SiteRefType   = "s"
	SiteTable     = "site"
	EquipRefType  = "e"
	EquipTable    = "equip"
	PointRefType  = "p"
	PointTable    = "point"
	BranchRefType = "branch"
	BranchTable   = "asset_tree_branch"
	TagRefType    = "g"
	TagRefTable   = "tag_ref"
	ALPHABET = "abcdefghijklmnopqrstuvwxyz0123456789"
	DEFAULT_LENGHT = 16
	DEFAULT_SPACEING = 8
	DEFAULT_SPACEING_CHAR = '-'
	DEFAULT_PREFIX_CHAR = '.'
)

var pattern = regexp.MustCompile(`^([a-z]{1,8}\.[a-z0-9]{8}-[a-z0-9]{8}([:]\d+)?)$`)

type Ref struct {
	Value string
}

//BeginsWith returns true if the current Ref HasPrefix of the provided value otherwise, false
func (r *Ref) BeginsWith(prefix string) bool {
	return strings.HasPrefix(r.Value, prefix)
}

//Valid returns true if the current Ref has a valid pattern
func (r *Ref) Valid() bool {
	return pattern.MatchString(r.Value)
}

//GetPrefix returns the all
func (r *Ref) GetPrefix() string {
	if r.Valid() == false {
		return ""
	}

	return strings.Split(r.Value, ".")[0]
}

//String returns the value
func (r Ref) String() string {
	return r.Value
}

//CreateRef creates and returns new ref for an asset
func (r Ref) CreateRef(assetType string) string {
	return r.Value
}
