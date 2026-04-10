This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector ProvisionedCapacity

Details about a connector's provisioned capacity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "McuCount" : Integer,
  "WorkerCount" : Integer
}

```

### YAML

```yaml

  McuCount: Integer
  WorkerCount: Integer

```

## Properties

`McuCount`

The number of microcontroller units (MCUs) allocated to each connector worker. The valid
values are 1,2,4,8.

_Required_: Yes

_Type_: Integer

_Allowed values_: `1 | 2 | 4 | 8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerCount`

The number of workers that are allocated to the connector.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Plugin

S3LogDelivery

All content copied from https://docs.aws.amazon.com/.
