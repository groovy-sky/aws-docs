This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens StorageLensExpandedPrefixesDataExport

This resource specifies the properties of your S3 Storage Lens Expanded Prefixes metrics export.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3BucketDestination" : S3BucketDestination,
  "StorageLensTableDestination" : StorageLensTableDestination
}

```

### YAML

```yaml

  S3BucketDestination:
    S3BucketDestination
  StorageLensTableDestination:
    StorageLensTableDestination

```

## Properties

`S3BucketDestination`

This property specifies the general purpose bucket where the S3 Storage Lens Expanded Prefixes metrics export files are located. At least one export destination must be specified.

_Required_: No

_Type_: [S3BucketDestination](aws-properties-s3-storagelens-s3bucketdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageLensTableDestination`

This property configures S3 Storage Lens Expanded Prefixes metrics report to read-only S3 table buckets.

_Required_: No

_Type_: [StorageLensTableDestination](aws-properties-s3-storagelens-storagelenstabledestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensConfiguration

StorageLensGroupLevel

All content copied from https://docs.aws.amazon.com/.
