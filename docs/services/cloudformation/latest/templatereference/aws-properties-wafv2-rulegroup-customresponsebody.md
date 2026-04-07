This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup CustomResponseBody

The response body to use in a custom response to a web request. This is referenced by
key from [CustomResponse](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-blockaction.html#cfn-wafv2-webacl-blockaction-customresponse) `CustomResponseBodyKey`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : String,
  "ContentType" : String
}

```

### YAML

```yaml

  Content: String
  ContentType: String

```

## Properties

`Content`

The payload of the custom response.

You can use JSON escape strings in JSON content. To do this, you must specify JSON
content in the `ContentType` setting.

For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html)
in the _AWS WAF Developer Guide_.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentType`

The type of content in the payload that you are defining in the `Content`
string.

_Required_: Yes

_Type_: String

_Allowed values_: `TEXT_PLAIN | TEXT_HTML | APPLICATION_JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomResponse

FieldToMatch
