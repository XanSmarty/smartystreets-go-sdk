package main

import (
	"context"
	"fmt"
	international_autocomplete "github.com/smartystreets/smartystreets-go-sdk/international-autocomplete-api"
	"github.com/smartystreets/smartystreets-go-sdk/wireup"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ltime | log.Llongfile)

	client := wireup.BuildInternationalAutocompleteAPIClient(
		wireup.WebsiteKeyCredential(os.Getenv("SMARTY_AUTH_WEB"), os.Getenv("SMARTY_AUTH_REFERER")),
		//wireup.SecretKeyCredential(os.Getenv("SMARTY_AUTH_ID"), os.Getenv("SMARTY_AUTH_TOKEN")),
		// The appropriate license values to be used for your subscriptions
		// can be found on the Subscriptions page the account dashboard.
		// https://www.smartystreets.com/docs/cloud/licensing
		wireup.WithLicenses("international-autocomplete-cloud"),
	)

	// Documentation for input fields can be found at:
	// https://smartystreets.com/docs/cloud/us-autocomplete-api#http-request-input-fields

	lookup := &international_autocomplete.Lookup{
		Country:  "FRA",
		Search:   "Louis",
		Locality: "Paris",
	}

	if err := client.SendLookupWithContext(context.Background(), lookup); err != nil {
		log.Fatal("Error sending batch:", err)
	}

	fmt.Printf("Results for input: [%s]\n", lookup.Search)
	for s, candidate := range lookup.Result.Candidates {
		fmt.Printf("#%d: %#v\n", s, candidate)
	}

	log.Println("OK")
}
