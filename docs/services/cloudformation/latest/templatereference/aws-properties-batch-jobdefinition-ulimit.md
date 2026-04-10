This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition Ulimit

The `ulimit` settings to pass to the container. For more information, see
[Ulimit](../../../../reference/amazonecs/latest/apireference/api-ulimit.md).

###### Note

This object isn't applicable to jobs that are running on Fargate resources.

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

The hard limit for the `ulimit` type.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The `type` of the `ulimit`. Valid values are: `core` \|
`cpu` \| `data` \| `fsize` \| `locks` \|
`memlock` \| `msgqueue` \| `nice` \| `nofile` \|
`nproc` \| `rss` \| `rtprio` \| `rttime` \|
`sigpending` \| `stack`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SoftLimit`

The soft limit for the `ulimit` type.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tmpfs

Volume

All content copied from https://docs.aws.amazon.com/.
