package scraper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Ruler represents some kind of ruling person in Ancient Rome. The inteval the ruler lived
// is given in an ISO 8601 date range. Mostly only the year is given.
type Ruler struct {
	Name    string
	Lived   string
	Offices []Office
}

// Office represents a single role a Ruler has had over a time period. The inteval that the
// Ruler held the Office is given in ISO 8601 date range.
type Office struct {
	Office string
	Held   string
}

// Converts Ruler into string the form "name, birth/death, (office, start/end), (office, start/end)".
func (r Ruler) toString() string {
	s := r.Name

	if r.Lived != "" {
		s = s + ", " + r.Lived
	}

	for _, o := range r.Offices {
		s = s + ", (" + o.toString() + ")"
	}

	return s
}

// Converts Office into string of the form "office, start, end".
func (o Office) toString() string {
	return o.Office + ", " + o.Held
}

// ConsulYear represents the Consuls that were present in a given year.
type ConsulYear struct {
	Consuls []string
	Year    int
}

var (
	e1  Ruler = Ruler{"Augustus", "-0064/0014", []Office{{"Emperor", "-0032/0014"}}}
	e2  Ruler = Ruler{"Tiberius", "-0042/0037", []Office{{"Emperor", "0014/0037"}}}
	e3  Ruler = Ruler{"Caligula", "0012/0041", []Office{{"Emperor", "0037/0041"}}}
	e4  Ruler = Ruler{"Claudius", "-0011/0054", []Office{{"Emperor", "0041/0054"}}}
	e5  Ruler = Ruler{"Nero", "0037/0068", []Office{{"Emperor", "0054/0068"}}}
	e6  Ruler = Ruler{"Otho", "003B/0069", []Office{{"Emperor", "0068/0069"}}}
	e7  Ruler = Ruler{"Aulus Vitellius", "-0016/0069", []Office{{"Emperor", "0069/0069"}}}
	e8  Ruler = Ruler{"Vespasian", "0009/0079", []Office{{"Emperor", "0069/0079"}}}
	e9  Ruler = Ruler{"Titus", "0039/0081", []Office{{"Emperor", "0079/0081"}}}
	e10 Ruler = Ruler{"Domitian", "0051/0096", []Office{{"Emperor", "0081/0096"}}}
	e11 Ruler = Ruler{"Nerva", "0030/0098", []Office{{"Emperor", "0096/0098"}}}
	e12 Ruler = Ruler{"Trajan", "0053/0117", []Office{{"Emperor", "0098/0117"}}}
	e13 Ruler = Ruler{"Hadrian", "0076/0138", []Office{{"Emperor", "0117/0138"}}}
	e14 Ruler = Ruler{"Antoninus Pius", "0086/0161", []Office{{"Emperor", "0138/0161"}}}
	e15 Ruler = Ruler{"Marcus Aurelius", "0121/0180", []Office{{"Emperor", "0161/0180"}}}
	e16 Ruler = Ruler{"Lucius Verus", "0130/0169", []Office{{"Emperor", "0161/0169"}}}
	e17 Ruler = Ruler{"Commodus", "0161/0192", []Office{{"Emperor", "0177/0192"}}}
	e18 Ruler = Ruler{"Publius Helvius Pertinax", "0126/0193", []Office{{"Emperor", "0193/0193"}}}
	e19 Ruler = Ruler{"Marcus Didius Severus Julianus", "0133/0193", []Office{{"Emperor", "0193/0193"}}}
	e20 Ruler = Ruler{"Septimius Severus", "0145/0211", []Office{{"Emperor", "0193/0211"}}}
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
