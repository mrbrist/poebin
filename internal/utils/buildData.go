package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// TODO: This needs cleaning up and fixing things that are not being unmarshalled correctly

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
	i.ItemMap = make(map[uint16]Item)
	i.ItemList = []Item{}
	i.ItemSets = []ItemSet{}

	for {
		tok, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch elem := tok.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "active_item_set":
				var active uint16
				if err := d.DecodeElement(&active, &elem); err == nil {
					i.ActiveItemSet = &active
				}
			case "Item":
				var item Item
				if err := d.DecodeElement(&item, &elem); err != nil {
					return err
				}
				i.ItemList = append(i.ItemList, item)
				i.ItemMap[*item.ID] = item
			case "ItemSet":
				var set ItemSet
				if err := d.DecodeElement(&set, &elem); err != nil {
					return err
				}
				i.ItemSets = append(i.ItemSets, set)
			}
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
	ID    uint16 `xml:"id,attr"`
	Title string `xml:"title,attr,omitempty"`
	Gear  Gear   `xml:"Gear"`
}

func (is *ItemSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			var id uint64
			if _, err := fmt.Sscanf(attr.Value, "%d", &id); err == nil {
				is.ID = uint16(id)
			}
		case "title":
			is.Title = attr.Value
		}
	}

	for {
		tok, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch elem := tok.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			case "Gear":
				if err := d.DecodeElement(&is.Gear, &elem); err != nil {
					return err
				}

			case "Slot":
				var slot struct {
					Name   string `xml:"name,attr"`
					ItemID string `xml:"itemId,attr"`
				}
				if err := d.DecodeElement(&slot, &elem); err != nil {
					return err
				}

				// Skip empty slots
				if slot.ItemID == "" || slot.ItemID == "0" {
					continue
				}

				// Assign to Gear
				switch slot.Name {
				case "Weapon 1":
					is.Gear.Weapon1 = slot.ItemID
				case "Weapon 2":
					is.Gear.Weapon2 = slot.ItemID
				case "Weapon 1 Swap":
					is.Gear.Weapon1Swap = slot.ItemID
				case "Weapon 2 Swap":
					is.Gear.Weapon2Swap = slot.ItemID
				case "Helmet":
					is.Gear.Helmet = slot.ItemID
				case "Body Armour":
					is.Gear.BodyArmour = slot.ItemID
				case "Gloves":
					is.Gear.Gloves = slot.ItemID
				case "Boots":
					is.Gear.Boots = slot.ItemID
				case "Amulet":
					is.Gear.Amulet = slot.ItemID
				case "Ring 1":
					is.Gear.Ring1 = slot.ItemID
				case "Ring 2":
					is.Gear.Ring2 = slot.ItemID
				case "Belt":
					is.Gear.Belt = slot.ItemID
				case "Flask 1":
					is.Gear.Flask1 = slot.ItemID
				case "Flask 2":
					is.Gear.Flask2 = slot.ItemID
				case "Flask 3":
					is.Gear.Flask3 = slot.ItemID
				case "Flask 4":
					is.Gear.Flask4 = slot.ItemID
				case "Flask 5":
					is.Gear.Flask5 = slot.ItemID
				case "Charm 1":
					is.Gear.Charm1 = slot.ItemID
				case "Charm 2":
					is.Gear.Charm2 = slot.ItemID
				case "Charm 3":
					is.Gear.Charm3 = slot.ItemID
				default:
					is.Gear.Sockets = append(is.Gear.Sockets, slot.ItemID)
				}

			case "SocketIdURL":
				var socket struct {
					ItemID string `xml:"itemPbURL,attr"` // you can adjust attr if needed
				}
				if err := d.DecodeElement(&socket, &elem); err == nil && socket.ItemID != "" {
					is.Gear.Sockets = append(is.Gear.Sockets, socket.ItemID)
				}
			}

		case xml.EndElement:
			if elem.Name.Local == start.Name.Local {
				return nil
			}
		}
	}

	return nil
}

