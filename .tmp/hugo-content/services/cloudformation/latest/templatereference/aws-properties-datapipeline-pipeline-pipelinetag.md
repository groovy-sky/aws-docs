This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataPipeline::Pipeline PipelineTag

A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you
can use to control permissions. For more information, see [Controlling Access to\
Pipelines and Resources](../../../datapipeline/latest/developerguide/dp-control-access.md) in the
_AWS Data Pipeline Developer Guide_.

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

The key name of a tag.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to associate with the key name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipelineObject

Next

All content copied from https://docs.aws.amazon.com/.
