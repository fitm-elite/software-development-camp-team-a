package timezone_test

import (
	"testing"
	"time"

	"github.com/fitm-elite/elebs/packages/timezone"
)

// TestTimeZoneIsAsiaBangkok is a test function for timezone that's Asia/Bangkok?.
func TestTimeZoneIsAsiaBangkok(t *testing.T) {
	t.Parallel()

	time.Local = timezone.NewAsiaBangkok()

	t.Run("TestTimeZoneIsAsiaBangkok", func(t *testing.T) {
		t.Parallel()

		got := time.Local.String()
		want := "Asia/Bangkok"

		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}
	})
}
