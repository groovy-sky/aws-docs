This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::MLTransform TransformParameters

The algorithm-specific parameters that are associated with the machine learning
transform.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FindMatchesParameters" : FindMatchesParameters,
  "TransformType" : String
}

```

### YAML

```yaml

  FindMatchesParameters:
    FindMatchesParameters
  TransformType: String

```

## Properties

`FindMatchesParameters`

The parameters for the find matches algorithm.

_Required_: No

_Type_: [FindMatchesParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-mltransform-findmatchesparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformType`

The type of machine learning transform. `FIND_MATCHES` is the only option.

For information about the types of machine learning transforms, see [Working with machine learning transforms](https://docs.aws.amazon.com/glue/latest/dg/console-machine-learning-transforms.html).

_Required_: Yes

_Type_: String

_Allowed values_: `FIND_MATCHES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransformEncryption

AWS::Glue::Partition
