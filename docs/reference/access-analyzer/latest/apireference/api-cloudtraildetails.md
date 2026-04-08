# CloudTrailDetails

Contains information about CloudTrail access.

## Contents

**accessRole**

The ARN of the service role that IAM Access Analyzer uses to access your CloudTrail trail and
service last accessed information.

Type: String

Pattern: `arn:[^:]*:iam::[^:]*:role/.{1,576}`

Required: Yes

**startTime**

The start of the time range for which IAM Access Analyzer reviews your CloudTrail events. Events
with a timestamp before this time are not considered to generate a policy.

Type: Timestamp

Required: Yes

**trails**

A `Trail` object that contains settings for a trail.

Type: Array of [Trail](api-trail.md) objects

Required: Yes

**endTime**

The end of the time range for which IAM Access Analyzer reviews your CloudTrail events. Events with
a timestamp after this time are not considered to generate a policy. If this is not
included in the request, the default value is the current time.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/cloudtraildetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/cloudtraildetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/cloudtraildetails.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ArchiveRuleSummary

CloudTrailProperties
