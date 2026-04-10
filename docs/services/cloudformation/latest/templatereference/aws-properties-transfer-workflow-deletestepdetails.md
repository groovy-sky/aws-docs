This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow DeleteStepDetails

An object that contains the name and file location for a file being deleted by a
workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "SourceFileLocation" : String
}

```

### YAML

```yaml

  Name: String
  SourceFileLocation: String

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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecryptStepDetails

EfsInputFileLocation

All content copied from https://docs.aws.amazon.com/.
