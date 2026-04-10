This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityIntegration UpdateGitHubIntegrationDetail

Contains details required to update an integration with GitHub.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "code" : String,
  "installationId" : String
}

```

### YAML

```yaml

  code: String
  installationId: String

```

## Properties

`code`

The authorization code received from GitHub to update the integration.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`installationId`

The installation ID of the GitHub App associated with the integration.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateDetails

UpdateGitLabSelfManagedIntegrationDetail

All content copied from https://docs.aws.amazon.com/.
