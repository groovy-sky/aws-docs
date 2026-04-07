This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImageRecipe ComponentConfiguration

Configuration details of the component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentArn" : String,
  "Parameters" : [ ComponentParameter, ... ]
}

```

### YAML

```yaml

  ComponentArn: String
  Parameters:
    - ComponentParameter

```

## Properties

`ComponentArn`

The Amazon Resource Name (ARN) of the component.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:imagebuilder:[^:]+:(?:[0-9]{12}|aws(?:-[a-z-]+)?):component/[a-z0-9-_]+/(?:(?:([0-9]+|x)\.([0-9]+|x)\.([0-9]+|x))|(?:[0-9]+\.[0-9]+\.[0-9]+/[0-9]+))$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

A group of parameter settings that Image Builder uses to configure the component for a specific
recipe.

_Required_: No

_Type_: Array of [ComponentParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagerecipe-componentparameter.html)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AdditionalInstanceConfiguration

ComponentParameter
