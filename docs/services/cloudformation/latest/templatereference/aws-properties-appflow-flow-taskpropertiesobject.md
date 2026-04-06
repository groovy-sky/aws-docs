This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow TaskPropertiesObject

A map used to store task-related information. The execution service looks for particular
information based on the `TaskType`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The task property key.

_Required_: Yes

_Type_: String

_Allowed values_: `VALUE | VALUES | DATA_TYPE | UPPER_BOUND | LOWER_BOUND | SOURCE_DATA_TYPE | DESTINATION_DATA_TYPE | VALIDATION_ACTION | MASK_VALUE | MASK_LENGTH | TRUNCATE_LENGTH | MATH_OPERATION_FIELDS_ORDER | CONCAT_FORMAT | SUBFIELD_CATEGORY_MAP | EXCLUDE_SOURCE_FIELDS_LIST | INCLUDE_NEW_FIELDS | ORDERED_PARTITION_KEYS_LIST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The task property value.

_Required_: Yes

_Type_: String

_Pattern_: `.+`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Task

TrendmicroSourceProperties
