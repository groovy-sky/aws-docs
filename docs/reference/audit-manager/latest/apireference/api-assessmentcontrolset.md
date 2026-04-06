# AssessmentControlSet

Represents a set of controls in an Audit Manager assessment.

## Contents

**controls**

The list of controls that's contained with the control set.

Type: Array of [AssessmentControl](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_AssessmentControl.html) objects

Required: No

**delegations**

The delegations that are associated with the control set.

Type: Array of [Delegation](api-delegation.md) objects

Required: No

**description**

The description for the control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

Required: No

**id**

The identifier of the control set in the assessment. This is the control set name in a
plain string format.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**manualEvidenceCount**

The total number of evidence objects that are uploaded manually to the control set.

Type: Integer

Required: No

**roles**

The roles that are associated with the control set.

Type: Array of [Role](api-role.md) objects

Required: No

**status**

The current status of the control set.

Type: String

Valid Values: `ACTIVE | UNDER_REVIEW | REVIEWED`

Required: No

**systemEvidenceCount**

The total number of evidence objects that are retrieved automatically for the control
set.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/auditmanager-2017-07-25/AssessmentControlSet)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/auditmanager-2017-07-25/AssessmentControlSet)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/auditmanager-2017-07-25/AssessmentControlSet)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssessmentControl

AssessmentEvidenceFolder
