This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Segment SourceSegments

Specifies the base segment to build the segment on. A base segment, also called a
_source segment_, defines the initial population of endpoints for
a segment. When you add dimensions to the segment, Amazon Pinpoint filters the base
segment by using the dimensions that you specify.

You can specify more than one dimensional segment or only one imported segment. If you
specify an imported segment, the segment size estimate that displays on the Amazon
Pinpoint console indicates the size of the imported segment without any filters applied
to it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Version" : Integer
}

```

### YAML

```yaml

  Id: String
  Version: Integer

```

## Properties

`Id`

The unique identifier for the source segment.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version number of the source segment.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SetDimension

AWS::Pinpoint::SMSChannel

All content copied from https://docs.aws.amazon.com/.
