This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer DateTimeConverter

This processor converts a datetime string into a format that you specify.

For more information about this processor including examples, see [datetimeConverter](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-datetimeConverter) in the _CloudWatch Logs User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Locale" : String,
  "MatchPatterns" : [ String, ... ],
  "Source" : String,
  "SourceTimezone" : String,
  "Target" : String,
  "TargetFormat" : String,
  "TargetTimezone" : String
}

```

### YAML

```yaml

  Locale: String
  MatchPatterns:
    - String
  Source: String
  SourceTimezone: String
  Target: String
  TargetFormat: String
  TargetTimezone: String

```

## Properties

`Locale`

The locale of the source field. If you omit this, the default of
`locale.ROOT` is used.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchPatterns`

A list of patterns to match against the `source` field.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The key to apply the date conversion to.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceTimezone`

The time zone of the source field. If you omit this, the default used is the UTC
zone.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

The JSON field to store the result in.

_Required_: Yes

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetFormat`

The datetime format to use for the converted data in the target field.

If you omit this, the default of ` yyyy-MM-dd'T'HH:mm:ss.SSS'Z` is
used.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTimezone`

The time zone of the target field. If you omit this, the default used is the UTC
zone.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Csv

DeleteKeys

All content copied from https://docs.aws.amazon.com/.
