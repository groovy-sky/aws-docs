# AssessmentFrameworkMetadata

The metadata that's associated with a standard framework or a custom framework.

## Contents

**arn**

The Amazon Resource Name (ARN) of the framework.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:auditmanager:.*`

Required: No

**complianceType**

The compliance type that the new custom framework supports, such as CIS or HIPAA.

Type: String

Length Constraints: Maximum length of 100.

Pattern: `^[\w\W\s\S]*$`

Required: No

**controlsCount**

The number of controls that are associated with the framework.

Type: Integer

Required: No

**controlSetsCount**

The number of control sets that are associated with the framework.

Type: Integer

Required: No

**createdAt**

The time when the framework was created.

Type: Timestamp

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

**type**

The framework type, such as a standard framework or a custom framework.

Type: String

Valid Values: `Standard | Custom`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/assessmentframeworkmetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/assessmentframeworkmetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/assessmentframeworkmetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssessmentFramework

AssessmentFrameworkShareRequest

All content copied from https://docs.aws.amazon.com/.
