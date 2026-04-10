This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Recipe Action

Represents a transformation and associated parameters that are used to apply a change
to an AWS Glue DataBrew dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operation" : String,
  "Parameters" : RecipeParameters
}

```

### YAML

```yaml

  Operation: String
  Parameters:
    RecipeParameters

```

## Properties

`Operation`

The name of a valid DataBrew transformation to be performed on the
data.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

Contextual parameters for the transformation.

_Required_: No

_Type_: [RecipeParameters](aws-properties-databrew-recipe-recipeparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataBrew::Recipe

ConditionExpression

All content copied from https://docs.aws.amazon.com/.
