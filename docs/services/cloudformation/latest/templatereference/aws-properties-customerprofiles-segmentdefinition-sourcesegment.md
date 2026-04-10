This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition SourceSegment

The source segments to build off of.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SegmentDefinitionName" : String
}

```

### YAML

```yaml

  SegmentDefinitionName: String

```

## Properties

`SegmentDefinitionName`

The name of the source segment.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SortAttribute

Tag

All content copied from https://docs.aws.amazon.com/.
