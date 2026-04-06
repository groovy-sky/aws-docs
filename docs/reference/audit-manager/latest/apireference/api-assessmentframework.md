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

Type: Array of [AssessmentControlSet](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_AssessmentControlSet.html) objects

Required: No

**id**

The unique identifier for the framework.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**metadata**

The metadata of a framework, such as the name, ID, or description.

Type: [FrameworkMetadata](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_FrameworkMetadata.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/auditmanager-2017-07-25/AssessmentFramework)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/auditmanager-2017-07-25/AssessmentFramework)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/auditmanager-2017-07-25/AssessmentFramework)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssessmentEvidenceFolder

AssessmentFrameworkMetadata
