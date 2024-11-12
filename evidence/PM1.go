package evidence

import (
	"strconv"
	"strings"
)

var (
	pm1Ext   = 3
	pm1Count = 2
	oe       float32
)

// PM1
func CheckPM1(item map[string]string) string {
	if PP2PM1_special[item["entrez_id"]] || !ismissense.MatchString(item["Function"]) {
		return "0"
	}
	if item["entrezID_oe"] == "-" {
		return "0"
	}
	id_oe_pairs := strings.Split(item["entrezID_oe"], ",") //xxx_0.31693,2475_0.30241
	for _, pair := range id_oe_pairs {
		pairs := strings.Split(pair, "_")
		gene_id := pairs[0]
		oe, _ := strconv.ParseFloat(pairs[1], 32)

		if gene_id == item["entrez_id"] {
			//log.Printf("num, geneid,entrezID_oe: %d: %s:%f, %s", len(id_oe_pairs), item["entrez_id"], oe, item["entrezID_oe"])
			if oe <= 0.2112 {
				return "1"
			} else if oe <= 0.3747 {
				return "Supporting"
			} else {
				return "0"
			}
		}
	}

	return "0"

}

func ComparePM1(item map[string]string) {
	rule := "PM1"
	val := CheckPM1(item)
	if val != item[rule] {
		PrintConflict(item, rule, val, "Interpro_domain", "pfamId")
	}
}
