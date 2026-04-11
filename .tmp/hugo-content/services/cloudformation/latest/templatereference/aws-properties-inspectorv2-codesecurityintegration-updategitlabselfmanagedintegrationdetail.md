This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityIntegration UpdateGitLabSelfManagedIntegrationDetail

Contains details required to update an integration with a self-managed GitLab
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "authCode" : String
}

```

### YAML

```yaml

  authCode: String

```

## Properties

`authCode`

The authorization code received from the self-managed GitLab instance to update the
integration.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateGitHubIntegrationDetail

AWS::InspectorV2::CodeSecurityScanConfiguration

All content copied from https://docs.aws.amazon.com/.
