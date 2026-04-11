This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow TagStepDetails

Details for a step that creates one or more tags.

You specify one or more tags. Each tag contains a key-value pair.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "SourceFileLocation" : String,
  "Tags" : [ S3Tag, ... ]
}

```

### YAML

```yaml

  Name: String
  SourceFileLocation: String
  Tags:
    - S3Tag

```

## Properties

`Name`

The name of the step, used as an identifier.

_Required_: No

_Type_: String

_Pattern_: `^[\w-]*$`

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceFileLocation`

Specifies which file to use as input to the workflow step: either the output from the previous step, or the originally uploaded file
for the workflow.

- To use the previous file as the input, enter `${previous.file}`.
In this case, this workflow step uses the output file from the previous workflow step as input.
This is the default value.

- To use the originally uploaded file location as input for this step, enter `${original.file}`.

_Required_: No

_Type_: String

_Pattern_: `^\$\{(\w+.)+\w+\}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Array that contains from 1 to 10 key/value pairs.

_Required_: No

_Type_: Array of [S3Tag](aws-properties-transfer-workflow-s3tag.md)

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

WorkflowStep

All content copied from https://docs.aws.amazon.com/.
