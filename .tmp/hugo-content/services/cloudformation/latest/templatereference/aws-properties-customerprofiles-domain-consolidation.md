This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain Consolidation

A list of matching attributes that represent matching criteria. If two profiles meet
at least one of the requirements in the matching attributes list, they will be
merged.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MatchingAttributesList" : [ [ , ... ], ... ]
}

```

### YAML

```yaml

  MatchingAttributesList:
    -
    -

```

## Properties

`MatchingAttributesList`

A list of matching criteria.

_Required_: Yes

_Type_: Array of Array

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConflictResolution

DataStore

All content copied from https://docs.aws.amazon.com/.
