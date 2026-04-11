This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance CreditSpecification

Specifies the credit option for CPU usage of a T instance.

`CreditSpecification` is a property of the [AWS::EC2::Instance](../userguide/aws-properties-ec2-instance.md) resource.

For more information, see [Burstable performance instances](../../../ec2/latest/userguide/burstable-performance-instances.md) in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CPUCredits" : String
}

```

### YAML

```yaml

  CPUCredits: String

```

## Properties

`CPUCredits`

The credit option for CPU usage of the instance.

Valid values: `standard` \| `unlimited`

T3 instances with `host` tenancy do not support the `unlimited`
CPU credit option.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CpuOptions

Ebs

All content copied from https://docs.aws.amazon.com/.
