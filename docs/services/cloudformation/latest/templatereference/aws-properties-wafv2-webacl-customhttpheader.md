This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL CustomHTTPHeader

A custom header for custom request and response handling. This is used in [CustomResponse](../userguide/aws-properties-wafv2-webacl-blockaction.md#cfn-wafv2-webacl-blockaction-customresponse) and [CustomRequestHandling](../userguide/aws-properties-wafv2-webacl-countaction.md#cfn-wafv2-webacl-countaction-customrequesthandling).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the custom header.

For custom request header insertion, when AWS WAF inserts the header into the request,
it prefixes this name `x-amzn-waf-`, to avoid confusion with the headers that
are already in the request. For example, for the header name `sample`, AWS WAF
inserts the header `x-amzn-waf-sample`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the custom header.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CountAction

CustomRequestHandling

All content copied from https://docs.aws.amazon.com/.
