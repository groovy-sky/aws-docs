This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule AdvancedFieldSelector

Defines criteria for selecting resources based on field values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndsWith" : [ String, ... ],
  "Equals" : [ String, ... ],
  "Field" : String,
  "NotEndsWith" : [ String, ... ],
  "NotEquals" : [ String, ... ],
  "NotStartsWith" : [ String, ... ],
  "StartsWith" : [ String, ... ]
}

```

### YAML

```yaml

  EndsWith:
    - String
  Equals:
    - String
  Field: String
  NotEndsWith:
    - String
  NotEquals:
    - String
  NotStartsWith:
    - String
  StartsWith:
    - String

```

## Properties

`EndsWith`

Matches if the field value ends with the specified value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Equals`

Matches if the field value equals the specified value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Field`

The name of the field to use for selection.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotEndsWith`

Matches if the field value does not end with the specified value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotEquals`

Matches if the field value does not equal the specified value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotStartsWith`

Matches if the field value does not start with the specified value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartsWith`

Matches if the field value starts with the specified value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdvancedEventSelector

CloudtrailParameters

All content copied from https://docs.aws.amazon.com/.
