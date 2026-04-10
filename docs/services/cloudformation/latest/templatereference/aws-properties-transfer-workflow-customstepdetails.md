This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow CustomStepDetails

Details for a step that invokes an AWS Lambda function.

Consists of the Lambda function's name, target, and timeout (in seconds).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "SourceFileLocation" : String,
  "Target" : String,
  "TimeoutSeconds" : Integer
}

```

### YAML

```yaml

  Name: String
  SourceFileLocation: String
  Target: String
  TimeoutSeconds: Integer

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

`Target`

The ARN for the Lambda function that is being called.

_Required_: No

_Type_: String

_Pattern_: `arn:[a-z-]+:lambda:.*$`

_Minimum_: `0`

_Maximum_: `170`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeoutSeconds`

Timeout, in seconds, for the step.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1800`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CopyStepDetails

DecryptStepDetails

All content copied from https://docs.aws.amazon.com/.
