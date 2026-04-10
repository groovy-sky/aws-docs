This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector Capacity

Information about the capacity of the connector, whether it is auto scaled or
provisioned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScaling" : AutoScaling,
  "ProvisionedCapacity" : ProvisionedCapacity
}

```

### YAML

```yaml

  AutoScaling:
    AutoScaling
  ProvisionedCapacity:
    ProvisionedCapacity

```

## Properties

`AutoScaling`

Information about the auto scaling parameters for the connector.

_Required_: No

_Type_: [AutoScaling](aws-properties-kafkaconnect-connector-autoscaling.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedCapacity`

Details about a fixed capacity allocated to a connector.

_Required_: No

_Type_: [ProvisionedCapacity](aws-properties-kafkaconnect-connector-provisionedcapacity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScaling

CloudWatchLogsLogDelivery

All content copied from https://docs.aws.amazon.com/.
