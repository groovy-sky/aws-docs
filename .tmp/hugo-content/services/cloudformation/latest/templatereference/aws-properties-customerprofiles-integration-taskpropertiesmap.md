This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration TaskPropertiesMap

A map used to store task-related information. The execution service looks for
particular information based on the `TaskType`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OperatorPropertyKey" : String,
  "Property" : String
}

```

### YAML

```yaml

  OperatorPropertyKey: String
  Property: String

```

## Properties

`OperatorPropertyKey`

The task property key.

_Required_: Yes

_Type_: String

_Allowed values_: `VALUE | VALUES | DATA_TYPE | UPPER_BOUND | LOWER_BOUND | SOURCE_DATA_TYPE | DESTINATION_DATA_TYPE | VALIDATION_ACTION | MASK_VALUE | MASK_LENGTH | TRUNCATE_LENGTH | MATH_OPERATION_FIELDS_ORDER | CONCAT_FORMAT | SUBFIELD_CATEGORY_MAP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Property`

The task property value.

_Required_: Yes

_Type_: String

_Pattern_: `.+`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Task

TriggerConfig

All content copied from https://docs.aws.amazon.com/.
