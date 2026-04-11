# AssessmentControl

The control entity that represents a standard control or a custom control in an Audit Manager assessment.

## Contents

**assessmentReportEvidenceCount**

The amount of evidence in the assessment report.

Type: Integer

Required: No

**comments**

The list of comments that's attached to the control.

Type: Array of [ControlComment](api-controlcomment.md) objects

Required: No

**description**

_This member has been deprecated._

The description of the control.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**evidenceCount**

The amount of evidence that's collected for the control.

Type: Integer

Required: No

**evidenceSources**

The list of data sources for the evidence.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `.*\S.*`

Required: No

**id**

The identifier for the control.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**name**

The name of the control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**response**

The response of the control.

Type: String

Valid Values: `MANUAL | AUTOMATE | DEFER | IGNORE`

Required: No

**status**

The status of the control.

Type: String

Valid Values: `UNDER_REVIEW | REVIEWED | INACTIVE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/assessmentcontrol.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/assessmentcontrol.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/assessmentcontrol.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Assessment

AssessmentControlSet

All content copied from https://docs.aws.amazon.com/.
