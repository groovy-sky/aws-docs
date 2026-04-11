This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet AcceleratorCountRequest

The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips)
on an instance. To exclude accelerator-enabled instance types, set `Max` to
`0`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Max" : Integer,
  "Min" : Integer
}

```

### YAML

```yaml

  Max: Integer
  Min: Integer

```

## Properties

`Max`

The maximum number of accelerators. To specify no maximum limit, omit this
parameter. To exclude accelerator-enabled instance types, set `Max` to
`0`.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Min`

The minimum number of accelerators. To specify no minimum limit, omit this
parameter.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::EC2Fleet

AcceleratorTotalMemoryMiBRequest

All content copied from https://docs.aws.amazon.com/.
