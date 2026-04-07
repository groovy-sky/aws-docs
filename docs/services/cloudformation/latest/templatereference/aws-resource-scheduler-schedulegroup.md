This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::ScheduleGroup

A _schedule group_ is an Amazon EventBridge Scheduler resource you use to organize your schedules.

Your AWS account comes with a `default` scheduler group. You associate a new schedule with the `default` group or with schedule groups that
you create and manage. You can create up to [500 schedule groups](https://docs.aws.amazon.com/scheduler/latest/UserGuide/scheduler-quotas.html) in your AWS account.
With EventBridge Scheduler, you apply [tags](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) to schedule groups, not to individual schedules to organize your resources.

For more information about managing schedule groups, see [Managing a schedule group](https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-schedule-group.html)
in the _EventBridge Scheduler User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Scheduler::ScheduleGroup",
  "Properties" : {
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Scheduler::ScheduleGroup
Properties:
  Name: String
  Tags:
    - Tag

```

## Properties

`Name`

The name of the schedule group.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-zA-Z-_.]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-scheduler-schedulegroup-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `Name` attribute of the schedule group.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes that `Fn::GetAtt` returns.
For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

`Arn`

The Amazon Resource Name (ARN) of the schedule group.

`CreationDate`

The date and time at which the schedule group was created.

`LastModificationDate`

The time at which the schedule group was last modified.

`State`

Specifies the state of the schedule group.

_Allowed Values_: `ACTIVE` \| `DELETING`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Target

Tag
