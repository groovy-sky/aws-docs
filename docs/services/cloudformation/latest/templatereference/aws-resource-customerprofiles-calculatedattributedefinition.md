This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::CalculatedAttributeDefinition

A calculated attribute definition for Customer Profiles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CustomerProfiles::CalculatedAttributeDefinition",
  "Properties" : {
      "AttributeDetails" : AttributeDetails,
      "CalculatedAttributeName" : String,
      "Conditions" : Conditions,
      "Description" : String,
      "DisplayName" : String,
      "DomainName" : String,
      "Statistic" : String,
      "Tags" : [ Tag, ... ],
      "UseHistoricalData" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::CustomerProfiles::CalculatedAttributeDefinition
Properties:
  AttributeDetails:
    AttributeDetails
  CalculatedAttributeName: String
  Conditions:
    Conditions
  Description: String
  DisplayName: String
  DomainName: String
  Statistic: String
  Tags:
    - Tag
  UseHistoricalData: Boolean

```

## Properties

`AttributeDetails`

Mathematical expression and a list of attribute items specified in that
expression.

_Required_: Yes

_Type_: [AttributeDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-calculatedattributedefinition-attributedetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalculatedAttributeName`

The name of an attribute defined in a profile object type.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_][a-zA-Z_0-9-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Conditions`

The conditions including range, object count, and threshold for the calculated
attribute.

_Required_: No

_Type_: [Conditions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-calculatedattributedefinition-conditions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the calculated attribute.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The display name of the calculated attribute.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z_][a-zA-Z_0-9-\s]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The unique name of the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Statistic`

The aggregation operation to perform for the calculated attribute.

_Required_: Yes

_Type_: String

_Allowed values_: `FIRST_OCCURRENCE | LAST_OCCURRENCE | COUNT | SUM | MINIMUM | MAXIMUM | AVERAGE | MAX_OCCURRENCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-calculatedattributedefinition-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseHistoricalData`

Whether historical data ingested before the Calculated Attribute was created should be
included in calculations.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The timestamp of when the calculated attribute definition was created.

`LastUpdatedAt`

The timestamp of when the calculated attribute definition was most recently
edited.

`Status`

Status of the Calculated Attribute creation (whether all historical data has been
indexed.)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Connect Customer Profiles

AttributeDetails
