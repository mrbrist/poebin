package utils

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// ---------------------------
// PathOfBuildingVersion
// ---------------------------
type PathOfBuildingVersion struct {
	PathOfExileOne *PathOfBuilding `xml:"PathOfBuilding,omitempty"`
	PathOfExileTwo *PathOfBuilding `xml:"PathOfBuilding2,omitempty"`
}

// ---------------------------
// PathOfBuilding
// ---------------------------
type PathOfBuilding struct {
	Build  Build  `xml:"Build"`
	Skills Skills `xml:"Skills"`
	Tree   Tree   `xml:"Tree"`
	Items  Items  `xml:"Items"`
	Notes  string `xml:"Notes,omitempty"`
	Config Config `xml:"Config"`
}

// ---------------------------
// Build
// ---------------------------
type Build struct {
	Level            uint8      `xml:"level,attr"`
	ClassName        string     `xml:"className,attr"`
	AscendClassName  *string    `xml:"ascendClassName,attr,omitempty"`
	Bandit           *string    `xml:"bandit,attr,omitempty"`
	PantheonMajorGod *string    `xml:"pantheonMajorGod,attr,omitempty"`
	PantheonMinorGod *string    `xml:"pantheonMinorGod,attr,omitempty"`
	MainSocketGroup  uint8      `xml:"mainSocketGroup,attr"`
	Stats            []StatType `xml:"PlayerStat"`
}

type StatType struct {
	Name  string `xml:"stat,attr"`
	Value string `xml:"value,attr"`
}

type BuildStat struct {
	Name  string `xml:"stat"`
	Value string `xml:"value"`
}

// ---------------------------
// Skills
// ---------------------------
type Skills struct {
	SkillSets []SkillSet `xml:"SkillSet"`
}

type SkillSet struct {
	Skills []Skill `xml:"Skill"`
}

type Skill struct {
	Enabled bool   `xml:"enabled,attr"`
	Slot    string `xml:"slot,attr"`
	Gems    []Gem  `xml:"Gem"`
}

func (s Skill) SupportGems() []Gem {
	var result []Gem
	for _, g := range s.Gems {
		if g.IsSupport() {
			result = append(result, g)
		}
	}
	return result
}

// ---------------------------
// Gem
// ---------------------------
type Gem struct {
	Name      string  `xml:"nameSpec"`
	SkillID   *string `xml:"skill_id,omitempty"`
	GemID     *string `xml:"gem_id,omitempty"`
	QualityID *string `xml:"quality_id,omitempty"`
	Enabled   bool    `xml:"enabled"`
	Level     uint8   `xml:"level"`
	Quality   uint8   `xml:"quality"`
}

func (g Gem) IsSupport() bool {
	if g.GemID != nil {
		return strings.HasPrefix(*g.GemID, "Metadata/Items/Gems/Support")
	}
	if g.SkillID != nil {
		return strings.HasPrefix(*g.SkillID, "Support") || strings.HasSuffix(*g.SkillID, "Support")
	}
	return strings.Contains(g.Name, "Support")
}

func (g Gem) IsActive() bool {
	return !g.IsSupport()
}

func (g Gem) IsVaal() bool {
	return strings.HasPrefix(g.Name, "Vaal ")
}

func (g Gem) NonVaalName() string {
	return strings.TrimPrefix(g.Name, "Vaal ")
}

func (g *Gem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Inner struct {
		Name      string  `xml:"nameSpec"`
		SkillID   *string `xml:"skill_id"`
		GemID     *string `xml:"gem_id"`
		QualityID *string `xml:"quality_id"`
		Enabled   *bool   `xml:"enabled"`
		Level     *uint8  `xml:"level"`
		Quality   *uint8  `xml:"quality"`
	}
	var inner Inner
	if err := d.DecodeElement(&inner, &start); err != nil {
		return err
	}
	g.Name = inner.Name
	if g.Name == "" && inner.SkillID != nil {
		g.Name = *inner.SkillID
	}
	g.SkillID = inner.SkillID
	g.GemID = inner.GemID
	g.QualityID = inner.QualityID
	g.Enabled = true
	if inner.Enabled != nil {
		g.Enabled = *inner.Enabled
	}
	if inner.Level != nil {
		g.Level = *inner.Level
	}
	if inner.Quality != nil {
		g.Quality = *inner.Quality
	}
	return nil
}

// ---------------------------
// Tree
// ---------------------------
type Tree struct {
	Specs []Spec `xml:"Spec"`
}

type Spec struct {
	Title   string   `xml:"title,attr"`
	Nodes   string   `xml:"nodes,attr"`
	Sockets []Socket `xml:"Sockets>Socket"`
}

type Socket struct {
	NodeId int     `xml:"nodeId,attr"`
	ItemId *uint16 `xml:"itemId,attr"` // pointer to allow "nil"
}

type Override struct {
	Name   string `xml:"dn,omitempty"`
	NodeID uint32 `xml:"nodeId,omitempty"`
	Effect string `xml:",chardata"`
}

// ---------------------------
// Items
// ---------------------------
type Items struct {
	ActiveItemSet *uint16         `xml:"active_item_set,omitempty"`
	ItemList      []Item          `xml:"Item"`
	ItemMap       map[uint16]Item `xml:"-"` // build after unmarshal
	ItemSets      []ItemSet       `xml:"ItemSet"`
}

