This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSetEventDestination CloudWatchDestination

An object that defines an Amazon CloudWatch destination for email events. You can use Amazon CloudWatch to
monitor and gain insights on your email sending metrics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DimensionConfigurations" : [ DimensionConfiguration, ... ]
}

```

### YAML

```yaml

  DimensionConfigurations:
    - DimensionConfiguration

```

## Properties

`DimensionConfigurations`

An array of objects that define the dimensions to use when you send email events to
Amazon CloudWatch.

_Required_: No

_Type_: Array of [DimensionConfiguration](aws-properties-pinpointemail-configurationseteventdestination-dimensionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::PinpointEmail::ConfigurationSetEventDestination

DimensionConfiguration

All content copied from https://docs.aws.amazon.com/.
