This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow DataTransferApi

The API of the connector application that Amazon AppFlow uses to transfer your
data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  Name: String
  Type: String

```

## Properties

`Name`

The name of the connector application API.

_Required_: Yes

_Type_: String

_Pattern_: `[\w/-]+`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

You can specify one of the following types:

AUTOMATIC

The default. Optimizes a flow for datasets that fluctuate in size from small to
large. For each flow run, Amazon AppFlow chooses to use the SYNC or ASYNC API type
based on the amount of data that the run transfers.

SYNC

A synchronous API. This type of API optimizes a flow for small to medium-sized
datasets.

ASYNC

An asynchronous API. This type of API optimizes a flow for large datasets.

_Required_: Yes

_Type_: String

_Allowed values_: `SYNC | ASYNC | AUTOMATIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DatadogSourceProperties

DestinationConnectorProperties
