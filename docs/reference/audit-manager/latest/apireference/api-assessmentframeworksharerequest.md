# AssessmentFrameworkShareRequest

Represents a share request for a custom framework in AWS Audit Manager.

## Contents

**comment**

An optional comment from the sender about the share request.

Type: String

Length Constraints: Maximum length of 500.

Pattern: `^[\w\W\s\S]*$`

Required: No

**complianceType**

The compliance type that the shared custom framework supports, such as CIS or
HIPAA.

Type: String

Length Constraints: Maximum length of 100.

Pattern: `^[\w\W\s\S]*$`

Required: No

**creationTime**

The time when the share request was created.

Type: Timestamp

Required: No

**customControlsCount**

The number of custom controls that are part of the shared custom framework.

Type: Integer

Required: No

**destinationAccount**

The AWS account of the recipient.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]{12}$`

Required: No

**destinationRegion**

The AWS Region of the recipient.

Type: String

Pattern: `^[a-z]{2}-[a-z]+-[0-9]{1}$`

Required: No

**expirationTime**

The time when the share request expires.

Type: Timestamp

Required: No

**frameworkDescription**

The description of the shared custom framework.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**frameworkId**

The unique identifier for the shared custom framework.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**frameworkName**

The name of the custom framework that the share request is for.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**id**

The unique identifier for the share request.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**lastUpdated**

Specifies when the share request was last updated.

Type: Timestamp

Required: No

**sourceAccount**

The AWS account of the sender.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]{12}$`

Required: No

**standardControlsCount**

The number of standard controls that are part of the shared custom framework.

Type: Integer

Required: No

**status**

The status of the share request.

Type: String

Valid Values: `ACTIVE | REPLICATING | SHARED | EXPIRING | FAILED | EXPIRED | DECLINED | REVOKED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/assessmentframeworksharerequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/assessmentframeworksharerequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/assessmentframeworksharerequest.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssessmentFrameworkMetadata

AssessmentMetadata

All content copied from https://docs.aws.amazon.com/.
