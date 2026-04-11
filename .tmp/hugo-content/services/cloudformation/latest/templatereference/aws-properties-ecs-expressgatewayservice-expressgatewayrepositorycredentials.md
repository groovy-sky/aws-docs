This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::ExpressGatewayService ExpressGatewayRepositoryCredentials

The repository credentials for private registry authentication to pass to the container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CredentialsParameter" : String
}

```

### YAML

```yaml

  CredentialsParameter: String

```

## Properties

`CredentialsParameter`

The Amazon Resource Name (ARN) of the secret containing the private repository credentials.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExpressGatewayContainer

ExpressGatewayScalingTarget

All content copied from https://docs.aws.amazon.com/.
