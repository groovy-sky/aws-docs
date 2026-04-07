This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityIntegration CreateDetails

Contains details required to create a code security integration with a specific
repository provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "gitlabSelfManaged" : CreateGitLabSelfManagedIntegrationDetail
}

```

### YAML

```yaml

  gitlabSelfManaged:
    CreateGitLabSelfManagedIntegrationDetail

```

## Properties

`gitlabSelfManaged`

Details specific to creating an integration with a self-managed GitLab
instance.

_Required_: Yes

_Type_: [CreateGitLabSelfManagedIntegrationDetail](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-codesecurityintegration-creategitlabselfmanagedintegrationdetail.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::InspectorV2::CodeSecurityIntegration

CreateGitLabSelfManagedIntegrationDetail
