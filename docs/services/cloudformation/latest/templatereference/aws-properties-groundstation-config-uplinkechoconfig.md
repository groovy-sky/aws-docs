This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config UplinkEchoConfig

Provides information about how AWS Ground Station should echo back uplink transmissions to a dataflow endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AntennaUplinkConfigArn" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  AntennaUplinkConfigArn: String
  Enabled: Boolean

```

## Properties

`AntennaUplinkConfigArn`

Defines the ARN of the uplink config to echo back to a dataflow endpoint.

_Required_: No

_Type_: String

_Pattern_: `^(arn:(aws[a-zA-Z-]*)?:[a-z0-9-.]+:.*)|()$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Whether or not uplink echo is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create an UplinkEchoConfig

The following example creates a Ground Station `UplinkEchoConfig`

#### JSON

```json

{
  "UplinkEchoConfig": {
    "AntennaUplinkConfigArn": "arn:aws:groundstation:us-east-2:012345678910:config/antenna-uplink/11111111-1111-1111-1111-111111111111",
    "Enabled": true
  }
}
```

#### YAML

```yaml

UplinkEchoConfig:
  AntennaUplinkConfigArn: arn:aws:groundstation:us-east-2:012345678910:config/antenna-uplink/11111111-1111-1111-1111-111111111111
  Enabled: true
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrackingConfig

UplinkSpectrumConfig

All content copied from https://docs.aws.amazon.com/.
