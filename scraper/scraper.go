package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Ruler represents some kind of ruling person in Ancient Rome. Birth and Death are strings are
// given when known and applicable
type Ruler struct {
	Name    string
	Birth   string
	Death   string
	Offices []Office
}

// Office represents a single role a Ruler has had over a time period.
type Office struct {
	Office string
	Start  string
	End    string
}

// Converts Ruler into string the form "name, birth, death, (office, start, end), (office, start, end)".
func (r Ruler) toString() string {
	s := r.Name

	if r.Birth != "" {
		s = s + ", " + r.Birth
	}

	if r.Death != "" {
		s = s + ", " + r.Death
	}

	for _, o := range r.Offices {
		s = s + ", (" + o.toString() + ")"
	}

	return s
}

// Converts Office into string of the form "office, start, end"
func (o Office) toString() string {
	return o.Office + ", " + o.Start + ", " + o.End
}

// ConsulYear represents the Consuls that were present in a given year.
type ConsulYear struct {
	Consuls []string
	Year    int
}

var (
	e1  Ruler = Ruler{"Augustus", "63BCE", "14CE", []Office{{"Emperor", "31BCE", "14CE"}}}
	e2  Ruler = Ruler{"Tiberius", "42BCE", "37CE", []Office{{"Emperor", "14CE", "37CE"}}}
	e3  Ruler = Ruler{"Caligula", "12CE", "41CE", []Office{{"Emperor", "37CE", "41CE"}}}
	e4  Ruler = Ruler{"Claudius", "10BCE", "54CE", []Office{{"Emperor", "41CE", "54CE"}}}
	e5  Ruler = Ruler{"Nero", "37CE", "68CE", []Office{{"Emperor", "54CE", "68CE"}}}
	e6  Ruler = Ruler{"Otho", "3BCE", "69CE", []Office{{"Emperor", "68CE", "69CE"}}}
	e7  Ruler = Ruler{"Aulus Vitellius", "15BCE", "69CE", []Office{{"Emperor", "69CE", "69CE"}}}
	e8  Ruler = Ruler{"Vespasian", "9CE", "79CE", []Office{{"Emperor", "69CE", "79CE"}}}
	e9  Ruler = Ruler{"Titus", "39CE", "81CE", []Office{{"Emperor", "79CE", "81CE"}}}
	e10 Ruler = Ruler{"Domitian", "51CE", "96CE", []Office{{"Emperor", "81CE", "96CE"}}}
	e11 Ruler = Ruler{"Nerva", "30CE", "98CE", []Office{{"Emperor", "96CE", "98CE"}}}
	e12 Ruler = Ruler{"Trajan", "53CE", "117CE", []Office{{"Emperor", "98CE", "117CE"}}}
	e13 Ruler = Ruler{"Hadrian", "76CE", "138CE", []Office{{"Emperor", "117CE", "138CE"}}}
	e14 Ruler = Ruler{"Antoninus Pius", "86CE", "161CE", []Office{{"Emperor", "138CE", "161CE"}}}
	e15 Ruler = Ruler{"Marcus Aurelius", "121CE", "180CE", []Office{{"Emperor", "161CE", "180CE"}}}
	e16 Ruler = Ruler{"Lucius Verus", "130CE", "169CE", []Office{{"Emperor", "161CE", "169CE"}}}
	e17 Ruler = Ruler{"Commodus", "161CE", "192CE", []Office{{"Emperor", "177CE", "192CE"}}}
	e18 Ruler = Ruler{"Publius Helvius Pertinax", "126CE", "193CE", []Office{{"Emperor", "193CE", "193CE"}}}
	e19 Ruler = Ruler{"Marcus Didius Severus Julianus", "133CE", "193CE", []Office{{"Emperor", "193CE", "193CE"}}}
	e20 Ruler = Ruler{"Septimius Severus", "145CE", "211CE", []Office{{"Emperor", "193CE", "211CE"}}}
)

// Emperors is a list of the 20 emperors from 31 BCE to 211 CE.
var Emperors []Ruler = []Ruler{e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16, e17, e18, e19, e20}

// GetConsuls uses Colly to scrap a list of Consuls from
// https://en.wikipedia.org/wiki/List_of_Roman_consuls from 200BCE - 43CE.
func GetConsuls() []Ruler {

	c := colly.NewCollector()

	rulers := []Ruler{}

	// Get's the table for 200BCE - 100BCE
	c.OnHTML("#mw-content-text > div.mw-parser-output > table:nth-child(55)", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			// This is concaternated string of 'td' results
			elem := h.ChildText("td")

			// Split out each element by \n
			elems := strings.Split(elem, "\n")

			if len(elems) >= 3 {
				// Only accept years and not 'suff.'
				if year, err := strconv.Atoi(elems[0]); err == nil {
					consulYear := ConsulYear{elems[1:3], year}
					twoRulers := genConsulRuler(consulYear)
					rulers = append(rulers, twoRulers[0], twoRulers[1])
				}
			}
		})
	})

	// Get's the table for 100BCE - 43BCE
	c.OnHTML("#mw-content-text > div.mw-parser-output > table:nth-child(57)", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			// This is concaternated string of 'td' results
			elem := h.ChildText("td")

			// Split out each element by \n
			elems := strings.Split(elem, "\n")

			if len(elems) >= 3 {
				// Only accept years and not 'suff.' and stops at 46BCE
				year, err := strconv.Atoi(elems[0])
				if err == nil || year < 43 {
					consulYear := ConsulYear{elems[1:3], year}
					twoRulers := genConsulRuler(consulYear)
					rulers = append(rulers, twoRulers[0], twoRulers[1])
				}
			}
		})
	})

	c.Visit("https://en.wikipedia.org/wiki/List_of_Roman_consuls")

	return rulers
}

func genConsulRuler(c ConsulYear) []Ruler {
	var rulers []Ruler
	for _, consul := range c.Consuls {
		office := Office{"Consul of Rome", fmt.Sprint(c.Year), fmt.Sprint(c.Year)}
		ruler := Ruler{
			Name:    consul,
			Offices: []Office{office},
		}
		rulers = append(rulers, ruler)
	}
	return rulers
}
