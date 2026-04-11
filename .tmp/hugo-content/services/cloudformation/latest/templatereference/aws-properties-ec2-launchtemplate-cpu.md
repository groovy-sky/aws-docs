This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate Cpu

Specifies the CPU performance to consider when using an instance family as the baseline reference.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "References" : [ Reference, ... ]
}

```

### YAML

```yaml

  References:
    - Reference

```

## Properties

`References`

The instance family to use as the baseline reference for CPU performance. All
instance types that match your specified attributes are compared against the CPU
performance of the referenced instance family, regardless of CPU manufacturer or
architecture differences.

_Required_: No

_Type_: Array of [Reference](aws-properties-ec2-launchtemplate-reference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionTrackingSpecification

CpuOptions

All content copied from https://docs.aws.amazon.com/.
