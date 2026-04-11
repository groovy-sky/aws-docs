This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::AgentSpace GitHubRepositoryResource

Represents a GitHub repository resource used in an integration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Owner" : String
}

```

### YAML

```yaml

  Name: String
  Owner: String

```

## Properties

`Name`

The name of the GitHub repository.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Owner`

The owner of the GitHub repository.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitHubCapabilitiesResource

IntegratedResource

All content copied from https://docs.aws.amazon.com/.
