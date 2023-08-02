package domain

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	SiteRefType            = "s"
	SiteTable              = "site"
	EquipRefType           = "e"
	EquipTable             = "equip"
	PointRefType           = "p"
	PointTable             = "point"
	BranchRefType          = "branch"
	BranchTable            = "asset_tree_branch"
	TagRefType             = "g"
	TagRefTable            = "tag_ref"
	FormControlTable       = "j"
	FormControlRefTable    = "k"
	alphabet               = "abcdefghijklmnopqrstuvwxyz0123456789"
	defaultLength          = 16
	defaultSpacing         = 8
	defaultSpacingChar     = "-"
	defaultPrefixChar      = "."
	defaultPrefixMaxLength = 50
)

var pattern = regexp.MustCompile(`^([a-z]{1,8}\.[a-z0-9]{8}-[a-z0-9]{8}([:]\d+)?)$`)

type Ref struct {
	Value string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// BeginsWith returns true if the current Ref HasPrefix of the provided value otherwise, false
func (r *Ref) BeginsWith(prefix string) bool {
	return strings.HasPrefix(r.Value, prefix)
}

// Valid returns true if the current Ref has a valid pattern
func (r *Ref) Valid() bool {
	return pattern.MatchString(r.Value)
}

// GetPrefix returns the all
func (r *Ref) GetPrefix() string {
	if r.Valid() == false {
		return ""
	}

	return strings.Split(r.Value, ".")[0]
}

// String returns the value
func (r Ref) String() string {
	return r.Value
}

func randomChar() string {
	return string(alphabet[rand.Intn(len(alphabet))])
}

func RandomWithPrefix(prefix string) (string, error) {
	if len(prefix) < 1 || len(prefix) > defaultPrefixMaxLength {
		return "", fmt.Errorf("invalid prefix length")
	}
	return fmt.Sprintf("%s%s%s", prefix, defaultPrefixChar, random(defaultLength, defaultSpacing, defaultSpacingChar)), nil
}

// NewRef generates a new reference string with the given prefix. The prefix is truncated
// to a length of defaultPrefixMaxLength runes if it is longer. The remaining reference
// string is generated using the RandomWithPrefix function.
func NewRef(prefix string) string {
	i := 0
	n := defaultPrefixMaxLength
	for i < len(prefix) && n > 0 {
		_, size := utf8.DecodeRuneInString(prefix[i:])
		i += size
		n--
	}
	ref, _ := RandomWithPrefix(prefix[:i])
	return ref
}

func RandomWithoutPrefix() string {
	return random(defaultLength, defaultSpacing, defaultSpacingChar)
}

func random(length int, spacing int, spacerChar string) string {
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
