package cedato

import (
	"testing"
	"text/template"

	"github.com/prebid/prebid-server/privacy"
	"github.com/stretchr/testify/assert"
)

/// TBD

func TestCedatoSyncer(t *testing.T) {
	syncURL := "http://test-bh.ybp.yahoo.com/sync/appnexuspbs?gdpr={{.GDPR}}&euconsent={{.GDPRConsent}}&us_privacy={{.USPrivacy}}&url=localhost%2Fsetuid%3Fbidder%3Dcedato%26gdpr%3D{{.GDPR}}%26gdpr_consent%3D{{.GDPRConsent}}%26uid%3D%24%7BUID%7D"
	syncURLTemplate := template.Must(
		template.New("sync-template").Parse(syncURL),
	)

	syncer := NewCedatoSyncer(syncURLTemplate)
	syncInfo, err := syncer.GetUsersyncInfo(privacy.Policies{})

	assert.NoError(t, err)
	assert.Equal(t, "http://test-bh.ybp.yahoo.com/sync/appnexuspbs?gdpr=&euconsent=&us_privacy=&url=localhost%2Fsetuid%3Fbidder%3Dcedato%26gdpr%3D%26gdpr_consent%3D%26uid%3D%24%7BUID%7D", syncInfo.URL)
	assert.Equal(t, "redirect", syncInfo.Type)
	assert.EqualValues(t, 672, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
}
