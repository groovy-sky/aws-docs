This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard TrainingMetric

A result from a SageMaker AI training job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Notes" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Name: String
  Notes: String
  Value: Number

```

## Properties

`Name`

The name of the result from the SageMaker AI training job.

_Required_: Yes

_Type_: String

_Pattern_: `.{1,255}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Notes`

Any additional notes describing the result of the training job.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of a result from the SageMaker AI training job.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrainingJobDetails

UserContext

All content copied from https://docs.aws.amazon.com/.
