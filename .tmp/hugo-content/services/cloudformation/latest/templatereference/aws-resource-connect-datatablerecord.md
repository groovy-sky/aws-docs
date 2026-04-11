This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::DataTableRecord

The `AWS::Connect::DataTableRecord` resource Property description not available. for Connect.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::DataTableRecord",
  "Properties" : {
      "DataTableArn" : String,
      "DataTableRecord" : DataTableRecord,
      "InstanceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::DataTableRecord
Properties:
  DataTableArn: String
  DataTableRecord:
    DataTableRecord
  InstanceArn: String

```

## Properties

`DataTableArn`

The Amazon Resource Name (ARN) for the data table. Does not include version aliases.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataTableRecord`

Property description not available.

_Required_: No

_Type_: [DataTableRecord](aws-properties-connect-datatablerecord-datatablerecord.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RecordId`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Validation

DataTableRecord

All content copied from https://docs.aws.amazon.com/.
