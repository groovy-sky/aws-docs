This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NeptuneGraph::Graph Tag

A key-value pair to associate with the Neptune Analytics graph.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

**Key** (string)   –
A key is the required name of the tag. The string value can be
from 1 to 128 Unicode characters in length. It can't be
prefixed with `AWS:` and can only contain the
set of Unicode characters specified by this Java regular expression:
`"^([\p{L}\p{Z}\p{N}_.:/=+\-]*)$")`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

**Value** (string)   –
A value is the optional value of the tag. The string value can
be from 1 to 256 Unicode characters in length. It can't be
prefixed with `AWS:` and can only contain the
set of Unicode characters specified by this Java regular expression:
`"^([\p{L}\p{Z}\p{N}_.:/=+\-]*)$")`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::NeptuneGraph::Graph

VectorSearchConfiguration

All content copied from https://docs.aws.amazon.com/.
