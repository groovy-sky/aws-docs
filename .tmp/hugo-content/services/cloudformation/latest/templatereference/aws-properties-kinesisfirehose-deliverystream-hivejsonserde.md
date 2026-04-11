This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream HiveJsonSerDe

The native Hive / HCatalog JsonSerDe. Used by Firehose for deserializing
data, which means converting it from the JSON format in preparation for serializing it to
the Parquet or ORC format. This is one of two deserializers you can choose, depending on
which one offers the functionality you need. The other option is the OpenX SerDe.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimestampFormats" : [ String, ... ]
}

```

### YAML

```yaml

  TimestampFormats:
    - String

```

## Properties

`TimestampFormats`

Indicates how you want Firehose to parse the date and timestamps that
may be present in your input data JSON. To specify these format strings, follow the pattern
syntax of JodaTime's DateTimeFormat format strings. For more information, see [Class DateTimeFormat](https://www.joda.org/joda-time/apidocs/org/joda/time/format/DateTimeFormat.html). You can also use the special value `millis` to
parse timestamps in epoch milliseconds. If you don't specify a format, Firehose uses `java.sql.Timestamp::valueOf` by default.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExtendedS3DestinationConfiguration

HttpEndpointCommonAttribute

All content copied from https://docs.aws.amazon.com/.
