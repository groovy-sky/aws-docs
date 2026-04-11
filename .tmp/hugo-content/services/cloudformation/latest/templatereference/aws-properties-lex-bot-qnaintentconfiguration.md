This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot QnAIntentConfiguration

Details about the the configuration of the built-in `Amazon.QnAIntent`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BedrockModelConfiguration" : BedrockModelSpecification,
  "DataSourceConfiguration" : DataSourceConfiguration
}

```

### YAML

```yaml

  BedrockModelConfiguration:
    BedrockModelSpecification
  DataSourceConfiguration:
    DataSourceConfiguration

```

## Properties

`BedrockModelConfiguration`

Property description not available.

_Required_: Yes

_Type_: [BedrockModelSpecification](aws-properties-lex-bot-bedrockmodelspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSourceConfiguration`

Contains details about the configuration of the data source used for the `AMAZON.QnAIntent`.

_Required_: Yes

_Type_: [DataSourceConfiguration](aws-properties-lex-bot-datasourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QInConnectIntentConfiguration

QnAKendraConfiguration

All content copied from https://docs.aws.amazon.com/.
