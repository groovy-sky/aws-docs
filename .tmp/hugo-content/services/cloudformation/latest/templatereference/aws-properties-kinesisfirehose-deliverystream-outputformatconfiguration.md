This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream OutputFormatConfiguration

Specifies the serializer that you want Firehose to use to convert the
format of your data before it writes it to Amazon S3. This parameter is required if
`Enabled` is set to true.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Serializer" : Serializer
}

```

### YAML

```yaml

  Serializer:
    Serializer

```

## Properties

`Serializer`

Specifies which serializer to use. You can choose either the ORC SerDe or the Parquet
SerDe. If both are non-null, the server rejects the request.

_Required_: No

_Type_: [Serializer](aws-properties-kinesisfirehose-deliverystream-serializer.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OrcSerDe

ParquetSerDe

All content copied from https://docs.aws.amazon.com/.
