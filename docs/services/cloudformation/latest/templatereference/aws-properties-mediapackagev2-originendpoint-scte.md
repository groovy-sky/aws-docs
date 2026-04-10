This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint Scte

The SCTE-35 configuration associated with the origin endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ScteFilter" : [ String, ... ],
  "ScteInSegments" : String
}

```

### YAML

```yaml

  ScteFilter:
    - String
  ScteInSegments: String

```

## Properties

`ScteFilter`

The filter associated with the SCTE-35 configuration.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScteInSegments`

Controls whether SCTE-35 messages are included in segment files.

- None – SCTE-35 messages are not included in segments (default)

- All – SCTE-35 messages are embedded in segment data

For DASH manifests, when set to `All`, an `InbandEventStream` tag signals that SCTE messages are present in segments. This setting works independently of manifest ad markers.

_Required_: No

_Type_: String

_Allowed values_: `NONE | ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MssManifestConfiguration

ScteDash

All content copied from https://docs.aws.amazon.com/.
