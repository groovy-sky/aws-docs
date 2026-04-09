# EnhancedImageScanFinding

The details of an enhanced image scan. This is returned when enhanced scanning is
enabled for your private registry.

## Contents

**awsAccountId**

The AWS account ID associated with the image.

Type: String

Pattern: `[0-9]{12}`

Required: No

**description**

The description of the finding.

Type: String

Required: No

**exploitAvailable**

If a finding discovered in your environment has an exploit available.

Type: String

Required: No

**findingArn**

The Amazon Resource Number (ARN) of the finding.

Type: String

Required: No

**firstObservedAt**

The date and time that the finding was first observed.

Type: Timestamp

Required: No

**fixAvailable**

Details on whether a fix is available through a version update. This value can be
`YES`, `NO`, or `PARTIAL`. A `PARTIAL`
fix means that some, but not all, of the packages identified in the finding have fixes
available through updated versions.

Type: String

Required: No

**lastObservedAt**

The date and time that the finding was last observed.

Type: Timestamp

Required: No

**packageVulnerabilityDetails**

An object that contains the details of a package vulnerability finding.

Type: [PackageVulnerabilityDetails](api-packagevulnerabilitydetails.md) object

Required: No

**remediation**

An object that contains the details about how to remediate a finding.

Type: [Remediation](api-remediation.md) object

Required: No

**resources**

Contains information on the resources involved in a finding.

Type: Array of [Resource](api-resource.md) objects

Required: No

**score**

The Amazon Inspector score given to the finding.

Type: Double

Required: No

**scoreDetails**

An object that contains details of the Amazon Inspector score.

Type: [ScoreDetails](api-scoredetails.md) object

Required: No

**severity**

The severity of the finding.

Type: String

Required: No

**status**

The status of the finding.

Type: String

Required: No

**title**

The title of the finding.

Type: String

Required: No

**type**

The type of the finding.

Type: String

Required: No

**updatedAt**

The date and time the finding was last updated at.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/enhancedimagescanfinding.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/enhancedimagescanfinding.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/enhancedimagescanfinding.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfigurationForRepositoryCreationTemplate

Image

All content copied from https://docs.aws.amazon.com/.
