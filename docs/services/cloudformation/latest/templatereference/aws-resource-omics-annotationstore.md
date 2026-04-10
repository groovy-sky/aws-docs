This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::AnnotationStore

###### Important

AWS HealthOmics variant stores and annotation stores are no longer open to new customers.
Existing customers can continue to use the service as normal. For more information, see
[AWS HealthOmics variant store and annotation store availability change](../../../omics/latest/dev/variant-store-availability-change.md).

Creates an annotation store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Omics::AnnotationStore",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Reference" : ReferenceItem,
      "SseConfig" : SseConfig,
      "StoreFormat" : String,
      "StoreOptions" : StoreOptions,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Omics::AnnotationStore
Properties:
  Description: String
  Name: String
  Reference:
    ReferenceItem
  SseConfig:
    SseConfig
  StoreFormat: String
  StoreOptions:
    StoreOptions
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

The name of the Annotation Store.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-z]){1}([a-z0-9_]){2,254}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Reference`

The genome reference for the store's annotations.

_Required_: No

_Type_: [ReferenceItem](aws-properties-omics-annotationstore-referenceitem.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseConfig`

The store's server-side encryption (SSE) settings.

_Required_: No

_Type_: [SseConfig](aws-properties-omics-annotationstore-sseconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StoreFormat`

The annotation file format of the store.

_Required_: Yes

_Type_: String

_Allowed values_: `GFF | TSV | VCF`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StoreOptions`

File parsing options for the annotation store.

_Required_: No

_Type_: [StoreOptions](aws-properties-omics-annotationstore-storeoptions.md)

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

`{ "Ref": "AnnotationStore.Id" }` `Ref` returns the id for the annotation store.

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

AWS HealthOmics

ReferenceItem

All content copied from https://docs.aws.amazon.com/.
