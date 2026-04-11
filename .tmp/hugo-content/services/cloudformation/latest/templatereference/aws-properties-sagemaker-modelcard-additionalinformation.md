This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard AdditionalInformation

Additional information about the model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CaveatsAndRecommendations" : String,
  "CustomDetails" : {Key: Value, ...},
  "EthicalConsiderations" : String
}

```

### YAML

```yaml

  CaveatsAndRecommendations: String
  CustomDetails:
    Key: Value
  EthicalConsiderations: String

```

## Properties

`CaveatsAndRecommendations`

Caveats and recommendations for those who might use this model in their applications.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDetails`

Any additional information to document about the model.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_][a-zA-Z0-9_]*`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EthicalConsiderations`

Any ethical considerations documented by the model card author.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::ModelCard

BusinessDetails

All content copied from https://docs.aws.amazon.com/.
