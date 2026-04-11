# AssessmentReport

A finalized document that's generated from an Audit Manager assessment. These
reports summarize the relevant evidence that was collected for your audit, and link to the
relevant evidence folders. These evidence folders are named and organized according to the
controls that are specified in your assessment.

## Contents

**assessmentId**

The identifier for the specified assessment.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**assessmentName**

The name of the associated assessment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**author**

The name of the user who created the assessment report.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `^[a-zA-Z0-9-_()\s\+=,.@]+$`

Required: No

**awsAccountId**

The identifier for the specified AWS account.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]{12}$`

Required: No

**creationTime**

Specifies when the assessment report was created.

Type: Timestamp

Required: No

**description**

The description of the specified assessment report.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**id**

The unique identifier for the assessment report.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**name**

The name that's given to the assessment report.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[a-zA-Z0-9-_\.]+$`

Required: No

**status**

The current status of the specified assessment report.

Type: String

Valid Values: `COMPLETE | IN_PROGRESS | FAILED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/assessmentreport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/assessmentreport.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/assessmentreport.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssessmentMetadataItem

AssessmentReportEvidenceError

All content copied from https://docs.aws.amazon.com/.
