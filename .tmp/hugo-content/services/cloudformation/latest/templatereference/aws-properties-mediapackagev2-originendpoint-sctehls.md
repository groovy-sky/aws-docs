This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint ScteHls

The SCTE-35 HLS configuration associated with the origin endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdMarkerHls" : String
}

```

### YAML

```yaml

  AdMarkerHls: String

```

## Properties

`AdMarkerHls`

The SCTE-35 HLS ad-marker configuration.

_Required_: No

_Type_: String

_Allowed values_: `DATERANGE | SCTE35_ENHANCED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScteDash

Segment

All content copied from https://docs.aws.amazon.com/.
