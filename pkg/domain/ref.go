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
	PersonRefType            = "c"
	PersonAssetType          = "Person"
	EquipTypeRefType         = "d"
	EquipTypeAssetType       = "Equip Type"
	PersonFavoriteRefType    = "f"
	PersonFavoriteAssetType  = "Person Favorite"
	TextRefRefType           = "h"
	TextRefAssetType         = "Text Ref"
	IssueRefType             = "ir"
	IssueAssetType           = "Issue"
	RuleRefType              = "l"
	RuleAssetType            = "Rule"
	NotificationRefType      = "m"
	NotificationAssetType    = "Notification"
	PersonAssetTreeRefType   = "n"
	PersonAssetTreeType      = "Person Asset Tree"
	ReportRefType            = "q"
	ReportAssetType          = "Report"
	ProfileRefType           = "r"
	ProfileAssetType         = "Profile"
	EventSubscriptionRefType = "sub"
	EventSubscriptionType    = "Event Subscription"
	TagRefType               = "t"
	TagAssetType             = "Tag"
	PointUridRefType         = "u"
	PointUridAssetType       = "Point Urid"
	BlobRefType              = "v"
	BlobAssetType            = "Blob"
	RoleGroupRefType         = "w"
	RoleGroupAssetType       = "Role Group"
	RoleGroupRoleRefType     = "x"
	RoleGroupRoleAssetType   = "Role Group Role"
	EquipTypeConfigRefType   = "y"
	EquipTypeConfigAssetType = "Equip Type Config"
	EventRefType             = "eh"
	EventAssetType           = "Event"
	SiteRefType              = "s"
	SiteAssetType            = "Site"
	SiteTable                = "site"
	EquipRefType             = "e"
	EquipAssetType           = "Equip"
	EquipTable               = "equip"
	PointRefType             = "p"
	PointAssetType           = "Point"
	PointTable               = "point"
	BranchRefType            = "branch"
	BranchAssetType          = "Branch"
	BranchTable              = "asset_tree_branch"
	TagRefRefType            = "g"
	TagRefAssetType          = "Tag Ref"
	TagRefTable              = "tag_ref"
	FormControlRefType       = "j"
	FormControlAssetType     = "Form Control"
	FormControlTable         = "form_control"
	DashboardRefType         = "a"
	DashboardAssetType       = "Dashboard"
	FormControlRefRefType    = "k"
	FormControlRefAssetType  = "Form Control Ref"
	alphabet                 = "abcdefghijklmnopqrstuvwxyz0123456789"
	defaultLength            = 16
	defaultSpacing           = 8
	defaultSpacingChar       = "-"
	defaultPrefixChar        = "."
	defaultPrefixMaxLength   = 50
)

var (
	assetTypeMap = map[string]string{
		PersonRefType:            PersonAssetType,
		EquipTypeRefType:         EquipTypeAssetType,
		PersonFavoriteRefType:    PersonFavoriteAssetType,
		TextRefRefType:           TextRefAssetType,
		IssueRefType:             IssueAssetType,
		RuleRefType:              RuleAssetType,
		NotificationRefType:      NotificationAssetType,
		PersonAssetTreeRefType:   PersonAssetTreeType,
		ReportRefType:            ReportAssetType,
		ProfileRefType:           ProfileAssetType,
		EventSubscriptionRefType: EventSubscriptionType,
		TagRefType:               TagAssetType,
		PointUridRefType:         PointUridAssetType,
		BlobRefType:              BlobAssetType,
		RoleGroupRefType:         RoleGroupAssetType,
		RoleGroupRoleRefType:     RoleGroupRoleAssetType,
		EquipTypeConfigRefType:   EquipTypeConfigAssetType,
		EventRefType:             EventAssetType,
		SiteRefType:              SiteAssetType,
		EquipRefType:             EquipAssetType,
		PointRefType:             PointAssetType,
		BranchRefType:            BranchAssetType,
		TagRefRefType:            TagRefAssetType,
		FormControlRefType:       FormControlAssetType,
		FormControlRefRefType:    FormControlRefAssetType,
		DashboardRefType:         DashboardAssetType,
	}
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

func GetAssetType(ref string) string {
	return assetTypeMap[strings.Split(ref, ".")[0]]
}
