This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens StorageLensTableDestination

This resource configures your S3 Storage Lens reports to export to read-only S3 table
buckets. With this resource, you can store your Storage Lens metrics in S3 Tables that are created in a read-only S3 table bucket called aws-s3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encryption" : Encryption,
  "IsEnabled" : Boolean
}

```

### YAML

```yaml

  Encryption:
    Encryption
  IsEnabled: Boolean

```

## Properties

`Encryption`

This resource configures your data encryption settings for Storage Lens metrics in
read-only S3 table buckets.

_Required_: No

_Type_: [Encryption](aws-properties-s3-storagelens-encryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsEnabled`

This property indicates whether the export to read-only S3 table buckets is enabled
for your S3 Storage Lens configuration. When set to true, Storage Lens reports are
automatically exported to tables in addition to other configured destinations.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensGroupSelectionCriteria

Tag

All content copied from https://docs.aws.amazon.com/.
