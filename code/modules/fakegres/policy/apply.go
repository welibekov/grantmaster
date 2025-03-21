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

// Apply applies a set of policies by removing revoked ones and saving the current policies.
// It retrieves existing policies, calculates which ones to revoke, and then
// removes revoked policies from the file system before saving the new set of policies.
func (f *FGPolicy) Apply(_ context.Context, policies []types.Policy) error {
	// Retrieve existing policies from the system.
	existingPolicies, err := f.Get(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get existing policies: %v", err)
	}
	debug.OutputMarshal(existingPolicies, "existing policies")

	// Determine which policies should be revoked by finding the difference
	// between the current set of policies and the existing ones.
	revokePolicies := utils.Diff(policies, existingPolicies)
	debug.OutputMarshal(revokePolicies, "revoke policies")
	debug.OutputMarshal(policies, "apply policies")

	// Remove revoked policies from the file system.
	err = fgUtils.Remove[types.Policy](revokePolicies,
		func(item types.Policy) string {
			return filepath.Join(f.policyDir, item.Username)
		})

	if err != nil {
		return fmt.Errorf("couldn't remove policies: %v", err)
	}

	// Save the new set of policies to the file system.
	err = fgUtils.Save(policies,
		func(item types.Policy) string {
			return filepath.Join(f.policyDir, item.Username)
		})

	if err != nil {
		return fmt.Errorf("couldn't save policies: %v", err)
	}

	return nil // Return nil if everything went smooth.
}