func (i *Items) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias Items
	var alias Alias
	if err := d.DecodeElement(&alias, &start); err != nil {
		return err
	}
	i.ActiveItemSet = alias.ActiveItemSet
	i.ItemList = alias.ItemList
	i.ItemSets = alias.ItemSets
	i.ItemMap = make(map[uint16]Item, len(alias.ItemList))
	for _, item := range alias.ItemList {
		if item.ID != nil {
			i.ItemMap[*item.ID] = item
		}
	}
	return nil
}

type Item struct {
	ID       *uint16    `xml:"id,attr"` // pointer allows nil
	Content  string     `xml:",chardata"`
	ModRange []ModRange `xml:"ModRange"`
}

func (i *Item) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			if attr.Value == "nil" {
				i.ID = nil
			} else {
				var val uint16
				_, err := fmt.Sscanf(attr.Value, "%d", &val)
				if err != nil {
					return fmt.Errorf("invalid Item ID: %w", err)
				}
				i.ID = &val
			}
		}
	}
	type Alias Item
	var tmp Alias
	if err := d.DecodeElement(&tmp, &start); err != nil {
		return err
	}
	i.Content = tmp.Content
	i.ModRange = tmp.ModRange
	return nil
}

type ModRange struct {
	ID    int     `xml:"id,attr"`
	Range float64 `xml:"range,attr"`
}

// ---------------------------
// ItemSet & Gear
// ---------------------------
type ItemSet struct {
	ID    uint16  `xml:"id,attr"`
	Title *string `xml:"title,attr,omitempty"`
	Gear  Gear    `xml:"Gear"`
}

type Gear struct {
	Weapon1     *uint16 `xml:"Weapon 1,omitempty"`
	Weapon2     *uint16 `xml:"Weapon 2,omitempty"`
	Weapon1Swap *uint16 `xml:"Weapon 1 Swap,omitempty"`
	Weapon2Swap *uint16 `xml:"Weapon 2 Swap,omitempty"`
	Helmet      *uint16 `xml:"Helmet,omitempty"`
	BodyArmour  *uint16 `xml:"Body Armour,omitempty"`
	Gloves      *uint16 `xml:"Gloves,omitempty"`
	Boots       *uint16 `xml:"Boots,omitempty"`
	Amulet      *uint16 `xml:"Amulet,omitempty"`
	Ring1       *uint16 `xml:"Ring 1,omitempty"`
	Ring2       *uint16 `xml:"Ring 2,omitempty"`
	Belt        *uint16 `xml:"Belt,omitempty"`
	Flask1      *uint16 `xml:"Flask 1,omitempty"`
	Flask2      *uint16 `xml:"Flask 2,omitempty"`
	Flask3      *uint16 `xml:"Flask 3,omitempty"`
	Flask4      *uint16 `xml:"Flask 4,omitempty"`
	Flask5      *uint16 `xml:"Flask 5,omitempty"`
	Charm1      *uint16 `xml:"Charm 1,omitempty"`
	Charm2      *uint16 `xml:"Charm 2,omitempty"`
	Charm3      *uint16 `xml:"Charm 3,omitempty"`
	Sockets     []uint16
}

// UnmarshalXML for Gear handles "nil" IDs safely
func (g *Gear) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Slot struct {
		ItemID string `xml:"itemId"`
		Name   string `xml:"name"`
	}
	for {
		tok, err := d.Token()
		if err != nil {
			break
		}
		switch elem := tok.(type) {
		case xml.StartElement:
			if elem.Name.Local == "Slot" {
				var s Slot
				if err := d.DecodeElement(&s, &elem); err != nil {
					return err
				}
				if s.ItemID == "nil" || s.ItemID == "" {
					continue
				}
				var val uint16
				_, err := fmt.Sscanf(s.ItemID, "%d", &val)
				if err != nil {
					continue
				}
				switch s.Name {
				case "Weapon 1":
					g.Weapon1 = &val
				case "Weapon 2":
					g.Weapon2 = &val
				case "Weapon 1 Swap":
					g.Weapon1Swap = &val
				case "Weapon 2 Swap":
					g.Weapon2Swap = &val
				case "Helmet":
					g.Helmet = &val
				case "Body Armour":
					g.BodyArmour = &val
				case "Gloves":
					g.Gloves = &val
				case "Boots":
					g.Boots = &val
				case "Amulet":
					g.Amulet = &val
				case "Ring 1":
					g.Ring1 = &val
				case "Ring 2":
					g.Ring2 = &val
				case "Belt":
					g.Belt = &val
				case "Flask 1":
					g.Flask1 = &val
				case "Flask 2":
					g.Flask2 = &val
				case "Flask 3":
					g.Flask3 = &val
				case "Flask 4":
					g.Flask4 = &val
				case "Flask 5":
					g.Flask5 = &val
				case "Charm 1":
					g.Charm1 = &val
				case "Charm 2":
					g.Charm2 = &val
				case "Charm 3":
					g.Charm3 = &val
				default:
					g.Sockets = append(g.Sockets, val)
				}
			}
		}
	}
	return nil
}

// ---------------------------
// Config & Input
// ---------------------------
type Config struct {
	ActiveConfigSet *string     `xml:"activeConfigSet,omitempty"`
	ConfigSets      []ConfigSet `xml:"ConfigSet"`
	Input           []Input     `xml:"Input"`
}

type ConfigSet struct {
	ID    *string `xml:"id,omitempty"`
	Input []Input `xml:"Input"`
}

type Input struct {
	Name    string   `xml:"name"`
	String  *string  `xml:"string,omitempty"`
	Boolean *bool    `xml:"boolean,omitempty"`
	Number  *float32 `xml:"number,omitempty"`
}
