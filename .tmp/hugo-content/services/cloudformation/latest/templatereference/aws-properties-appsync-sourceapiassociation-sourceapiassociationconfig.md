This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::SourceApiAssociation SourceApiAssociationConfig

Describes properties used to specify configurations related to a source API. This is a
property of the `AWS:AppSync:SourceApiAssociation` type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MergeType" : String
}

```

### YAML

```yaml

  MergeType: String

```

## Properties

`MergeType`

The property that indicates which merging option is enabled in the source API
association.

Valid merge types are `MANUAL_MERGE` (default) and `AUTO_MERGE`.
Manual merges are the default behavior and require the user to trigger any changes from
the source APIs to the merged API manually. Auto merges subscribe the merged API to the
changes performed on the source APIs so that any change in the source APIs are also made
to the merged API. Auto merges use `MergedApiExecutionRoleArn` to perform
merge operations.

The following values are valid:

`MANUAL_MERGE | AUTO_MERGE`

_Required_: No

_Type_: String

_Allowed values_: `AUTO_MERGE | MANUAL_MERGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppSync::SourceApiAssociation

Next

All content copied from https://docs.aws.amazon.com/.
