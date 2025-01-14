// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package environment

// DefaultResourceGroupName returns the resource group name used by default for projects.
// Returns environment variable value by default
// otherwise uses convention, `{AZURE_ENV_NAME}-rg`.
func DefaultResourceGroupName(env *Environment) string {
	resourceGroupName, ok := env.Values[ResourceGroupEnvVarName]
	if ok {
		return resourceGroupName
	}

	return env.GetEnvName() + "-rg"
}
