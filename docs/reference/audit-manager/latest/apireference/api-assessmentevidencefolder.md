# AssessmentEvidenceFolder

The folder where AWS Audit Manager stores evidence for an assessment.

## Contents

**assessmentId**

The identifier for the assessment.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**assessmentReportSelectionCount**

The total count of evidence that's included in the assessment report.

Type: Integer

Required: No

**author**

The name of the user who created the evidence folder.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**controlId**

The unique identifier for the control.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**controlName**

The name of the control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**controlSetId**

The identifier for the control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**dataSource**

The AWS service that the evidence was collected from.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**date**

The date when the first evidence was added to the evidence folder.

Type: Timestamp

Required: No

**evidenceAwsServiceSourceCount**

The total number of AWS resources that were assessed to generate the
evidence.

Type: Integer

Required: No

**evidenceByTypeComplianceCheckCount**

The number of evidence that falls under the compliance check category. This evidence is
collected from AWS Config or AWS Security Hub CSPM.

Type: Integer

Required: No

**evidenceByTypeComplianceCheckIssuesCount**

The total number of issues that were reported directly from AWS Security Hub CSPM,
AWS Config, or both.

Type: Integer

Required: No

**evidenceByTypeConfigurationDataCount**

The number of evidence that falls under the configuration data category. This evidence
is collected from configuration snapshots of other AWS services such as
Amazon EC2, Amazon S3, or IAM.

Type: Integer

Required: No

**evidenceByTypeManualCount**

The number of evidence that falls under the manual category. This evidence is imported
manually.

Type: Integer

Required: No

**evidenceByTypeUserActivityCount**

The number of evidence that falls under the user activity category. This evidence is
collected from AWS CloudTrail logs.

Type: Integer

Required: No

**evidenceResourcesIncludedCount**

The amount of evidence that's included in the evidence folder.

Type: Integer

Required: No

**id**

The identifier for the folder that the evidence is stored in.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**name**

The name of the evidence folder.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**totalEvidence**

The total amount of evidence in the evidence folder.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/assessmentevidencefolder.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/assessmentevidencefolder.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/assessmentevidencefolder.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssessmentControlSet

AssessmentFramework

All content copied from https://docs.aws.amazon.com/.
