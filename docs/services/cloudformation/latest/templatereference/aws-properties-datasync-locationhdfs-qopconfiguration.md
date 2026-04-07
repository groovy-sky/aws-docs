This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationHDFS QopConfiguration

The Quality of Protection (QOP) configuration specifies the Remote Procedure Call
(RPC) and data transfer privacy settings configured on the Hadoop Distributed File
System (HDFS) cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataTransferProtection" : String,
  "RpcProtection" : String
}

```

### YAML

```yaml

  DataTransferProtection: String
  RpcProtection: String

```

## Properties

`DataTransferProtection`

The data transfer protection setting configured on the HDFS cluster. This setting
corresponds to your `dfs.data.transfer.protection` setting in the
`hdfs-site.xml` file on your Hadoop cluster.

_Required_: No

_Type_: String

_Allowed values_: `AUTHENTICATION | INTEGRITY | PRIVACY | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RpcProtection`

The Remote Procedure Call (RPC) protection setting configured on the HDFS cluster.
This setting corresponds to your `hadoop.rpc.protection` setting in your
`core-site.xml` file on your Hadoop cluster.

_Required_: No

_Type_: String

_Allowed values_: `AUTHENTICATION | INTEGRITY | PRIVACY | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NameNode

Tag
