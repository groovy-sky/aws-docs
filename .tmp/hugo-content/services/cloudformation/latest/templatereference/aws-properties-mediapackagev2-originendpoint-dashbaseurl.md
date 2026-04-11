This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint DashBaseUrl

The base URLs to use for retrieving segments. You can specify multiple locations and indicate the priority and weight for when each should be used, for use in mutli-CDN workflows.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DvbPriority" : Integer,
  "DvbWeight" : Integer,
  "ServiceLocation" : String,
  "Url" : String
}

```

### YAML

```yaml

  DvbPriority: Integer
  DvbWeight: Integer
  ServiceLocation: String
  Url: String

```

## Properties

`DvbPriority`

For use with DVB-DASH profiles only. The priority of this location for servings segments. The lower the number, the higher the priority.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `15000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DvbWeight`

For use with DVB-DASH profiles only. The weighting for source locations that have the same priority.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `15000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceLocation`

The name of the source location.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

A source location for segments.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaPackageV2::OriginEndpoint

DashDvbFontDownload

All content copied from https://docs.aws.amazon.com/.
