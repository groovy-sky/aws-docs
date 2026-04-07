This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Scheduler::Schedule FlexibleTimeWindow

Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumWindowInMinutes" : Number,
  "Mode" : String
}

```

### YAML

```yaml

  MaximumWindowInMinutes: Number
  Mode: String

```

## Properties

`MaximumWindowInMinutes`

The maximum time window during which a schedule can be invoked.

_Minimum_: `1`

_Maximum_: `1440`

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `1440`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

Determines whether the schedule is invoked within a flexible time window. You must use quotation marks when you specify this value in your JSON or YAML template.

_Allowed Values_: `"OFF"` \| `"FLEXIBLE"`

_Required_: Yes

_Type_: String

_Allowed values_: `OFF | FLEXIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventBridgeParameters

KinesisParameters
