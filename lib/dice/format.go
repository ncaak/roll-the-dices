package dice

import (
	"fmt"
	"strings"
)

// Entry point to get roll replies. If nothing is set previously just stick with default formatted reply
func (r *Roller) GetReply() string {
	if r.reply == "" {
		r.setDefaultReply()
	}

	return r.reply
}

// Appends checks results with a format (dice)d(faces)[(comma separated valus)]
func (r *Roller) appendChecksResults(builder *strings.Builder) {
	for index, check := range r.checks {
		// From the first item, following ones are included as a multiple roll
		if index > 0 {
			builder.WriteString("+")
		}
		// Slices are represented with square brackets giving the following format: 1d20[1]
		fmt.Fprintf(builder, "%dd%d[%s]", check.dice, check.faces, getStringValues(check.results))
	}
}

// Appends the sum of the bonus if it's different from 0
func (r *Roller) appendBonusTotal(builder *strings.Builder) {
	var total = getSliceItemsTotal(r.bonus)
	if total > 0 {
		fmt.Fprintf(builder, "+%d", total)
	} else if total < 0 {
		fmt.Fprintf(builder, "%d", total)
	}
}

// Generates a string with verbose line and totals for the roll
func (r *Roller) getDistReplyComp(tag string) string {
	var str strings.Builder
	fmt.Fprintf(&str, "_%s_ (%d)` = ", tag, r.total)

	// Formats the checks into a human reading format for each roll check
	r.appendChecksResults(&str)
	// Append the bonus if any (adding + symbol on positive bonus)
	r.appendBonusTotal(&str)

	return str.String() + "`\n"
}

// Generates a text line with the roll resolution
func (r Roller) getRepeatReplyComp() string {
	var str strings.Builder
	// Formats the checks into a human reading format for each roll check
	r.appendChecksResults(&str)
	// Append the bonus if any (adding + symbol on positive bonus)
	r.appendBonusTotal(&str)

	fmt.Fprintf(&str, " = %d", r.total)

	return str.String() + "\n"
}

// Generates a String with a verbose result of the roll
// * Retrieves every result of every check done
// * Retrieves every bonus
func (r *Roller) setDefaultReply() {
	var str strings.Builder
	if strings.TrimSpace(r.command) != "" {
		fmt.Fprintf(&str, "%s: ", strings.TrimSpace(r.command))
	}
	// Formats the checks into a human reading format for each roll check
	r.appendChecksResults(&str)
	// Finds every bonus and writes it after the dice
	r.appendBonusTotal(&str)
	// Append equals symbol and the total sum of the roll
	fmt.Fprintf(&str, "= %d", r.total)

	r.reply = str.String()
}

// Generates a string with distributed results with markdown for an expected complex roll
// * Retrieves every result of every check an creates a line with the subtotal
// * Retrieves every bonus and makes an overall subtotal in a new line
func (r *Roller) setGroupedReply() {
	var str strings.Builder
	// Finds every check and results to write it verbosely in a separated line
	for _, ch := range r.checks {
		// Formats the checks into a human reading format
		fmt.Fprintf(&str, "`%dd%d[%s]` : %d\n", ch.dice, ch.faces, getStringValues(ch.results), ch.total)
	}
	// Puts a subtotal with the bonuses in a separated line
	fmt.Fprintf(&str, "_Bonus_ : %d\n", getSliceItemsTotal(r.bonus))
	// Writes tag for the roll at the beginning of the ending line with the total of the roll
	var tag = "Total"
	if strings.TrimSpace(r.command) != "" {
		tag = strings.TrimSpace(r.command)
	}
	fmt.Fprintf(&str, "*%s: %d*", tag, r.total)

	r.reply = str.String()
}
