This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration ManifestProcessingRules

The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdMarkerPassthrough" : AdMarkerPassthrough
}

```

### YAML

```yaml

  AdMarkerPassthrough:
    AdMarkerPassthrough

```

## Properties

`AdMarkerPassthrough`

For HLS, when set to `true`, MediaTailor passes through `EXT-X-CUE-IN`, `EXT-X-CUE-OUT`, and `EXT-X-SPLICEPOINT-SCTE35` ad markers from the origin manifest to the MediaTailor personalized manifest.

No logic is applied to these ad markers. For example, if `EXT-X-CUE-OUT` has a value of `60`, but no ads are filled for that ad break, MediaTailor will not set the value to `0`.

_Required_: No

_Type_: [AdMarkerPassthrough](aws-properties-mediatailor-playbackconfiguration-admarkerpassthrough.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogConfiguration

ManifestServiceInteractionLog

All content copied from https://docs.aws.amazon.com/.
