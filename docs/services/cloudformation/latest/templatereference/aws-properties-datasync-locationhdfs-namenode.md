This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationHDFS NameNode

The NameNode of the Hadoop Distributed File System (HDFS). The NameNode manages the
file system's namespace and performs operations such as opening, closing, and renaming
files and directories. The NameNode also contains the information to map blocks of data
to the DataNodes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hostname" : String,
  "Port" : Integer
}

```

### YAML

```yaml

  Hostname: String
  Port: Integer

```

## Properties

`Hostname`

The hostname of the NameNode in the HDFS cluster. This value is the IP address or Domain
Name Service (DNS) name of the NameNode. An agent that's installed on-premises uses this
hostname to communicate with the NameNode in the network.

_Required_: Yes

_Type_: String

_Pattern_: `^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port that the NameNode uses to listen to client requests.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedSecretConfig

QopConfiguration

All content copied from https://docs.aws.amazon.com/.
