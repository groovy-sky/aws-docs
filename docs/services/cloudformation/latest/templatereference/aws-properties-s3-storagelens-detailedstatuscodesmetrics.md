This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens DetailedStatusCodesMetrics

This resource enables Amazon S3 Storage Lens detailed status code metrics.
Detailed status code metrics generate metrics for HTTP status codes, such as `200
        OK`, `403 Forbidden`, `503 Service Unavailable` and others.

For more information, see [Assessing your storage activity and usage\
with S3 Storage Lens](../../../s3/latest/userguide/storage-lens.md) in the _Amazon S3 User Guide_. For a
complete list of metrics, see [S3 Storage Lens metrics\
glossary](../../../s3/latest/userguide/storage-lens-metrics-glossary.md) in the _Amazon S3 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsEnabled" : Boolean
}

```

### YAML

```yaml

  IsEnabled: Boolean

```

## Properties

`IsEnabled`

Indicates whether detailed status code metrics are enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataExport

Encryption

All content copied from https://docs.aws.amazon.com/.