type Gear struct {
	Weapon1     string   `xml:"Weapon 1,omitempty"`
	Weapon2     string   `xml:"Weapon 2,omitempty"`
	Weapon1Swap string   `xml:"Weapon 1 Swap,omitempty"`
	Weapon2Swap string   `xml:"Weapon 2 Swap,omitempty"`
	Helmet      string   `xml:"Helmet,omitempty"`
	BodyArmour  string   `xml:"Body Armour,omitempty"`
	Gloves      string   `xml:"Gloves,omitempty"`
	Boots       string   `xml:"Boots,omitempty"`
	Amulet      string   `xml:"Amulet,omitempty"`
	Ring1       string   `xml:"Ring 1,omitempty"`
	Ring2       string   `xml:"Ring 2,omitempty"`
	Belt        string   `xml:"Belt,omitempty"`
	Flask1      string   `xml:"Flask 1,omitempty"`
	Flask2      string   `xml:"Flask 2,omitempty"`
	Flask3      string   `xml:"Flask 3,omitempty"`
	Flask4      string   `xml:"Flask 4,omitempty"`
	Flask5      string   `xml:"Flask 5,omitempty"`
	Charm1      string   `xml:"Charm 1,omitempty"`
	Charm2      string   `xml:"Charm 2,omitempty"`
	Charm3      string   `xml:"Charm 3,omitempty"`
	Sockets     []string // generic sockets
}

func (g *Gear) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch elem := tok.(type) {
		case xml.StartElement:
			if elem.Name.Local == "Slot" {
				var slot struct {
					Name   string `xml:"name,attr"`
					ItemID string `xml:"itemId,attr"`
				}
				if err := d.DecodeElement(&slot, &elem); err != nil {
					return err
				}

				// Skip itemId="0" unless you want them stored
				if slot.ItemID == "" || slot.ItemID == "0" {
					continue
				}

				switch slot.Name {
				// Standard slots
				case "Weapon 1":
					g.Weapon1 = slot.ItemID
				case "Weapon 2":
					g.Weapon2 = slot.ItemID
				case "Weapon 1 Swap":
					g.Weapon1Swap = slot.ItemID
				case "Weapon 2 Swap":
					g.Weapon2Swap = slot.ItemID
				case "Helmet":
					g.Helmet = slot.ItemID
				case "Body Armour":
					g.BodyArmour = slot.ItemID
				case "Gloves":
					g.Gloves = slot.ItemID
				case "Boots":
					g.Boots = slot.ItemID
				case "Amulet":
					g.Amulet = slot.ItemID
				case "Ring 1":
					g.Ring1 = slot.ItemID
				case "Ring 2":
					g.Ring2 = slot.ItemID
				case "Belt":
					g.Belt = slot.ItemID
				case "Flask 1":
					g.Flask1 = slot.ItemID
				case "Flask 2":
					g.Flask2 = slot.ItemID
				case "Flask 3":
					g.Flask3 = slot.ItemID
				case "Flask 4":
					g.Flask4 = slot.ItemID
				case "Flask 5":
					g.Flask5 = slot.ItemID
				case "Charm 1":
					g.Charm1 = slot.ItemID
				case "Charm 2":
					g.Charm2 = slot.ItemID
				case "Charm 3":
					g.Charm3 = slot.ItemID
				default:
					// Everything else goes into Sockets (Abyssal Sockets, unknowns)
					g.Sockets = append(g.Sockets, slot.ItemID)
				}
			}

			// Optional: capture SocketIdURL elements as generic sockets
			if elem.Name.Local == "SocketIdURL" {
				var socket struct {
					ItemID string `xml:"itemPbURL,attr"` // sometimes itemId is in different attr
				}
				if err := d.DecodeElement(&socket, &elem); err == nil && socket.ItemID != "" {
					g.Sockets = append(g.Sockets, socket.ItemID)
				}
			}

		case xml.EndElement:
			if elem.Name.Local == start.Name.Local {
				return nil
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
