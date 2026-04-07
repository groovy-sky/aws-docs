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

_Type_: Array of [ParameterAttribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datapipeline-pipeline-parameterattribute.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the parameter object.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ParameterAttribute

ParameterValue
