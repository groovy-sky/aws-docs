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

_Type_: [AutoScaling](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kafkaconnect-connector-autoscaling.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedCapacity`

Details about a fixed capacity allocated to a connector.

_Required_: No

_Type_: [ProvisionedCapacity](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kafkaconnect-connector-provisionedcapacity.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AutoScaling

CloudWatchLogsLogDelivery
