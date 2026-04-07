This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::OriginEndpoint MssManifestConfiguration

The `MssManifestConfiguration` property type specifies Property description not available. for an [AWS::MediaPackageV2::OriginEndpoint](aws-resource-mediapackagev2-originendpoint.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilterConfiguration" : FilterConfiguration,
  "ManifestLayout" : String,
  "ManifestName" : String,
  "ManifestWindowSeconds" : Integer
}

```

### YAML

```yaml

  FilterConfiguration:
    FilterConfiguration
  ManifestLayout: String
  ManifestName: String
  ManifestWindowSeconds: Integer

```

## Properties

`FilterConfiguration`

Property description not available.

_Required_: No

_Type_: [FilterConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediapackagev2-originendpoint-filterconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestLayout`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `FULL | COMPACT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestName`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestWindowSeconds`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LowLatencyHlsManifestConfiguration

Scte
