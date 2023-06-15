package main

import (
	"fmt"
)

// The main structure we want to populate
type Ruler struct {
	Name    string
	Birth   string
	Death   string
	Offices []Office
}

type Office struct {
	Office string
	Start  string
	End    string
}

// The 20 emperors from 31 BCE to 211 CE
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

var Emperors = []Ruler{e1, e2, e3, e4, e5, e6, e7, e8, e9, e10, e11, e12, e13, e14, e15, e16, e17, e18, e19, e20}

func main() {
	fmt.Println("Running scrapper...")
}
