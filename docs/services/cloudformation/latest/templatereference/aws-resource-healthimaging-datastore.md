This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::HealthImaging::Datastore

Create a data store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::HealthImaging::Datastore",
  "Properties" : {
      "DatastoreName" : String,
      "KmsKeyArn" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::HealthImaging::Datastore
Properties:
  DatastoreName: String
  KmsKeyArn: String
  Tags:
    Key: Value

```

## Properties

`DatastoreName`

The data store name.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9._/#-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyArn`

The Amazon Resource Name (ARN) assigned to the Key Management Service (KMS) key for accessing encrypted data.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags provided when creating a data store.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp when the data store was created.

`DatastoreArn`

The Amazon Resource Name (ARN) for the data store.

`DatastoreId`

The data store identifier.

`DatastoreStatus`

The data store status.

`UpdatedAt`

The timestamp when the data store was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS HealthImaging

Next

All content copied from https://docs.aws.amazon.com/.
