This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::DataTableAttribute

Represents an attribute (column) in a data table. Attributes define the schema and validation rules for values
that can be stored in the table. They specify the data type, constraints, and whether the attribute is used as a
primary key for record identification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::DataTableAttribute",
  "Properties" : {
      "DataTableArn" : String,
      "Description" : String,
      "InstanceArn" : String,
      "Name" : String,
      "Primary" : Boolean,
      "Validation" : Validation,
      "ValueType" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::DataTableAttribute
Properties:
  DataTableArn: String
  Description: String
  InstanceArn: String
  Name: String
  Primary: Boolean
  Validation:
    Validation
  ValueType: String

```

## Properties

`DataTableArn`

The Amazon Resource Name (ARN) of the data table that contains this attribute.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

An optional description explaining the purpose and usage of this attribute.

_Required_: No

_Type_: String

_Pattern_: `^[\P{C}
	]+$`

_Minimum_: `0`

_Maximum_: `250`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The human-readable name of the attribute. Must be unique within the data table and conform to Connect naming
standards.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}\p{Z}\p{N}\-_.:=@'|]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Primary`

Boolean indicating whether this attribute is used as a primary key for record identification. Primary attributes
must have unique value combinations and cannot contain expressions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Validation`

The validation rules applied to values of this attribute. Based on JSON Schema Draft 2020-12 with additional
Connect-specific validations for data integrity.

_Required_: No

_Type_: [Validation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-datatableattribute-validation.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueType`

The type of value allowed for this attribute. Must be one of TEXT, TEXT\_LIST, NUMBER, NUMBER\_LIST, or BOOLEAN.
Determines how values are validated and processed.

_Required_: No

_Type_: String

_Allowed values_: `TEXT | NUMBER | BOOLEAN | TEXT_LIST | NUMBER_LIST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AttributeId`

The unique identifier for the attribute within the data table.

`LastModifiedRegion`

The AWS Region where this attribute was last modified, used for region replication.

`LastModifiedTime`

The timestamp when this attribute was last modified.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Enum
