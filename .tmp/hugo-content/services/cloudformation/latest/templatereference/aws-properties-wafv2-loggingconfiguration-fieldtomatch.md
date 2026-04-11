This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::LoggingConfiguration FieldToMatch

The parts of the request that you want to keep out of the logs. This is used in the
logging configuration `RedactedFields` specification.

Example JSON for a `QueryString` field to match:

` "FieldToMatch": { "QueryString": {} }`

Example JSON for a `Method` field to match specification:

` "FieldToMatch": { "Method": { "Name": "DELETE" } }`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Method" : Json,
  "QueryString" : Json,
  "SingleHeader" : SingleHeader,
  "UriPath" : Json
}

```

### YAML

```yaml

  Method: Json
  QueryString: Json
  SingleHeader:
    SingleHeader
  UriPath: Json

```

## Properties

`Method`

Redact the indicated HTTP method. The method indicates the type of operation that the
request is asking the origin to perform.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

Redact the query string. This is the part of a URL that appears after a `?`
character, if any.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleHeader`

Redact a single header. Provide the name of the header to inspect, for example,
`User-Agent` or `Referer`. This setting isn't case
sensitive.

Example JSON: `"SingleHeader": { "Name": "haystack" }`

_Required_: No

_Type_: [SingleHeader](aws-properties-wafv2-loggingconfiguration-singleheader.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UriPath`

Redact the request URI path. This is the part of the web request that identifies a
resource, for example, `/images/daily-ad.jpg`.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Condition

Filter

All content copied from https://docs.aws.amazon.com/.
