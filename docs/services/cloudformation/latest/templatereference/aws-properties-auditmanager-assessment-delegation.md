This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AuditManager::Assessment Delegation

The `Delegation` property type specifies the assignment of a control set to a delegate for review.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssessmentId" : String,
  "AssessmentName" : String,
  "Comment" : String,
  "ControlSetId" : String,
  "CreatedBy" : String,
  "CreationTime" : Number,
  "Id" : String,
  "LastUpdated" : Number,
  "RoleArn" : String,
  "RoleType" : String,
  "Status" : String
}

```

### YAML

```yaml

  AssessmentId: String
  AssessmentName: String
  Comment: String
  ControlSetId: String
  CreatedBy: String
  CreationTime: Number
  Id: String
  LastUpdated: Number
  RoleArn: String
  RoleType: String
  Status: String

```

## Properties

`AssessmentId`

The identifier for the assessment that's associated with the delegation.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssessmentName`

The name of the assessment that's associated with the delegation.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_\.]+$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Comment`

The comment that's related to the delegation.

_Required_: No

_Type_: String

_Pattern_: `^[\w\W\s\S]*$`

_Maximum_: `350`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ControlSetId`

The identifier for the control set that's associated with the delegation.

_Required_: No

_Type_: String

_Pattern_: `^[\w\W\s\S]*$`

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedBy`

The user or role that created the delegation.

_Minimum_: `1`

_Maximum_: `100`

_Pattern_: `^[a-zA-Z0-9-_()\\[\\]\\s]+$`

_Required_: No

_Type_: String

_Pattern_: `^arn:.*:*:.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreationTime`

Specifies when the delegation was created.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The unique identifier for the delegation.

_Required_: No

_Type_: String

_Pattern_: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastUpdated`

Specifies when the delegation was last updated.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*:iam:.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleType`

The type of customer persona.

###### Note

In `CreateAssessment`, `roleType` can only be
`PROCESS_OWNER`.

In `UpdateSettings`, `roleType` can only be
`PROCESS_OWNER`.

In `BatchCreateDelegationByAssessment`, `roleType` can only be
`RESOURCE_OWNER`.

_Required_: No

_Type_: String

_Allowed values_: `PROCESS_OWNER | RESOURCE_OWNER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the delegation.

_Required_: No

_Type_: String

_Allowed values_: `IN_PROGRESS | UNDER_REVIEW | COMPLETE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Delegation](../../../../reference/audit-manager/latest/apireference/api-delegation.md)
in the _AWS Audit Manager API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSService

Role

All content copied from https://docs.aws.amazon.com/.
