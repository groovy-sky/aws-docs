This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis AssetOptions

An array of analysis level configurations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Timezone" : String,
  "WeekStart" : String
}

```

### YAML

```yaml

  Timezone: String
  WeekStart: String

```

## Properties

`Timezone`

Determines the timezone for the analysis.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeekStart`

Determines the week start day for an analysis.

_Required_: No

_Type_: String

_Allowed values_: `SUNDAY | MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArcOptions

AttributeAggregationFunction

All content copied from https://docs.aws.amazon.com/.
