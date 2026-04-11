This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate CpuOptions

Specifies the CPU options for an instance. For more information, see [Optimize\
CPU options](../../../ec2/latest/userguide/instance-optimize-cpu.md) in the _Amazon Elastic Compute Cloud User_
_Guide_.

`CpuOptions` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmdSevSnp" : String,
  "CoreCount" : Integer,
  "ThreadsPerCore" : Integer
}

```

### YAML

```yaml

  AmdSevSnp: String
  CoreCount: Integer
  ThreadsPerCore: Integer

```

## Properties

`AmdSevSnp`

Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported
with M6a, R6a, and C6a instance types only. For more information, see [AMD SEV-SNP for\
Amazon EC2 instances](../../../ec2/latest/userguide/sev-snp.md).

_Required_: No

_Type_: String

_Allowed values_: `enabled | disabled`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CoreCount`

The number of CPU cores for the instance.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThreadsPerCore`

The number of threads per CPU core. To disable multithreading for the instance,
specify a value of `1`. Otherwise, specify the default value of
`2`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Optimize CPU\
options](../../../ec2/latest/userguide/instance-optimize-cpu.md) in the _Amazon EC2 User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Cpu

CreditSpecification

All content copied from https://docs.aws.amazon.com/.
