This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::KnowledgeBase SourceConfiguration

Configuration information about the external data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppIntegrations" : AppIntegrationsConfiguration,
  "ManagedSourceConfiguration" : ManagedSourceConfiguration
}

```

### YAML

```yaml

  AppIntegrations:
    AppIntegrationsConfiguration
  ManagedSourceConfiguration:
    ManagedSourceConfiguration

```

## Properties

`AppIntegrations`

Configuration information for Amazon AppIntegrations to automatically ingest content.

_Required_: No

_Type_: [AppIntegrationsConfiguration](aws-properties-wisdom-knowledgebase-appintegrationsconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedSourceConfiguration`

Source configuration for managed resources.

_Required_: No

_Type_: [ManagedSourceConfiguration](aws-properties-wisdom-knowledgebase-managedsourceconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerSideEncryptionConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
