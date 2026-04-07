This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Fleet ComputeConfiguration

Contains compute attributes. These attributes only need be specified when your project's or fleet's `computeType` is set to `ATTRIBUTE_BASED_COMPUTE` or `CUSTOM_INSTANCE_TYPE`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "disk" : Integer,
  "instanceType" : String,
  "machineType" : String,
  "memory" : Integer,
  "vCpu" : Integer
}

```

### YAML

```yaml

  disk: Integer
  instanceType: String
  machineType: String
  memory: Integer
  vCpu: Integer

```

## Properties

`disk`

The amount of disk space of the instance type included in your fleet.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`instanceType`

The EC2 instance type to be launched in your fleet.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`machineType`

The machine type of the instance type included in your fleet.

_Required_: No

_Type_: String

_Allowed values_: `GENERAL | NVME`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`memory`

The amount of memory of the instance type included in your fleet.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`vCpu`

The number of vCPUs of the instance type included in your fleet.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CodeBuild::Fleet

FleetProxyRule
