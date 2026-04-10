This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppMesh::Route MatchRange

An object that represents the range of values to match on. The first character of the range is included in the range, though the last character is not. For example, if the range specified were 1-100, only values 1-99 would be matched.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "End" : Integer,
  "Start" : Integer
}

```

### YAML

```yaml

  End: Integer
  Start: Integer

```

## Properties

`End`

The end of the range.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Start`

The start of the range.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpTimeout

QueryParameter

All content copied from https://docs.aws.amazon.com/.
