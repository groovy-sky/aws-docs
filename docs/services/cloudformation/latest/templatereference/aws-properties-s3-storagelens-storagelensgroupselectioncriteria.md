This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens StorageLensGroupSelectionCriteria

This resource indicates which Storage Lens group ARNs to include or exclude in the Storage
Lens group aggregation. You can only attach Storage Lens groups to your dashboard if they're
included in your Storage Lens group aggregation. If this value is left null, then all Storage
Lens groups are selected.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Exclude" : [ String, ... ],
  "Include" : [ String, ... ]
}

```

### YAML

```yaml

  Exclude:
    - String
  Include:
    - String

```

## Properties

`Exclude`

This property indicates which Storage Lens group ARNs to exclude from the Storage Lens
group aggregation.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Include`

This property indicates which Storage Lens group ARNs to include in the Storage Lens group
aggregation.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageLensGroupLevel

StorageLensTableDestination

All content copied from https://docs.aws.amazon.com/.
