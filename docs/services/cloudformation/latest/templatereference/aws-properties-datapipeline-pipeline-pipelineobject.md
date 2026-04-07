This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataPipeline::Pipeline PipelineObject

PipelineObject is property of the AWS::DataPipeline::Pipeline resource that contains
information about a pipeline object. This can be a logical, physical, or physical attempt
pipeline object. The complete set of components of a pipeline defines the pipeline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Fields" : [ Field, ... ],
  "Id" : String,
  "Name" : String
}

```

### YAML

```yaml

  Fields:
    - Field
  Id: String
  Name: String

```

## Properties

`Fields`

Key-value pairs that define the properties of the object.

_Required_: Yes

_Type_: Array of [Field](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datapipeline-pipeline-field.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the object.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\n\t]*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the object.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\n\t]*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ParameterValue

PipelineTag
