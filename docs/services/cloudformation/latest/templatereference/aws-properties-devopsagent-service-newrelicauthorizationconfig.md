This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service NewRelicAuthorizationConfig

The authorization configuration for a New Relic service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKey" : NewRelicApiKeyConfig
}

```

### YAML

```yaml

  ApiKey:
    NewRelicApiKeyConfig

```

## Properties

`ApiKey`

The API key configuration for authenticating with New Relic.

_Required_: Yes

_Type_: [NewRelicApiKeyConfig](aws-properties-devopsagent-service-newrelicapikeyconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NewRelicApiKeyConfig

NewRelicServiceDetails

All content copied from https://docs.aws.amazon.com/.
