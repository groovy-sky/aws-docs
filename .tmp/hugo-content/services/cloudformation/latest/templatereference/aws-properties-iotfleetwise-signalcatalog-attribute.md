This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::SignalCatalog Attribute

A signal that represents static information about the vehicle, such as engine type or
manufacturing date.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedValues" : [ String, ... ],
  "AssignedValue" : String,
  "DataType" : String,
  "DefaultValue" : String,
  "Description" : String,
  "FullyQualifiedName" : String,
  "Max" : Number,
  "Min" : Number,
  "Unit" : String
}

```

### YAML

```yaml

  AllowedValues:
    - String
  AssignedValue: String
  DataType: String
  DefaultValue: String
  Description: String
  FullyQualifiedName: String
  Max: Number
  Min: Number
  Unit: String

```

## Properties

`AllowedValues`

A list of possible values an attribute can be assigned.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssignedValue`

A specified value for the attribute.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

The specified data type of the attribute.

_Required_: Yes

_Type_: String

_Allowed values_: `INT8 | UINT8 | INT16 | UINT16 | INT32 | UINT32 | INT64 | UINT64 | BOOLEAN | FLOAT | DOUBLE | STRING | UNIX_TIMESTAMP | INT8_ARRAY | UINT8_ARRAY | INT16_ARRAY | UINT16_ARRAY | INT32_ARRAY | UINT32_ARRAY | INT64_ARRAY | UINT64_ARRAY | BOOLEAN_ARRAY | FLOAT_ARRAY | DOUBLE_ARRAY | STRING_ARRAY | UNIX_TIMESTAMP_ARRAY | UNKNOWN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The default value of the attribute.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A brief description of the attribute.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000-\u001F\u007F]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FullyQualifiedName`

The fully qualified name of the attribute. For example, the fully qualified name of an
attribute might be `Vehicle.Body.Engine.Type`.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_.]+`

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Max`

The specified possible maximum value of the attribute.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Min`

The specified possible minimum value of the attribute.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The scientific unit for the attribute.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actuator

Branch

All content copied from https://docs.aws.amazon.com/.
