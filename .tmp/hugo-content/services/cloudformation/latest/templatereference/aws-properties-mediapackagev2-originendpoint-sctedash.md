This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint ScteDash

The SCTE configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdMarkerDash" : String
}

```

### YAML

```yaml

  AdMarkerDash: String

```

## Properties

`AdMarkerDash`

Choose how ad markers are included in the packaged content. If you include ad markers in the content stream in your upstream encoders, then you need to inform MediaPackage what to do with the ad markers in the output.

Value description:

- `Binary` \- The SCTE-35 marker is expressed as a hex-string (Base64 string) rather than full XML.

- `XML` \- The SCTE marker is expressed fully in XML.

_Required_: No

_Type_: String

_Allowed values_: `BINARY | XML`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Scte

ScteHls

All content copied from https://docs.aws.amazon.com/.
