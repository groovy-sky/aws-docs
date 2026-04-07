This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::Service InstanceConfiguration

Describes the runtime configuration of an AWS App Runner service instance (scaling unit).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cpu" : String,
  "InstanceRoleArn" : String,
  "Memory" : String
}

```

### YAML

```yaml

  Cpu: String
  InstanceRoleArn: String
  Memory: String

```

## Properties

`Cpu`

The number of CPU units reserved for each instance of your App Runner service.

Default: `1 vCPU`

_Required_: No

_Type_: String

_Pattern_: `256|512|1024|2048|4096|(0.25|0.5|1|2|4) vCPU`

_Minimum_: `3`

_Maximum_: `9`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceRoleArn`

The Amazon Resource Name (ARN) of an IAM role that provides permissions to your App Runner service. These are permissions that your code needs when it calls
any AWS APIs.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):iam::[0-9]{12}:role/[\w+=,.@-]{1,64}`

_Minimum_: `29`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Memory`

The amount of memory, in MB or GB, reserved for each instance of your App Runner service.

Default: `2 GB`

_Required_: No

_Type_: String

_Pattern_: `512|1024|2048|3072|4096|6144|8192|10240|12288|(0.5|1|2|3|4|6|8|10|12) GB`

_Minimum_: `3`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IngressConfiguration

KeyValuePair
