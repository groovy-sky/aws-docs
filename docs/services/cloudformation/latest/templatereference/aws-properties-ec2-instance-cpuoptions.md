This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance CpuOptions

Specifies the CPU options for the instance. When you specify CPU options, you must
specify both the number of CPU cores and threads per core.

Modifying the CPU options for an instance results in instance [replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement).

For more information, see [Optimize CPU options](../../../ec2/latest/userguide/instance-optimize-cpu.md) in
the _Amazon Elastic Compute Cloud User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CoreCount" : Integer,
  "ThreadsPerCore" : Integer
}

```

### YAML

```yaml

  CoreCount: Integer
  ThreadsPerCore: Integer

```

## Properties

`CoreCount`

The number of CPU cores for the instance.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThreadsPerCore`

The number of threads per CPU core.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BlockDeviceMapping

CreditSpecification
