This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project ScopeConfiguration

Contains configuration information about the scope for a webhook.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Domain" : String,
  "Name" : String,
  "Scope" : String
}

```

### YAML

```yaml

  Domain: String
  Name: String
  Scope: String

```

## Properties

`Domain`

The domain of the GitHub Enterprise organization or the GitLab Self Managed group. Note that this parameter is only required if your project's source type is GITHUB\_ENTERPRISE or GITLAB\_SELF\_MANAGED.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of either the enterprise or organization that will send webhook events to CodeBuild, depending on if the webhook is a global or organization webhook respectively.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The type of scope for a GitHub or GitLab webhook. The scope default is GITHUB\_ORGANIZATION.

_Required_: No

_Type_: String

_Allowed values_: `GITHUB_ORGANIZATION | GITHUB_GLOBAL | GITLAB_GROUP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3LogsConfig

Source

All content copied from https://docs.aws.amazon.com/.
