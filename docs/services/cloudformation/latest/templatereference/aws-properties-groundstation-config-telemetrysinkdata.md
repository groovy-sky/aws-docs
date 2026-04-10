This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config TelemetrySinkData

Contains configuration data for a telemetry sink. The specific data structure depends on the sink type being configured.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KinesisDataStreamData" : KinesisDataStreamData
}

```

### YAML

```yaml

  KinesisDataStreamData:
    KinesisDataStreamData

```

## Properties

`KinesisDataStreamData`

Configuration data for delivering telemetry to a Kinesis Data Stream stream.

_Required_: No

_Type_: [KinesisDataStreamData](aws-properties-groundstation-config-kinesisdatastreamdata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create TelemetrySinkData for Kinesis

The following example creates `TelemetrySinkData` configured for Amazon Kinesis Data Stream delivery.

#### JSON

```json

{
  "TelemetrySinkData": {
    "KinesisDataStreamData": {
      "KinesisDataStreamArn": "arn:aws:kinesis:us-east-2:123456789012:stream/my-telemetry-stream",
      "KinesisRoleArn": "arn:aws:iam::123456789012:role/GroundStationKinesisRole"
    }
  }
}
```

#### YAML

```yaml

TelemetrySinkData:
  KinesisDataStreamData:
    KinesisDataStreamArn: arn:aws:kinesis:us-east-2:123456789012:stream/my-telemetry-stream
    KinesisRoleArn: arn:aws:iam::123456789012:role/GroundStationKinesisRole

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TelemetrySinkConfig

TrackingConfig

All content copied from https://docs.aws.amazon.com/.
