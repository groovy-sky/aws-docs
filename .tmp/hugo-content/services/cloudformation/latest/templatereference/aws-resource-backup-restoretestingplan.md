This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::RestoreTestingPlan

Creates a restore testing plan.

The first of two steps to create a restore testing
plan. After this request is successful, finish the procedure using
CreateRestoreTestingSelection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Backup::RestoreTestingPlan",
  "Properties" : {
      "RecoveryPointSelection" : RestoreTestingRecoveryPointSelection,
      "RestoreTestingPlanName" : String,
      "ScheduleExpression" : String,
      "ScheduleExpressionTimezone" : String,
      "StartWindowHours" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Backup::RestoreTestingPlan
Properties:
  RecoveryPointSelection:
    RestoreTestingRecoveryPointSelection
  RestoreTestingPlanName: String
  ScheduleExpression: String
  ScheduleExpressionTimezone: String
  StartWindowHours: Integer
  Tags:
    - Tag

```

## Properties

`RecoveryPointSelection`

The specified criteria to assign a set of resources, such as
recovery point types or backup vaults.

_Required_: Yes

_Type_: [RestoreTestingRecoveryPointSelection](aws-properties-backup-restoretestingplan-restoretestingrecoverypointselection.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestoreTestingPlanName`

The RestoreTestingPlanName is a unique string that is the name
of the restore testing plan. This cannot be changed after creation,
and it must consist of only alphanumeric characters and underscores.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScheduleExpression`

A CRON expression in specified timezone when a restore
testing plan is executed. When no CRON expression is provided, AWS Backup will use the default
expression `cron(0 5 ? * * *)`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpressionTimezone`

Optional. This is the timezone in which the schedule
expression is set. By default, ScheduleExpressions are in UTC.
You can modify this to a specified timezone.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartWindowHours`

Defaults to 24 hours.

A value in hours after a
restore test is scheduled before a job will be canceled if it
doesn't start successfully. This value is optional. If this value
is included, this parameter has a maximum value of 168 hours
(one week).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional tags to include. A tag is a key-value pair you can use to manage, filter, and
search for your resources. Allowed characters include UTF-8 letters,numbers, spaces, and
the following characters: `+ - = . _ : /.`

_Required_: No

_Type_: Array of [Tag](aws-properties-backup-restoretestingplan-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`RestoreTestingPlanArn`

An Amazon Resource Name (ARN) that uniquely identifies
a restore testing plan.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

RestoreTestingRecoveryPointSelection

All content copied from https://docs.aws.amazon.com/.
