This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImagePipeline AutoDisablePolicy

Defines the rules by which an image pipeline is automatically disabled when it fails.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FailureCount" : Integer
}

```

### YAML

```yaml

  FailureCount: Integer

```

## Properties

`FailureCount`

The number of consecutive scheduled image pipeline executions that must fail before Image Builder
automatically disables the pipeline.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ImageBuilder::ImagePipeline

EcrConfiguration

All content copied from https://docs.aws.amazon.com/.
