# Assessment

An entity that defines the scope of audit evidence collected by AWS Audit Manager.
An Audit Manager assessment is an implementation of an Audit Manager framework.

## Contents

**arn**

The Amazon Resource Name (ARN) of the assessment.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:auditmanager:.*`

Required: No

**awsAccount**

The AWS account that's associated with the assessment.

Type: [AWSAccount](api-awsaccount.md) object

Required: No

**framework**

The framework that the assessment was created from.

Type: [AssessmentFramework](api-assessmentframework.md) object

Required: No

**metadata**

The metadata for the assessment.

Type: [AssessmentMetadata](api-assessmentmetadata.md) object

Required: No

**tags**

The tags that are associated with the assessment.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Value Pattern: `.{0,255}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/assessment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/assessment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/assessment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Data Types

AssessmentControl
