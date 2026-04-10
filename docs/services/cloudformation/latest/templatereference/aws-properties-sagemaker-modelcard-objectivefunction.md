This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard ObjectiveFunction

The function that is optimized during model training.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Function" : Function,
  "Notes" : String
}

```

### YAML

```yaml

  Function:
    Function
  Notes: String

```

## Properties

`Function`

A function object that details optimization direction, metric, and additional descriptions.

_Required_: No

_Type_: [Function](aws-properties-sagemaker-modelcard-function.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Notes`

Notes about the object function, including other considerations for possible objective functions.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ModelPackageDetails

SecurityConfig

All content copied from https://docs.aws.amazon.com/.
