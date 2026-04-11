This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline GitPullRequestFilter

The event criteria for the pull request trigger configuration, such as the lists of
branches or file paths to include and exclude.

The following are valid values for the events for this filter:

- CLOSED

- OPEN

- UPDATED

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Branches" : GitBranchFilterCriteria,
  "Events" : [ String, ... ],
  "FilePaths" : GitFilePathFilterCriteria
}

```

### YAML

```yaml

  Branches:
    GitBranchFilterCriteria
  Events:
    - String
  FilePaths:
    GitFilePathFilterCriteria

```

## Properties

`Branches`

The field that specifies to filter on branches for the pull request trigger
configuration.

_Required_: No

_Type_: [GitBranchFilterCriteria](aws-properties-codepipeline-pipeline-gitbranchfiltercriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Events`

The field that specifies which pull request events to filter on (OPEN, UPDATED,
CLOSED) for the trigger configuration.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilePaths`

The field that specifies to filter on file paths for the pull request trigger
configuration.

_Required_: No

_Type_: [GitFilePathFilterCriteria](aws-properties-codepipeline-pipeline-gitfilepathfiltercriteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitFilePathFilterCriteria

GitPushFilter

All content copied from https://docs.aws.amazon.com/.
