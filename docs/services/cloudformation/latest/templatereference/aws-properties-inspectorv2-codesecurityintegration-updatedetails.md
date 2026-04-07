This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityIntegration UpdateDetails

Contains details required to update a code security integration with a specific
repository provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "github" : UpdateGitHubIntegrationDetail,
  "gitlabSelfManaged" : UpdateGitLabSelfManagedIntegrationDetail
}

```

### YAML

```yaml

  github:
    UpdateGitHubIntegrationDetail
  gitlabSelfManaged:
    UpdateGitLabSelfManagedIntegrationDetail

```

## Properties

`github`

Details specific to updating an integration with GitHub.

_Required_: No

_Type_: [UpdateGitHubIntegrationDetail](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-codesecurityintegration-updategithubintegrationdetail.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`gitlabSelfManaged`

Details specific to updating an integration with a self-managed GitLab
instance.

_Required_: No

_Type_: [UpdateGitLabSelfManagedIntegrationDetail](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-codesecurityintegration-updategitlabselfmanagedintegrationdetail.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateGitLabSelfManagedIntegrationDetail

UpdateGitHubIntegrationDetail
