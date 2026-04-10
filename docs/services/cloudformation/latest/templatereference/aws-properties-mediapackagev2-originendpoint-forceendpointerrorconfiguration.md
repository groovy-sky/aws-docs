This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint ForceEndpointErrorConfiguration

The failover settings for the endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndpointErrorConditions" : [ String, ... ]
}

```

### YAML

```yaml

  EndpointErrorConditions:
    - String

```

## Properties

`EndpointErrorConditions`

The failover conditions for the endpoint. The options are:

- `STALE_MANIFEST` \- The manifest stalled and there are no new segments or parts.

- `INCOMPLETE_MANIFEST` \- There is a gap in the manifest.

- `MISSING_DRM_KEY` \- Key rotation is enabled but we're unable to fetch the key for the current key period.

- `SLATE_INPUT` \- The segments which contain slate content are considered to be missing content.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterConfiguration

HlsManifestConfiguration

All content copied from https://docs.aws.amazon.com/.
