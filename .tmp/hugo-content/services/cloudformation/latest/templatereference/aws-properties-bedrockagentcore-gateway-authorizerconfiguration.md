This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway AuthorizerConfiguration

Represents inbound authorization configuration options used to authenticate incoming requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomJWTAuthorizer" : CustomJWTAuthorizerConfiguration
}

```

### YAML

```yaml

  CustomJWTAuthorizer:
    CustomJWTAuthorizerConfiguration

```

## Properties

`CustomJWTAuthorizer`

The inbound JWT-based authorization, specifying how incoming requests should be authenticated.

_Required_: Yes

_Type_: [CustomJWTAuthorizerConfiguration](aws-properties-bedrockagentcore-gateway-customjwtauthorizerconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BedrockAgentCore::Gateway

AuthorizingClaimMatchValueType

All content copied from https://docs.aws.amazon.com/.
