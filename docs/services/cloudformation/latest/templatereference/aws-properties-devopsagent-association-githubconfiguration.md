This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association GitHubConfiguration

Configuration for GitHub repository integration. Defines the repository name, numeric repository ID, owner name,
and owner type (user or organization) required for the Agent Space to access and interact with the GitHub
repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Owner" : String,
  "OwnerType" : String,
  "RepoId" : String,
  "RepoName" : String
}

```

### YAML

```yaml

  Owner: String
  OwnerType: String
  RepoId: String
  RepoName: String

```

## Properties

`Owner`

Repository owner.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnerType`

The type of repository owner.

_Allowed Values_: `organization` \| `user`

_Required_: Yes

_Type_: String

_Allowed values_: `organization | user`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepoId`

Associated Github repo ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepoName`

Associated Github repo name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventChannelConfiguration

GitLabConfiguration

All content copied from https://docs.aws.amazon.com/.
