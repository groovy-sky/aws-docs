This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition

A segment definition resource of Amazon Connect Customer Profiles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CustomerProfiles::SegmentDefinition",
  "Properties" : {
      "Description" : String,
      "DisplayName" : String,
      "DomainName" : String,
      "SegmentDefinitionName" : String,
      "SegmentGroups" : SegmentGroup,
      "SegmentSort" : SegmentSort,
      "SegmentSqlQuery" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CustomerProfiles::SegmentDefinition
Properties:
  Description: String
  DisplayName: String
  DomainName: String
  SegmentDefinitionName: String
  SegmentGroups:
    SegmentGroup
  SegmentSort:
    SegmentSort
  SegmentSqlQuery: String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the segment definition.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

Display name of the segment definition.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainName`

The name of the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SegmentDefinitionName`

Name of the segment definition.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SegmentGroups`

Contains all groups of the segment definition.

_Required_: No

_Type_: [SegmentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-segmentdefinition-segmentgroup.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SegmentSort`

Property description not available.

_Required_: No

_Type_: [SegmentSort](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-segmentdefinition-segmentsort.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentSqlQuery`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `50000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags belonging to the segment definition.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-segmentdefinition-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

When the segment definition was created.

`SegmentDefinitionArn`

The arn of the segment definition.

`SegmentType`

The segment type.

Classic : Segments created using traditional SegmentGroup structure

Enhanced : Segments created using SQL queries

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TrainingMetrics

AddressDimension
