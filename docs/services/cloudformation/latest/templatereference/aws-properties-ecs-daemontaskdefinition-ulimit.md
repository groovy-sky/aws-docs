This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::DaemonTaskDefinition Ulimit

The `ulimit` settings to pass to the container.

Amazon ECS tasks hosted on AWS Fargate use the default resource limit values
set by the operating system with the exception of the `nofile` resource limit
parameter which AWS Fargate overrides. The `nofile` resource limit sets a
restriction on the number of open files that a container can use. The default
`nofile` soft limit is ` 65535` and the default hard limit is
`65535`.

You can specify the `ulimit` settings for a container in a task
definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HardLimit" : Integer,
  "Name" : String,
  "SoftLimit" : Integer
}

```

### YAML

```yaml

  HardLimit: Integer
  Name: String
  SoftLimit: Integer

```

## Properties

`HardLimit`

The hard limit for the `ulimit` type. The value can be specified in bytes,
seconds, or as a count, depending on the `type` of the
`ulimit`.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The `type` of the `ulimit`.

_Required_: Yes

_Type_: String

_Allowed values_: `core | cpu | data | fsize | locks | memlock | msgqueue | nice | nofile | nproc | rss | rtprio | rttime | sigpending | stack`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SoftLimit`

The soft limit for the `ulimit` type. The value can be specified in bytes,
seconds, or as a count, depending on the `type` of the
`ulimit`.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tmpfs

Volume

All content copied from https://docs.aws.amazon.com/.
