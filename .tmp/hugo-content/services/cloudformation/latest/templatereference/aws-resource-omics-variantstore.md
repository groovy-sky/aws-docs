This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::VariantStore

Create a store for variant data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Omics::VariantStore",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Reference" : ReferenceItem,
      "SseConfig" : SseConfig,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Omics::VariantStore
Properties:
  Description: String
  Name: String
  Reference:
    ReferenceItem
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

_Minimum_: `0`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the store.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-z]){1}([a-z0-9_]){2,254}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Reference`

The genome reference for the store's variants.

_Required_: Yes

_Type_: [ReferenceItem](aws-properties-omics-variantstore-referenceitem.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseConfig`

Server-side encryption (SSE) settings for the store.

_Required_: No

_Type_: [SseConfig](aws-properties-omics-variantstore-sseconfig.md)

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

`{ "Ref": "VariantStore.Status" }`

For the Amazon Omics resource `VariantStore.Status`, `Ref` returns
the status of the variant store.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

When the store was created.

`Id`

The store's ID.

`Status`

The store's status.

`StatusMessage`

The store's status message.

`StoreArn`

The store's ARN.

`StoreSizeBytes`

The store's size in bytes.

`UpdateTime`

When the store was updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SseConfig

ReferenceItem

All content copied from https://docs.aws.amazon.com/.
