This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream Serializer

The serializer that you want Firehose to use to convert data to the target
format before writing it to Amazon S3. Firehose supports two types of
serializers: the ORC SerDe and the Parquet SerDe.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OrcSerDe" : OrcSerDe,
  "ParquetSerDe" : ParquetSerDe
}

```

### YAML

```yaml

  OrcSerDe:
    OrcSerDe
  ParquetSerDe:
    ParquetSerDe

```

## Properties

`OrcSerDe`

A serializer to use for converting data to the ORC format before storing it in Amazon
S3. For more information, see [Apache\
ORC](https://orc.apache.org/docs).

_Required_: No

_Type_: [OrcSerDe](aws-properties-kinesisfirehose-deliverystream-orcserde.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParquetSerDe`

A serializer to use for converting data to the Parquet format before storing it in
Amazon S3. For more information, see [Apache Parquet](https://parquet.apache.org/docs/contribution-guidelines).

_Required_: No

_Type_: [ParquetSerDe](aws-properties-kinesisfirehose-deliverystream-parquetserde.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecretsManagerConfiguration

SnowflakeBufferingHints

All content copied from https://docs.aws.amazon.com/.
