This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens StorageLensGroupLevel

This resource determines the scope of Storage Lens group data that is displayed in the
Storage Lens dashboard.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StorageLensGroupSelectionCriteria" : StorageLensGroupSelectionCriteria
}

```

### YAML

```yaml

  StorageLensGroupSelectionCriteria:
    StorageLensGroupSelectionCriteria

```

## Properties

`StorageLensGroupSelectionCriteria`

This property indicates which Storage Lens group ARNs to include or exclude in the Storage
Lens group aggregation. If this value is left null, then all Storage Lens groups are
selected.

_Required_: No

_Type_: [StorageLensGroupSelectionCriteria](aws-properties-s3-storagelens-storagelensgroupselectioncriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensExpandedPrefixesDataExport

StorageLensGroupSelectionCriteria

All content copied from https://docs.aws.amazon.com/.
