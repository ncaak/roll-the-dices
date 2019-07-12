package dice

import (
	"fmt"
	"strings"
)


// Generates a String with a verbose result of the roll
// * Retrieves every result of every check done
// * Retrieves every bonus
func (r *Roller) FormatReply() string {
	var fmtReply strings.Builder
	if strings.TrimSpace(r.command) != "" {
		fmtReply.WriteString(fmt.Sprintf("%s: ", strings.TrimSpace(r.command)))
	}
	// Finds every check and results to write it verbosely
	for index, check := range r.checks {
		// From the first item, following ones are included as a multiple roll
		if index > 0 {
			fmtReply.WriteString("+")
		}
		// Slices are represented with square brackets giving the following format: 1d20[1]
		fmtReply.WriteString(fmt.Sprintf("%dd%d%d", check.dice, check.faces, check.results))

	}
	// Finds every bonus and writes it after the dice
	for _, bonus := range r.bonus {
		// Negative integers have the '-' symbol included, but positives one need to be appended to '+' symbol
		if bonus > 0 {
			fmtReply.WriteString("+")
		}
		fmtReply.WriteString(fmt.Sprintf("%d", bonus))
	}
	// Append equals symbol and the total sum of the roll
	fmtReply.WriteString(fmt.Sprintf("= %d", r.total))

	return fmtReply.String()
}

// Generates a string with distributed results with markdown for an expected complex roll
// * Retrieves every result of every check an creates a line with the subtotal
// * Retrieves every bonus and makes an overall subtotal in a new line
func (r* Roller) RichReply() string {
	var fmtReply strings.Builder
	// Finds every check and results to write it verbosely in a separated line
	for _, ch := range r.checks {
		// Slices are represented with square brackets giving the following format: 1d20[1]
		fmtReply.WriteString(fmt.Sprintf("`%dd%d%d` : %d\n", ch.dice, ch.faces, ch.results, ch.total))
	}
	// Puts a subtotal with the bonuses in a separated line
	fmtReply.WriteString(fmt.Sprintf("_Bonus_ : %d\n", sum(r.bonus)))
	// Writes tag for the roll at the beginning of the ending line with the total of the roll
	var tag = "Total"
	if strings.TrimSpace(r.command) != "" {
		tag = strings.TrimSpace(r.command)
	}
	fmtReply.WriteString(fmt.Sprintf("*%s: %d*", tag, r.total))

	return fmtReply.String()
}
