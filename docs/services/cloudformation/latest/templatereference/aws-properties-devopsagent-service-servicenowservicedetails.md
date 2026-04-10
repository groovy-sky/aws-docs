This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service ServiceNowServiceDetails

Configuration details for registering a ServiceNow service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationConfig" : ServiceNowAuthorizationConfig,
  "InstanceUrl" : String
}

```

### YAML

```yaml

  AuthorizationConfig:
    ServiceNowAuthorizationConfig
  InstanceUrl: String

```

## Properties

`AuthorizationConfig`

The authorization configuration for the ServiceNow service.

_Required_: No

_Type_: [ServiceNowAuthorizationConfig](aws-properties-devopsagent-service-servicenowauthorizationconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceUrl`

The ServiceNow instance URL. Must be an HTTPS URL matching the format `https://<instance-name>.service-now.com`.

_Required_: Yes

_Type_: String

_Pattern_: `^https://[a-zA-Z0-9-]+\.service-now\.com/?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceNowAuthorizationConfig

Tag

All content copied from https://docs.aws.amazon.com/.
