This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataPipeline::Pipeline ParameterObject

Contains information about a parameter object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : [ ParameterAttribute, ... ],
  "Id" : String
}

```

### YAML

```yaml

  Attributes:
    - ParameterAttribute
  Id: String

```

## Properties

`Attributes`

The attributes of the parameter object.

_Required_: Yes

_Type_: Array of [ParameterAttribute](aws-properties-datapipeline-pipeline-parameterattribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the parameter object.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterAttribute

ParameterValue

All content copied from https://docs.aws.amazon.com/.
