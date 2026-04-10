This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataPipeline::Pipeline Field

A key-value pair that describes a property of a `PipelineObject`. The value is specified as either a string value ( `StringValue`) or a reference to another object ( `RefValue`) but not as both.
To view fields for a
data pipeline object, see [Pipeline Object Reference](../../../datapipeline/latest/developerguide/dp-pipeline-objects.md) in the
_AWS Data Pipeline Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "RefValue" : String,
  "StringValue" : String
}

```

### YAML

```yaml

  Key: String
  RefValue: String
  StringValue:
    String

```

## Properties

`Key`

Specifies the name of a field for a particular object. To view valid values for a
particular field, see [Pipeline\
Object Reference](../../../datapipeline/latest/developerguide/dp-pipeline-objects.md) in the _AWS Data Pipeline Developer Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefValue`

A field value that you specify as an identifier of another object in the same
pipeline definition.

###### Note

You can specify the field value as either a string value
( `StringValue`) or a reference to another object
( `RefValue`), but not both.

Required if the key that you are using requires it.

_Required_: Conditional

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringValue`

A field value that you specify as a string. To view valid values for a
particular field, see [Pipeline\
Object Reference](../../../datapipeline/latest/developerguide/dp-pipeline-objects.md) in the _AWS Data Pipeline Developer Guide_.

###### Note

You can specify the field value as either a string value
( `StringValue`) or a reference to another object
( `RefValue`), but not both.

Required if the key that you are using requires it.

_Required_: Conditional

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataPipeline::Pipeline

ParameterAttribute

All content copied from https://docs.aws.amazon.com/.
