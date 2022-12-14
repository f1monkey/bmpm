// GENERATED CODE. DO NOT EDIT!
package beidermorse

import "regexp"

type ashLang int64

const (
	ashany ashLang = 1 << iota
	ashcyrillic
	ashenglish
	ashfrench
	ashgerman
	ashhebrew
	ashhungarian
	ashpolish
	ashromanian
	ashrussian
	ashspanish
)

func (l ashLang) String() string {
	switch l {
	case ashany:
		return "any"
	case ashcyrillic:
		return "cyrillic"
	case ashenglish:
		return "english"
	case ashfrench:
		return "french"
	case ashgerman:
		return "german"
	case ashhebrew:
		return "hebrew"
	case ashhungarian:
		return "hungarian"
	case ashpolish:
		return "polish"
	case ashromanian:
		return "romanian"
	case ashrussian:
		return "russian"
	case ashspanish:
		return "spanish"
	}
	return ""
}

const ashAll = ashcyrillic +
	ashenglish +
	ashfrench +
	ashgerman +
	ashhebrew +
	ashhungarian +
	ashpolish +
	ashromanian +
	ashrussian +
	ashspanish

var ashRules = map[ashLang][]rule{
	ashany: []rule{
		{
			pattern: "yna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(in[512]|ina)",
		},
		{
			pattern: "ina",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(in[512]|ina)",
		},
		{
			pattern: "liova",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lof[512]|lef[512]|lova)",
		},
		{
			pattern: "lova",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lof[512]|lef[512]|lova)",
		},
		{
			pattern: "ova",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(of[512]|ova)",
		},
		{
			pattern: "eva",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ef[512]|eva)",
		},
		{
			pattern: "aia",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(aja|i[512])",
		},
		{
			pattern: "aja",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(aja|i[512])",
		},
		{
			pattern: "aya",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(aja|i[512])",
		},
		{
			pattern: "lowa",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lova|lof[128]|l[128]|el[128])",
		},
		{
			pattern: "kowa",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(kova|kof[128]|k[128]|ek[128])",
		},
		{
			pattern: "owa",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ova|of[128]|)",
		},
		{
			pattern: "lowna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lovna|levna|l[128]|el[128])",
		},
		{
			pattern: "kowna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(kovna|k[128]|ek[128])",
		},
		{
			pattern: "owna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ovna|[128])",
		},
		{
			pattern: "l??wna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(l|el[128])",
		},
		{
			pattern: "k??wna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(k|ek[128])",
		},
		{
			pattern: "??wna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern: "a",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(a|i[128])",
		},
		{
			pattern: "rh",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "r",
		},
		{
			pattern:  "ssch",
			phonetic: "S",
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
		},
		{
			pattern:  "tsch",
			phonetic: "tS",
		},
		{
			pattern: "sch",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(sk[256]|S|StS[512])",
		},
		{
			pattern:  "sch",
			phonetic: "(S|StS[512])",
		},
		{
			pattern:  "ssh",
			phonetic: "S",
		},
		{
			pattern: "sh",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[??????]"),
			},
			phonetic: "sh",
		},
		{
			pattern: "sh",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phonetic: "(S[516]|sh)",
		},
		{
			pattern:  "sh",
			phonetic: "S",
		},
		{
			pattern:  "kh",
			phonetic: "(x[516]|kh)",
		},
		{
			pattern:  "chs",
			phonetic: "(ks[16]|xs|tSs[516])",
		},
		{
			pattern: "ch",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(x|k[256]|tS[516])",
		},
		{
			pattern:  "ch",
			phonetic: "(x|tS[516])",
		},
		{
			pattern:  "ck",
			phonetic: "(k|tsk[128])",
		},
		{
			pattern:  "czy",
			phonetic: "tSi",
		},
		{
			pattern: "cze",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(tSe|tSF)",
		},
		{
			pattern:  "ciewicz",
			phonetic: "(tsevitS|tSevitS)",
		},
		{
			pattern:  "siewicz",
			phonetic: "(sevitS|SevitS)",
		},
		{
			pattern:  "ziewicz",
			phonetic: "(zevitS|ZevitS)",
		},
		{
			pattern:  "riewicz",
			phonetic: "rjevitS",
		},
		{
			pattern:  "diewicz",
			phonetic: "djevitS",
		},
		{
			pattern:  "tiewicz",
			phonetic: "tjevitS",
		},
		{
			pattern:  "iewicz",
			phonetic: "evitS",
		},
		{
			pattern:  "ewicz",
			phonetic: "evitS",
		},
		{
			pattern:  "owicz",
			phonetic: "ovitS",
		},
		{
			pattern:  "icz",
			phonetic: "itS",
		},
		{
			pattern:  "cz",
			phonetic: "tS",
		},
		{
			pattern: "cia",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(tSB[128]|tsB)",
		},
		{
			pattern:  "cia",
			phonetic: "(tSa[128]|tsa)",
		},
		{
			pattern: "ci??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(tSom[128]|tsom)",
		},
		{
			pattern:  "ci??",
			phonetic: "(tSon[128]|tson)",
		},
		{
			pattern: "ci??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(tSem[128]|tsem)",
		},
		{
			pattern:  "ci??",
			phonetic: "(tSen[128]|tsen)",
		},
		{
			pattern: "cie",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(tSF[128]|tsF)",
		},
		{
			pattern:  "cie",
			phonetic: "(tSe[128]|tse)",
		},
		{
			pattern:  "cio",
			phonetic: "(tSo[128]|tso)",
		},
		{
			pattern:  "ciu",
			phonetic: "(tSu[128]|tsu)",
		},
		{
			pattern: "ci",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(tsi[128]|tSi[384]|tS[256]|si)",
		},
		{
			pattern:  "ci",
			phonetic: "(tsi[128]|tSi[384]|si)",
		},
		{
			pattern: "ce",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(tsF[128]|tSe[384]|se)",
		},
		{
			pattern:  "ce",
			phonetic: "(tSe[384]|tse[128]|se)",
		},
		{
			pattern:  "cy",
			phonetic: "(si|tsi[128])",
		},
		{
			pattern:  "ssz",
			phonetic: "S",
		},
		{
			pattern:  "sz",
			phonetic: "S",
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp[16]|sp)",
		},
		{
			pattern:  "sp",
			phonetic: "(Sp[16]|sp)",
		},
		{
			pattern:  "sst",
			phonetic: "(St[16]|st)",
		},
		{
			pattern:  "st",
			phonetic: "(St[16]|st)",
		},
		{
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern: "sia",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(SB[128]|sB[128]|sja)",
		},
		{
			pattern:  "sia",
			phonetic: "(Sa[128]|sja)",
		},
		{
			pattern: "si??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Som[128]|som)",
		},
		{
			pattern:  "si??",
			phonetic: "(Son[128]|son)",
		},
		{
			pattern: "si??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Sem[128]|sem)",
		},
		{
			pattern:  "si??",
			phonetic: "(Sen[128]|sen)",
		},
		{
			pattern: "sie",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(SF[128]|sF|zi[16])",
		},
		{
			pattern:  "sie",
			phonetic: "(se|Se[128]|zi[16])",
		},
		{
			pattern:  "sio",
			phonetic: "(So[128]|so)",
		},
		{
			pattern:  "siu",
			phonetic: "(Su[128]|sju)",
		},
		{
			pattern:  "si",
			phonetic: "(Si[128]|si|zi[16])",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiou??????]"),
			},
			phonetic: "(s|z[16])",
		},
		{
			pattern:  "gue",
			phonetic: "ge",
		},
		{
			pattern:  "gui",
			phonetic: "gi",
		},
		{
			pattern:  "guy",
			phonetic: "gi",
		},
		{
			pattern: "gh",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(g[256]|gh)",
		},
		{
			pattern: "gauz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "haus",
		},
		{
			pattern: "gaus",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "haus",
		},
		{
			pattern: "gol'ts",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "golts",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "gol'tz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern:  "goltz",
			phonetic: "holts",
		},
		{
			pattern: "gol'ts",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "golts",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "gol'tz",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "goltz",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "gendler",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hendler",
		},
		{
			pattern: "gejmer",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajmer",
		},
		{
			pattern: "gejm",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajm",
		},
		{
			pattern: "geymer",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajmer",
		},
		{
			pattern: "geym",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajm",
		},
		{
			pattern: "geimer",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajmer",
		},
		{
			pattern: "geim",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajm",
		},
		{
			pattern: "gof",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hof",
		},
		{
			pattern: "ger",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ger",
		},
		{
			pattern: "gen",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "gen",
		},
		{
			pattern: "gin",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "gin",
		},
		{
			pattern: "gie",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ge|gi[16]|ji[8])",
		},
		{
			pattern:  "gie",
			phonetic: "ge",
		},
		{
			pattern: "ge",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phonetic: "(gE|xe[1024]|dZe[260])",
		},
		{
			pattern: "gi",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[yaeiou]$"),
			},
			phonetic: "(gI|xi[1024]|dZi[260])",
		},
		{
			pattern:  "ge",
			phonetic: "(gE|dZe[260]|hE[512]|xe[1024])",
		},
		{
			pattern:  "gi",
			phonetic: "(gI|dZi[260]|hI[512]|xi[1024])",
		},
		{
			pattern: "gy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou????????????????]"),
			},
			phonetic: "(gi|dj[64])",
		},
		{
			pattern:  "gy",
			phonetic: "(gi|d[64])",
		},
		{
			pattern: "g",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[jyaeiou]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouyei]"),
			},
			phonetic: "g",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aouei]"),
			},
			phonetic: "(g|h[512])",
		},
		{
			pattern:  "ej",
			phonetic: "(aj|eZ[264]|ex[1024])",
		},
		{
			pattern:  "ej",
			phonetic: "aj",
		},
		{
			pattern: "ly",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phonetic: "(l|lj)",
		},
		{
			pattern: "li",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phonetic: "(l|lj)",
		},
		{
			pattern: "lj",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phonetic: "(l|lj)",
		},
		{
			pattern:  "lio",
			phonetic: "(lo|le[512])",
		},
		{
			pattern:  "lyo",
			phonetic: "(lo|le[512])",
		},
		{
			pattern:  "ll",
			phonetic: "(l|J[1024])",
		},
		{
			pattern: "j",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aoeiuy]"),
			},
			phonetic: "(j|dZ[4]|x[1024]|Z[264])",
		},
		{
			pattern:  "j",
			phonetic: "(j|x[1024])",
		},
		{
			pattern:  "pf",
			phonetic: "(pf|p|f)",
		},
		{
			pattern:  "ph",
			phonetic: "(ph|f)",
		},
		{
			pattern:  "qu",
			phonetic: "(kv[16]|k)",
		},
		{
			pattern: "rze",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phonetic: "(Se[128]|re)",
		},
		{
			pattern:  "rze",
			phonetic: "(rze|rtsE[16]|Ze[128]|re[128]|rZe[128])",
		},
		{
			pattern: "rzy",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phonetic: "(Si[128]|ri)",
		},
		{
			pattern:  "rzy",
			phonetic: "(Zi[128]|ri[128]|rZi)",
		},
		{
			pattern: "rz",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phonetic: "(S[128]|r)",
		},
		{
			pattern:  "rz",
			phonetic: "(rz|rts[16]|Z[128]|r[128]|rZ[128])",
		},
		{
			pattern: "tz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ts|tS[20])",
		},
		{
			pattern: "tz",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ts|tS[20])",
		},
		{
			pattern:  "tz",
			phonetic: "(ts[532]|tz)",
		},
		{
			pattern:  "zh",
			phonetic: "(Z|zh[128]|tsh[16])",
		},
		{
			pattern: "zia",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(ZB[128]|zB[128]|zja)",
		},
		{
			pattern:  "zia",
			phonetic: "(Za[128]|zja)",
		},
		{
			pattern: "zi??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Zom[128]|zom)",
		},
		{
			pattern:  "zi??",
			phonetic: "(Zon[128]|zon)",
		},
		{
			pattern: "zi??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Zem[128]|zem)",
		},
		{
			pattern:  "zi??",
			phonetic: "(Zen[128]|zen)",
		},
		{
			pattern: "zie",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(ZF[128]|zF[128]|ze|tsi[16])",
		},
		{
			pattern:  "zie",
			phonetic: "(ze|Ze[128]|tsi[16])",
		},
		{
			pattern:  "zio",
			phonetic: "(Zo[128]|zo)",
		},
		{
			pattern:  "ziu",
			phonetic: "(Zu[128]|zju)",
		},
		{
			pattern:  "zi",
			phonetic: "(Zi[128]|zi|tsi[16])",
		},
		{
			pattern: "thal",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "tal",
		},
		{
			pattern: "th",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "t",
		},
		{
			pattern: "th",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiou]"),
			},
			phonetic: "(t[16]|th)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern:  "vogel",
			phonetic: "(vogel|fogel[16])",
		},
		{
			pattern: "v",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(v|f[16])",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy??????]$"),
			},
			phonetic: "",
		},
		{
			pattern:  "h",
			phonetic: "(h|x[384])",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(h|H[20])",
		},
		{
			pattern: "yi",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile(" $"),
			},
			phonetic: "i",
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phonetic: "i",
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phonetic: "i",
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phonetic: "i",
		},
		{
			pattern: "e",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "in",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(e|[8])",
		},
		{
			pattern: "yj",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
		},
		{
			pattern: "ij",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "oue",
			phonetic: "oue",
		},
		{
			pattern:  "au",
			phonetic: "(au|o[8])",
		},
		{
			pattern:  "ou",
			phonetic: "(ou|u[8])",
		},
		{
			pattern:  "ue",
			phonetic: "(Q|uje[512])",
		},
		{
			pattern:  "ae",
			phonetic: "(Y[16]|aje[512]|ae)",
		},
		{
			pattern:  "oe",
			phonetic: "(Y[16]|oje[512]|oe)",
		},
		{
			pattern:  "ee",
			phonetic: "(i[4]|aje[512]|e)",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:  "eu",
			phonetic: "(aj[16]|oj[16]|eu)",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "ie",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(i[16]|e[128]|ije[512]|je)",
		},
		{
			pattern:  "ie",
			phonetic: "(i[16]|e[128]|ije[512]|je)",
		},
		{
			pattern:  "ye",
			phonetic: "(je|ije[512])",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "io",
			phonetic: "(jo|e[512])",
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e[512])",
		},
		{
			pattern:  "ea",
			phonetic: "(ea|ja[256])",
		},
		{
			pattern: "e",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(e|je[512])",
		},
		{
			pattern:  "oo",
			phonetic: "(u[4]|o)",
		},
		{
			pattern:  "uu",
			phonetic: "u",
		},
		{
			pattern:  "??",
			phonetic: "(tS[128]|ts)",
		},
		{
			pattern:  "??",
			phonetic: "l",
		},
		{
			pattern:  "??",
			phonetic: "n",
		},
		{
			pattern:  "??",
			phonetic: "(n|nj[1024])",
		},
		{
			pattern:  "??",
			phonetic: "(S[128]|s)",
		},
		{
			pattern:  "??",
			phonetic: "S",
		},
		{
			pattern:  "??",
			phonetic: "ts",
		},
		{
			pattern:  "??",
			phonetic: "Z",
		},
		{
			pattern:  "??",
			phonetic: "(Z[128]|z)",
		},
		{
			pattern:  "o??",
			phonetic: "u",
		},
		{
			pattern: "??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "om",
		},
		{
			pattern:  "??",
			phonetic: "on",
		},
		{
			pattern:  "??",
			phonetic: "(Y|e)",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "??",
			phonetic: "(e[256]|a)",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "??",
			phonetic: "e",
		},
		{
			pattern:  "??",
			phonetic: "e",
		},
		{
			pattern:  "??",
			phonetic: "e",
		},
		{
			pattern: "??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "em",
		},
		{
			pattern:  "??",
			phonetic: "en",
		},
		{
			pattern:  "??",
			phonetic: "i",
		},
		{
			pattern:  "??",
			phonetic: "i",
		},
		{
			pattern:  "??",
			phonetic: "Y",
		},
		{
			pattern:  "??",
			phonetic: "Y",
		},
		{
			pattern:  "??",
			phonetic: "(u[128]|o)",
		},
		{
			pattern:  "??",
			phonetic: "Q",
		},
		{
			pattern:  "??",
			phonetic: "Q",
		},
		{
			pattern:  "??",
			phonetic: "u",
		},
		{
			pattern:  "??",
			phonetic: "Q",
		},
		{
			pattern:  "??",
			phonetic: "s",
		},
		{
			pattern:  "'",
			phonetic: "",
		},
		{
			pattern:  "\"",
			phonetic: "",
		},
		{
			pattern: "a",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(A|B[128])",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(E|F[128])",
		},
		{
			pattern: "o",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bc??dgkl??mn??rs??twz????]"),
			},
			phonetic: "(O|P[128])",
		},
		{
			pattern:  "a",
			phonetic: "A",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "c",
			phonetic: "(k|ts[128])",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "O",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "U",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "z",
			phonetic: "(ts[16]|z)",
		},
	},
	ashcyrillic: []rule{
		{
			pattern:  "????",
			phonetic: "tsa",
		},
		{
			pattern:  "????",
			phonetic: "tsu",
		},
		{
			pattern:  "??????",
			phonetic: "tsa",
		},
		{
			pattern:  "??????",
			phonetic: "tse",
		},
		{
			pattern:  "??????",
			phonetic: "tso",
		},
		{
			pattern:  "??????",
			phonetic: "tsu",
		},
		{
			pattern:  "??????",
			phonetic: "se",
		},
		{
			pattern:  "??????",
			phonetic: "so",
		},
		{
			pattern:  "??????",
			phonetic: "ze",
		},
		{
			pattern:  "??????",
			phonetic: "zo",
		},
		{
			pattern: "????????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "haus",
		},
		{
			pattern: "????????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "haus",
		},
		{
			pattern: "??????????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "????????????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajmer",
		},
		{
			pattern: "????????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajm",
		},
		{
			pattern: "??????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hof",
		},
		{
			pattern: "??????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ger",
		},
		{
			pattern: "??????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "gen",
		},
		{
			pattern: "??????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "gin",
		},
		{
			pattern: "??",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(??|??|??|??|??|??|??|??|??|??)$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(??|??|??|??|??)"),
			},
			phonetic: "g",
		},
		{
			pattern: "??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(??|??|??|??|??)"),
			},
			phonetic: "(g|h)",
		},
		{
			pattern:  "????",
			phonetic: "la",
		},
		{
			pattern:  "????",
			phonetic: "lu",
		},
		{
			pattern:  "????",
			phonetic: "(le|lo)",
		},
		{
			pattern:  "??????",
			phonetic: "(le|lo)",
		},
		{
			pattern:  "????",
			phonetic: "(lE|lo)",
		},
		{
			pattern:  "??????",
			phonetic: "je",
		},
		{
			pattern:  "????",
			phonetic: "je",
		},
		{
			pattern:  "??????",
			phonetic: "je",
		},
		{
			pattern:  "????",
			phonetic: "je",
		},
		{
			pattern: "????",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(??|??|??)"),
			},
			phonetic: "j",
		},
		{
			pattern: "????",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^(??|??|??)"),
			},
			phonetic: "j",
		},
		{
			pattern: "????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
		},
		{
			pattern: "????",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
		},
		{
			pattern:  "??",
			phonetic: "(e|jo)",
		},
		{
			pattern: "????",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(jaj|aj)",
		},
		{
			pattern: "??",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("(??|??|??|??)$"),
			},
			phonetic: "je",
		},
		{
			pattern: "??",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "je",
		},
		{
			pattern:  "????",
			phonetic: "aj",
		},
		{
			pattern:  "????",
			phonetic: "aj",
		},
		{
			pattern:  "??????",
			phonetic: "aue",
		},
		{
			pattern:  "??????",
			phonetic: "aue",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "??",
			phonetic: "b",
		},
		{
			pattern:  "??",
			phonetic: "v",
		},
		{
			pattern:  "??",
			phonetic: "g",
		},
		{
			pattern:  "??",
			phonetic: "d",
		},
		{
			pattern:  "??",
			phonetic: "E",
		},
		{
			pattern:  "??",
			phonetic: "Z",
		},
		{
			pattern:  "??",
			phonetic: "z",
		},
		{
			pattern:  "??",
			phonetic: "I",
		},
		{
			pattern:  "??",
			phonetic: "j",
		},
		{
			pattern:  "??",
			phonetic: "k",
		},
		{
			pattern:  "??",
			phonetic: "l",
		},
		{
			pattern:  "??",
			phonetic: "m",
		},
		{
			pattern:  "??",
			phonetic: "n",
		},
		{
			pattern:  "??",
			phonetic: "o",
		},
		{
			pattern:  "??",
			phonetic: "p",
		},
		{
			pattern:  "??",
			phonetic: "r",
		},
		{
			pattern: "??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "??",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern:  "??",
			phonetic: "s",
		},
		{
			pattern:  "??",
			phonetic: "t",
		},
		{
			pattern:  "??",
			phonetic: "u",
		},
		{
			pattern:  "??",
			phonetic: "f",
		},
		{
			pattern:  "??",
			phonetic: "x",
		},
		{
			pattern:  "??",
			phonetic: "ts",
		},
		{
			pattern:  "??",
			phonetic: "tS",
		},
		{
			pattern:  "??",
			phonetic: "S",
		},
		{
			pattern:  "??",
			phonetic: "StS",
		},
		{
			pattern:  "??",
			phonetic: "",
		},
		{
			pattern:  "??",
			phonetic: "I",
		},
		{
			pattern:  "??",
			phonetic: "",
		},
		{
			pattern:  "??",
			phonetic: "E",
		},
		{
			pattern:  "??",
			phonetic: "ju",
		},
		{
			pattern:  "??",
			phonetic: "ja",
		},
	},
	ashenglish: []rule{
		{
			pattern:  "tch",
			phonetic: "tS",
		},
		{
			pattern:  "ch",
			phonetic: "(tS|x)",
		},
		{
			pattern:  "ck",
			phonetic: "k",
		},
		{
			pattern: "cc",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phonetic: "ks",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "c",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phonetic: "s",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern: "gh",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "g",
		},
		{
			pattern:  "gh",
			phonetic: "(g|f|w)",
		},
		{
			pattern:  "gn",
			phonetic: "(gn|n)",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phonetic: "(g|dZ)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "ph",
			phonetic: "f",
		},
		{
			pattern:  "sch",
			phonetic: "(S|sk)",
		},
		{
			pattern:  "sh",
			phonetic: "S",
		},
		{
			pattern: "who",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hu",
		},
		{
			pattern: "wh",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "w",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "H",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "j",
			phonetic: "dZ",
		},
		{
			pattern: "kn",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "n",
		},
		{
			pattern: "mb",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "m",
		},
		{
			pattern: "ng",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(N|ng)",
		},
		{
			pattern: "pn",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(pn|n)",
		},
		{
			pattern: "ps",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ps|s)",
		},
		{
			pattern:  "qu",
			phonetic: "kw",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "tia",
			phonetic: "(So|Sa)",
		},
		{
			pattern:  "tio",
			phonetic: "So",
		},
		{
			pattern: "wr",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "r",
		},
		{
			pattern:  "w",
			phonetic: "(w|v)",
		},
		{
			pattern: "x",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "z",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern: "yi",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile(" $"),
			},
			phonetic: "i",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouy]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "oue",
			phonetic: "(aue|oue)",
		},
		{
			pattern:  "ai",
			phonetic: "(aj|e)",
		},
		{
			pattern:  "ay",
			phonetic: "aj",
		},
		{
			pattern: "a",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "aj",
		},
		{
			pattern:  "a",
			phonetic: "(e|o|a)",
		},
		{
			pattern:  "ei",
			phonetic: "(aj|i)",
		},
		{
			pattern:  "ey",
			phonetic: "(aj|i)",
		},
		{
			pattern:  "ear",
			phonetic: "ia",
		},
		{
			pattern:  "ea",
			phonetic: "(i|e)",
		},
		{
			pattern:  "ee",
			phonetic: "i",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "i",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(|E)",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "ie",
			phonetic: "i",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "aj",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "oa",
			phonetic: "ou",
		},
		{
			pattern:  "oi",
			phonetic: "oj",
		},
		{
			pattern:  "oo",
			phonetic: "u",
		},
		{
			pattern:  "ou",
			phonetic: "(u|ou)",
		},
		{
			pattern:  "oy",
			phonetic: "oj",
		},
		{
			pattern: "o",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "ou",
		},
		{
			pattern:  "o",
			phonetic: "(o|a)",
		},
		{
			pattern: "u",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[^aeiou]e"),
			},
			phonetic: "(ju|u)",
		},
		{
			pattern: "u",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "r",
				suffix:           "",
			},
			phonetic: "(e|u)",
		},
		{
			pattern:  "u",
			phonetic: "(u|a)",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	ashfrench: []rule{
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "ph",
			phonetic: "f",
		},
		{
			pattern:  "??",
			phonetic: "s",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "ch",
			phonetic: "S",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiy??????]"),
			},
			phonetic: "s",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "gn",
			phonetic: "(n|gn)",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phonetic: "Z",
		},
		{
			pattern: "gue",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "k",
		},
		{
			pattern: "gu",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phonetic: "g",
		},
		{
			pattern: "que",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "k",
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy??????]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouy??????]"),
			},
			phonetic: "z",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern:  "j",
			phonetic: "Z",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern: "ouh",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aioe]"),
			},
			phonetic: "(v|uh)",
		},
		{
			pattern: "ou",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phonetic: "v",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
		},
		{
			pattern: "u",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeio]"),
			},
			phonetic: "v",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "eau",
			phonetic: "o",
		},
		{
			pattern:  "ai",
			phonetic: "aj",
		},
		{
			pattern:  "ay",
			phonetic: "aj",
		},
		{
			pattern:  "??",
			phonetic: "e",
		},
		{
			pattern:  "??",
			phonetic: "e",
		},
		{
			pattern:  "??",
			phonetic: "e",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "o??",
			phonetic: "u",
		},
		{
			pattern:  "ou",
			phonetic: "u",
		},
		{
			pattern:  "oi",
			phonetic: "oj",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(e|)",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aoeu]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "yi",
			phonetic: "i",
		},
		{
			pattern:  "ii",
			phonetic: "i",
		},
		{
			pattern:  "yy",
			phonetic: "i",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	ashgerman: []rule{
		{
			pattern:  "ziu",
			phonetic: "tsu",
		},
		{
			pattern:  "zia",
			phonetic: "tsa",
		},
		{
			pattern:  "zio",
			phonetic: "tso",
		},
		{
			pattern:  "ssch",
			phonetic: "S",
		},
		{
			pattern:  "chsch",
			phonetic: "xS",
		},
		{
			pattern: "ewitsch",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "evitS",
		},
		{
			pattern: "owitsch",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ovitS",
		},
		{
			pattern: "evitsch",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "evitS",
		},
		{
			pattern: "ovitsch",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ovitS",
		},
		{
			pattern: "witsch",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "vitS",
		},
		{
			pattern: "vitsch",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "vitS",
		},
		{
			pattern:  "sch",
			phonetic: "S",
		},
		{
			pattern:  "chs",
			phonetic: "ks",
		},
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern:  "ck",
			phonetic: "k",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[eiy]"),
			},
			phonetic: "ts",
		},
		{
			pattern: "sp",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "Sp",
		},
		{
			pattern: "st",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "St",
		},
		{
			pattern:  "ssp",
			phonetic: "(Sp|sp)",
		},
		{
			pattern:  "sp",
			phonetic: "(Sp|sp)",
		},
		{
			pattern:  "sst",
			phonetic: "(St|st)",
		},
		{
			pattern:  "st",
			phonetic: "(St|st)",
		},
		{
			pattern:  "pf",
			phonetic: "(pf|p|f)",
		},
		{
			pattern:  "ph",
			phonetic: "(ph|f)",
		},
		{
			pattern:  "qu",
			phonetic: "kv",
		},
		{
			pattern: "ewitz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(evits|evitS)",
		},
		{
			pattern: "ewiz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(evits|evitS)",
		},
		{
			pattern: "evitz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(evits|evitS)",
		},
		{
			pattern: "eviz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(evits|evitS)",
		},
		{
			pattern: "owitz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ovits|ovitS)",
		},
		{
			pattern: "owiz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ovits|ovitS)",
		},
		{
			pattern: "ovitz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ovits|ovitS)",
		},
		{
			pattern: "oviz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ovits|ovitS)",
		},
		{
			pattern: "witz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(vits|vitS)",
		},
		{
			pattern: "wiz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(vits|vitS)",
		},
		{
			pattern: "vitz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(vits|vitS)",
		},
		{
			pattern: "viz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(vits|vitS)",
		},
		{
			pattern:  "tz",
			phonetic: "ts",
		},
		{
			pattern: "thal",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "tal",
		},
		{
			pattern: "th",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "t",
		},
		{
			pattern: "th",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[??????aeiou]"),
			},
			phonetic: "(t|th)",
		},
		{
			pattern:  "th",
			phonetic: "t",
		},
		{
			pattern: "rh",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "r",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy??????]$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "H",
		},
		{
			pattern:  "ss",
			phonetic: "s",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[??????aeiouy]"),
			},
			phonetic: "(z|s)",
		},
		{
			pattern: "s",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aeiouy??????j]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeiouy??????]"),
			},
			phonetic: "z",
		},
		{
			pattern:  "??",
			phonetic: "s",
		},
		{
			pattern: "ij",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "ue",
			phonetic: "Q",
		},
		{
			pattern:  "ae",
			phonetic: "Y",
		},
		{
			pattern:  "oe",
			phonetic: "Y",
		},
		{
			pattern:  "??",
			phonetic: "Q",
		},
		{
			pattern:  "??",
			phonetic: "(Y|e)",
		},
		{
			pattern:  "??",
			phonetic: "Y",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:  "eu",
			phonetic: "(aj|oj)",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern:  "ie",
			phonetic: "I",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aoeu]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "??",
			phonetic: "n",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "??",
			phonetic: "o",
		},
		{
			pattern:  "??",
			phonetic: "u",
		},
		{
			pattern:  "??",
			phonetic: "s",
		},
		{
			pattern:  "a",
			phonetic: "A",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "O",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "U",
		},
		{
			pattern:  "v",
			phonetic: "(f|v)",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "z",
			phonetic: "ts",
		},
	},
	ashhebrew: []rule{
		{
			pattern:  "????",
			phonetic: "i",
		},
		{
			pattern:  "????",
			phonetic: "i",
		},
		{
			pattern:  "????",
			phonetic: "VV",
		},
		{
			pattern:  "????",
			phonetic: "VV",
		},
		{
			pattern:  "????",
			phonetic: "Z",
		},
		{
			pattern:  "????",
			phonetic: "dZ",
		},
		{
			pattern:  "??",
			phonetic: "L",
		},
		{
			pattern:  "??",
			phonetic: "b",
		},
		{
			pattern:  "??",
			phonetic: "g",
		},
		{
			pattern:  "??",
			phonetic: "d",
		},
		{
			pattern: "??",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "1",
		},
		{
			pattern: "??",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "1",
		},
		{
			pattern:  "??",
			phonetic: "",
		},
		{
			pattern:  "????",
			phonetic: "V",
		},
		{
			pattern:  "????",
			phonetic: "WW",
		},
		{
			pattern:  "??",
			phonetic: "W",
		},
		{
			pattern:  "??",
			phonetic: "z",
		},
		{
			pattern:  "??",
			phonetic: "X",
		},
		{
			pattern:  "??",
			phonetic: "T",
		},
		{
			pattern:  "????",
			phonetic: "i",
		},
		{
			pattern:  "??",
			phonetic: "i",
		},
		{
			pattern:  "??",
			phonetic: "X",
		},
		{
			pattern: "??",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "K",
		},
		{
			pattern:  "??",
			phonetic: "k",
		},
		{
			pattern:  "??",
			phonetic: "l",
		},
		{
			pattern:  "??",
			phonetic: "m",
		},
		{
			pattern:  "??",
			phonetic: "m",
		},
		{
			pattern:  "??",
			phonetic: "n",
		},
		{
			pattern:  "??",
			phonetic: "n",
		},
		{
			pattern:  "??",
			phonetic: "s",
		},
		{
			pattern:  "??",
			phonetic: "L",
		},
		{
			pattern:  "??",
			phonetic: "f",
		},
		{
			pattern:  "??",
			phonetic: "f",
		},
		{
			pattern:  "??",
			phonetic: "C",
		},
		{
			pattern:  "??",
			phonetic: "C",
		},
		{
			pattern:  "??",
			phonetic: "K",
		},
		{
			pattern:  "??",
			phonetic: "r",
		},
		{
			pattern:  "??",
			phonetic: "s",
		},
		{
			pattern:  "??",
			phonetic: "TB",
		},
	},
	ashhungarian: []rule{
		{
			pattern:  "sz",
			phonetic: "s",
		},
		{
			pattern:  "zs",
			phonetic: "Z",
		},
		{
			pattern:  "cs",
			phonetic: "tS",
		},
		{
			pattern:  "ay",
			phonetic: "(oj|aj)",
		},
		{
			pattern:  "ai",
			phonetic: "(oj|aj)",
		},
		{
			pattern:  "aj",
			phonetic: "(oj|aj)",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[??o]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[??o]$"),
			},
			phonetic: "j",
		},
		{
			pattern:  "ee",
			phonetic: "(aj|e)",
		},
		{
			pattern:  "ely",
			phonetic: "(aj|eli)",
		},
		{
			pattern:  "ly",
			phonetic: "(j|li)",
		},
		{
			pattern: "gy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou????????????????]"),
			},
			phonetic: "dj",
		},
		{
			pattern:  "gy",
			phonetic: "(d|gi)",
		},
		{
			pattern: "ny",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou????????????????]"),
			},
			phonetic: "nj",
		},
		{
			pattern:  "ny",
			phonetic: "(n|ni)",
		},
		{
			pattern: "ty",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou????????????????]"),
			},
			phonetic: "tj",
		},
		{
			pattern:  "ty",
			phonetic: "(t|ti)",
		},
		{
			pattern:  "qu",
			phonetic: "(ku|kv)",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "??",
			phonetic: "e",
		},
		{
			pattern:  "??",
			phonetic: "i",
		},
		{
			pattern:  "??",
			phonetic: "o",
		},
		{
			pattern:  "??",
			phonetic: "Y",
		},
		{
			pattern:  "??",
			phonetic: "Y",
		},
		{
			pattern:  "??",
			phonetic: "u",
		},
		{
			pattern:  "??",
			phonetic: "Q",
		},
		{
			pattern:  "??",
			phonetic: "Q",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "c",
			phonetic: "ts",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "(S|s)",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	ashpolish: []rule{
		{
			pattern: "ska",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ski",
		},
		{
			pattern: "cka",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "tski",
		},
		{
			pattern: "lowa",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lova|lof|l|el)",
		},
		{
			pattern: "kowa",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(kova|kof|k|ek)",
		},
		{
			pattern: "owa",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ova|of|)",
		},
		{
			pattern: "lowna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lovna|levna|l|el)",
		},
		{
			pattern: "kowna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(kovna|k|ek)",
		},
		{
			pattern: "owna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ovna|)",
		},
		{
			pattern: "l??wna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(l|el)",
		},
		{
			pattern: "k??wna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(k|ek)",
		},
		{
			pattern: "??wna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern: "a",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(a|i)",
		},
		{
			pattern:  "czy",
			phonetic: "tSi",
		},
		{
			pattern: "cze",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(tSe|tSF)",
		},
		{
			pattern:  "ciewicz",
			phonetic: "(tsevitS|tSevitS)",
		},
		{
			pattern:  "siewicz",
			phonetic: "(sevitS|SevitS)",
		},
		{
			pattern:  "ziewicz",
			phonetic: "(zevitS|ZevitS)",
		},
		{
			pattern:  "riewicz",
			phonetic: "rjevitS",
		},
		{
			pattern:  "diewicz",
			phonetic: "djevitS",
		},
		{
			pattern:  "tiewicz",
			phonetic: "tjevitS",
		},
		{
			pattern:  "iewicz",
			phonetic: "evitS",
		},
		{
			pattern:  "ewicz",
			phonetic: "evitS",
		},
		{
			pattern:  "owicz",
			phonetic: "ovitS",
		},
		{
			pattern:  "icz",
			phonetic: "itS",
		},
		{
			pattern:  "cz",
			phonetic: "tS",
		},
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern: "cia",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(tSB|tsB)",
		},
		{
			pattern:  "cia",
			phonetic: "(tSa|tsa)",
		},
		{
			pattern: "ci??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(tSom|tsom)",
		},
		{
			pattern:  "ci??",
			phonetic: "(tSon|tson)",
		},
		{
			pattern: "ci??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(tSem|tsem)",
		},
		{
			pattern:  "ci??",
			phonetic: "(tSen|tsen)",
		},
		{
			pattern: "cie",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(tSF|tsF)",
		},
		{
			pattern:  "cie",
			phonetic: "(tSe|tse)",
		},
		{
			pattern:  "cio",
			phonetic: "(tSo|tso)",
		},
		{
			pattern:  "ciu",
			phonetic: "(tSu|tsu)",
		},
		{
			pattern:  "ci",
			phonetic: "(tSi|tsI)",
		},
		{
			pattern:  "??",
			phonetic: "(tS|ts)",
		},
		{
			pattern:  "ssz",
			phonetic: "S",
		},
		{
			pattern:  "sz",
			phonetic: "S",
		},
		{
			pattern: "sia",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(SB|sB|sja)",
		},
		{
			pattern:  "sia",
			phonetic: "(Sa|sja)",
		},
		{
			pattern: "si??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Som|som)",
		},
		{
			pattern:  "si??",
			phonetic: "(Son|son)",
		},
		{
			pattern: "si??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Sem|sem)",
		},
		{
			pattern:  "si??",
			phonetic: "(Sen|sen)",
		},
		{
			pattern: "sie",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(SF|sF|se)",
		},
		{
			pattern:  "sie",
			phonetic: "(Se|se)",
		},
		{
			pattern:  "sio",
			phonetic: "(So|so)",
		},
		{
			pattern:  "siu",
			phonetic: "(Su|sju)",
		},
		{
			pattern:  "si",
			phonetic: "(Si|sI)",
		},
		{
			pattern:  "??",
			phonetic: "(S|s)",
		},
		{
			pattern: "zia",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(ZB|zB|zja)",
		},
		{
			pattern:  "zia",
			phonetic: "(Za|zja)",
		},
		{
			pattern: "zi??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Zom|zom)",
		},
		{
			pattern:  "zi??",
			phonetic: "(Zon|zon)",
		},
		{
			pattern: "zi??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "(Zem|zem)",
		},
		{
			pattern:  "zi??",
			phonetic: "(Zen|zen)",
		},
		{
			pattern: "zie",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(ZF|zF)",
		},
		{
			pattern:  "zie",
			phonetic: "(Ze|ze)",
		},
		{
			pattern:  "zio",
			phonetic: "(Zo|zo)",
		},
		{
			pattern:  "ziu",
			phonetic: "(Zu|zju)",
		},
		{
			pattern:  "zi",
			phonetic: "(Zi|zI)",
		},
		{
			pattern: "??e",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(Ze|ZF)",
		},
		{
			pattern: "??e",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(Ze|ZF|ze|zF)",
		},
		{
			pattern:  "??e",
			phonetic: "Ze",
		},
		{
			pattern:  "??e",
			phonetic: "(Ze|ze)",
		},
		{
			pattern:  "??y",
			phonetic: "Zi",
		},
		{
			pattern:  "??i",
			phonetic: "(Zi|zi)",
		},
		{
			pattern:  "??",
			phonetic: "Z",
		},
		{
			pattern:  "??",
			phonetic: "(Z|z)",
		},
		{
			pattern: "rze",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phonetic: "(Se|re)",
		},
		{
			pattern:  "rze",
			phonetic: "(Ze|re|rZe)",
		},
		{
			pattern: "rzy",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phonetic: "(Si|ri)",
		},
		{
			pattern:  "rzy",
			phonetic: "(Zi|ri|rZi)",
		},
		{
			pattern: "rz",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "t",
			},
			phonetic: "(S|r)",
		},
		{
			pattern:  "rz",
			phonetic: "(Z|r|rZ)",
		},
		{
			pattern:  "lio",
			phonetic: "(lo|le)",
		},
		{
			pattern:  "??",
			phonetic: "l",
		},
		{
			pattern:  "??",
			phonetic: "n",
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "s",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern:  "??",
			phonetic: "(u|o)",
		},
		{
			pattern: "??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "om",
		},
		{
			pattern: "??",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bp]"),
			},
			phonetic: "em",
		},
		{
			pattern:  "??",
			phonetic: "on",
		},
		{
			pattern:  "??",
			phonetic: "en",
		},
		{
			pattern:  "ije",
			phonetic: "je",
		},
		{
			pattern:  "yje",
			phonetic: "je",
		},
		{
			pattern:  "iie",
			phonetic: "je",
		},
		{
			pattern:  "yie",
			phonetic: "je",
		},
		{
			pattern:  "iye",
			phonetic: "je",
		},
		{
			pattern:  "yye",
			phonetic: "je",
		},
		{
			pattern: "ij",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yj",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "rie",
			phonetic: "rje",
		},
		{
			pattern:  "die",
			phonetic: "dje",
		},
		{
			pattern:  "tie",
			phonetic: "tje",
		},
		{
			pattern: "ie",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "F",
		},
		{
			pattern:  "ie",
			phonetic: "e",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "au",
			phonetic: "au",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:  "ej",
			phonetic: "aj",
		},
		{
			pattern:  "ai",
			phonetic: "aj",
		},
		{
			pattern:  "ay",
			phonetic: "aj",
		},
		{
			pattern:  "aj",
			phonetic: "aj",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[ou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "a",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "B",
		},
		{
			pattern: "e",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bcdgkpstwz??]"),
			},
			phonetic: "(E|F)",
		},
		{
			pattern: "o",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bc??dgkl??mn??rs??twz????]"),
			},
			phonetic: "P",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "c",
			phonetic: "ts",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "(h|x)",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "y",
			phonetic: "I",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	ashromanian: []rule{
		{
			pattern:  "j",
			phonetic: "Z",
		},
		{
			pattern:  "ce",
			phonetic: "tSe",
		},
		{
			pattern:  "ci",
			phonetic: "(tSi|tS)",
		},
		{
			pattern: "ch",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "k",
		},
		{
			pattern:  "ch",
			phonetic: "x",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "gi",
			phonetic: "(dZi|dZ)",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "dZ",
		},
		{
			pattern:  "gh",
			phonetic: "g",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeou]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "??",
			phonetic: "ts",
		},
		{
			pattern:  "??",
			phonetic: "S",
		},
		{
			pattern:  "h",
			phonetic: "(x|h)",
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "y",
			phonetic: "i",
		},
		{
			pattern:  "??",
			phonetic: "i",
		},
		{
			pattern:  "ea",
			phonetic: "ja",
		},
		{
			pattern:  "??",
			phonetic: "(e|a)",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	ashrussian: []rule{
		{
			pattern: "yna",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(in|ina)",
		},
		{
			pattern: "ina",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(in|ina)",
		},
		{
			pattern: "liova",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lof|lef)",
		},
		{
			pattern: "lova",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(lof|lef|lova)",
		},
		{
			pattern: "ova",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(of|ova)",
		},
		{
			pattern: "eva",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(ef|ova)",
		},
		{
			pattern: "aia",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(aja|i)",
		},
		{
			pattern: "aja",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(aja|i)",
		},
		{
			pattern: "aya",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(aja|i)",
		},
		{
			pattern:  "tsya",
			phonetic: "tsa",
		},
		{
			pattern:  "tsyu",
			phonetic: "tsu",
		},
		{
			pattern:  "tsia",
			phonetic: "tsa",
		},
		{
			pattern:  "tsie",
			phonetic: "tse",
		},
		{
			pattern:  "tsio",
			phonetic: "tso",
		},
		{
			pattern:  "tsye",
			phonetic: "tse",
		},
		{
			pattern:  "tsyo",
			phonetic: "tso",
		},
		{
			pattern:  "tsiu",
			phonetic: "tsu",
		},
		{
			pattern:  "sie",
			phonetic: "se",
		},
		{
			pattern:  "sio",
			phonetic: "so",
		},
		{
			pattern:  "zie",
			phonetic: "ze",
		},
		{
			pattern:  "zio",
			phonetic: "zo",
		},
		{
			pattern:  "sye",
			phonetic: "se",
		},
		{
			pattern:  "syo",
			phonetic: "so",
		},
		{
			pattern:  "zye",
			phonetic: "ze",
		},
		{
			pattern:  "zyo",
			phonetic: "zo",
		},
		{
			pattern: "gauz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "haus",
		},
		{
			pattern: "gaus",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "haus",
		},
		{
			pattern: "gol'ts",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "golts",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "gol'tz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "goltz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "holts",
		},
		{
			pattern: "gejmer",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajmer",
		},
		{
			pattern: "gejm",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajm",
		},
		{
			pattern: "geimer",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajmer",
		},
		{
			pattern: "geim",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajm",
		},
		{
			pattern: "geymer",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajmer",
		},
		{
			pattern: "geym",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hajm",
		},
		{
			pattern: "gendler",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hendler",
		},
		{
			pattern: "gof",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hof",
		},
		{
			pattern: "gojf",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hojf",
		},
		{
			pattern: "goyf",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hojf",
		},
		{
			pattern: "goif",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "hojf",
		},
		{
			pattern: "ger",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ger",
		},
		{
			pattern: "gen",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "gen",
		},
		{
			pattern: "gin",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "gin",
		},
		{
			pattern:  "gg",
			phonetic: "g",
		},
		{
			pattern: "kog",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phonetic: "(kog|koh)",
		},
		{
			pattern: "kag",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phonetic: "(kag|kah)",
		},
		{
			pattern: "g",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[jaeoiuy]$"),
			},
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phonetic: "g",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aeoiu]"),
			},
			phonetic: "(g|h)",
		},
		{
			pattern:  "kh",
			phonetic: "x",
		},
		{
			pattern:  "ch",
			phonetic: "(tS|x)",
		},
		{
			pattern:  "sch",
			phonetic: "(StS|S)",
		},
		{
			pattern:  "ssh",
			phonetic: "S",
		},
		{
			pattern:  "sh",
			phonetic: "S",
		},
		{
			pattern:  "zh",
			phonetic: "Z",
		},
		{
			pattern: "tz",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "ts",
		},
		{
			pattern:  "tz",
			phonetic: "(ts|tz)",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[iey]"),
			},
			phonetic: "s",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "qu",
			phonetic: "(kv|k)",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern: "s",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "s",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "lya",
			phonetic: "la",
		},
		{
			pattern:  "lyu",
			phonetic: "lu",
		},
		{
			pattern:  "lia",
			phonetic: "la",
		},
		{
			pattern:  "liu",
			phonetic: "lu",
		},
		{
			pattern:  "lja",
			phonetic: "la",
		},
		{
			pattern:  "lju",
			phonetic: "lu",
		},
		{
			pattern:  "le",
			phonetic: "(lo|lE)",
		},
		{
			pattern:  "lyo",
			phonetic: "(lo|le)",
		},
		{
			pattern:  "lio",
			phonetic: "(lo|le)",
		},
		{
			pattern:  "ije",
			phonetic: "je",
		},
		{
			pattern:  "ie",
			phonetic: "je",
		},
		{
			pattern:  "iye",
			phonetic: "je",
		},
		{
			pattern:  "iie",
			phonetic: "je",
		},
		{
			pattern:  "yje",
			phonetic: "je",
		},
		{
			pattern:  "ye",
			phonetic: "je",
		},
		{
			pattern:  "yye",
			phonetic: "je",
		},
		{
			pattern:  "yie",
			phonetic: "je",
		},
		{
			pattern: "ij",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yj",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aou]"),
			},
			phonetic: "j",
		},
		{
			pattern:  "io",
			phonetic: "(jo|e)",
		},
		{
			pattern: "i",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phonetic: "j",
		},
		{
			pattern: "i",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phonetic: "j",
		},
		{
			pattern:  "ei",
			phonetic: "aj",
		},
		{
			pattern:  "ey",
			phonetic: "aj",
		},
		{
			pattern:  "ej",
			phonetic: "aj",
		},
		{
			pattern:  "yo",
			phonetic: "(jo|e)",
		},
		{
			pattern: "y",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[au]"),
			},
			phonetic: "j",
		},
		{
			pattern: "y",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aiou]$"),
			},
			phonetic: "j",
		},
		{
			pattern: "ii",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phonetic: "i",
		},
		{
			pattern: "iy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phonetic: "i",
		},
		{
			pattern: "yy",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phonetic: "i",
		},
		{
			pattern: "yi",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^ "),
			},
			phonetic: "i",
		},
		{
			pattern: "yj",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
		},
		{
			pattern: "ij",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "i",
		},
		{
			pattern: "e",
			leftContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "(je|E)",
		},
		{
			pattern:  "ee",
			phonetic: "(aje|i)",
		},
		{
			pattern: "e",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[aou]$"),
			},
			phonetic: "je",
		},
		{
			pattern:  "y",
			phonetic: "I",
		},
		{
			pattern:  "oo",
			phonetic: "(oo|u)",
		},
		{
			pattern:  "'",
			phonetic: "",
		},
		{
			pattern:  "\"",
			phonetic: "",
		},
		{
			pattern:  "aue",
			phonetic: "aue",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "b",
			phonetic: "b",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "j",
			phonetic: "j",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
		{
			pattern:  "v",
			phonetic: "v",
		},
		{
			pattern:  "z",
			phonetic: "z",
		},
	},
	ashspanish: []rule{
		{
			pattern:  "??",
			phonetic: "(n|nj)",
		},
		{
			pattern:  "ch",
			phonetic: "(tS|dZ)",
		},
		{
			pattern: "h",
			leftContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("[bdgt]$"),
			},
			phonetic: "",
		},
		{
			pattern: "h",
			rightContext: &ruleMatcher{
				matchEmptyString: true,
				contains:         "",
				prefix:           "",
				suffix:           "",
			},
			phonetic: "",
		},
		{
			pattern:  "j",
			phonetic: "x",
		},
		{
			pattern:  "x",
			phonetic: "ks",
		},
		{
			pattern:  "ll",
			phonetic: "(l|Z)",
		},
		{
			pattern:  "w",
			phonetic: "v",
		},
		{
			pattern:  "v",
			phonetic: "(b|v)",
		},
		{
			pattern:  "b",
			phonetic: "(b|v)",
		},
		{
			pattern: "m",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[bpvf]"),
			},
			phonetic: "(m|n)",
		},
		{
			pattern: "c",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "s",
		},
		{
			pattern:  "c",
			phonetic: "k",
		},
		{
			pattern:  "z",
			phonetic: "(z|s)",
		},
		{
			pattern: "gu",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(g|gv)",
		},
		{
			pattern: "g",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[ei]"),
			},
			phonetic: "(x|g)",
		},
		{
			pattern:  "qu",
			phonetic: "k",
		},
		{
			pattern:  "q",
			phonetic: "k",
		},
		{
			pattern:  "uo",
			phonetic: "(vo|o)",
		},
		{
			pattern: "u",
			rightContext: &ruleMatcher{
				matchEmptyString: false,
				contains:         "",
				prefix:           "",
				suffix:           "",
				pattern:          regexp.MustCompile("^[aei]"),
			},
			phonetic: "v",
		},
		{
			pattern:  "y",
			phonetic: "(i|j|S|Z)",
		},
		{
			pattern:  "??",
			phonetic: "v",
		},
		{
			pattern:  "??",
			phonetic: "a",
		},
		{
			pattern:  "??",
			phonetic: "e",
		},
		{
			pattern:  "??",
			phonetic: "i",
		},
		{
			pattern:  "??",
			phonetic: "o",
		},
		{
			pattern:  "??",
			phonetic: "u",
		},
		{
			pattern:  "a",
			phonetic: "a",
		},
		{
			pattern:  "d",
			phonetic: "d",
		},
		{
			pattern:  "e",
			phonetic: "E",
		},
		{
			pattern:  "f",
			phonetic: "f",
		},
		{
			pattern:  "g",
			phonetic: "g",
		},
		{
			pattern:  "h",
			phonetic: "h",
		},
		{
			pattern:  "i",
			phonetic: "I",
		},
		{
			pattern:  "k",
			phonetic: "k",
		},
		{
			pattern:  "l",
			phonetic: "l",
		},
		{
			pattern:  "m",
			phonetic: "m",
		},
		{
			pattern:  "n",
			phonetic: "n",
		},
		{
			pattern:  "o",
			phonetic: "o",
		},
		{
			pattern:  "p",
			phonetic: "p",
		},
		{
			pattern:  "r",
			phonetic: "r",
		},
		{
			pattern:  "s",
			phonetic: "s",
		},
		{
			pattern:  "t",
			phonetic: "t",
		},
		{
			pattern:  "u",
			phonetic: "u",
		},
	},
}

