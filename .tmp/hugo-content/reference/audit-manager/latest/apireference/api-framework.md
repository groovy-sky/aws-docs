# Framework

The file that's used to structure and automate AWS Audit Manager assessments for a
given compliance standard.

## Contents

**arn**

The Amazon Resource Name (ARN) of the framework.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:auditmanager:.*`

Required: No

**complianceType**

The compliance type that the framework supports, such as CIS or HIPAA.

Type: String

Length Constraints: Maximum length of 100.

Pattern: `^[\w\W\s\S]*$`

Required: No

**controlSets**

The control sets that are associated with the framework.

###### Note

The `Controls` object returns a partial response when called through
Framework APIs. For a complete `Controls` object, use
`GetControl`.

Type: Array of [ControlSet](api-controlset.md) objects

Array Members: Minimum number of 1 item.

Required: No

**controlSources**

_This member has been deprecated._

The control data sources where Audit Manager collects evidence from.

###### Important

This API parameter is no longer supported.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z_0-9-\s.,]+$`

Required: No

**createdAt**

The time when the framework was created.

Type: Timestamp

Required: No

**createdBy**

The user or role that created the framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z0-9\s-_()\[\]]+$`

Required: No

**description**

The description of the framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**id**

The unique identifier for the framework.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**lastUpdatedAt**

The time when the framework was most recently updated.

Type: Timestamp

Required: No

**lastUpdatedBy**

The user or role that most recently updated the framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z0-9\s-_()\[\]]+$`

Required: No

**logo**

The logo that's associated with the framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[\w,\s-]+\.[A-Za-z]+$`

Required: No

**name**

The name of the framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**tags**

The tags that are associated with the framework.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Value Pattern: `.{0,255}`

Required: No

**type**

Specifies whether the framework is a standard framework or a custom framework.

Type: String

Valid Values: `Standard | Custom`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/framework.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/framework.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/framework.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvidenceInsights

FrameworkMetadata

All content copied from https://docs.aws.amazon.com/.
