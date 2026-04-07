This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration CdnConfiguration

The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdSegmentUrlPrefix" : String,
  "ContentSegmentUrlPrefix" : String
}

```

### YAML

```yaml

  AdSegmentUrlPrefix: String
  ContentSegmentUrlPrefix: String

```

## Properties

`AdSegmentUrlPrefix`

A non-default content delivery network (CDN) to serve ad segments. By default, AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN for the origin ads.mediatailor. _<region>_.amazonaws.com. Then specify the rule's name in this `AdSegmentUrlPrefix`. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for ad segments.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentSegmentUrlPrefix`

A content delivery network (CDN) to cache content segments, so that content requests don’t always have to go to the origin server. First, create a rule in your CDN for the content segment origin server. Then specify the rule's name in this `ContentSegmentUrlPrefix`. When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for content segments.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Bumper

DashConfiguration
