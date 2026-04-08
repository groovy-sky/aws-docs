# AnalyzerSummary

Contains information about the analyzer.

## Contents

**arn**

The ARN of the analyzer.

Type: String

Pattern: `[^:]*:[^:]*:[^:]*:[^:]*:[^:]*:analyzer/.{1,255}`

Required: Yes

**createdAt**

A timestamp for the time at which the analyzer was created.

Type: Timestamp

Required: Yes

**name**

The name of the analyzer.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[A-Za-z][A-Za-z0-9_.-]*`

Required: Yes

**status**

The status of the analyzer. An `Active` analyzer successfully monitors
supported resources and generates new findings. The analyzer is `Disabled` when
a user action, such as removing trusted access for AWS Identity and Access Management Access Analyzer from AWS Organizations, causes
the analyzer to stop generating new findings. The status is `Creating` when the
analyzer creation is in progress and `Failed` when the analyzer creation has
failed.

Type: String

Valid Values: `ACTIVE | CREATING | DISABLED | FAILED`

Required: Yes

**type**

The type represents the zone of trust or scope for the analyzer.

Type: String

Valid Values: `ACCOUNT | ORGANIZATION | ACCOUNT_UNUSED_ACCESS | ORGANIZATION_UNUSED_ACCESS | ACCOUNT_INTERNAL_ACCESS | ORGANIZATION_INTERNAL_ACCESS`

Required: Yes

**configuration**

Specifies if the analyzer is an external access, unused access, or internal access
analyzer. The [GetAnalyzer](api-getanalyzer.md)
action includes this property in its response if a configuration is specified, while the
[ListAnalyzers](api-listanalyzers.md)
action omits it.

Type: [AnalyzerConfiguration](api-analyzerconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**lastResourceAnalyzed**

The resource that was most recently analyzed by the analyzer.

Type: String

Required: No

**lastResourceAnalyzedAt**

The time at which the most recently analyzed resource was analyzed.

Type: Timestamp

Required: No

**statusReason**

The `statusReason` provides more details about the current status of the
analyzer. For example, if the creation for the analyzer fails, a `Failed` status
is returned. For an analyzer with organization as the type, this failure can be due to an
issue with creating the service-linked roles required in the member accounts of the AWS
organization.

Type: [StatusReason](api-statusreason.md) object

Required: No

**tags**

An array of key-value pairs applied to the analyzer. The key-value pairs consist of the
set of Unicode letters, digits, whitespace, `_`, `.`, `/`,
`=`, `+`, and `-`.

The tag key is a value that is 1 to 128 characters in length and cannot be prefixed with
`aws:`.

The tag value is a value that is 0 to 256 characters in length.

Type: String to string map

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/analyzersummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/analyzersummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/analyzersummary.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AnalyzerConfiguration

ArchiveRuleSummary
