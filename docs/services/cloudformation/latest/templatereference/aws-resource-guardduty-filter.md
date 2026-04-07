This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Filter

The `AWS::GuardDuty::Filter` resource specifies a new filter defined by the
provided `findingCriteria`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GuardDuty::Filter",
  "Properties" : {
      "Action" : String,
      "Description" : String,
      "DetectorId" : String,
      "FindingCriteria" : FindingCriteria,
      "Name" : String,
      "Rank" : Integer,
      "Tags" : [ TagItem, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GuardDuty::Filter
Properties:
  Action: String
  Description: String
  DetectorId: String
  FindingCriteria:
    FindingCriteria
  Name: String
  Rank: Integer
  Tags:
    - TagItem

```

## Properties

`Action`

Specifies the action that is to be applied to the findings that match the filter.

_Required_: No

_Type_: String

_Allowed values_: `NOOP | ARCHIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the filter. Valid characters include alphanumeric characters, and
special characters such as hyphen, period, colon, underscore, parentheses ( `{ }`,
`[ ]`, and `( )`), forward slash, horizontal tab, vertical tab,
newline, form feed, return, and whitespace.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetectorId`

The detector ID associated with the GuardDuty account for which you want to create a filter.

To find the `detectorId` in the current Region, see the
Settings page in the GuardDuty console, or run the [ListDetectors](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_ListDetectors.html) API.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FindingCriteria`

Represents the criteria to be used in the filter for querying findings.

_Required_: Yes

_Type_: [FindingCriteria](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-filter-findingcriteria.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the filter. Valid characters include period (.), underscore (\_), dash (-), and
alphanumeric characters. A whitespace is considered to be an invalid character.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rank`

Specifies the position of the filter in the list of current filters. Also specifies the
order in which this filter is applied to the findings. The minimum value for this property
is 1 and the maximum is 100.

By default, filters may not be created in the same order as they are ranked. To ensure
that the filters are created in the expected order, you can use an optional attribute,
[DependsOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html),
with the following syntax: `"DependsOn":[ "ObjectName" ]`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be added to a new filter resource. Each tag consists of a key and an
optional value, both of which you define.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [TagItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-filter-tagitem.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the filter, such as
`SampleFilter`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Declare a Filter Resource

The following example shows how to declare a GuardDuty `Filter` resource:

#### JSON

```json

{
    "Type": "AWS::GuardDuty::Filter",
    "Properties": {
        "Action": "ARCHIVE",
        "Description": "SampleFilter",
        "DetectorId": "a12abc34d567e8fa901bc2d34e56789f0",
        "FindingCriteria": {
            "Criterion": {
                "updatedAt": {
                "Gte": 0
                },
                "severity": {
                "Gte": 0
        }
    },
    "Rank": 1,
    "Name": "SampleFilter"
    }
}
```

#### YAML

```yaml

Type: "AWS::GuardDuty::Filter"
Properties:
    Action : "ARCHIVE"
    Description : "SampleFilter"
    DetectorId : "a12abc34d567e8fa901bc2d34e56789f0"
    FindingCriteria :
        Criterion:
            "updatedAt":
                Gte: 0
            "severity":
                Gte: 0
    Rank : 1
    Name : "SampleFilter"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagItem

Condition
