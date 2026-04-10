This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::App JobConfig

Describes the configuration details that apply to the jobs for an Amplify app.

Use `JobConfig` to apply configuration to jobs, such as customizing the build instance size when you create or
update an Amplify app. For more information about customizable build
instances, see [Custom build instances](../../../amplify/latest/userguide/custom-build-instance.md) in the _Amplify User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BuildComputeType" : String
}

```

### YAML

```yaml

  BuildComputeType: String

```

## Properties

`BuildComputeType`

Specifies the size of the build instance. Amplify supports three
instance sizes: `STANDARD_8GB`, `LARGE_16GB`, and
`XLARGE_72GB`. If you don't specify a value, Amplify uses
the `STANDARD_8GB` default.

The following list describes the CPU, memory, and storage capacity for each build
instance type:

STANDARD\_8GB

- vCPUs: 4

- Memory: 8 GiB

- Disk space: 128 GB

LARGE\_16GB

- vCPUs: 8

- Memory: 16 GiB

- Disk space: 128 GB

XLARGE\_72GB

- vCPUs: 36

- Memory: 72 GiB

- Disk space: 256 GB

_Required_: Yes

_Type_: String

_Allowed values_: `STANDARD_8GB | LARGE_16GB | XLARGE_72GB`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentVariable

Tag

All content copied from https://docs.aws.amazon.com/.
