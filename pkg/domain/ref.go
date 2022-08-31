package domain

import (
	"fmt"
	"math/rand"
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
	Alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	DefaultLength = 16
	DefaultSpacing = 8
	DefaultSpacingChar = "-"
	DefaultPrefixChar = "."
	DefaultPrefixMaxLength = 50
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

func randomChar() string{
	return string(Alphabet[rand.Intn(len(Alphabet))])
}

func RandomWithPrefix(prefix string) (string, error){
	if len(prefix) < 1 ||  len(prefix) > DefaultPrefixMaxLength {
		return "", fmt.Errorf("invalid prefix length")
	}
	return fmt.Sprintf("%s%s%s", prefix, DefaultPrefixChar, random(DefaultLength, DefaultSpacing, DefaultSpacingChar)), nil
}

func RandomWithoutPrefix() string{
	return random(DefaultLength, DefaultSpacing, DefaultSpacingChar)
}

func random(length int, spacing int, spacerChar string) string{
	var sb strings.Builder
	spacer := 0
	for length > 0 {
		if spacer == spacing {
			sb.WriteString(spacerChar)
			spacer = 0
		}
		length--
		spacer++
		sb.WriteString(randomChar())
	}
	return sb.String()
}
