This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL AssociationConfig

Specifies custom configurations for the associations between the web ACL and protected resources.

Use this to customize the maximum size of the request body that your protected resources forward to AWS WAF for inspection. You can
customize this setting for CloudFront, API Gateway, Amazon Cognito, App Runner, or Verified Access resources. The default setting is 16 KB (16,384 bytes).

###### Note

You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://aws.amazon.com/waf/pricing).

For Application Load Balancer and AWS AppSync, the limit is fixed at 8 KB (8,192 bytes).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RequestBody" : {Key: Value, ...}
}

```

### YAML

```yaml

  RequestBody:
    Key: Value

```

## Properties

`RequestBody`

Customizes the maximum size of the request body that your protected CloudFront, API Gateway, Amazon Cognito, App Runner, and Verified Access resources forward to AWS WAF for inspection. The default size is 16 KB (16,384 bytes). You can change the setting for any of the available resource types.

###### Note

You are charged additional fees when your protected resources forward body sizes that are larger than the default. For more information, see [AWS WAF Pricing](https://aws.amazon.com/waf/pricing).

Example JSON: ` {
    "API_GATEWAY": "KB_48",
    "APP_RUNNER_SERVICE": "KB_32"
    }`

For Application Load Balancer and AWS AppSync, the limit is fixed at 8 KB (8,192 bytes).

_Required_: No

_Type_: Object of [RequestBodyAssociatedResourceTypeConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-requestbodyassociatedresourcetypeconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AsnMatchStatement

AWSManagedRulesACFPRuleSet
