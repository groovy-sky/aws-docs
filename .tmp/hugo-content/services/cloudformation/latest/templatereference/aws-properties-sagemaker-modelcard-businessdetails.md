This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard BusinessDetails

Information about how the model supports business goals.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BusinessProblem" : String,
  "BusinessStakeholders" : String,
  "LineOfBusiness" : String
}

```

### YAML

```yaml

  BusinessProblem: String
  BusinessStakeholders: String
  LineOfBusiness: String

```

## Properties

`BusinessProblem`

The specific business problem that the model is trying to solve.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BusinessStakeholders`

The relevant stakeholders for the model.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LineOfBusiness`

The broader business need that the model is serving.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdditionalInformation

Container

All content copied from https://docs.aws.amazon.com/.
