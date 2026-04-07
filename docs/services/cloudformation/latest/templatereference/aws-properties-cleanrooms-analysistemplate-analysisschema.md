This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate AnalysisSchema

A relation within an analysis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReferencedTables" : [ String, ... ]
}

```

### YAML

```yaml

  ReferencedTables:
    - String

```

## Properties

`ReferencedTables`

The tables referenced in the analysis schema.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Update requires_: Updates are not supported.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnalysisParameter

AnalysisSource
