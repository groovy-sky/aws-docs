This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline GitPushFilter

The event criteria that specify when a specified repository event will start the
pipeline for the specified trigger configuration, such as the lists of Git tags to
include and exclude.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Branches" : GitBranchFilterCriteria,
  "FilePaths" : GitFilePathFilterCriteria,
  "Tags" : GitTagFilterCriteria
}

```

### YAML

```yaml

  Branches:
    GitBranchFilterCriteria
  FilePaths:
    GitFilePathFilterCriteria
  Tags:
    GitTagFilterCriteria

```

## Properties

`Branches`

The field that specifies to filter on branches for the push trigger
configuration.

_Required_: No

_Type_: [GitBranchFilterCriteria](aws-properties-codepipeline-pipeline-gitbranchfiltercriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilePaths`

The field that specifies to filter on file paths for the push trigger
configuration.

_Required_: No

_Type_: [GitFilePathFilterCriteria](aws-properties-codepipeline-pipeline-gitfilepathfiltercriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The field that contains the details for the Git tags trigger
configuration.

_Required_: No

_Type_: [GitTagFilterCriteria](aws-properties-codepipeline-pipeline-gittagfiltercriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitPullRequestFilter

GitTagFilterCriteria

All content copied from https://docs.aws.amazon.com/.
