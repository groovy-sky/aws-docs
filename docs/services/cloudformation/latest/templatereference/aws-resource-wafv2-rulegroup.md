This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup

###### Note

This is the latest version of **AWS WAF**, named AWS WAFV2, released in November, 2019. For
information, including how to migrate your AWS WAF resources from the
prior release, see the [AWS WAF developer guide](../../../waf/latest/developerguide/waf-chapter.md).

Use an AWS::WAFv2::RuleGroup to define a collection of rules for inspecting and
controlling web requests. You use a rule group in an [AWS::WAFv2::WebACL](aws-resource-wafv2-webacl.md) by providing
its Amazon Resource Name (ARN) to the rule statement
`RuleGroupReferenceStatement`, when you add rules to the web ACL.

When you create a rule group, you define an immutable capacity limit. If you update a
rule group, you must stay within the capacity. This allows others to reuse the rule group
with confidence in its capacity requirements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFv2::RuleGroup",
  "Properties" : {
      "AvailableLabels" : [ LabelSummary, ... ],
      "Capacity" : Integer,
      "ConsumedLabels" : [ LabelSummary, ... ],
      "CustomResponseBodies" : {Key: Value, ...},
      "Description" : String,
      "Name" : String,
      "Rules" : [ Rule, ... ],
      "Scope" : String,
      "Tags" : [ Tag, ... ],
      "VisibilityConfig" : VisibilityConfig
    }
}

```

### YAML

```yaml

Type: AWS::WAFv2::RuleGroup
Properties:
  AvailableLabels:
    - LabelSummary
  Capacity: Integer
  ConsumedLabels:
    - LabelSummary
  CustomResponseBodies:
    Key: Value
  Description: String
  Name: String
  Rules:
    - Rule
  Scope: String
  Tags:
    - Tag
  VisibilityConfig:
    VisibilityConfig

```

## Properties

`AvailableLabels`

The labels that one or more rules in this rule group add to matching web requests. These
labels are defined in the `RuleLabels` for a `Rule`.

_Required_: No

_Type_: Array of [LabelSummary](aws-properties-wafv2-rulegroup-labelsummary.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Capacity`

The web ACL capacity units (WCUs) required for this rule group.

When you create your own rule group, you define this, and you cannot change it after
creation. When you add or modify the rules in a rule group, AWS WAF enforces
this limit.

AWS WAF uses WCUs to calculate and control the operating resources that are
used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity
differently for each rule type, to reflect the relative cost of each rule. Simple rules
that cost little to run use fewer WCUs than more complex rules that use more processing
power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU
usage when they use a rule group. The WCU limit for web ACLs is 1,500.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConsumedLabels`

The labels that one or more rules in this rule group match against in label match statements. These labels are defined in a `LabelMatchStatement` specification, in the [Statement](../userguide/aws-properties-wafv2-webacl-notstatement.md#cfn-wafv2-webacl-notstatement-statement) definition of a rule.

_Required_: No

_Type_: Array of [LabelSummary](aws-properties-wafv2-rulegroup-labelsummary.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomResponseBodies`

A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the rule group, and then use them in the rules that you define in the rule group.

For information about customizing web requests and responses,
see [Customizing web requests and responses in AWS WAF](../../../waf/latest/developerguide/waf-custom-request-response.md)
in the _AWS WAF Developer Guide_.