var ashLangRules = []langRule{
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "zh",
			prefix:           "",
			suffix:           "",
		},
		langs:  660,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "eau",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("[aoeiu??????]h"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "vogel",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "vogel",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "witz",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "tz",
		},
		langs:  532,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "tz",
			suffix:           "",
		},
		langs:  516,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "g??e",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "g??i",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ghe",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ghi",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "vici",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "schi",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "chsch",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "tsch",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ssch",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "sch",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "sch",
			suffix:           "",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "rz",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "rz",
		},
		langs:  144,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("[^aoeiu??????]rz"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("rz[^aoeiu??????]"),
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cki",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ska",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cka",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ue",
			prefix:           "",
			suffix:           "",
		},
		langs:  528,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ae",
			prefix:           "",
			suffix:           "",
		},
		langs:  532,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "oe",
			prefix:           "",
			suffix:           "",
		},
		langs:  540,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "th",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "th",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("th[^aoeiu]"),
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "mann",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "cz",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "cy",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "niew",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "stein",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "heim",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "heimer",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ii",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "iy",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "yy",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "yi",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "yj",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ij",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gaus",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gauz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gauz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "goltz",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("gol'tz$"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "golts",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("gol'ts$"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "goltz",
			suffix:           "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("^gol'tz"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "golts",
			suffix:           "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("^gol'ts"),
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gendler",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gejmer",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gejm",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "geimer",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "geim",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "geymer",
			prefix:           "",
			suffix:           "",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "geym",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gof",
		},
		langs:  512,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "thal",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "zweig",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "ck",
		},
		langs:  20,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "c",
		},
		langs:  448,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sz",
			prefix:           "",
			suffix:           "",
		},
		langs:  192,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gue",
			prefix:           "",
			suffix:           "",
		},
		langs:  1032,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gui",
			prefix:           "",
			suffix:           "",
		},
		langs:  1032,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "guy",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "cs",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "cs",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "dzs",
			prefix:           "",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "zs",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "zs",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "wl",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "wr",
			suffix:           "",
		},
		langs:  148,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "gy",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("gy[aeou]"),
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gy",
			prefix:           "",
			suffix:           "",
		},
		langs:  576,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ly",
			prefix:           "",
			suffix:           "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ny",
			prefix:           "",
			suffix:           "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ty",
			prefix:           "",
			suffix:           "",
		},
		langs:  704,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  264,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  1096,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  264,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  1024,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  1216,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  80,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  80,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  8,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  64,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  1088,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  2,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "??",
			prefix:           "",
			suffix:           "",
		},
		langs:  32,
		accept: true,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "a",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "o",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "e",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "i",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "y",
			prefix:           "",
			suffix:           "",
		},
		langs:  290,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "u",
			prefix:           "",
			suffix:           "",
		},
		langs:  34,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("v[^aoeiu??????]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("y[^aoeiu]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "",
			suffix:           "",
			pattern:          regexp.MustCompile("c[^aohk]"),
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "dzi",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ou",
			prefix:           "",
			suffix:           "",
		},
		langs:  16,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "aj",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ej",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "oj",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "uj",
			prefix:           "",
			suffix:           "",
		},
		langs:  28,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "k",
			prefix:           "",
			suffix:           "",
		},
		langs:  256,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "v",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "ky",
			prefix:           "",
			suffix:           "",
		},
		langs:  128,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "eu",
			prefix:           "",
			suffix:           "",
		},
		langs:  640,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "w",
			prefix:           "",
			suffix:           "",
		},
		langs:  1864,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "kie",
			prefix:           "",
			suffix:           "",
		},
		langs:  1032,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "gie",
			prefix:           "",
			suffix:           "",
		},
		langs:  1288,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "q",
			prefix:           "",
			suffix:           "",
		},
		langs:  960,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "sch",
			prefix:           "",
			suffix:           "",
		},
		langs:  1224,
		accept: false,
	},
	{
		match: ruleMatcher{
			matchEmptyString: false,
			contains:         "",
			prefix:           "h",
			suffix:           "",
		},
		langs:  512,
		accept: false,
	},
}

