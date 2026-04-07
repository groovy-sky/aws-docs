This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::VodSource

The VOD source configuration parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaTailor::VodSource",
  "Properties" : {
      "HttpPackageConfigurations" : [ HttpPackageConfiguration, ... ],
      "SourceLocationName" : String,
      "Tags" : [ Tag, ... ],
      "VodSourceName" : String
    }
}

```

### YAML

```yaml

Type: AWS::MediaTailor::VodSource
Properties:
  HttpPackageConfigurations:
    - HttpPackageConfiguration
  SourceLocationName: String
  Tags:
    - Tag
  VodSourceName: String

```

## Properties

`HttpPackageConfigurations`

The HTTP package configurations for the VOD source.

_Required_: Yes

_Type_: Array of [HttpPackageConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-vodsource-httppackageconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceLocationName`

The name of the source location that the VOD source is associated with.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the VOD source. Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-vodsource-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VodSourceName`

The name of the VOD source.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

HttpPackageConfiguration
