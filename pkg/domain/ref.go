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
