# UnusedPermissionDetails

Contains information about an unused access finding for a permission. IAM Access Analyzer
charges for unused access analysis based on the number of IAM roles and users analyzed
per month. For more details on pricing, see [IAM Access Analyzer\
pricing](https://aws.amazon.com/iam/access-analyzer/pricing).

## Contents

**serviceNamespace**

The namespace of the AWS service that contains the unused actions.

Type: String

Required: Yes

**actions**

A list of unused actions for which the unused access finding was generated.

Type: Array of [UnusedAction](api-unusedaction.md) objects

Required: No

**lastAccessed**

The time at which the permission was last accessed.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/unusedpermissiondetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/unusedpermissiondetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/unusedpermissiondetails.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnusedIamUserPasswordDetails

UnusedPermissionsRecommendedStep
