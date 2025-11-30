package utils

import "encoding/xml"

type BuildData struct {
	XMLName xml.Name `xml:"PathOfBuilding"`
	Text    string   `xml:",chardata"`
	Build   struct {
		Text                   string `xml:",chardata"`
		ViewMode               string `xml:"viewMode,attr"`
		AscendClassName        string `xml:"ascendClassName,attr"`
		Bandit                 string `xml:"bandit,attr"`
		PantheonMinorGod       string `xml:"pantheonMinorGod,attr"`
		ClassName              string `xml:"className,attr"`
		CharacterLevelAutoMode string `xml:"characterLevelAutoMode,attr"`
		Level                  string `xml:"level,attr"`
		MainSocketGroup        string `xml:"mainSocketGroup,attr"`
		TargetVersion          string `xml:"targetVersion,attr"`
		PantheonMajorGod       string `xml:"pantheonMajorGod,attr"`
		PlayerStat             []struct {
			Text  string `xml:",chardata"`
			Stat  string `xml:"stat,attr"`
			Value string `xml:"value,attr"`
		} `xml:"PlayerStat"`
		TimelessData struct {
			Text               string `xml:",chardata"`
			SearchList         string `xml:"searchList,attr"`
			SearchListFallback string `xml:"searchListFallback,attr"`
			DevotionVariant1   string `xml:"devotionVariant1,attr"`
			DevotionVariant2   string `xml:"devotionVariant2,attr"`
		} `xml:"TimelessData"`
	} `xml:"Build"`
	Calcs struct {
		Text  string `xml:",chardata"`
		Input []struct {
			Text   string `xml:",chardata"`
			String string `xml:"string,attr"`
			Name   string `xml:"name,attr"`
			Number string `xml:"number,attr"`
		} `xml:"Input"`
		Section []struct {
			Text       string `xml:",chardata"`
			Collapsed  string `xml:"collapsed,attr"`
			Subsection string `xml:"subsection,attr"`
			ID         string `xml:"id,attr"`
		} `xml:"Section"`
	} `xml:"Calcs"`
	Notes string `xml:"Notes"`
	Tree  struct {
		Text       string `xml:",chardata"`
		ActiveSpec string `xml:"activeSpec,attr"`
		Spec       struct {
			Text                   string `xml:",chardata"`
			Nodes                  string `xml:"nodes,attr"`
			TreeVersion            string `xml:"treeVersion,attr"`
			AscendClassId          string `xml:"ascendClassId,attr"`
			MasteryEffects         string `xml:"masteryEffects,attr"`
			ClassId                string `xml:"classId,attr"`
			SecondaryAscendClassId string `xml:"secondaryAscendClassId,attr"`
			URL                    string `xml:"URL"`
			Sockets                struct {
				Text   string `xml:",chardata"`
				Socket []struct {
					Text   string `xml:",chardata"`
					ItemId string `xml:"itemId,attr"`
					NodeId string `xml:"nodeId,attr"`
				} `xml:"Socket"`
			} `xml:"Sockets"`
			Overrides struct {
				Text     string `xml:",chardata"`
				Override []struct {
					Text              string `xml:",chardata"`
					Dn                string `xml:"dn,attr"`
					NodeId            string `xml:"nodeId,attr"`
					ActiveEffectImage string `xml:"activeEffectImage,attr"`
					Icon              string `xml:"icon,attr"`
				} `xml:"Override"`
			} `xml:"Overrides"`
		} `xml:"Spec"`
	} `xml:"Tree"`
	Config struct {
		Text            string `xml:",chardata"`
		ActiveConfigSet string `xml:"activeConfigSet,attr"`
		ConfigSet       struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id,attr"`
			Input []struct {
				Text    string `xml:",chardata"`
				Number  string `xml:"number,attr"`
				Name    string `xml:"name,attr"`
				String  string `xml:"string,attr"`
				Boolean string `xml:"boolean,attr"`
			} `xml:"Input"`
			Placeholder []struct {
				Text   string `xml:",chardata"`
				Number string `xml:"number,attr"`
				Name   string `xml:"name,attr"`
			} `xml:"Placeholder"`
		} `xml:"ConfigSet"`
	} `xml:"Config"`
	Skills struct {
		Text                string `xml:",chardata"`
		SortGemsByDPS       string `xml:"sortGemsByDPS,attr"`
		DefaultGemLevel     string `xml:"defaultGemLevel,attr"`
		ActiveSkillSet      string `xml:"activeSkillSet,attr"`
		SortGemsByDPSField  string `xml:"sortGemsByDPSField,attr"`
		ShowAltQualityGems  string `xml:"showAltQualityGems,attr"`
		ShowSupportGemTypes string `xml:"showSupportGemTypes,attr"`
		DefaultGemQuality   string `xml:"defaultGemQuality,attr"`
		SkillSet            struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id,attr"`
			Skill []struct {
				Text                 string `xml:",chardata"`
				Slot                 string `xml:"slot,attr"`
				Enabled              string `xml:"enabled,attr"`
				MainActiveSkillCalcs string `xml:"mainActiveSkillCalcs,attr"`
				MainActiveSkill      string `xml:"mainActiveSkill,attr"`
				IncludeInFullDPS     string `xml:"includeInFullDPS,attr"`
				Label                string `xml:"label,attr"`
				Gem                  []struct {
					Text           string `xml:",chardata"`
					QualityId      string `xml:"qualityId,attr"`
					Count          string `xml:"count,attr"`
					GemId          string `xml:"gemId,attr"`
					Quality        string `xml:"quality,attr"`
					SkillId        string `xml:"skillId,attr"`
					NameSpec       string `xml:"nameSpec,attr"`
					Enabled        string `xml:"enabled,attr"`
					VariantId      string `xml:"variantId,attr"`
					SkillPartCalcs string `xml:"skillPartCalcs,attr"`
					SkillPart      string `xml:"skillPart,attr"`
					EnableGlobal1  string `xml:"enableGlobal1,attr"`
					Level          string `xml:"level,attr"`
					EnableGlobal2  string `xml:"enableGlobal2,attr"`
				} `xml:"Gem"`
			} `xml:"Skill"`
		} `xml:"SkillSet"`
	} `xml:"Skills"`
	Import struct {
		Text              string `xml:",chardata"`
		LastAccountHash   string `xml:"lastAccountHash,attr"`
		LastCharacterHash string `xml:"lastCharacterHash,attr"`
		LastRealm         string `xml:"lastRealm,attr"`
		ExportParty       string `xml:"exportParty,attr"`
	} `xml:"Import"`
	Items struct {
		Text                string `xml:",chardata"`
		ActiveItemSet       string `xml:"activeItemSet,attr"`
		ShowStatDifferences string `xml:"showStatDifferences,attr"`
		UseSecondWeaponSet  string `xml:"useSecondWeaponSet,attr"`
		Item                []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			ModRange []struct {
				Text  string `xml:",chardata"`
				ID    string `xml:"id,attr"`
				Range string `xml:"range,attr"`
			} `xml:"ModRange"`
		} `xml:"Item"`
		ItemSet struct {
			Text               string `xml:",chardata"`
			ID                 string `xml:"id,attr"`
			UseSecondWeaponSet string `xml:"useSecondWeaponSet,attr"`
			SocketIdURL        []struct {
				Text      string `xml:",chardata"`
				ItemPbURL string `xml:"itemPbURL,attr"`
				Name      string `xml:"name,attr"`
				NodeId    string `xml:"nodeId,attr"`
			} `xml:"SocketIdURL"`
			Slot []struct {
				Text      string `xml:",chardata"`
				ItemId    string `xml:"itemId,attr"`
				Name      string `xml:"name,attr"`
				ItemPbURL string `xml:"itemPbURL,attr"`
				Active    string `xml:"active,attr"`
			} `xml:"Slot"`
		} `xml:"ItemSet"`
	} `xml:"Items"`
	Party struct {
		Text             string `xml:",chardata"`
		Append           string `xml:"append,attr"`
		Destination      string `xml:"destination,attr"`
		ShowAdvanceTools string `xml:"ShowAdvanceTools,attr"`
	} `xml:"Party"`
	TreeView struct {
		Text                string `xml:",chardata"`
		ZoomLevel           string `xml:"zoomLevel,attr"`
		ZoomY               string `xml:"zoomY,attr"`
		ZoomX               string `xml:"zoomX,attr"`
		SearchStr           string `xml:"searchStr,attr"`
		ShowStatDifferences string `xml:"showStatDifferences,attr"`
	} `xml:"TreeView"`
}
