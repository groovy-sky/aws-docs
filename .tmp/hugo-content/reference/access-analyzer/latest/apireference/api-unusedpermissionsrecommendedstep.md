# UnusedPermissionsRecommendedStep

Contains information about the action to take for a policy in an unused permissions
finding.

## Contents

**recommendedAction**

A recommendation of whether to create or detach a policy for an unused permissions
finding.

Type: String

Valid Values: `CREATE_POLICY | DETACH_POLICY`

Required: Yes

**existingPolicyId**

If the recommended action for the unused permissions finding is to detach a policy, the
ID of an existing policy to be detached.

Type: String

Required: No

**policyUpdatedAt**

The time at which the existing policy for the unused permissions finding was last
updated.

Type: Timestamp

Required: No

**recommendedPolicy**

If the recommended action for the unused permissions finding is to replace the existing
policy, the contents of the recommended policy to replace the policy specified in the
`existingPolicyId` field.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/unusedpermissionsrecommendedstep.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/unusedpermissionsrecommendedstep.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/unusedpermissionsrecommendedstep.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnusedPermissionDetails

ValidatePolicyFinding
