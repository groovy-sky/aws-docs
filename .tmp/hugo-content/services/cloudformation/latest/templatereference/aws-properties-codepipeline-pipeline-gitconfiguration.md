This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline GitConfiguration

A type of trigger configuration for Git-based source actions.

###### Note

You can specify the Git configuration trigger type for all third-party
Git-based source actions that are supported by the
`CodeStarSourceConnection` action type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PullRequest" : [ GitPullRequestFilter, ... ],
  "Push" : [ GitPushFilter, ... ],
  "SourceActionName" : String
}

```

### YAML

```yaml

  PullRequest:
    - GitPullRequestFilter
  Push:
    - GitPushFilter
  SourceActionName: String

```

## Properties

`PullRequest`

The field where the repository event that will start the pipeline is specified as
pull requests.

_Required_: No

_Type_: Array of [GitPullRequestFilter](aws-properties-codepipeline-pipeline-gitpullrequestfilter.md)

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Push`

The field where the repository event that will start the pipeline, such as pushing
Git tags, is specified with details.

_Required_: No

_Type_: Array of [GitPushFilter](aws-properties-codepipeline-pipeline-gitpushfilter.md)

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceActionName`

The name of the pipeline source action where the trigger configuration, such as Git
tags, is specified. The trigger configuration will start the pipeline upon the specified
change only.

###### Note

You can only specify one trigger configuration per source action.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitBranchFilterCriteria

GitFilePathFilterCriteria

All content copied from https://docs.aws.amazon.com/.
