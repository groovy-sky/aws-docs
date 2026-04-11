This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream Deserializer

The deserializer you want Kinesis Data Firehose to use for converting the input data
from JSON. Kinesis Data Firehose then serializes the data to its final format using the
`Serializer`. Kinesis Data Firehose supports two types of deserializers: the
[Apache Hive JSON SerDe](https://cwiki.apache.org/confluence/display/Hive/LanguageManual+DDL) and the [OpenX JSON SerDe](https://github.com/rcongiu/Hive-JSON-Serde).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HiveJsonSerDe" : HiveJsonSerDe,
  "OpenXJsonSerDe" : OpenXJsonSerDe
}

```

### YAML

```yaml

  HiveJsonSerDe:
    HiveJsonSerDe
  OpenXJsonSerDe:
    OpenXJsonSerDe

```

## Properties

`HiveJsonSerDe`

The native Hive / HCatalog JsonSerDe. Used by Firehose for deserializing
data, which means converting it from the JSON format in preparation for serializing it to
the Parquet or ORC format. This is one of two deserializers you can choose, depending on
which one offers the functionality you need. The other option is the OpenX SerDe.

_Required_: No

_Type_: [HiveJsonSerDe](aws-properties-kinesisfirehose-deliverystream-hivejsonserde.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenXJsonSerDe`

The OpenX SerDe. Used by Firehose for deserializing data, which means
converting it from the JSON format in preparation for serializing it to the Parquet or ORC
format. This is one of two deserializers you can choose, depending on which one offers the
functionality you need. The other option is the native Hive / HCatalog JsonSerDe.

_Required_: No

_Type_: [OpenXJsonSerDe](aws-properties-kinesisfirehose-deliverystream-openxjsonserde.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeliveryStreamEncryptionConfigurationInput

DestinationTableConfiguration

All content copied from https://docs.aws.amazon.com/.
