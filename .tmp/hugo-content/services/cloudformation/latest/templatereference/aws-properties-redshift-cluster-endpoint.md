This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::Cluster Endpoint

Describes a connection endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Address" : String,
  "Port" : String
}

```

### YAML

```yaml

  Address: String
  Port: String

```

## Properties

`Address`

The DNS address of the cluster. This property is read only.

_Required_: No

_Type_: String

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port that the database engine is listening on. This property is read only.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Redshift::Cluster

LoggingProperties

All content copied from https://docs.aws.amazon.com/.
