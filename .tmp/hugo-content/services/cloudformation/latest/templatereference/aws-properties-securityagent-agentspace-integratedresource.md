This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::AgentSpace IntegratedResource

Represents an integrated resource from a third-party provider. This is a union type that contains provider-specific resource information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Integration" : String,
  "ProviderResources" : [ ProviderResource, ... ]
}

```

### YAML

```yaml

  Integration: String
  ProviderResources:
    - ProviderResource

```

## Properties

`Integration`

The unique identifier of the integration that provides access to the resource.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderResources`

The metadata for the integrated resource.

_Required_: Yes

_Type_: Array of [ProviderResource](aws-properties-securityagent-agentspace-providerresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitHubRepositoryResource

ProviderResource

All content copied from https://docs.aws.amazon.com/.
