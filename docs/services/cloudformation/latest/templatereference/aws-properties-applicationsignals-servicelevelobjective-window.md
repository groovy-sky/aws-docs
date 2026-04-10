This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective Window

The start and end time of the time exclusion window.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Duration" : Integer,
  "DurationUnit" : String
}

```

### YAML

```yaml

  Duration: Integer
  DurationUnit: String

```

## Properties

`Duration`

The start and end time of the time exclusion window.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DurationUnit`

The unit of measurement to use during the time window exclusion.

_Required_: Yes

_Type_: String

_Allowed values_: `MINUTE | HOUR | DAY | MONTH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Next

All content copied from https://docs.aws.amazon.com/.
