This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AlarmMuteRule

Creates or updates an alarm mute rule.

Alarm mute rules automatically mute alarm actions during predefined time windows. When a mute rule is active, targeted alarms continue to evaluate metrics and transition between states, but their configured actions (such as Amazon SNS notifications or Auto Scaling actions) are muted.

You can create mute rules with recurring schedules using `cron` expressions or one-time mute windows using `at` expressions. Each mute rule can target up to 100 specific alarms by name.

If you specify a rule name that already exists, this operation updates the existing rule with the new configuration.

**Permissions**

To create or update a mute rule, you must have the `cloudwatch:PutAlarmMuteRule` permission on two types of resources: the alarm mute rule resource itself, and each alarm that the rule targets.

For example, If you want to allow a user to create mute rules that target only specific alarms named "WebServerCPUAlarm" and "DatabaseConnectionAlarm", you would create an IAM policy with one statement granting `cloudwatch:PutAlarmMuteRule` on the alarm mute rule resource ( `arn:aws:cloudwatch:[REGION]:123456789012:alarm-mute-rule:*`), and another statement granting `cloudwatch:PutAlarmMuteRule` on the targeted alarm resources ( `arn:aws:cloudwatch:[REGION]:123456789012:alarm:WebServerCPUAlarm` and `arn:aws:cloudwatch:[REGION]:123456789012:alarm:DatabaseConnectionAlarm`).

You can also use IAM policy conditions to allow targeting alarms based on resource tags. For example, you can restrict users to create/update mute rules to only target alarms that have a specific tag key-value pair, such as `Team=TeamA`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudWatch::AlarmMuteRule",
  "Properties" : {
      "Description" : String,
      "ExpireDate" : String,
      "MuteTargets" : MuteTargets,
      "Name" : String,
      "Rule" : Rule,
      "StartDate" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudWatch::AlarmMuteRule
Properties:
  Description: String
  ExpireDate: String
  MuteTargets:
    MuteTargets
  Name: String
  Rule:
    Rule
  StartDate: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the alarm mute rule that helps you identify its purpose.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpireDate`

The date and time when the mute rule expires and is no longer evaluated, specified in ISO 8601 format (for example, `2026-12-31T23:59`). After this time, the rule status becomes EXPIRED. This date and time is interpreted according to the schedule timezone, or UTC if no timezone is specified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MuteTargets`

Specifies which alarms this rule applies to.

_Required_: No

_Type_: [MuteTargets](aws-properties-cloudwatch-alarmmuterule-mutetargets.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the alarm mute rule. This name must be unique within your AWS account and region.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rule`

The configuration that defines when and how long alarms should be muted.

_Required_: Yes

_Type_: [Rule](aws-properties-cloudwatch-alarmmuterule-rule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartDate`

The date and time when the mute rule becomes active, specified in ISO 8601 format (for example, `2026-04-15T08:00`). If omitted, the rule becomes active immediately upon creation. This date and time is interpreted according to the schedule timezone, or UTC if no timezone is specified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs to associate with the alarm mute rule. You can use tags to categorize and manage your mute rules.

_Required_: No

_Type_: Array of [Tag](aws-properties-cloudwatch-alarmmuterule-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The name of the alarm mute rule.

### Fn::GetAtt

Retrieves details for a specific alarm mute rule.

This operation returns complete information about the mute rule, including its configuration, status, targeted alarms, and metadata.

The returned status indicates the current state of the mute rule:

- **SCHEDULED**: The mute rule is configured and will become active in the future

- **ACTIVE**: The mute rule is currently muting alarm actions

- **EXPIRED**: The mute rule has passed its expiration date and will no longer become active

**Permissions**

To retrieve details for a mute rule, you need the `cloudwatch:GetAlarmMuteRule` permission on the alarm mute rule resource.

`Arn`

The Amazon Resource Name (ARN) of the alarm mute rule.

`LastUpdatedTimestamp`

The date and time when the mute rule was last updated.

`MuteType`

Indicates whether the mute rule is one-time or recurring. Valid values are `ONE_TIME` or `RECURRING`.

`Status`

The current status of the alarm mute rule. Valid values are `SCHEDULED`, `ACTIVE`, or `EXPIRED`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

MuteTargets

All content copied from https://docs.aws.amazon.com/.
