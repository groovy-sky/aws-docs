# AssessmentMetadata

The metadata that's associated with the specified assessment.

## Contents

**assessmentReportsDestination**

The destination that evidence reports are stored in for the assessment.

Type: [AssessmentReportsDestination](api-assessmentreportsdestination.md) object

Required: No

**complianceType**

The name of the compliance standard that's related to the assessment, such as PCI-DSS.

Type: String

Length Constraints: Maximum length of 100.

Pattern: `^[\w\W\s\S]*$`

Required: No

**creationTime**

Specifies when the assessment was created.

Type: Timestamp

Required: No

**delegations**

The delegations that are associated with the assessment.

Type: Array of [Delegation](api-delegation.md) objects

Required: No

**description**

The description of the assessment.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**id**

The unique identifier for the assessment.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**lastUpdated**

The time of the most recent update.

Type: Timestamp

Required: No

**name**

The name of the assessment.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**roles**

The roles that are associated with the assessment.

Type: Array of [Role](api-role.md) objects

Required: No

**scope**

The wrapper of AWS accounts and services that are in scope for the
assessment.

Type: [Scope](api-scope.md) object

Required: No

**status**

The overall status of the assessment.

Type: String

Valid Values: `ACTIVE | INACTIVE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/auditmanager-2017-07-25/AssessmentMetadata)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/auditmanager-2017-07-25/AssessmentMetadata)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/auditmanager-2017-07-25/AssessmentMetadata)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssessmentFrameworkShareRequest

AssessmentMetadataItem
