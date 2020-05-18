package cedato

import (
	"testing"

	"github.com/prebid/prebid-server/adapters/adapterstest"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "cedatotest", NewCedatoBidder("https://h.cedatoplayer.com/hb"))
}
