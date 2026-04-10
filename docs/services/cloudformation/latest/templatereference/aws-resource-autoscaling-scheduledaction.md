This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::ScheduledAction

The `AWS::AutoScaling::ScheduledAction` resource specifies an Amazon EC2 Auto
Scaling scheduled action so that the Auto Scaling group can change the number of instances
available for your application in response to predictable load changes.

When you update a stack with an Auto Scaling group and scheduled action, CloudFormation
always sets the min size, max size, and desired capacity properties of your group to the
values that are defined in the `AWS::AutoScaling::AutoScalingGroup` section of your
template. However, you might not want CloudFormation to do that when you have a scheduled
action in effect. You can use an [UpdatePolicy\
attribute](../userguide/aws-attribute-updatepolicy.md) to prevent CloudFormation from changing the min size, max size, or desired
capacity property values during a stack update unless you modified the individual values in
your template. If you have rolling updates enabled, before you can update the Auto Scaling
group, you must suspend scheduled actions by specifying an [UpdatePolicy\
attribute](../userguide/aws-attribute-updatepolicy.md) for the Auto Scaling group. You can find a sample update policy for
rolling updates in [Configure Amazon EC2\
Auto Scaling resources](../userguide/quickref-ec2-auto-scaling.md).

For more information, see [Scheduled scaling](../../../autoscaling/ec2/userguide/schedule-time.md) and [Suspending and resuming scaling processes](../../../autoscaling/ec2/userguide/as-suspend-resume-processes.md) in the _Amazon EC2 Auto Scaling_
_User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AutoScaling::ScheduledAction",
  "Properties" : {
      "AutoScalingGroupName" : String,
      "DesiredCapacity" : Integer,
      "EndTime" : String,
      "MaxSize" : Integer,
      "MinSize" : Integer,
      "Recurrence" : String,
      "StartTime" : String,
      "TimeZone" : String
    }
}

```

### YAML

```yaml

Type: AWS::AutoScaling::ScheduledAction
Properties:
  AutoScalingGroupName: String
  DesiredCapacity: Integer
  EndTime: String
  MaxSize: Integer
  MinSize: Integer
  Recurrence: String
  StartTime: String
  TimeZone: String

```

## Properties

`AutoScalingGroupName`

The name of the Auto Scaling group.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DesiredCapacity`

The desired capacity is the initial capacity of the Auto Scaling group after the scheduled
action runs and the capacity it attempts to maintain. It can scale beyond this capacity
if you add more scaling conditions.

###### Note

You must specify at least one of the following properties: `MaxSize`,
`MinSize`, or `DesiredCapacity`.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndTime`

The date and time for the recurring schedule to end, in UTC. For example,
`"2021-06-01T00:00:00Z"`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxSize`

The maximum size of the Auto Scaling group.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinSize`

The minimum size of the Auto Scaling group.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Recurrence`

