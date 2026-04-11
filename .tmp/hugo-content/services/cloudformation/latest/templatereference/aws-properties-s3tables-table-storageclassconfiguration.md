This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Tables::Table StorageClassConfiguration

The configuration details for the storage class of tables or table buckets. This allows you to optimize storage costs by selecting the appropriate storage class based on your access patterns and performance requirements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StorageClass" : String
}

```

### YAML

```yaml

  StorageClass: String

```

## Properties

`StorageClass`

The storage class for the table or table bucket. Valid values include storage classes optimized for different access patterns and cost profiles.

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | INTELLIGENT_TIERING`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnapshotManagement

Tag

All content copied from https://docs.aws.amazon.com/.
