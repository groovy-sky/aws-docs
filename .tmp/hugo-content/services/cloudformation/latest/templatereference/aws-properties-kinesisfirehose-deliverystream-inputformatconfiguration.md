This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream InputFormatConfiguration

Specifies the deserializer you want to use to convert the format of the input data.
This parameter is required if `Enabled` is set to true.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Deserializer" : Deserializer
}

```

### YAML

```yaml

  Deserializer:
    Deserializer

```

## Properties

`Deserializer`

Specifies which deserializer to use. You can choose either the Apache Hive JSON SerDe
or the OpenX JSON SerDe. If both are non-null, the server rejects the request.

_Required_: No

_Type_: [Deserializer](aws-properties-kinesisfirehose-deliverystream-deserializer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcebergDestinationConfiguration

KinesisStreamSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
