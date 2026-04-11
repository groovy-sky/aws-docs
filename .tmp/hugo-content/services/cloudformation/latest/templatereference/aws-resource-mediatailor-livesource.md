This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::LiveSource

Live source configuration parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaTailor::LiveSource",
  "Properties" : {
      "HttpPackageConfigurations" : [ HttpPackageConfiguration, ... ],
      "LiveSourceName" : String,
      "SourceLocationName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaTailor::LiveSource
Properties:
  HttpPackageConfigurations:
    - HttpPackageConfiguration
  LiveSourceName: String
  SourceLocationName: String
  Tags:
    - Tag

```

## Properties

`HttpPackageConfigurations`

The HTTP package configurations for the live source.

_Required_: Yes

_Type_: Array of [HttpPackageConfiguration](aws-properties-mediatailor-livesource-httppackageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LiveSourceName`

The name that's used to refer to a live source.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceLocationName`

The name of the source location.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the live source. Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](../../../mediatailor/latest/ug/tagging.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-mediatailor-livesource-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaTailor::ChannelPolicy

HttpPackageConfiguration

All content copied from https://docs.aws.amazon.com/.
