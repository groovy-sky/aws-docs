This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain ConflictResolution

Determines how the auto-merging process should resolve conflicts between different
profiles. For example, if Profile A and Profile B have the same `FirstName`
and `LastName`, `ConflictResolution` specifies which
`EmailAddress` should be used.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConflictResolvingModel" : String,
  "SourceName" : String
}

```

### YAML

```yaml

  ConflictResolvingModel: String
  SourceName: String

```

## Properties

`ConflictResolvingModel`

How the auto-merging process should resolve conflicts between different
profiles.

_Required_: Yes

_Type_: String

_Allowed values_: `RECENCY | SOURCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceName`

The `ObjectType` name that is used to resolve profile merging conflicts
when choosing `SOURCE` as the `ConflictResolvingModel`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoMerging

Consolidation

All content copied from https://docs.aws.amazon.com/.