var ashFinalRules = finalRules{
	approx: finalRule{
		first: []rule{
			{
				pattern: "h",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phonetic: "p",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "p",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "p",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[gdZz]"),
				},
				phonetic: "b",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "b",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phonetic: "f",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "f",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "f",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgdZz]"),
				},
				phonetic: "v",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "v",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phonetic: "k",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "k",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "k",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdZz]"),
				},
				phonetic: "g",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "g",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phonetic: "t",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "t",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "t",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgZz]"),
				},
				phonetic: "d",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "d",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "dZ",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "tS",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phonetic: "s",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern:  "jnm",
				phonetic: "jm",
			},
			{
				pattern: "ji",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "i",
			},
			{
				pattern: "jI",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "I",
			},
			{
				pattern: "a",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[aAB]"),
				},
				phonetic: "",
			},
			{
				pattern: "a",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[AB]$"),
				},
				phonetic: "",
			},
			{
				pattern: "A",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "A",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "B",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "B",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "b",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "d",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "f",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "g",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "k",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "l",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "l",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "m",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "m",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "n",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "n",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "p",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "r",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "r",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "t",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "v",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "z",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "n",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bp]"),
				},
				phonetic: "m",
			},
			{
				pattern: "kAg",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phonetic: "(kOg|kO[512])",
			},
			{
				pattern: "kOg",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phonetic: "(kAg|kA[512])",
			},
			{
				pattern: "kog",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phonetic: "(kog|ko[512])",
			},
			{
				pattern: "kag",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[AEOIUaeoiu]"),
				},
				phonetic: "(kag|ka[512])",
			},
			{
				pattern:  "h",
				phonetic: "",
			},
			{
				pattern:  "H",
				phonetic: "(x|)",
			},
			{
				pattern: "F",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]h"),
				},
				phonetic: "e",
			},
			{
				pattern: "F",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]x"),
				},
				phonetic: "e",
			},
			{
				pattern: "B",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]h"),
				},
				phonetic: "a",
			},
			{
				pattern: "B",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdgkpstvzZ]x"),
				},
				phonetic: "a",
			},
			{
				pattern: "e",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "i",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "E",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "I",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "F",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "Q",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "Y",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln]$"),
				},
				phonetic: "",
			},
			{
				pattern: "e",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(e|)",
			},
			{
				pattern: "i",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(i|)",
			},
			{
				pattern: "E",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(E|)",
			},
			{
				pattern: "I",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(I|)",
			},
			{
				pattern: "F",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(F|)",
			},
			{
				pattern: "Q",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(Q|)",
			},
			{
				pattern: "Y",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgklmnprsStvzZ]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[ln][bdfgklmnprsStvzZ]"),
				},
				phonetic: "(Y|)",
			},
			{
				pattern:  "lEs",
				phonetic: "(lEs|lz)",
			},
			{
				pattern: "lE",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[bdfgkmnprStvzZ]$"),
				},
				phonetic: "(lE|l)",
			},
			{
				pattern:  "aue",
				phonetic: "D",
			},
			{
				pattern:  "oue",
				phonetic: "D",
			},
			{
				pattern:  "AvE",
				phonetic: "(D|AvE)",
			},
			{
				pattern:  "Ave",
				phonetic: "(D|Ave)",
			},
			{
				pattern:  "avE",
				phonetic: "(D|avE)",
			},
			{
				pattern:  "ave",
				phonetic: "(D|ave)",
			},
			{
				pattern:  "OvE",
				phonetic: "(D|OvE)",
			},
			{
				pattern:  "Ove",
				phonetic: "(D|Ove)",
			},
			{
				pattern:  "ovE",
				phonetic: "(D|ovE)",
			},
			{
				pattern:  "ove",
				phonetic: "(D|ove)",
			},
			{
				pattern:  "ea",
				phonetic: "(D|ea)",
			},
			{
				pattern:  "EA",
				phonetic: "(D|EA)",
			},
			{
				pattern:  "Ea",
				phonetic: "(D|Ea)",
			},
			{
				pattern:  "eA",
				phonetic: "(D|eA)",
			},
			{
				pattern:  "aji",
				phonetic: "D",
			},
			{
				pattern:  "ajI",
				phonetic: "D",
			},
			{
				pattern:  "aje",
				phonetic: "D",
			},
			{
				pattern:  "ajE",
				phonetic: "D",
			},
			{
				pattern:  "Aji",
				phonetic: "D",
			},
			{
				pattern:  "AjI",
				phonetic: "D",
			},
			{
				pattern:  "Aje",
				phonetic: "D",
			},
			{
				pattern:  "AjE",
				phonetic: "D",
			},
			{
				pattern:  "oji",
				phonetic: "D",
			},
			{
				pattern:  "ojI",
				phonetic: "D",
			},
			{
				pattern:  "oje",
				phonetic: "D",
			},
			{
				pattern:  "ojE",
				phonetic: "D",
			},
			{
				pattern:  "Oji",
				phonetic: "D",
			},
			{
				pattern:  "OjI",
				phonetic: "D",
			},
			{
				pattern:  "Oje",
				phonetic: "D",
			},
			{
				pattern:  "OjE",
				phonetic: "D",
			},
			{
				pattern:  "eji",
				phonetic: "D",
			},
			{
				pattern:  "ejI",
				phonetic: "D",
			},
			{
				pattern:  "eje",
				phonetic: "D",
			},
			{
				pattern:  "ejE",
				phonetic: "D",
			},
			{
				pattern:  "Eji",
				phonetic: "D",
			},
			{
				pattern:  "EjI",
				phonetic: "D",
			},
			{
				pattern:  "Eje",
				phonetic: "D",
			},
			{
				pattern:  "EjE",
				phonetic: "D",
			},
			{
				pattern:  "uji",
				phonetic: "D",
			},
			{
				pattern:  "ujI",
				phonetic: "D",
			},
			{
				pattern:  "uje",
				phonetic: "D",
			},
			{
				pattern:  "ujE",
				phonetic: "D",
			},
			{
				pattern:  "Uji",
				phonetic: "D",
			},
			{
				pattern:  "UjI",
				phonetic: "D",
			},
			{
				pattern:  "Uje",
				phonetic: "D",
			},
			{
				pattern:  "UjE",
				phonetic: "D",
			},
			{
				pattern:  "iji",
				phonetic: "D",
			},
			{
				pattern:  "ijI",
				phonetic: "D",
			},
			{
				pattern:  "ije",
				phonetic: "D",
			},
			{
				pattern:  "ijE",
				phonetic: "D",
			},
			{
				pattern:  "Iji",
				phonetic: "D",
			},
			{
				pattern:  "IjI",
				phonetic: "D",
			},
			{
				pattern:  "Ije",
				phonetic: "D",
			},
			{
				pattern:  "IjE",
				phonetic: "D",
			},
			{
				pattern:  "aja",
				phonetic: "D",
			},
			{
				pattern:  "ajA",
				phonetic: "D",
			},
			{
				pattern:  "ajo",
				phonetic: "D",
			},
			{
				pattern:  "ajO",
				phonetic: "D",
			},
			{
				pattern:  "aju",
				phonetic: "D",
			},
			{
				pattern:  "ajU",
				phonetic: "D",
			},
			{
				pattern:  "Aja",
				phonetic: "D",
			},
			{
				pattern:  "AjA",
				phonetic: "D",
			},
			{
				pattern:  "Ajo",
				phonetic: "D",
			},
			{
				pattern:  "AjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "oja",
				phonetic: "D",
			},
			{
				pattern:  "ojA",
				phonetic: "D",
			},
			{
				pattern:  "ojo",
				phonetic: "D",
			},
			{
				pattern:  "ojO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "Oja",
				phonetic: "D",
			},
			{
				pattern:  "OjA",
				phonetic: "D",
			},
			{
				pattern:  "Ojo",
				phonetic: "D",
			},
			{
				pattern:  "OjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "eja",
				phonetic: "D",
			},
			{
				pattern:  "ejA",
				phonetic: "D",
			},
			{
				pattern:  "ejo",
				phonetic: "D",
			},
			{
				pattern:  "ejO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "Eja",
				phonetic: "D",
			},
			{
				pattern:  "EjA",
				phonetic: "D",
			},
			{
				pattern:  "Ejo",
				phonetic: "D",
			},
			{
				pattern:  "EjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "uja",
				phonetic: "D",
			},
			{
				pattern:  "ujA",
				phonetic: "D",
			},
			{
				pattern:  "ujo",
				phonetic: "D",
			},
			{
				pattern:  "ujO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "Uja",
				phonetic: "D",
			},
			{
				pattern:  "UjA",
				phonetic: "D",
			},
			{
				pattern:  "Ujo",
				phonetic: "D",
			},
			{
				pattern:  "UjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "ija",
				phonetic: "D",
			},
			{
				pattern:  "ijA",
				phonetic: "D",
			},
			{
				pattern:  "ijo",
				phonetic: "D",
			},
			{
				pattern:  "ijO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "Ija",
				phonetic: "D",
			},
			{
				pattern:  "IjA",
				phonetic: "D",
			},
			{
				pattern:  "Ijo",
				phonetic: "D",
			},
			{
				pattern:  "IjO",
				phonetic: "D",
			},
			{
				pattern:  "Aju",
				phonetic: "D",
			},
			{
				pattern:  "AjU",
				phonetic: "D",
			},
			{
				pattern:  "j",
				phonetic: "i",
			},
			{
				pattern: "lYndEr",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lander",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lAndEr",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lAnder",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "landEr",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lender",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lEndEr",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lendEr",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "lEnder",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "lYnder",
			},
			{
				pattern: "bUrk",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "burk",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "bUrg",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "burg",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(burk|berk)",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phonetic: "z",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[rmnl]"),
				},
				phonetic: "z",
			},
			{
				pattern: "s",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				phonetic: "z",
			},
			{
				pattern: "S",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[rmnl]$"),
				},
				phonetic: "z",
			},
			{
				pattern: "dS",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "S",
			},
			{
				pattern: "dZ",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "S",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "S",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(S|s)",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "(S|s)",
			},
			{
				pattern:  "S",
				phonetic: "s",
			},
			{
				pattern:  "dZ",
				phonetic: "z",
			},
			{
				pattern:  "Z",
				phonetic: "z",
			},
		},
		second: map[int64][]rule{
			int64(ashany): []rule{
				{
					pattern:  "b",
					phonetic: "(b|v[1024])",
				},
				{
					pattern:  "J",
					phonetic: "z",
				},
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "AiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "OiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "UiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "AiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "OiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "UiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(o|om[128]|im[128])",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(a|o|on[128]|in[128])",
				},
				{
					pattern:  "B",
					phonetic: "(a|o)",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "AiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "OiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "UiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "AiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "OiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "UiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(i|im[128]|om[128])",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(i|in[128]|on[128])",
				},
				{
					pattern:  "F",
					phonetic: "i",
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiouAEIBFOUQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phonetic: "(Q[16]|i|D[4])",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk[16])",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts[16])",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern:  "I",
					phonetic: "(Q[16]|i)",
				},
				{
					pattern: "lE",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(li|il[4])",
				},
				{
					pattern: "lE",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(li|il[4]|lY[16])",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "Ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "Oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "Ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "ei",
					phonetic: "(D|i)",
				},
				{
					pattern:  "Ei",
					phonetic: "(D|i)",
				},
				{
					pattern: "iA",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ia|io)",
				},
				{
					pattern:  "iA",
					phonetic: "(ia|io|iY[16])",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aeiouAEBFIOU]e"),
					},
					phonetic: "(a|o|Y[16]|D[4])",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("i[^aeiouAEIOU]$"),
					},
					phonetic: "(i|Y[16]|[4])",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("a[^aeiouAEIOU]$"),
					},
					phonetic: "(i|Y[16]|[4])",
				},
				{
					pattern: "e",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "e",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "e",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "e",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "e",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "e",
					phonetic: "(i|Y[16])",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuAOIUQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(i|Y[16])",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprstv]$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					phonetic: "o",
				},
				{
					pattern:  "O",
					phonetic: "(o|Y[16])",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[oeiuQY]$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern:  "A",
					phonetic: "(a|o|Y[16])",
				},
				{
					pattern: "U",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "u",
				},
				{
					pattern: "U",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DoiuQY]$"),
					},
					phonetic: "u",
				},
				{
					pattern: "U",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "u",
				},
				{
					pattern: "Uk",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(uk|Qk[16])",
				},
				{
					pattern: "Uk",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "uk",
				},
				{
					pattern: "sUts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(suts|sQts[16])",
				},
				{
					pattern: "Uts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "uts",
				},
				{
					pattern:  "U",
					phonetic: "(u|Q[16])",
				},
			},
			int64(ashrussian): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					phonetic: "i",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern: "om",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(om|im)",
				},
				{
					pattern: "on",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(on|in)",
				},
				{
					pattern: "em",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(im|om)",
				},
				{
					pattern: "en",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(in|on)",
				},
				{
					pattern: "Em",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(im|Ym|om)",
				},
				{
					pattern: "En",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(in|Yn|on)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			int64(ashcyrillic): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiEIou]$"),
					},
					phonetic: "i",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern: "om",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(om|im)",
				},
				{
					pattern: "on",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(on|in)",
				},
				{
					pattern: "em",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(im|om)",
				},
				{
					pattern: "en",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(in|on)",
				},
				{
					pattern: "Em",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(im|Ym|om)",
				},
				{
					pattern: "En",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(in|Yn|on)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			int64(ashenglish): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^aEIeiou]e"),
					},
					phonetic: "(Q|i|D)",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern: "lE",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[bdfgkmnprsStvzZ]$"),
					},
					phonetic: "(il|li|lY)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("D[^aeiEIou]$"),
					},
					phonetic: "(i|)",
				},
				{
					pattern: "e",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("D[^aeiEIou]$"),
					},
					phonetic: "(i|)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiEuQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQY]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
			},
			int64(ashfrench): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			int64(ashgerman): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiAEIOUouQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern:  "I",
					phonetic: "(Q|i)",
				},
				{
					pattern:  "AU",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "aU",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "Au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "OU",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "oU",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "Ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "Ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "Oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "Ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoAOUiuQY]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoAOQY]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "o",
				},
				{
					pattern: "O",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoAOUeiuQY]$"),
					},
					phonetic: "o",
				},
				{
					pattern:  "O",
					phonetic: "(o|Y)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern: "A",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoeOUiuQY]$"),
					},
					phonetic: "(a|o)",
				},
				{
					pattern:  "A",
					phonetic: "(a|o|Y)",
				},
				{
					pattern: "U",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "u",
				},
				{
					pattern: "U",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuUQY]$"),
					},
					phonetic: "u",
				},
				{
					pattern: "U",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "u",
				},
				{
					pattern: "Uk",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(uk|Qk)",
				},
				{
					pattern: "Uk",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "uk",
				},
				{
					pattern: "sUts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(suts|sQts)",
				},
				{
					pattern: "Uts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "uts",
				},
				{
					pattern:  "U",
					phonetic: "(u|Q)",
				},
			},
			int64(ashhebrew): []rule{},
			int64(ashhungarian): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			int64(ashpolish): []rule{
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(o|om|im)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(o|on|in)",
				},
				{
					pattern:  "B",
					phonetic: "o",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(i|im|om)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(i|in|on)",
				},
				{
					pattern:  "F",
					phonetic: "i",
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phonetic: "i",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			int64(ashromanian): []rule{
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiB",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(o|om|im)",
				},
				{
					pattern: "B",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(o|on|in)",
				},
				{
					pattern:  "B",
					phonetic: "o",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(D|Dm)",
				},
				{
					pattern: "aiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "oiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "uiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "eiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "EiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "iiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "IiF",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(D|Dn)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[bp]"),
					},
					phonetic: "(i|im|om)",
				},
				{
					pattern: "F",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[dgkstvz]"),
					},
					phonetic: "(i|in|on)",
				},
				{
					pattern:  "F",
					phonetic: "i",
				},
				{
					pattern:  "P",
					phonetic: "(o|u)",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aeiAEBFIou]$"),
					},
					phonetic: "i",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprst]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[DaoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
			int64(ashspanish): []rule{
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aEIeiou]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "I",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[^k]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "Ik",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[lr]$"),
					},
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(ik|Qk)",
				},
				{
					pattern: "Ik",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "ik",
				},
				{
					pattern: "sIts",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "(sits|sQts)",
				},
				{
					pattern: "Its",
					rightContext: &ruleMatcher{
						matchEmptyString: true,
						contains:         "",
						prefix:           "",
						suffix:           "",
					},
					phonetic: "its",
				},
				{
					pattern:  "I",
					phonetic: "(i|Q)",
				},
				{
					pattern:  "au",
					phonetic: "(D|a|u)",
				},
				{
					pattern:  "ou",
					phonetic: "(D|o|u)",
				},
				{
					pattern:  "ai",
					phonetic: "(D|a|i)",
				},
				{
					pattern:  "oi",
					phonetic: "(D|o|i)",
				},
				{
					pattern:  "ui",
					phonetic: "(D|u|i)",
				},
				{
					pattern:  "a",
					phonetic: "(a|o)",
				},
				{
					pattern:  "e",
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[fklmnprsStv]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^ts$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					leftContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("[aoiuQ]$"),
					},
					phonetic: "i",
				},
				{
					pattern: "E",
					rightContext: &ruleMatcher{
						matchEmptyString: false,
						contains:         "",
						prefix:           "",
						suffix:           "",
						pattern:          regexp.MustCompile("^[aoQ]"),
					},
					phonetic: "i",
				},
				{
					pattern:  "E",
					phonetic: "(Y|i)",
				},
			},
		},
	},
	exact: finalRule{
		first: []rule{
			{
				pattern: "h",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[fktSs]"),
				},
				phonetic: "p",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "p",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "p",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[gdZz]"),
				},
				phonetic: "b",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "b",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pktSs]"),
				},
				phonetic: "f",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "f",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "f",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgdZz]"),
				},
				phonetic: "v",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "v",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pftSs]"),
				},
				phonetic: "k",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "k",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "k",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bdZz]"),
				},
				phonetic: "g",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "g",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSs]"),
				},
				phonetic: "t",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "t",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "t",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgZz]"),
				},
				phonetic: "d",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "d",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "dZ",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "tS",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkSt]"),
				},
				phonetic: "s",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "s",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[sSzZ]"),
				},
				phonetic: "",
			},
			{
				pattern:  "jnm",
				phonetic: "jm",
			},
			{
				pattern: "ji",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "i",
			},
			{
				pattern: "jI",
				leftContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "I",
			},
			{
				pattern: "a",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[aAB]"),
				},
				phonetic: "",
			},
			{
				pattern: "a",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[AB]$"),
				},
				phonetic: "",
			},
			{
				pattern: "A",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "A",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "B",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "B",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "b",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "b",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "d",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "d",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "f",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "f",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "g",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "g",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "k",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "k",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "l",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "l",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "m",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "m",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "n",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "n",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "p",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "p",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "r",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "r",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "t",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "t",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "v",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "v",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "z",
					suffix:           "",
				},
				phonetic: "",
			},
			{
				pattern:  "H",
				phonetic: "h",
			},
			{
				pattern: "s",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[^t]$"),
				},
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgZd]"),
				},
				phonetic: "z",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[pfkst]"),
				},
				phonetic: "S",
			},
			{
				pattern: "Z",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "S",
			},
			{
				pattern: "S",
				rightContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("^[bgzd]"),
				},
				phonetic: "Z",
			},
			{
				pattern: "z",
				rightContext: &ruleMatcher{
					matchEmptyString: true,
					contains:         "",
					prefix:           "",
					suffix:           "",
				},
				phonetic: "s",
			},
			{
				pattern: "ji",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phonetic: "j",
			},
			{
				pattern: "jI",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phonetic: "j",
			},
			{
				pattern: "je",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phonetic: "j",
			},
			{
				pattern: "jE",
				leftContext: &ruleMatcher{
					matchEmptyString: false,
					contains:         "",
					prefix:           "",
					suffix:           "",
					pattern:          regexp.MustCompile("[aAoOeEiIuU]$"),
				},
				phonetic: "j",
			},
		},
		second: map[int64][]rule{
			int64(ashany): []rule{
				{
					pattern:  "A",
					phonetic: "a",
				},
				{
					pattern:  "B",
					phonetic: "a",
				},
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "F",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
				{
					pattern:  "O",
					phonetic: "o",
				},
				{
					pattern:  "P",
					phonetic: "o",
				},
				{
					pattern:  "U",
					phonetic: "u",
				},
				{
					pattern:  "J",
					phonetic: "l",
				},
			},
			int64(ashrussian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			int64(ashcyrillic): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			int64(ashenglish): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			int64(ashfrench): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			int64(ashgerman): []rule{
				{
					pattern:  "A",
					phonetic: "a",
				},
				{
					pattern:  "B",
					phonetic: "a",
				},
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "F",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
				{
					pattern:  "O",
					phonetic: "o",
				},
				{
					pattern:  "P",
					phonetic: "o",
				},
				{
					pattern:  "U",
					phonetic: "u",
				},
				{
					pattern:  "J",
					phonetic: "l",
				},
			},
			int64(ashhebrew): []rule{},
			int64(ashhungarian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			int64(ashpolish): []rule{
				{
					pattern:  "B",
					phonetic: "a",
				},
				{
					pattern:  "F",
					phonetic: "e",
				},
				{
					pattern:  "P",
					phonetic: "o",
				},
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			int64(ashromanian): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
			int64(ashspanish): []rule{
				{
					pattern:  "E",
					phonetic: "e",
				},
				{
					pattern:  "I",
					phonetic: "i",
				},
			},
		},
	},
}

var ashDiscards = []string{
	"ben",
	"bar",
	"ha",
}
