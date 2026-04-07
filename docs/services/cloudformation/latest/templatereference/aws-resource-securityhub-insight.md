This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::Insight

The `AWS::SecurityHub::Insight` resource creates a custom insight in AWS Security Hub CSPM. An insight
is a collection of findings that relate to a security issue that requires attention or remediation. For more information,
see [Insights in AWS Security Hub CSPM](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-insights.html) in the _AWS Security Hub CSPM User Guide_.

Tags aren't supported for this resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::Insight",
  "Properties" : {
      "Filters" : AwsSecurityFindingFilters,
      "GroupByAttribute" : String,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::Insight
Properties:
  Filters:
    AwsSecurityFindingFilters
  GroupByAttribute: String
  Name: String

```

## Properties

`Filters`

One or more attributes used to filter the findings included in the insight. The insight
only includes findings that match the criteria defined in the filters. You can filter by up to ten finding attributes.
For each attribute, you can provide up to 20 filter values.

_Required_: Yes

_Type_: [AwsSecurityFindingFilters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-awssecurityfindingfilters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupByAttribute`

The grouping attribute for the insight's findings. Indicates how to group the matching
findings, and identifies the type of item that the insight applies to. For example, if an
insight is grouped by resource identifier, then the insight produces a list of resource
identifiers.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of a Security Hub CSPM insight.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of a Security Hub CSPM insight. For example, `arn:aws:securityhub:us-west-1:123456789012:insight/123456789012/custom/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`InsightArn`

The ARN of a Security Hub CSPM insight.

## Examples

### Creating a Security Hub CSPM insight

The following example creates a custom Security Hub CSPM insight. The insight includes findings that match the specified filters.

#### JSON

```json

{
    "Description": "Example template to create a Security Hub insight",
    "Resources": {
        "SecurityHubInsight": {
            "Type": "AWS::SecurityHub::Insight",
            "Properties": {
                "Name": "Example insight name",
                "GroupByAttribute": "ResourceId",
                "Filters": {
                   "CompanyName": [
                       {
                           "Comparison": "EQUALS",
                           "Value": "AWS"
                       }
                   ],
                   "CreatedAt": [
                       {
                           "DateRange": {
                               "Unit": "DAYS",
                               "Value": 5
                           }
                       }
                   ],
                   "Criticality": [
                       {
                           "Gte": 50,
                           "Lte": 95
                       }
                   ],
                   "Id": [
                       {
                           "Comparison": "EQUALS",
                           "Value": "example-id"
                       }
                   ],
                   "RecordState": [
                       {
                           "Comparison": "EQUALS",
                           "Value": "ACTIVE"
                       }
                   ],
                   "ResourceTags": [
                       {
                           "Comparison": "EQUALS",
                           "Key": "department",
                           "Value": "security"
                       },
                       {
                           "Comparison": "EQUALS",
                           "Key": "department",
                           "Value": "operations"
                       }
                   ],
                   "SeverityLabel": [
                       {
                           "Comparison": "EQUALS",
                           "Value": "LOW"
                       },
                       {
                           "Comparison": "EQUALS",
                           "Value": "HIGH"
                       }
                   ],
                   "UpdatedAt": [
                       {
                           "Start": "2023-04-25T17:05:54.832Z",
                           "End": "2023-05-25T17:05:54.832Z"
                       }
                   ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

Description: Example template to create a Security Hub insight
Resources:
  SecurityHubInsight:
    Type: "AWS::SecurityHub::Insight"
    Properties:
      Name: "Example insight name"
      GroupByAttribute: "ResourceId"
      Filters:
        CompanyName:
          - Comparison: EQUALS
            Value: AWS
        CreatedAt:
          - DateRange:
              Unit: DAYS
              Value: 5
        Criticality:
          - Gte: 50
            Lte: 95
        Id:
          - Comparison: EQUALS
            Value: example-id
        RecordState:
          - Comparison: EQUALS
            Value: ACTIVE
        ResourceTags:
          - Comparison: EQUALS
            Key: department
            Value: security
          - Comparison: EQUALS
            Key: department
            Value: operations
        SeverityLabel:
          - Comparison: EQUALS
            Value: LOW
          - Comparison: EQUALS
            Value: HIGH
        UpdatedAt:
          - Start: "2023-04-25T17:05:54.832Z"
            End: "2023-05-25T17:05:54.832Z"

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SecurityHub::HubV2

AwsSecurityFindingFilters
