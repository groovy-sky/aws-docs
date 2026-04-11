This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::ScheduledAction

Creates a scheduled action. A scheduled action contains a schedule and an Amazon Redshift API action.
For example, you can create a schedule of when to run the `ResizeCluster` API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::ScheduledAction",
  "Properties" : {
      "Enable" : Boolean,
      "EndTime" : String,
      "IamRole" : String,
      "Schedule" : String,
      "ScheduledActionDescription" : String,
      "ScheduledActionName" : String,
      "StartTime" : String,
      "TargetAction" : ScheduledActionType
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::ScheduledAction
Properties:
  Enable: Boolean
  EndTime: String
  IamRole: String
  Schedule: String
  ScheduledActionDescription: String
  ScheduledActionName: String
  StartTime: String
  TargetAction:
    ScheduledActionType

```

## Properties

`Enable`

If true, the schedule is enabled. If false, the scheduled action does not trigger.
For more information about `state` of the scheduled action, see [AWS::Redshift::ScheduledAction](aws-resource-redshift-scheduledaction.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndTime`

The end time in UTC when the schedule is no longer active. After this time, the scheduled action does not trigger.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRole`

The IAM role to assume to run the scheduled action.
This IAM role must have permission to run the Amazon Redshift API operation in the scheduled action.
This IAM role must allow the Amazon Redshift scheduler (Principal scheduler.redshift.amazonaws.com) to assume permissions on your behalf.

For more information about the IAM role to use with the Amazon Redshift scheduler, see
[Using Identity-Based Policies for Amazon Redshift](../../../redshift/latest/mgmt/redshift-iam-access-control-identity-based.md)
in the _Amazon Redshift Cluster Management Guide_.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The schedule for a one-time (at format) or recurring (cron format) scheduled action.
Schedule invocations must be separated by at least one hour.

Format of at expressions is " `at(yyyy-mm-ddThh:mm:ss)`". For example, " `at(2016-03-04T17:27:00)`".

Format of cron expressions is " `cron(Minutes Hours Day-of-month Month Day-of-week Year)`".
For example, " `cron(0 10 ? * MON *)`". For more information, see
[Cron Expressions](../../../amazoncloudwatch/latest/events/scheduledevents.md#CronExpressions)
in the _Amazon CloudWatch Events User Guide_.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduledActionDescription`

The description of the scheduled action.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduledActionName`

The name of the scheduled action.

_Required_: Yes

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StartTime`

The start time in UTC when the schedule is active. Before this time, the scheduled action does not trigger.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetAction`

A JSON format string of the Amazon Redshift API operation with input parameters.

" `{\"ResizeCluster\":{\"NodeType\":\"ra3.4xlarge\",\"ClusterIdentifier\":\"my-test-cluster\",\"NumberOfNodes\":3}}`".

_Required_: No

_Type_: [ScheduledActionType](aws-properties-redshift-scheduledaction-scheduledactiontype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myScheduledActionName" }`

For the Amazon Redshift `ScheduledActionName`, Ref returns the name of the scheduled action.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`NextInvocations`

List of times when the scheduled action will run.

`State`

The state of the scheduled action. For example, `DISABLED`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

PauseClusterMessage

All content copied from https://docs.aws.amazon.com/.
