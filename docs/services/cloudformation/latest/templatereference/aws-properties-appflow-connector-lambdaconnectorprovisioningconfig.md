This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Connector LambdaConnectorProvisioningConfig

Contains information about the configuration of the lambda which is being registered as
the connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LambdaArn" : String
}

```

### YAML

```yaml

  LambdaArn: String

```

## Properties

`LambdaArn`

Lambda ARN of the connector being registered.

_Required_: Yes

_Type_: String

_Pattern_: `arn:*:.*:.*:[0-9]+:.*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectorProvisioningConfig

AWS::AppFlow::ConnectorProfile
