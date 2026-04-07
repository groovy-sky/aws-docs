This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL

###### Note

This is the latest version of **AWS WAF**, named AWS WAFV2, released in November, 2019. For
information, including how to migrate your AWS WAF resources from the
prior release, see the [AWS WAF developer guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).

Use an AWS::WAFv2::WebACL to define a collection of rules to use to inspect and
control web requests. Each rule in a web ACL has a statement that defines what to look for
in web requests and an action that AWS WAF applies to requests that match
the statement. In the web ACL, you assign a default action to take (allow, block) for any
request that doesn't match any of the rules.

The rules in a web ACL can be a combination of explicitly defined rules and rule groups
that you reference from the web ACL. The rule groups can be rule groups that you manage or
rule groups that are managed by others.

You can associate a web ACL with one or more AWS resources to protect.
The resources can be an Amazon CloudFront distribution, an Amazon API Gateway
REST API, an Application Load Balancer, an AWS AppSync GraphQL API , an Amazon Cognito user pool, an AWS App Runner service, an AWS Amplify application, or an AWS Verified Access instance.

For more information, see [Web access control lists (web ACLs)](https://docs.aws.amazon.com/waf/latest/developerguide/web-acl.html)
in the _AWS WAF developer guide_.

**Web ACLs used in AWS Shield Advanced automatic application**
**layer DDoS mitigation**

If you use Shield Advanced automatic application layer DDoS mitigation, the web
ACLs that you use with automatic mitigation have a rule group rule whose name starts with
`ShieldMitigationRuleGroup`. This rule is used for automatic mitigations and
it's managed for you in the web ACL by Shield Advanced and AWS WAF. You'll see the rule listed among the web ACL rules when you view the web ACL through
the AWS WAF interfaces.

When you manage the web ACL through CloudFormation interfaces, you won't see the
Shield Advanced rule. CloudFormation doesn't include this type of rule in
the stack drift status between the actual configuration of the web ACL and your web ACL
template.

Don't add the Shield Advanced rule group rule to your web ACL template. The rule
shouldn't be in your template. When you update the web ACL template in a stack, the Shield Advanced rule is maintained for you by AWS WAF in the resulting
web ACL.

For more information, see [Shield Advanced\
automatic application layer DDoS mitigation](https://docs.aws.amazon.com/waf/latest/developerguide/ddos-automatic-app-layer-response.html) in the _AWS Shield Advanced developer guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFv2::WebACL",
  "Properties" : {
      "ApplicationConfig" : ApplicationConfig,
      "AssociationConfig" : AssociationConfig,
      "CaptchaConfig" : CaptchaConfig,
      "ChallengeConfig" : ChallengeConfig,
      "CustomResponseBodies" : {Key: Value, ...},
      "DataProtectionConfig" : DataProtectionConfig,
      "DefaultAction" : DefaultAction,
      "Description" : String,
      "Name" : String,
      "OnSourceDDoSProtectionConfig" : OnSourceDDoSProtectionConfig,
      "Rules" : [ Rule, ... ],
      "Scope" : String,
      "Tags" : [ Tag, ... ],
      "TokenDomains" : [ String, ... ],
      "VisibilityConfig" : VisibilityConfig
    }
}

```

### YAML

```yaml

Type: AWS::WAFv2::WebACL
Properties:
  ApplicationConfig:
    ApplicationConfig
  AssociationConfig:
    AssociationConfig
  CaptchaConfig:
    CaptchaConfig
  ChallengeConfig:
    ChallengeConfig
  CustomResponseBodies:
    Key: Value
  DataProtectionConfig:
    DataProtectionConfig
  DefaultAction:
    DefaultAction
  Description: String
  Name: String
  OnSourceDDoSProtectionConfig:
    OnSourceDDoSProtectionConfig
  Rules:
    - Rule
  Scope: String
  Tags:
    - Tag
  TokenDomains:
    - String
  VisibilityConfig:
    VisibilityConfig

```

## Properties

`ApplicationConfig`

Returns a list of `ApplicationAttribute` s.

_Required_: No

_Type_: [ApplicationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-applicationconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociationConfig`

Specifies custom configurations for the associations between the web ACL and protected resources.

Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can
customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).

###### Note

You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://aws.amazon.com/waf/pricing).

For Application Load Balancer and AWS AppSync, the limit is fixed at 8 KB (8,192 bytes).

_Required_: No

