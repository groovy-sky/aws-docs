This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorSCEP::Challenge

For general-purpose connectors. Creates a _challenge password_ for the specified connector. The SCEP protocol uses a challenge password to authenticate a request before issuing a certificate from a certificate authority (CA). Your SCEP clients include the challenge password as part of their certificate request to Connector for SCEP. To retrieve the connector Amazon Resource Names (ARNs) for the connectors in your account, call [ListConnectors](https://docs.aws.amazon.com/pca-connector-scep/latest/APIReference/API_ListConnectors.html).

To create additional challenge passwords for the connector, call `CreateChallenge` again. We recommend frequently rotating your challenge passwords.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCAConnectorSCEP::Challenge",
  "Properties" : {
      "ConnectorArn" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::PCAConnectorSCEP::Challenge
Properties:
  ConnectorArn: String
  Tags:
    Key: Value

```

## Properties

`ConnectorArn`

The Amazon Resource Name (ARN) of the connector.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[a-z]+)*:pca-connector-scep:[a-z]+(-[a-z]+)+-[1-9]\d*:\d{12}:connector\/[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$`

_Minimum_: `5`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ChallengeArn`

The Amazon Resource Name (ARN) of the challenge.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Private Certificate Authority Connector for SCEP

AWS::PCAConnectorSCEP::Connector
