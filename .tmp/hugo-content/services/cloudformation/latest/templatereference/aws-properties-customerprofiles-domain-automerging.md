This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain AutoMerging

Configuration information about the auto-merging process.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConflictResolution" : ConflictResolution,
  "Consolidation" : Consolidation,
  "Enabled" : Boolean,
  "MinAllowedConfidenceScoreForMerging" : Number
}

```

### YAML

```yaml

  ConflictResolution:
    ConflictResolution
  Consolidation:
    Consolidation
  Enabled: Boolean
  MinAllowedConfidenceScoreForMerging: Number

```

## Properties

`ConflictResolution`

Determines how the auto-merging process should resolve conflicts between different
profiles. For example, if Profile A and Profile B have the same `FirstName`
and `LastName`, `ConflictResolution` specifies which
`EmailAddress` should be used.

_Required_: No

_Type_: [ConflictResolution](aws-properties-customerprofiles-domain-conflictresolution.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Consolidation`

A list of matching attributes that represent matching criteria. If two profiles meet
at least one of the requirements in the matching attributes list, they will be
merged.

_Required_: No

_Type_: [Consolidation](aws-properties-customerprofiles-domain-consolidation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

The flag that enables the auto-merging of duplicate profiles.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinAllowedConfidenceScoreForMerging`

A number between 0 and 1 that represents the minimum confidence score required for
profiles within a matching group to be merged during the auto-merge process. A higher
score means that a higher similarity is required to merge profiles.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributeTypesSelector

ConflictResolution

All content copied from https://docs.aws.amazon.com/.
