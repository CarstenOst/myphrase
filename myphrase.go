package myphrase

import "rsc.io/quote"

func PithySay(arrayIndex int) string {

	pithySayings := [4]string{quote.Glass(), quote.Go(), quote.Hello(), quote.Opt()}

	return pithySayings[arrayIndex]

}
