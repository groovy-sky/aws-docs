This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet AcceleratorCapabilities

Provides information about the GPU accelerators used for jobs processed by a
fleet.

###### Important

Accelerator capabilities cannot be used with wait-and-save fleets. If you specify accelerator capabilities, you must use either spot or on-demand instance market options.

###### Note

Each accelerator type maps to specific EC2 instance families:

- `t4`: Uses G4dn instance family

- `a10g`: Uses G5 instance family

- `l4`: Uses G6 and Gr6 instance families

- `l40s`: Uses G6e instance family

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Count" : AcceleratorCountRange,
  "Selections" : [ AcceleratorSelection, ... ]
}

```

### YAML

```yaml

  Count:
    AcceleratorCountRange
  Selections:
    - AcceleratorSelection

```

## Properties

`Count`

The number of GPU accelerators specified for worker hosts in this fleet.

###### Important

You must specify either `acceleratorCapabilities.count.max` or `allowedInstanceTypes` when using accelerator capabilities. If you don't specify a maximum count, AWS Deadline Cloud uses the instance types you specify in `allowedInstanceTypes` to determine the maximum number of accelerators.

_Required_: No

_Type_: [AcceleratorCountRange](aws-properties-deadline-fleet-acceleratorcountrange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Selections`

A list of accelerator capabilities requested for this fleet. Only Amazon Elastic Compute Cloud instances
that provide these capabilities will be used. For example, if you specify both L4 and T4
chips, AWS Deadline Cloud will use Amazon EC2 instances that have either the L4 or the T4 chip
installed.

###### Important

- You must specify at least one accelerator selection.

- You cannot specify the same accelerator name multiple times in the selections list.

- All accelerators in the selections must use the same runtime version.

_Required_: Yes

_Type_: Array of [AcceleratorSelection](aws-properties-deadline-fleet-acceleratorselection.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Deadline::Fleet

AcceleratorCountRange

All content copied from https://docs.aws.amazon.com/.
