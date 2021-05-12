package wotoUD

import (
	"fmt"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

//---------------------------------------------------------

func (ud *udQuery) GetId() int {
	return ud.origin.id
}

//---------------------------------------------------------

// GetDef will give you the definition in the specified page,
// using wotoMD.
// WARNING: page arg is zero indexed.
func (c *UrbanCollection) GetDef(page uint8) (str string) {
	if len(c.List) == wv.BaseIndex {
		// TODO: return not found or something.
		return listNAvailble
	}

	//text = wotoMD.GetNormal(ud.List[0].Def).ToString()
	normal := wotoMD.GetNormal(c.List[page].Def)
	ex := wotoMD.GetItalic(wv.N_ESCAPE + wv.N_ESCAPE + c.List[page].Example)
	actual := page + wv.BaseOneIndex
	end := fmt.Sprintf(footnote, actual, len(c.List))
	final := normal.Append(ex).Append(wotoMD.GetNormal(end))
	return final.ToString()
}

//---------------------------------------------------------
