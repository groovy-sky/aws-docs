This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service NewRelicServiceDetails

Configuration details for registering a New Relic service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationConfig" : NewRelicAuthorizationConfig
}

```

### YAML

```yaml

  AuthorizationConfig:
    NewRelicAuthorizationConfig

```

## Properties

`AuthorizationConfig`

The authorization configuration for the New Relic service.

_Required_: Yes

_Type_: [NewRelicAuthorizationConfig](aws-properties-devopsagent-service-newrelicauthorizationconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NewRelicAuthorizationConfig

OAuthClientDetails

All content copied from https://docs.aws.amazon.com/.
