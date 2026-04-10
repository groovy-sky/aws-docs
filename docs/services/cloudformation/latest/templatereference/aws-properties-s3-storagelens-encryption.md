This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens Encryption

This resource contains the type of server-side encryption used to encrypt an Amazon S3
Storage Lens metrics export. For valid values, see the [StorageLensDataExportEncryption](../../../s3/latest/api/api-control-storagelensdataexportencryption.md) in the _Amazon S3 API_
_Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SSEKMS" : SSEKMS,
  "SSES3" : Json
}

```

### YAML

```yaml

  SSEKMS:
    SSEKMS
  SSES3: Json

```

## Properties

`SSEKMS`

Specifies the use of AWS Key Management Service keys (SSE-KMS) to encrypt the S3 Storage Lens
metrics export file.

_Required_: No

_Type_: [SSEKMS](aws-properties-s3-storagelens-ssekms.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSES3`

Specifies the use of an Amazon S3-managed key (SSE-S3) to encrypt the S3 Storage Lens
metrics export file.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DetailedStatusCodesMetrics

PrefixLevel

All content copied from https://docs.aws.amazon.com/.
