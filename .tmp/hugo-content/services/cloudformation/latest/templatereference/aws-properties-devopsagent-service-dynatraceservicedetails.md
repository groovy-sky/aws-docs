This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service DynatraceServiceDetails

Configuration details for registering a Dynatrace service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountUrn" : String,
  "AuthorizationConfig" : DynatraceAuthorizationConfig
}

```

### YAML

```yaml

  AccountUrn: String
  AuthorizationConfig:
    DynatraceAuthorizationConfig

```

## Properties

`AccountUrn`

The Dynatrace account URN.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AuthorizationConfig`

The authorization configuration for the Dynatrace service.

_Required_: No

_Type_: [DynatraceAuthorizationConfig](aws-properties-devopsagent-service-dynatraceauthorizationconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynatraceAuthorizationConfig

GitLabDetails

All content copied from https://docs.aws.amazon.com/.