For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](../../../waf/latest/developerguide/limits.md)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: Object of [CustomResponseBody](aws-properties-wafv2-rulegroup-customresponsebody.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the rule group that helps with identification.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9=:#@/\-,.][a-zA-Z0-9+=:#@/\-,.\s]+[a-zA-Z0-9+=:#@/\-,.]{1,256}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the rule group. You cannot change the name of a rule group after you create it.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]{1,128}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rules`

The rule statements used to identify the web requests that you want to allow, block, or
count. Each rule includes one top-level statement that AWS WAF uses to
identify matching web requests, and parameters that govern how AWS WAF
handles them.

_Required_: No

_Type_: Array of [Rule](aws-properties-wafv2-rulegroup-rule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

Specifies whether this is for an Amazon CloudFront distribution or for a regional
application. For an AWS Amplify application, use `CLOUDFRONT`.
A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway
REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool,
an AWS App Runner service, or an AWS Verified Access instance. Valid Values are
`CLOUDFRONT` and `REGIONAL`.

###### Note

For `CLOUDFRONT`, you must create your WAFv2 resources in the US East (N.
Virginia) Region, `us-east-1`.

_Required_: Yes

_Type_: String

_Allowed values_: `CLOUDFRONT | REGIONAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key:value pairs associated with an AWS resource. The key:value pair can
be anything you define. Typically, the tag key represents a category (such as
"environment") and the tag value represents a specific value within that category (such as
"test," "development," or "production"). You can add up to 50 tags to each AWS resource.

###### Note

To modify tags on existing resources, use the AWS WAF APIs or
command line interface. With AWS CloudFormation, you can only add tags to AWS WAF resources during resource creation.

_Required_: No

_Type_: Array of [Tag](aws-properties-wafv2-rulegroup-tag.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisibilityConfig`

Defines and enables Amazon CloudWatch metrics and web request sample collection.

_Required_: Yes

_Type_: [VisibilityConfig](aws-properties-wafv2-rulegroup-visibilityconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The `Ref` for the resource, containing the resource name, physical ID, and
scope, formatted as follows: `name|id|scope`.

For example:
`my-webacl-name|1234a1a-a1b1-12a1-abcd-a123b123456|REGIONAL`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the rule group.

`Id`

The ID of the rule group.

`LabelNamespace`

The label namespace prefix for this rule group. All labels added by rules in this rule
group have this prefix.

The syntax for the label namespace prefix for a rule group is the following:
`awswaf:<account ID>:rule group:<rule group name>:`

When a rule with a label matches a web request, AWS WAF adds the fully
qualified label to the request. A fully qualified label is made up of the label namespace
from the rule group or web ACL where the rule is defined and the label from the rule,
separated by a colon.

## Examples

### Create a rule group

The following shows an example rule group specification.

#### YAML

```yaml

 SampleRuleGroup:
      Type: 'AWS::WAFv2::RuleGroup'
      Properties:
        Name: SampleRuleGroup
        Scope: REGIONAL
        Description: SampleRuleGroup
        VisibilityConfig:
          SampledRequestsEnabled: true
          CloudWatchMetricsEnabled: true
          MetricName: SampleRuleGroupMetrics
        CustomResponseBodies:
          CustomResponseBodyKey1:
            ContentType: TEXT_PLAIN
            Content: this is a plain text
          CustomResponseBodyKey2:
            ContentType: APPLICATION_JSON
            Content: '{"jsonfieldname": "jsonfieldvalue"}'
          CustomResponseBodyKey3:
            ContentType: TEXT_HTML
            Content: <html>HTML text content</html>
        Capacity: 1000
        Rules:
          - Name: RuleOne
            Priority: 1
            Action:
              Allow:
                CustomRequestHandling:
                  InsertHeaders:
                    - Name: AllowActionHeader1Name
                      Value: AllowActionHeader1Value
                    - Name: AllowActionHeader2Name
                      Value: AllowActionHeader2Value
            VisibilityConfig:
              SampledRequestsEnabled: true
              CloudWatchMetricsEnabled: true
              MetricName: RuleOneMetric
            Statement:
              ByteMatchStatement:
                FieldToMatch:
                  AllQueryArguments: {}
                PositionalConstraint: CONTAINS
                SearchString: testagent
                TextTransformations:
                  - Priority: 1
                    Type: HTML_ENTITY_DECODE
          - Name: RuleTwo
            Priority: 2
            Action:
              Block:
                CustomResponse:
                  ResponseCode: 503
                  CustomResponseBodyKey: CustomResponseBodyKey1
                  ResponseHeaders:
                    - Name: BlockActionHeader1Name
                      Value: BlockActionHeader1Value
                    - Name: BlockActionHeader2Name
                      Value: BlockActionHeader2Value
            VisibilityConfig:
              SampledRequestsEnabled: true
              CloudWatchMetricsEnabled: true
              MetricName: RuleTwoMetric
            Statement:
              ByteMatchStatement:
                FieldToMatch:
                  SingleHeader:
                    Name: haystack
                PositionalConstraint: CONTAINS
                SearchString: badbot
                TextTransformations:
                  - Priority: 0
                    Type: NONE
          - Name: RuleThree
            Priority: 3
            Action:
              Count:
                CustomRequestHandling:
                  InsertHeaders:
                    - Name: CountActionHeader1Name
                      Value: CountActionHeader1Value
                    - Name: CountActionHeader2Name
                      Value: CountActionHeader2Value
            VisibilityConfig:
              SampledRequestsEnabled: true
              CloudWatchMetricsEnabled: true
              MetricName: RuleThreeMetric
            Statement:
              ByteMatchStatement:
                FieldToMatch:
                  Body: {}
                PositionalConstraint: CONTAINS
                SearchString: RegionOne
                TextTransformations:
                  - Priority: 0
                    Type: HTML_ENTITY_DECODE
          - Name: RuleFour
            Priority: 4
            Action:
              Allow: {}
            VisibilityConfig:
              SampledRequestsEnabled: true
              CloudWatchMetricsEnabled: true
              MetricName: RuleFourMetric
            Statement:
              SizeConstraintStatement:
                ComparisonOperator: GT
                Size: 1000
                FieldToMatch:
                  UriPath: {}
                TextTransformations:
                  - Priority: 0
                    Type: NONE
```

#### JSON

```json

    "SampleRuleGroup": {
        "Type": "AWS::WAFv2::RuleGroup",
        "Properties": {
            "Name": "SampleRuleGroup",
            "Scope": "REGIONAL",
            "Description": "SampleRuleGroup",
            "VisibilityConfig": {
                "SampledRequestsEnabled": true,
                "CloudWatchMetricsEnabled": true,
                "MetricName": "SampleRuleGroupMetrics"
            },
            "CustomResponseBodies": {
                "CustomResponseBodyKey1": {
                    "ContentType": "TEXT_PLAIN",
                    "Content": "this is a plain text"
                },
                "CustomResponseBodyKey2": {
                    "ContentType": "APPLICATION_JSON",
                    "Content": "{\"jsonfieldname\": \"jsonfieldvalue\"}"
                },
                "CustomResponseBodyKey3": {
                    "ContentType": "TEXT_HTML",
                    "Content": "<html>HTML text content</html>"
                }
            },
            "Capacity": 1000,
            "Rules": [
                {
                    "Name": "RuleOne",
                    "Priority": 1,
                    "Action": {
                        "Allow": {
                            "CustomRequestHandling": {
                                "InsertHeaders": [
                                    {
                                        "Name": "AllowActionHeader1Name",
                                        "Value": "AllowActionHeader1Value"
                                    },
                                    {
                                        "Name": "AllowActionHeader2Name",
                                        "Value": "AllowActionHeader2Value"
                                    }
                                ]
                            }
                        }
                    },
                    "VisibilityConfig": {
                        "SampledRequestsEnabled": true,
                        "CloudWatchMetricsEnabled": true,
                        "MetricName": "RuleOneMetric"
                    },
                    "Statement": {
                        "ByteMatchStatement": {
                            "FieldToMatch": {
                                "AllQueryArguments": {}
                            },
                            "PositionalConstraint": "CONTAINS",
                            "SearchString": "testagent",
                            "TextTransformations": [
                                {
                                    "Priority": 1,
                                    "Type": "HTML_ENTITY_DECODE"
                                }
                            ]
                        }
                    }
                },
                {
                    "Name": "RuleTwo",
                    "Priority": 2,
                    "Action": {
                        "Block": {
                            "CustomResponse": {
                                "ResponseCode": 503,
                                "CustomResponseBodyKey": "CustomResponseBodyKey1",
                                "ResponseHeaders": [
                                    {
                                        "Name": "BlockActionHeader1Name",
                                        "Value": "BlockActionHeader1Value"
                                    },
                                    {
                                        "Name": "BlockActionHeader2Name",
                                        "Value": "BlockActionHeader2Value"
                                    }
                                ]
                            }
                        }
                    },
                    "VisibilityConfig": {
                        "SampledRequestsEnabled": true,
                        "CloudWatchMetricsEnabled": true,
                        "MetricName": "RuleTwoMetric"
                    },
                    "Statement": {
                        "ByteMatchStatement": {
                            "FieldToMatch": {
                                "SingleHeader": {
                                    "Name": "haystack"
                                }
                            },
                            "PositionalConstraint": "CONTAINS",
                            "SearchString": "badbot",
                            "TextTransformations": [
                                {
                                    "Priority": 0,
                                    "Type": "NONE"
                                }
                            ]
                        }
                    }
                },
                {
                    "Name": "RuleThree",
                    "Priority": 3,
                    "Action": {
                        "Count": {
                            "CustomRequestHandling": {
                                "InsertHeaders": [
                                    {
                                        "Name": "CountActionHeader1Name",
                                        "Value": "CountActionHeader1Value"
                                    },
                                    {
                                        "Name": "CountActionHeader2Name",
                                        "Value": "CountActionHeader2Value"
                                    }
                                ]
                            }
                        }
                    },
                    "VisibilityConfig": {
                        "SampledRequestsEnabled": true,
                        "CloudWatchMetricsEnabled": true,
                        "MetricName": "RuleThreeMetric"
                    },
                    "Statement": {
                        "ByteMatchStatement": {
                            "FieldToMatch": {
                                "Body": {}
                            },
                            "PositionalConstraint": "CONTAINS",
                            "SearchString": "RegionOne",
                            "TextTransformations": [
                                {
                                    "Priority": 0,
                                    "Type": "HTML_ENTITY_DECODE"
                                }
                            ]
                        }
                    }
                },
                {
                    "Name": "RuleFour",
                    "Priority": 4,
                    "Action": {
                        "Allow": {}
                    },
                    "VisibilityConfig": {
                        "SampledRequestsEnabled": true,
                        "CloudWatchMetricsEnabled": true,
                        "MetricName": "RuleFourMetric"
                    },
                    "Statement": {
                        "SizeConstraintStatement": {
                            "ComparisonOperator": "GT",
                            "Size": 1000,
                            "FieldToMatch": {
                                "UriPath": {}
                            },
                            "TextTransformations": [
                                {
                                    "Priority": 0,
                                    "Type": "NONE"
                                }
                            ]
                        }
                    }
                }
            ]
        }
    }
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AllowAction

All content copied from https://docs.aws.amazon.com/.
