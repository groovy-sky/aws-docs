This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service DynatraceAuthorizationConfig

The OAuth authorization configuration for a Dynatrace service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OAuthClientCredentials" : OAuthClientDetails
}

```

### YAML

```yaml

  OAuthClientCredentials:
    OAuthClientDetails

```

## Properties

`OAuthClientCredentials`

The OAuth client credentials for authenticating with Dynatrace.

_Required_: No

_Type_: [OAuthClientDetails](aws-properties-devopsagent-service-oauthclientdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BearerTokenDetails

DynatraceServiceDetails

All content copied from https://docs.aws.amazon.com/.
