This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline InputArtifact

Represents information about an artifact to be worked on, such as a test or build
artifact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

The name of the artifact to be worked on (for example, "My App").

Artifacts are the files that are worked on by actions in the pipeline. See the
action configuration for each action for details about artifact parameters. For example,
the S3 source action input artifact is a file name (or file path), and the files are
generally provided as a ZIP file. Example artifact name: SampleApp\_Windows.zip

The input artifact of an action must exactly match the output artifact declared in
a preceding action, but the input artifact does not have to be the next action in strict
sequence from the action that provided the output artifact. Actions in parallel can
declare different output artifacts, which are in turn consumed by different following
actions.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_\-]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GitTagFilterCriteria

OutputArtifact

All content copied from https://docs.aws.amazon.com/.
