This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime AuthorizerConfiguration

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

_Required_: No

_Type_: [CustomJWTAuthorizerConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-runtime-customjwtauthorizerconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AgentRuntimeArtifact

AuthorizingClaimMatchValueType
