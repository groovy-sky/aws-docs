This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::CalculatedAttributeDefinition ValueRange

A structure letting customers specify a relative time window over which over which
data is included in the Calculated Attribute. Use positive numbers to indicate that the
endpoint is in the past, and negative numbers to indicate it is in the future.
ValueRange overrides Value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "End" : Integer,
  "Start" : Integer
}

```

### YAML

```yaml

  End: Integer
  Start: Integer

```

## Properties

`End`

The ending point for this overridden range. Positive numbers indicate how many days in
the past data should be included, and negative numbers indicate how many days in the
future.

_Required_: Yes

_Type_: Integer

_Minimum_: `-2147483648`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Start`

The starting point for this overridden range. Positive numbers indicate how many days
in the past data should be included, and negative numbers indicate how many days in the
future.

_Required_: Yes

_Type_: Integer

_Minimum_: `-2147483648`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Threshold

AWS::CustomerProfiles::Domain

All content copied from https://docs.aws.amazon.com/.
