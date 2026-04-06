# TrailProperties

Contains details about the CloudTrail trail being analyzed to generate a policy.

## Contents

**cloudTrailArn**

Specifies the ARN of the trail. The format of a trail ARN is
`arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail`.

Type: String

Pattern: `arn:[^:]*:cloudtrail:[^:]*:[^:]*:trail/.{1,576}`

Required: Yes

**allRegions**

Possible values are `true` or `false`. If set to
`true`, IAM Access Analyzer retrieves CloudTrail data from all regions to analyze and
generate a policy.

Type: Boolean

Required: No

**regions**

A list of regions to get CloudTrail data from and analyze to generate a policy.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/trailproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/trailproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/trailproperties.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Trail

UnusedAccessConfiguration