_Type_: [AssociationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-associationconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaptchaConfig`

Specifies how AWS WAF should handle `CAPTCHA` evaluations for rules that don't have their own `CaptchaConfig` settings. If you don't specify this, AWS WAF uses its default settings for `CaptchaConfig`.

_Required_: No

_Type_: [CaptchaConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-captchaconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChallengeConfig`

Specifies how AWS WAF should handle challenge evaluations for rules that
don't have their own `ChallengeConfig` settings. If you don't specify this,
AWS WAF uses its default settings for `ChallengeConfig`.

_Required_: No

_Type_: [ChallengeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-challengeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomResponseBodies`

A map of custom response keys and content bodies. When you create a rule with a block action, you can send a custom response to the web request. You define these for the web ACL, and then use them in the rules and default actions that you define in the web ACL.

For information about customizing web requests and responses,
see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html)
in the _AWS WAF Developer Guide_.

For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: Object of [CustomResponseBody](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-customresponsebody.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataProtectionConfig`

Specifies data protection to apply to the web request data for the web ACL. This is a web ACL level data protection option.

The data protection that you configure for the web ACL alters the data that's available for any other data collection activity,
including your AWS WAF logging destinations, web ACL request sampling, and Amazon Security Lake data collection and management. Your other option for data protection is in the logging configuration, which only affects logging.

_Required_: No

_Type_: [DataProtectionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-dataprotectionconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultAction`

The action to perform if none of the `Rules` contained in the `WebACL` match.

_Required_: Yes

_Type_: [DefaultAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-defaultaction.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the web ACL that helps with identification.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9=:#@/\-,.][a-zA-Z0-9+=:#@/\-,.\s]+[a-zA-Z0-9+=:#@/\-,.]{1,256}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the web ACL. You cannot change the name of a web ACL after you create it.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]{1,128}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnSourceDDoSProtectionConfig`

Configures the level of DDoS protection that applies to web ACLs associated with Application Load Balancers.

_Required_: No

_Type_: [OnSourceDDoSProtectionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-onsourceddosprotectionconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rules`

The rule statements used to identify the web requests that you want to manage. Each rule
includes one top-level statement that AWS WAF uses to identify matching web
requests, and parameters that govern how AWS WAF handles them.

_Required_: No

_Type_: Array of [Rule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-rule.html)

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

For information about how to define the association of the web ACL with your resource,
see [AWS::WAFv2::WebACLAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-resource-wafv2-webaclassociation.html).

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

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-tag.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenDomains`

Specifies the domains that AWS WAF should accept in a web request token.
This enables the use of tokens across multiple protected websites. When AWS WAF
provides a token, it uses the domain of the AWS resource that the web ACL is
protecting. If you don't specify a list of token domains, AWS WAF accepts tokens
only for the domain of the protected resource. With a token domain list, AWS WAF
accepts the resource's host domain plus all domains in the token domain list, including
their prefixed subdomains.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisibilityConfig`

Defines and enables Amazon CloudWatch metrics and web request sample collection.

_Required_: Yes

_Type_: [VisibilityConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-visibilityconfig.html)

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

The Amazon Resource Name (ARN) of the web ACL.

`Capacity`

The web ACL capacity units (WCUs) currently being used by this web ACL.

AWS WAF uses WCUs to calculate and control the operating resources that
are used to run your rules, rule groups, and web ACLs. AWS WAF calculates
capacity differently for each rule type, to reflect the relative cost of each rule. Simple
rules that cost little to run use fewer WCUs than more complex rules that use more
processing power. Rule group capacity is fixed at creation, which helps users plan their
web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.

`Id`

The ID of the web ACL.

`LabelNamespace`

The label namespace prefix for this web ACL. All labels added by rules in this web ACL
have this prefix.

The syntax for the label namespace prefix for a web ACL is the following:
`awswaf:<account ID>:webacl:<web ACL name>:`

When a rule with a label matches a web request, AWS WAF adds the fully
qualified label to the request. A fully qualified label is made up of the label namespace
from the rule group or web ACL where the rule is defined and the label from the rule,
separated by a colon.

## Examples

- [Create a web ACL with a variety of rate-based rules](#aws-resource-wafv2-webacl--examples--Create_a_web_ACL_with_a_variety_of_rate-based_rules)

- [Create a web ACL with custom request and response handling](#aws-resource-wafv2-webacl--examples--Create_a_web_ACL_with_custom_request_and_response_handling)

- [Create a web ACL with JSON body parsing](#aws-resource-wafv2-webacl--examples--Create_a_web_ACL_with_JSON_body_parsing)

### Create a web ACL with a variety of rate-based rules

The following example web ACL specification shows a number of rate-based rules.
For information about each rule, see the examples under `AWS::WAFv2::WebACL
                  RateBasedStatement`.

#### YAML

```yaml

Name: exampleWebACL
Id: exampleWebACLExampleID
ARN: arn:aws:wafv2:us-east-1:ExampleAccountNumber:regional/webacl/exampleWebACL/exampleWebACLExampleID
DefaultAction:
  Allow: {}
Description: ''
Rules:
- Name: rbrCountAll
  Priority: 0
  Statement:
    RateBasedStatement:
      Limit: 100000
      AggregateKeyType: CONSTANT
      ScopeDownStatement:
        GeoMatchStatement:
          CountryCodes:
          - GB
          ForwardedIPConfig:
            HeaderName: X-Forwarded-For
            FallbackBehavior: MATCH
  Action:
    Block: {}
  VisibilityConfig:
    SampledRequestsEnabled: true
    CloudWatchMetricsEnabled: true
    MetricName: rbrCountAll
- Name: rbrNoCustomKeys
  Priority: 1
  Statement:
    RateBasedStatement:
      Limit: 1000
      AggregateKeyType: FORWARDED_IP
      ForwardedIPConfig:
        HeaderName: X-Forwarded-For
        FallbackBehavior: MATCH
  Action:
    Block: {}
  VisibilityConfig:
    SampledRequestsEnabled: true
    CloudWatchMetricsEnabled: true
    MetricName: rbrNoCustomKeys
- Name: rbrCustomKeysA
  Priority: 2
  Statement:
    RateBasedStatement:
      Limit: 2000
      AggregateKeyType: CUSTOM_KEYS
      ForwardedIPConfig:
        HeaderName: X-Forwarded-For
        FallbackBehavior: MATCH
      CustomKeys:
      - Header:
          Name: Content-Type
          TextTransformations:
          - Priority: 0
            Type: NONE
      - ForwardedIP: {}
  Action:
    Block: {}
  VisibilityConfig:
    SampledRequestsEnabled: true
    CloudWatchMetricsEnabled: true
    MetricName: rbrCustomKeysA
- Name: rbrCustomKeysB
  Priority: 3
  Statement:
    RateBasedStatement:
      Limit: 3000
      AggregateKeyType: CUSTOM_KEYS
      CustomKeys:
      - QueryString:
          TextTransformations:
          - Priority: 0
            Type: NONE
      - HTTPMethod: {}
      - UriPath:
          TextTransformations:
          - Priority: 0
            Type: NONE
  Action:
    Block: {}
  VisibilityConfig:
    SampledRequestsEnabled: true
    CloudWatchMetricsEnabled: true
    MetricName: rbrCustomKeysB
- Name: labelUSStates
  Priority: 4
  Statement:
    GeoMatchStatement:
      CountryCodes:
      - US
  Action:
    Count: {}
  VisibilityConfig:
    SampledRequestsEnabled: true
    CloudWatchMetricsEnabled: true
    MetricName: labelUSStates
- Name: rbrRequestsFromUSStates
  Priority: 5
  Statement:
    RateBasedStatement:
      Limit: 500
      AggregateKeyType: CUSTOM_KEYS
      ScopeDownStatement:
        GeoMatchStatement:
          CountryCodes:
          - US
      CustomKeys:
      - LabelNamespace:
          Namespace: 'awswaf:clientip:geo:region:'
  Action:
    Block: {}
  VisibilityConfig:
    SampledRequestsEnabled: true
    CloudWatchMetricsEnabled: true
    MetricName: rbrRequestsFromUSStates
VisibilityConfig:
  SampledRequestsEnabled: true
  CloudWatchMetricsEnabled: true
  MetricName: exampleWebACL
Capacity: 193
ManagedByFirewallManager: false
LabelNamespace: 'awswaf:ExampleAccountNumber:webacl:exampleWebACL:'
```

#### JSON

```json

 {
  "Name": "exampleWebACL",
  "Id": "exampleWebACLExampleID",
  "ARN": "arn:aws:wafv2:us-east-1:ExampleAccountNumber:regional/webacl/exampleWebACL/exampleWebACLExampleID",
  "DefaultAction": {
    "Allow": {}
  },
  "Description": "",
  "Rules": [
    {
      "Name": "rbrCountAll",
      "Priority": 0,
      "Statement": {
        "RateBasedStatement": {
          "Limit": 100000,
          "AggregateKeyType": "CONSTANT",
          "ScopeDownStatement": {
            "GeoMatchStatement": {
              "CountryCodes": [
                "GB"
              ],
              "ForwardedIPConfig": {
                "HeaderName": "X-Forwarded-For",
                "FallbackBehavior": "MATCH"
              }
            }
          }
        }
      },
      "Action": {
        "Block": {}
      },
      "VisibilityConfig": {
        "SampledRequestsEnabled": true,
        "CloudWatchMetricsEnabled": true,
        "MetricName": "rbrCountAll"
      }
    },
    {
      "Name": "rbrNoCustomKeys",
      "Priority": 1,
      "Statement": {
        "RateBasedStatement": {
          "Limit": 1000,
          "AggregateKeyType": "FORWARDED_IP",
          "ForwardedIPConfig": {
            "HeaderName": "X-Forwarded-For",
            "FallbackBehavior": "MATCH"
          }
        }
      },
      "Action": {
        "Block": {}
      },
      "VisibilityConfig": {
        "SampledRequestsEnabled": true,
        "CloudWatchMetricsEnabled": true,
        "MetricName": "rbrNoCustomKeys"
      }
    },
    {
      "Name": "rbrCustomKeysA",
      "Priority": 2,
      "Statement": {
        "RateBasedStatement": {
          "Limit": 2000,
          "AggregateKeyType": "CUSTOM_KEYS",
          "ForwardedIPConfig": {
            "HeaderName": "X-Forwarded-For",
            "FallbackBehavior": "MATCH"
          },
          "CustomKeys": [
            {
              "Header": {
                "Name": "Content-Type",
                "TextTransformations": [
                  {
                    "Priority": 0,
                    "Type": "NONE"
                  }
                ]
              }
            },
            {
              "ForwardedIP": {}
            }
          ]
        }
      },
      "Action": {
        "Block": {}
      },
      "VisibilityConfig": {
        "SampledRequestsEnabled": true,
        "CloudWatchMetricsEnabled": true,
        "MetricName": "rbrCustomKeysA"
      }
    },
    {
      "Name": "rbrCustomKeysB",
      "Priority": 3,
      "Statement": {
        "RateBasedStatement": {
          "Limit": 3000,
          "AggregateKeyType": "CUSTOM_KEYS",
          "CustomKeys": [
            {
              "QueryString": {
                "TextTransformations": [
                  {
                    "Priority": 0,
                    "Type": "NONE"
                  }
                ]
              }
            },
            {
              "HTTPMethod": {}
            },
            {
              "UriPath": {
                "TextTransformations": [
                  {
                    "Priority": 0,
                    "Type": "NONE"
                  }
                ]
              }
            }
          ]
        }
      },
      "Action": {
        "Block": {}
      },
      "VisibilityConfig": {
        "SampledRequestsEnabled": true,
        "CloudWatchMetricsEnabled": true,
        "MetricName": "rbrCustomKeysB"
      }
    },
    {
      "Name": "labelUSStates",
      "Priority": 4,
      "Statement": {
        "GeoMatchStatement": {
          "CountryCodes": [
            "US"
          ]
        }
      },
      "Action": {
        "Count": {}
      },
      "VisibilityConfig": {
        "SampledRequestsEnabled": true,
        "CloudWatchMetricsEnabled": true,
        "MetricName": "labelUSStates"
      }
    },
    {
      "Name": "rbrRequestsFromUSStates",
      "Priority": 5,
      "Statement": {
        "RateBasedStatement": {
          "Limit": 500,
          "AggregateKeyType": "CUSTOM_KEYS",
          "ScopeDownStatement": {
            "GeoMatchStatement": {
              "CountryCodes": [
                "US"
              ]
            }
          },
          "CustomKeys": [
            {
              "LabelNamespace": {
                "Namespace": "awswaf:clientip:geo:region:"
              }
            }
          ]
        }
      },
      "Action": {
        "Block": {}
      },
      "VisibilityConfig": {
        "SampledRequestsEnabled": true,
        "CloudWatchMetricsEnabled": true,
        "MetricName": "rbrRequestsFromUSStates"
      }
    }
  ],
  "VisibilityConfig": {
    "SampledRequestsEnabled": true,
    "CloudWatchMetricsEnabled": true,
    "MetricName": "exampleWebACL"
  },
  "Capacity": 193,
  "ManagedByFirewallManager": false,
  "LabelNamespace": "awswaf:ExampleAccountNumber:webacl:exampleWebACL:"
}
```

### Create a web ACL with custom request and response handling

The following shows an example web ACL specification. This example includes custom
request and response configurations.

#### YAML

```yaml

 ExampleWebACL:
    Type: 'AWS::WAFv2::WebACL'
    Properties:
      Name: ExampleWebACL1
      Scope: REGIONAL
      Description: This is an example WebACL
      DefaultAction:
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
        MetricName: ExampleWebACLMetric
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
      Rules:
        - Name: RuleWithAWSManagedRules
          Priority: 0
          OverrideAction:
            Count: {}
          VisibilityConfig:
            SampledRequestsEnabled: true
            CloudWatchMetricsEnabled: true
            MetricName: RuleWithAWSManagedRulesMetric
          Statement:
            ManagedRuleGroupStatement:
              VendorName: AWS
              Name: AWSManagedRulesCommonRuleSet
              ExcludedRules: []
        - Name: BlockXssAttack
          Priority: 1
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
            MetricName: BlockXssAttackMetric
          Statement:
            XssMatchStatement:
              FieldToMatch:
                AllQueryArguments: {}
              TextTransformations:
                - Priority: 1
                  Type: NONE
```

#### JSON

```json

 "ExampleWebACL": {
      "Type": "AWS::WAFv2::WebACL",
      "Properties": {
        "Name": "ExampleWebACL1",
        "Scope": "REGIONAL",
        "Description": "This is an example WebACL",
        "DefaultAction": {
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
          "MetricName": "ExampleWebACLMetric"
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
        "Rules": [
          {
            "Name": "RuleWithAWSManagedRules",
            "Priority": 0,
            "OverrideAction": {
              "Count": {}
            },
            "VisibilityConfig": {
              "SampledRequestsEnabled": true,
              "CloudWatchMetricsEnabled": true,
              "MetricName": "RuleWithAWSManagedRulesMetric"
            },
            "Statement": {
              "ManagedRuleGroupStatement": {
            "VendorName": "AWS",
                "Name": "AWSManagedRulesCommonRuleSet",
                "ExcludedRules": []
              }
            }
          },
          {
            "Name": "BlockXssAttack",
            "Priority": 1,
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
              "MetricName": "BlockXssAttackMetric"
            },
            "Statement": {
              "XssMatchStatement": {
                "FieldToMatch": {
                  "AllQueryArguments": {}
                },
                "TextTransformations": [
                  {
                    "Priority": 1,
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

### Create a web ACL with JSON body parsing

The following shows an example web ACL specification. This example includes
inspection of the web request body as JSON.

#### YAML

```yaml

  ExampleWebACL:
    Type: 'AWS::WAFv2::WebACL'
    Properties:
      Name: TestingJsonBody
      Scope: REGIONAL
      Description: WebACL for JsonBody Testing
      DefaultAction:
        Allow: {}
      VisibilityConfig:
        SampledRequestsEnabled: true
        CloudWatchMetricsEnabled: true
        MetricName: ExampleWebACLMetric
      Rules:
        - Name: TestJsonBodyRule
          Priority: 0
          Action:
            Block: {}
          VisibilityConfig:
            SampledRequestsEnabled: true
            CloudWatchMetricsEnabled: true
            MetricName: JsonBodyMatchMetric
          Statement:
            ByteMatchStatement:
              FieldToMatch:
                JsonBody:
                  MatchPattern:
                    IncludedPaths:
                      - /foo
                      - /bar
                  MatchScope: VALUE
                  InvalidFallbackBehavior: MATCH
              PositionalConstraint: EXACTLY
              SearchString: BadBot
              TextTransformations:
                - Type: NONE
                  Priority: 0
```

#### JSON

```json

  "ExampleWebACL": {
    "Type": "AWS::WAFv2::WebACL",
    "Properties": {
      "Name": "TestingJsonBody",
      "Scope": "REGIONAL",
      "DefaultAction":
       {
        "Allow":
         {}
       },
      "Description": "WebACL for JsonBody Testing",
      "Rules": [
      {
        "Name": "TestJsonBodyRule",
        "Priority": 1,
        "Statement": {
          "ByteMatchStatement": {
            "SearchString": "BadBot",
            "FieldToMatch": {
              "JsonBody": {
                "MatchPattern": {
                  "IncludedPaths": [
                    "/foo", "/bar"
                  ]
                },
                "MatchScope": "VALUE",
                "InvalidFallbackBehavior": "MATCH"
               }
            },
            "TextTransformations": [
              {
                "Priority": 1,
                "Type": "NONE"
              }
            ],
            "PositionalConstraint": "EXACTLY"
          }
        },
        "Action": {
          "Block": {}
        },
        "VisibilityConfig": {
             "SampledRequestsEnabled": true,
             "CloudWatchMetricsEnabled": true,
             "MetricName": "JsonBodyMatchMetric"
         }
      } ],
      "VisibilityConfig": {
         "SampledRequestsEnabled": true,
         "CloudWatchMetricsEnabled": true,
         "MetricName": "TestingJsonBodyMetric"
      }
    }
  }
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

XssMatchStatement

AllowAction
