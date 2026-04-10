This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget IamCredentialProvider

An IAM credential provider for gateway authentication. This structure contains the configuration for authenticating with the target endpoint using IAM credentials and SigV4 signing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Region" : String,
  "Service" : String
}

```

### YAML

```yaml

  Region: String
  Service: String

```

## Properties

`Region`

The AWS Region used for SigV4 signing. If not specified, defaults to the gateway's Region.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Service`

The target AWS service name used for SigV4 signing. This value identifies the service that the gateway authenticates with when making requests to the target endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CredentialProviderConfiguration

McpLambdaTargetConfiguration

All content copied from https://docs.aws.amazon.com/.
