This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe BatchArrayProperties

The array properties for the submitted job, such as the size of the array. The array size can be between 2 and 10,000.
If you specify array properties for a job, it becomes an array job. This parameter is used only if the target is an AWS Batch job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Size" : Integer
}

```

### YAML

```yaml

  Size: Integer

```

## Properties

`Size`

The size of the array, if this is an array batch job.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AwsVpcConfiguration

BatchContainerOverrides

All content copied from https://docs.aws.amazon.com/.
