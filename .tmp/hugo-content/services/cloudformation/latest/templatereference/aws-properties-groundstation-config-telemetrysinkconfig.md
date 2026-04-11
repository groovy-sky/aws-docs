This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config TelemetrySinkConfig

Provides information about how AWS Ground Station should deliver telemetry during contacts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TelemetrySinkData" : TelemetrySinkData,
  "TelemetrySinkType" : String
}

```

### YAML

```yaml

  TelemetrySinkData:
    TelemetrySinkData
  TelemetrySinkType: String

```

## Properties

`TelemetrySinkData`

Configuration data for the telemetry sink. The structure of this data depends on the telemetry sink type specified.

_Required_: Yes

_Type_: [TelemetrySinkData](aws-properties-groundstation-config-telemetrysinkdata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TelemetrySinkType`

Type of telemetry sink where telemetry is to be delivered. Currently, only `KINESIS_DATA_STREAM` is supported.

_Required_: Yes

_Type_: String

_Allowed values_: `KINESIS_DATA_STREAM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a TelemetrySinkConfig

The following example creates a Ground Station `TelemetrySinkConfig` that delivers telemetry to an Amazon Kinesis Data Stream.

#### JSON

```json

{
  "TelemetrySinkConfig": {
    "TelemetrySinkType": "KINESIS_DATA_STREAM",
    "TelemetrySinkData": {
      "KinesisDataStreamData": {
        "KinesisDataStreamArn": "arn:aws:kinesis:us-west-2:123456789012:stream/my-telemetry-stream",
        "KinesisRoleArn": "arn:aws:iam::123456789012:role/GroundStationKinesisRole"
      }
    }
  }
}
```

#### YAML

```yaml

TelemetrySinkConfig:
  TelemetrySinkType: KINESIS_DATA_STREAM
  TelemetrySinkData:
    KinesisDataStreamData:
      KinesisDataStreamArn: arn:aws:kinesis:us-west-2:123456789012:stream/my-telemetry-stream
      KinesisRoleArn: arn:aws:iam::123456789012:role/GroundStationKinesisRole

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TelemetrySinkData

All content copied from https://docs.aws.amazon.com/.
