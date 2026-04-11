This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens PrefixLevelStorageMetrics

This resource contains the details of the prefix-level storage metrics for Amazon S3
Storage Lens.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsEnabled" : Boolean,
  "SelectionCriteria" : SelectionCriteria
}

```

### YAML

```yaml

  IsEnabled: Boolean
  SelectionCriteria:
    SelectionCriteria

```

## Properties

`IsEnabled`

This property identifies whether the details of the prefix-level storage metrics for S3
Storage Lens are enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectionCriteria`

This property identifies whether the details of the prefix-level storage metrics for S3
Storage Lens are enabled.

_Required_: No

_Type_: [SelectionCriteria](aws-properties-s3-storagelens-selectioncriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrefixLevel

S3BucketDestination

All content copied from https://docs.aws.amazon.com/.
