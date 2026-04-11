This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline GitFilePathFilterCriteria

The Git repository file paths specified as filter criteria to start the
pipeline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Excludes" : [ String, ... ],
  "Includes" : [ String, ... ]
}

```

### YAML

```yaml

  Excludes:
    - String
  Includes:
    - String

```

## Properties

`Excludes`

The list of patterns of Git repository file paths that, when a commit is pushed,
are to be excluded from starting the pipeline.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Includes`

The list of patterns of Git repository file paths that, when a commit is pushed,
are to be included as criteria that starts the pipeline.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitConfiguration

GitPullRequestFilter

All content copied from https://docs.aws.amazon.com/.
