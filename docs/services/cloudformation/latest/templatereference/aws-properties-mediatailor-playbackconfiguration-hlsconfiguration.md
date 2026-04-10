This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration HlsConfiguration

The configuration for HLS content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ManifestEndpointPrefix" : String
}

```

### YAML

```yaml

  ManifestEndpointPrefix: String

```

## Properties

`ManifestEndpointPrefix`

The URL that is used to initiate a playback session for devices that support Apple HLS. The session uses server-side reporting.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DashConfiguration

HttpRequest

All content copied from https://docs.aws.amazon.com/.
