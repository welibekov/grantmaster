package policy

import (
	"context"
	"fmt"
	"path/filepath"

	fgUtils "github.com/welibekov/grantmaster/modules/fakegres/utils"
	"github.com/welibekov/grantmaster/modules/policy/types"
	"github.com/welibekov/grantmaster/modules/policy/utils"
	"github.com/welibekov/grantmaster/modules/utils/debug"
)

func (f *FGPolicy) Apply(_ context.Context, policies []types.Policy) error {
	existingPolicies, err := f.Get(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get exising policies: %v", err)
	}
	debug.OutputMarshal(existingPolicies, "existing policies")

	revokePolicies := utils.Diff(policies, existingPolicies)
	debug.OutputMarshal(revokePolicies, "revoke policies")
	debug.OutputMarshal(policies, "apply policies")

	// Remove revoked policies.
	err = fgUtils.Remove[types.Policy](revokePolicies,
		func(item types.Policy) string {
			return filepath.Join(f.policyDir, item.Username)
		})

	if err != nil {
		return fmt.Errorf("couldn't remove policies: %v", err)
	}

	err = fgUtils.Save(policies,
		func(item types.Policy) string {
			return filepath.Join(f.policyDir, item.Username)
		})

	if err != nil {
		return fmt.Errorf("couldn't save policies: %v", err)
	}

	return nil
}
