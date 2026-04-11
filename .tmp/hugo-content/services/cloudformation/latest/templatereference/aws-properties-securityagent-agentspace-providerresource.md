This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::AgentSpace ProviderResource

Represents an integrated resource from a third-party provider. This is a union type that contains provider-specific resource information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GitHubCapabilities" : GitHubCapabilitiesResource,
  "GitHubRepository" : GitHubRepositoryResource
}

```

### YAML

```yaml

  GitHubCapabilities:
    GitHubCapabilitiesResource
  GitHubRepository:
    GitHubRepositoryResource

```

## Properties

`GitHubCapabilities`

The capabilities enabled for a GitHub resource integration.

_Required_: No

_Type_: [GitHubCapabilitiesResource](aws-properties-securityagent-agentspace-githubcapabilitiesresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GitHubRepository`

Represents a GitHub repository resource used in an integration.

_Required_: No

_Type_: [GitHubRepositoryResource](aws-properties-securityagent-agentspace-githubrepositoryresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IntegratedResource

Tag

All content copied from https://docs.aws.amazon.com/.
