This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup Body

Inspect the body of the web request. The body immediately follows the request
headers.

This is used to indicate the web request component to inspect, in the [FieldToMatch](../userguide/aws-properties-wafv2-rulegroup-regexpatternsetreferencestatement.md#cfn-wafv2-rulegroup-regexpatternsetreferencestatement-fieldtomatch) specification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OversizeHandling" : String
}

```

### YAML

```yaml

  OversizeHandling: String

```

## Properties

`OversizeHandling`

What AWS WAF should do if the body is larger than AWS WAF can inspect.

AWS WAF does not support inspecting the entire contents of the web request body if the body
exceeds the limit for the resource type. When a web request body is larger than the limit, the underlying host service
only forwards the contents that are within the limit to AWS WAF for inspection.

- For Application Load Balancer and AWS AppSync, the limit is fixed at 8 KB (8,192 bytes).

- For CloudFront, API Gateway, Amazon Cognito, App Runner, and Verified Access, the default limit is 16 KB (16,384 bytes), and
you can increase the limit for each resource type in the web ACL `AssociationConfig`, for additional processing fees.

- For AWS Amplify, use the CloudFront limit.

The options for oversize handling are the following:

- `CONTINUE` \- Inspect the available body contents normally, according to the rule inspection criteria.

- `MATCH` \- Treat the web request as matching the rule statement. AWS WAF
applies the rule action to the request.

- `NO_MATCH` \- Treat the web request as not matching the rule
statement.

You can combine the `MATCH` or `NO_MATCH`
settings for oversize handling with your rule and web ACL action settings, so that you block any request whose body is over the limit.

Default: `CONTINUE`

_Required_: No

_Type_: String

_Allowed values_: `CONTINUE | MATCH | NO_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Set the Body specification

The following shows an example Body field to match specification.

#### YAML

```yaml

FieldToMatch:
  Body:
    OversizeHandling: MATCH
```

#### JSON

```json

"FieldToMatch": {
  "Body": {
    "OversizeHandling": "MATCH"
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlockAction

ByteMatchStatement

All content copied from https://docs.aws.amazon.com/.
