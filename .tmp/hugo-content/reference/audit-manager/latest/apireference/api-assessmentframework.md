# AssessmentFramework

The file used to structure and automate AWS Audit Manager assessments for a given
compliance standard.

## Contents

**arn**

The Amazon Resource Name (ARN) of the framework.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:auditmanager:.*`

Required: No

**controlSets**

The control sets that are associated with the framework.

Type: Array of [AssessmentControlSet](api-assessmentcontrolset.md) objects

Required: No

**id**

The unique identifier for the framework.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**metadata**

The metadata of a framework, such as the name, ID, or description.

Type: [FrameworkMetadata](api-frameworkmetadata.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/assessmentframework.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/assessmentframework.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/assessmentframework.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssessmentEvidenceFolder

AssessmentFrameworkMetadata

All content copied from https://docs.aws.amazon.com/.
