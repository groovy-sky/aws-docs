This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationNFS OnPremConfig

The AWS DataSync agents that can connect to your Network File System (NFS)
file server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentArns" : [ String, ... ]
}

```

### YAML

```yaml

  AgentArns:
    - String

```

## Properties

`AgentArns`

The Amazon Resource Names (ARNs) of the DataSync agents that can connect to
your NFS file server.

You can specify more than one agent. For more information, see [Using multiple DataSync agents](../../../datasync/latest/userguide/do-i-need-datasync-agent.md#multiple-agents).

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MountOptions

Tag

All content copied from https://docs.aws.amazon.com/.
