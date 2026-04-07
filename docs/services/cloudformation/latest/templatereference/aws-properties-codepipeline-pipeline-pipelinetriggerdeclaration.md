This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline PipelineTriggerDeclaration

Represents information about the specified trigger configuration, such as the
filter criteria and the source stage for the action that contains the trigger.

###### Note

This is only supported for the `CodeStarSourceConnection` action
type.

###### Note

When a trigger configuration is specified, default change detection for
repository and branch commits is disabled.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GitConfiguration" : GitConfiguration,
  "ProviderType" : String
}

```

### YAML

```yaml

  GitConfiguration:
    GitConfiguration
  ProviderType: String

```

## Properties

`GitConfiguration`

Provides the filter criteria and the source stage for the repository event that
starts the pipeline, such as Git tags.

_Required_: No

_Type_: [GitConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codepipeline-pipeline-gitconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderType`

The source provider for the event, such as connections configured for a repository
with Git tags, for the specified trigger configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `CodeStarSourceConnection`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OutputArtifact

RetryConfiguration
