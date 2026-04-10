This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::TableBucket UnreferencedFileRemoval

The unreferenced file removal settings for your table bucket. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots. For more information, see the [_Amazon S3 User Guide_](../../../s3/latest/userguide/s3-table-buckets-maintenance.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NoncurrentDays" : Integer,
  "Status" : String,
  "UnreferencedDays" : Integer
}

```

### YAML

```yaml

  NoncurrentDays: Integer
  Status: String
  UnreferencedDays: Integer

```

## Properties

`NoncurrentDays`

The number of days an object can be noncurrent before Amazon S3 deletes it.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the unreferenced file removal configuration for your table bucket.

_Required_: No

_Type_: String

_Allowed values_: `Enabled | Disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnreferencedDays`

The number of days an object must be unreferenced by your table before Amazon S3 marks the object as noncurrent.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::S3Tables::TableBucketPolicy

All content copied from https://docs.aws.amazon.com/.
