This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryRule FieldToMatch

Specifies a field in the request to redact from WAF logs, such as headers, query
parameters, or body content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Method" : String,
  "QueryString" : String,
  "SingleHeader" : SingleHeader,
  "UriPath" : String
}

```

### YAML

```yaml

  Method: String
  QueryString:
    String
  SingleHeader:
    SingleHeader
  UriPath: String

```

## Properties

`Method`

Redacts the HTTP method from WAF logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

Redacts the entire query string from WAF logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleHeader`

Redacts a specific header field by name from WAF logs.

_Required_: No

_Type_: [SingleHeader](aws-properties-observabilityadmin-telemetryrule-singleheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UriPath`

Redacts the URI path from WAF logs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ELBLoadBalancerLoggingParameters

Filter

All content copied from https://docs.aws.amazon.com/.
