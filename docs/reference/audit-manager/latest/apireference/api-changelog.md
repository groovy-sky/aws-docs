# ChangeLog

The record of a change within AWS Audit Manager. For example, this could be the
status change of an assessment or the delegation of a control set.

## Contents

**action**

The action that was performed.

Type: String

Valid Values: `CREATE | UPDATE_METADATA | ACTIVE | INACTIVE | DELETE | UNDER_REVIEW | REVIEWED | IMPORT_EVIDENCE`

Required: No

**createdAt**

The time when the action was performed and the changelog record was created.

Type: Timestamp

Required: No

**createdBy**

The user or role that performed the action.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:iam:.*`

Required: No

**objectName**

The name of the object that changed. This could be the name of an assessment, control,
or control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

Required: No

**objectType**

The object that was changed, such as an assessment, control, or control set.

Type: String

Valid Values: `ASSESSMENT | CONTROL_SET | CONTROL | DELEGATION | ASSESSMENT_REPORT`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/changelog.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/changelog.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/changelog.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchImportEvidenceToAssessmentControlError

Control

All content copied from https://docs.aws.amazon.com/.
