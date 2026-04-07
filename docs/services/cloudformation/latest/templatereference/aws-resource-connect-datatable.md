This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::DataTable

Represents a data table in Amazon Connect. A data table is a JSON-like data structure where attributes and
values are dynamically set by customers. Customers can reference table values within call flows, applications, views,
and workspaces to pinpoint dynamic configuration that changes their contact center's behavior in a predetermined and
safe way.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::DataTable",
  "Properties" : {
      "Description" : String,
      "InstanceArn" : String,
      "Name" : String,
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "TimeZone" : String,
      "ValueLockLevel" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::DataTable
Properties:
  Description: String
  InstanceArn: String
  Name: String
  Status: String
  Tags:
    - Tag
  TimeZone: String
  ValueLockLevel: String

```

## Properties

`Description`

An optional description of the data table's purpose and contents.

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

The human-readable name of the data table. Must be unique within the instance and conform to Connect naming
standards.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}\p{Z}\p{N}\-_.:=@'|]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The current status of the data table. One of PUBLISHED or SAVED.

_Required_: No

_Type_: String

_Allowed values_: `PUBLISHED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key-value pairs for attribute based access control (TBAC or ABAC) and organization.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-datatable-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The IANA timezone identifier used when resolving time based dynamic values. Required even if no time slices are
specified.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueLockLevel`

The data level that concurrent value edits are locked on. One of DATA\_TABLE, PRIMARY\_VALUE, ATTRIBUTE, VALUE,
and NONE. Determines how concurrent edits are handled when multiple users attempt to modify values
simultaneously.

_Required_: No

_Type_: String

_Allowed values_: `NONE | DATA_TABLE | PRIMARY_VALUE | ATTRIBUTE | VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the data table. Does not include version aliases.

`CreatedTime`

The timestamp when the data table was created.

`LastModifiedRegion`

The AWS Region where the data table was last modified, used for region replication.

`LastModifiedTime`

The timestamp when the data table or any of its properties were last modified.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Connect::ContactFlowVersion

LockVersion
