This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::ReferenceStore

Creates a reference store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Omics::ReferenceStore",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "SseConfig" : SseConfig,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Omics::ReferenceStore
Properties:
  Description: String
  Name: String
  SseConfig:
    SseConfig
  Tags:
    Key: Value

```

## Properties

`Description`

A description for the store.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A name for the store.

_Required_: Yes

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseConfig`

Server-side encryption (SSE) settings for the store.

_Required_: No

_Type_: [SseConfig](aws-properties-omics-referencestore-sseconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags for the store.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the details of this resource. For example:

`{ "Ref": "ReferenceStore.Arn" }` `Ref` returns the arn for the reference store.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn``CreationTime`

When the store was created.

`ReferenceStoreId`

The store's ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

SseConfig

All content copied from https://docs.aws.amazon.com/.
