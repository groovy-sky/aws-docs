This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Connection

The `AWS::Glue::Connection` resource specifies an AWS Glue connection to a
data source. For more information, see [Adding a Connection to Your Data Store](../../../glue/latest/dg/populate-add-connection.md)
and [Connection Structure](../../../glue/latest/dg/aws-glue-api-catalog-connections.md#aws-glue-api-catalog-connections-Connection) in the _AWS Glue Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Connection",
  "Properties" : {
      "CatalogId" : String,
      "ConnectionInput" : ConnectionInput
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Connection
Properties:
  CatalogId: String
  ConnectionInput:
    ConnectionInput

```

## Properties

`CatalogId`

The ID of the data catalog to create the catalog object in. Currently, this should be
the AWS account ID.

###### Note

To specify the account ID, you can use the `Ref` intrinsic function
with the `AWS::AccountId` pseudo parameter. For example: `!Ref
                    AWS::AccountId`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectionInput`

The connection that you want to create.

_Required_: Yes

_Type_: [ConnectionInput](aws-properties-glue-connection-connectioninput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the connection name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

XMLClassifier

AuthenticationConfigurationInput

All content copied from https://docs.aws.amazon.com/.
