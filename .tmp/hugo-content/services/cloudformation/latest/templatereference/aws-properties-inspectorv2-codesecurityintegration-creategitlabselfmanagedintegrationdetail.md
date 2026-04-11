This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityIntegration CreateGitLabSelfManagedIntegrationDetail

Contains details required to create an integration with a self-managed GitLab
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "accessToken" : String,
  "instanceUrl" : String
}

```

### YAML

```yaml

  accessToken: String
  instanceUrl: String

```

## Properties

`accessToken`

The personal access token used to authenticate with the self-managed GitLab instance.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`instanceUrl`

The URL of the self-managed GitLab instance.

_Required_: Yes

_Type_: String

_Pattern_: `^https://[-a-zA-Z0-9()@:%_+.~#?&//=]{1,1024}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateDetails

UpdateDetails

All content copied from https://docs.aws.amazon.com/.
