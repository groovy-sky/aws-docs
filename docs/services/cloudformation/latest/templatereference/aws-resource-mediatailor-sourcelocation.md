This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::SourceLocation

A source location is a container for sources. For more information about source locations, see [Working with source locations](../../../mediatailor/latest/ug/channel-assembly-source-locations.md) in the _MediaTailor User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaTailor::SourceLocation",
  "Properties" : {
      "AccessConfiguration" : AccessConfiguration,
      "DefaultSegmentDeliveryConfiguration" : DefaultSegmentDeliveryConfiguration,
      "HttpConfiguration" : HttpConfiguration,
      "SegmentDeliveryConfigurations" : [ SegmentDeliveryConfiguration, ... ],
      "SourceLocationName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaTailor::SourceLocation
Properties:
  AccessConfiguration:
    AccessConfiguration
  DefaultSegmentDeliveryConfiguration:
    DefaultSegmentDeliveryConfiguration
  HttpConfiguration:
    HttpConfiguration
  SegmentDeliveryConfigurations:
    - SegmentDeliveryConfiguration
  SourceLocationName: String
  Tags:
    - Tag

```

## Properties

`AccessConfiguration`

The access configuration for the source location.

_Required_: No

_Type_: [AccessConfiguration](aws-properties-mediatailor-sourcelocation-accessconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultSegmentDeliveryConfiguration`

The default segment delivery configuration.

_Required_: No

_Type_: [DefaultSegmentDeliveryConfiguration](aws-properties-mediatailor-sourcelocation-defaultsegmentdeliveryconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpConfiguration`

The HTTP configuration for the source location.

_Required_: Yes

_Type_: [HttpConfiguration](aws-properties-mediatailor-sourcelocation-httpconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentDeliveryConfigurations`

The segment delivery configurations for the source location.

_Required_: No

_Type_: Array of [SegmentDeliveryConfiguration](aws-properties-mediatailor-sourcelocation-segmentdeliveryconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceLocationName`

The name of the source location.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the source location. Tags are key-value pairs that you can associate with Amazon resources to help with organization, access control, and cost tracking. For more information, see [Tagging AWS Elemental MediaTailor Resources](../../../mediatailor/latest/ug/tagging.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-mediatailor-sourcelocation-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AccessConfiguration

All content copied from https://docs.aws.amazon.com/.