The recurring schedule for this action. This format consists of five fields separated
by white spaces: \[Minute\] \[Hour\] \[Day\_of\_Month\] \[Month\_of\_Year\] \[Day\_of\_Week\]. The value
must be in quotes (for example, `"30 0 1 1,6,12 *"`). For more information
about this format, see [Crontab](http://crontab.org/).

When `StartTime` and `EndTime` are specified with
`Recurrence`, they form the boundaries of when the recurring action
starts and stops.

Cron expressions use Universal Coordinated Time (UTC) by default.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The date and time for this action to start, in YYYY-MM-DDThh:mm:ssZ format in UTC/GMT
only and in quotes (for example, `"2021-06-01T00:00:00Z"`).

If you specify `Recurrence` and `StartTime`, Amazon EC2 Auto Scaling performs
the action at this time, and then performs the action based on the specified
recurrence.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

Specifies the time zone for a cron expression. If a time zone is not provided, UTC is
used by default.

Valid values are the canonical names of the IANA time zones, derived from the IANA
Time Zone Database (such as `Etc/GMT+9` or `Pacific/Tahiti`). For
more information, see [https://en.wikipedia.org/wiki/List\_of\_tz\_database\_time\_zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:
`mystack-myscheduledaction-NT5EUXTNTXXD`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ScheduledActionName`

Returns the name of a scheduled action.

## Examples

The following examples schedule scaling actions for an Auto Scaling group.

- [Scheduled actions that run on a recurring schedule](#aws-resource-autoscaling-scheduledaction--examples--Scheduled_actions_that_run_on_a_recurring_schedule)

- [Scheduled scaling action that occurs only once](#aws-resource-autoscaling-scheduledaction--examples--Scheduled_scaling_action_that_occurs_only_once)

### Scheduled actions that run on a recurring schedule

The following template snippet includes two scheduled actions that scale the number of
instances in an Auto Scaling group. The `ScheduledActionOut` action starts at 7
AM every day and sets the Auto Scaling group to a minimum of five Amazon EC2 instances
with a maximum of 10. The `ScheduledActionIn` action starts at 7 PM every day
and sets the Auto Scaling group to a minimum and maximum of one Amazon EC2 instance. The
time zone is not provided. As a result, these scheduled actions will recur in UTC time.

#### JSON

```json

{
  "Resources":{
    "ScheduledActionOut":{
      "Type":"AWS::AutoScaling::ScheduledAction",
      "Properties":{
        "AutoScalingGroupName":{
          "Ref":"myASG"
        },
        "MaxSize":"10",
        "MinSize":"5",
        "Recurrence":"0 7 * * *"
      }
    },
    "ScheduledActionIn":{
      "Type":"AWS::AutoScaling::ScheduledAction",
      "Properties":{
        "AutoScalingGroupName":{
          "Ref":"myASG"
        },
        "MaxSize":"1",
        "MinSize":"1",
        "Recurrence":"0 19 * * *"
      }
    }
  }
}
```

#### YAML

```yaml

---
Resources:
  ScheduledActionOut:
    Type: AWS::AutoScaling::ScheduledAction
    Properties:
      AutoScalingGroupName: !Ref myASG
      MaxSize: '10'
      MinSize: '5'
      Recurrence: "0 7 * * *"
  ScheduledActionIn:
    Type: AWS::AutoScaling::ScheduledAction
    Properties:
      AutoScalingGroupName: !Ref myASG
      MaxSize: '1'
      MinSize: '1'
      Recurrence: "0 19 * * *"
```

### Scheduled scaling action that occurs only once

The following template snippet includes a one-time scheduled action. At the date and
time specified for `StartTime` (4:00 PM UTC on March 31, 2021), if the group
currently has more than 1 instance, it scales in to 1 instance. If the group currently has
no instances, it scales out to 1 instance.

#### JSON

```json

{
  "Resources":{
    "OneTimeScheduledAction":{
      "Type":"AWS::AutoScaling::ScheduledAction",
      "Properties":{
        "AutoScalingGroupName":{
          "Ref":"myASG"
        },
        "DesiredCapacity":"1",
        "StartTime":"2021-03-31T16:00:00Z"
      }
    }
  }
}
```

#### YAML

```yaml

---
Resources:
  OneTimeScheduledAction:
    Type: AWS::AutoScaling::ScheduledAction
    Properties:
      AutoScalingGroupName: !Ref myASG
      DesiredCapacity: '1'
      StartTime: '2021-03-31T16:00:00Z'
```

## See also

- You can find additional useful snippets in the following sections of the _AWS CloudFormation User Guide_:

- For examples of Auto Scaling groups, see [Configure\
Amazon EC2 Auto Scaling resources](../userguide/quickref-ec2-auto-scaling.md).

- For examples of launch templates, see [Create\
launch templates](../userguide/quickref-ec2-launch-templates.md).

- [PutScheduledUpdateGroupAction](../../../../reference/autoscaling/ec2/apireference/api-putscheduledupdategroupaction.md) in the _Amazon EC2 Auto Scaling API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetTrackingMetricStat

AWS::AutoScaling::WarmPool

All content copied from https://docs.aws.amazon.com/.
