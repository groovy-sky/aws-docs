This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function LambdaManagedInstancesCapacityProviderConfig

Configuration for Lambda-managed instances used by the capacity provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityProviderArn" : String,
  "ExecutionEnvironmentMemoryGiBPerVCpu" : Number,
  "PerExecutionEnvironmentMaxConcurrency" : Integer
}

```

### YAML

```yaml

  CapacityProviderArn: String
  ExecutionEnvironmentMemoryGiBPerVCpu: Number
  PerExecutionEnvironmentMaxConcurrency: Integer

```

## Properties

`CapacityProviderArn`

The Amazon Resource Name (ARN) of the capacity provider.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z-]*:lambda:(eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}:\d{12}:capacity-provider:[a-zA-Z0-9-_]+$`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionEnvironmentMemoryGiBPerVCpu`

The amount of memory in GiB allocated per vCPU for execution environments.

_Required_: No

_Type_: Number

_Minimum_: `2`

_Maximum_: `8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerExecutionEnvironmentMaxConcurrency`

The maximum number of concurrent executions that can run on each execution environment.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageConfig

LoggingConfig

All content copied from https://docs.aws.amazon.com/.
