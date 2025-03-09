package utils

//import (
//	"github.com/welibekov/grantmaster/modules/policy/types"
//	"github.com/welibekov/grantmaster/modules/utils"
//)
//
//// WhatToGrant function to compare two slices of Policy and return unique roles in A not found in B
//func WhatToGrant(A, B []types.Policy) []types.Policy {
//	// Create a map to hold roles from policies in B
//	roleMap := make(map[string]map[string]struct{})
//
//	// Populate the roleMap with roles for each user in B
//	for _, policy := range B {
//		if roleMap[policy.Username] == nil {
//			roleMap[policy.Username] = make(map[string]struct{})
//		}
//		for _, role := range policy.Roles {
//			roleMap[policy.Username][role] = struct{}{}
//		}
//	}
//
//	var result []types.Policy
//
//	// Iterate through policies in A and find unique roles
//	for _, policyA := range A {
//		var uniqueRoles []string
//
//		for _, roleA := range policyA.Roles {
//			// Check if the role does not exist in the corresponding user in B
//			if rolesInB, exists := roleMap[policyA.Username]; !exists || !utils.InMap(rolesInB, roleA) {
//				uniqueRoles = append(uniqueRoles, roleA)
//			}
//		}
//
//		// If we have unique roles, create a new Policy to add to the result
//		if len(uniqueRoles) > 0 {
//			result = append(result, types.Policy{
//				Username: policyA.Username,
//				Roles:    uniqueRoles,
//			})
//		}
//	}
//
//	return result
//}
