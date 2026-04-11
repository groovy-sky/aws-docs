This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob AppSpecification

Configuration to run a processing job in a specified container image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerArguments" : [ String, ... ],
  "ContainerEntrypoint" : [ String, ... ],
  "ImageUri" : String
}

```

### YAML

```yaml

  ContainerArguments:
    - String
  ContainerEntrypoint:
    - String
  ImageUri: String

```

## Properties

`ContainerArguments`

The arguments for a container used to run a processing job.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 0`

_Maximum_: `256 | 100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContainerEntrypoint`

The entrypoint for a container used to run a processing job.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 0`

_Maximum_: `256 | 100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageUri`

The container image to be run by the processing job.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::ProcessingJob

AthenaDatasetDefinition

All content copied from https://docs.aws.amazon.com/.
