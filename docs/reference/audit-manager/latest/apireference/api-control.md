# Control

A control in AWS Audit Manager.

## Contents

**actionPlanInstructions**

The recommended actions to carry out if the control isn't fulfilled.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**actionPlanTitle**

The title of the action plan for remediating the control.

Type: String

Length Constraints: Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**arn**

The Amazon Resource Name (ARN) of the control.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:auditmanager:.*`

Required: No

**controlMappingSources**

The data mapping sources for the control.

Type: Array of [ControlMappingSource](api-controlmappingsource.md) objects

Array Members: Minimum number of 1 item.

Required: No

**controlSources**

The data source types that determine where Audit Manager collects evidence from for
the control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z_0-9-\s.,]+$`

Required: No

**createdAt**

The time when the control was created.

Type: Timestamp

Required: No

**createdBy**

The user or role that created the control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z0-9\s-_()\[\]]+$`

Required: No

**description**

The description of the control.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**id**

The unique identifier for the control.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**lastUpdatedAt**

The time when the control was most recently updated.

Type: Timestamp

Required: No

**lastUpdatedBy**

The user or role that most recently updated the control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z0-9\s-_()\[\]]+$`

Required: No

**name**

The name of the control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

**state**

The state of the control. The `END_OF_SUPPORT` state is applicable to
standard controls only. This state indicates that the standard control can still be used to
collect evidence, but Audit Manager is no longer updating or maintaining that
control.

Type: String

Valid Values: `ACTIVE | END_OF_SUPPORT`

Required: No

**tags**

The tags associated with the control.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Value Pattern: `.{0,255}`

Required: No

**testingInformation**

The steps that you should follow to determine if the control has been satisfied.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**type**

Specifies whether the control is a standard control or a custom control.

Type: String

Valid Values: `Standard | Custom | Core`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/control.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/control.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/control.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChangeLog

ControlComment

All content copied from https://docs.aws.amazon.com/.
