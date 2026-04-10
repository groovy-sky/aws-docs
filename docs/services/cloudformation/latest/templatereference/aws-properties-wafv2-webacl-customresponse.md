This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL CustomResponse

A custom response to send to the client. You can define a custom response for rule
actions and default web ACL actions that are set to the block action.

For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](../../../waf/latest/developerguide/waf-custom-request-response.md) in the [AWS WAF developer guide](../../../waf/latest/developerguide/waf-chapter.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomResponseBodyKey" : String,
  "ResponseCode" : Integer,
  "ResponseHeaders" : [ CustomHTTPHeader, ... ]
}

```

### YAML

```yaml

  CustomResponseBodyKey: String
  ResponseCode: Integer
  ResponseHeaders:
    - CustomHTTPHeader

```

## Properties

`CustomResponseBodyKey`

References the response body that you want AWS WAF to return to the web request
client. You can define a custom response for a rule action or a default web ACL action that
is set to block. To do this, you first define the response body key and value in the
`CustomResponseBodies` setting for the [AWS::WAFv2::WebACL](aws-resource-wafv2-webacl.md) or [AWS::WAFv2::RuleGroup](aws-resource-wafv2-rulegroup.md) where you want to use it. Then, in the rule action or web ACL
default action `BlockAction` setting, you reference the response body using this
key.

_Required_: No

_Type_: String

_Pattern_: `^[\w\-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseCode`

The HTTP status code to return to the client.

For a list of status codes that you can use in your custom responses, see [Supported status codes for custom response](../../../waf/latest/developerguide/customizing-the-response-status-codes.md)
in the _AWS WAF Developer Guide_.

_Required_: Yes

_Type_: Integer

_Minimum_: `200`

_Maximum_: `599`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseHeaders`

The HTTP headers to use in the response. You can specify any header name except for `content-type`. Duplicate header names are not allowed.

For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](../../../waf/latest/developerguide/limits.md)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: Array of [CustomHTTPHeader](aws-properties-wafv2-webacl-customhttpheader.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomRequestHandling

CustomResponseBody

All content copied from https://docs.aws.amazon.com/.
