This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule RuleCondition

Specifies a condition for a listener rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Field" : String,
  "HostHeaderConfig" : HostHeaderConfig,
  "HttpHeaderConfig" : HttpHeaderConfig,
  "HttpRequestMethodConfig" : HttpRequestMethodConfig,
  "PathPatternConfig" : PathPatternConfig,
  "QueryStringConfig" : QueryStringConfig,
  "RegexValues" : [ String, ... ],
  "SourceIpConfig" : SourceIpConfig,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Field: String
  HostHeaderConfig:
    HostHeaderConfig
  HttpHeaderConfig:
    HttpHeaderConfig
  HttpRequestMethodConfig:
    HttpRequestMethodConfig
  PathPatternConfig:
    PathPatternConfig
  QueryStringConfig:
    QueryStringConfig
  RegexValues:
    - String
  SourceIpConfig:
    SourceIpConfig
  Values:
    - String

```

## Properties

`Field`

The field in the HTTP request. The following are the possible values:

- `http-header`

- `http-request-method`

- `host-header`

- `path-pattern`

- `query-string`

- `source-ip`

_Required_: No

_Type_: String

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostHeaderConfig`

Information for a host header condition. Specify only when `Field` is
`host-header`.

_Required_: No

_Type_: [HostHeaderConfig](aws-properties-elasticloadbalancingv2-listenerrule-hostheaderconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpHeaderConfig`

Information for an HTTP header condition. Specify only when `Field` is
`http-header`.

_Required_: Conditional

_Type_: [HttpHeaderConfig](aws-properties-elasticloadbalancingv2-listenerrule-httpheaderconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpRequestMethodConfig`

Information for an HTTP method condition. Specify only when `Field` is
`http-request-method`.

_Required_: Conditional

_Type_: [HttpRequestMethodConfig](aws-properties-elasticloadbalancingv2-listenerrule-httprequestmethodconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PathPatternConfig`

Information for a path pattern condition. Specify only when `Field` is
`path-pattern`.

_Required_: No

_Type_: [PathPatternConfig](aws-properties-elasticloadbalancingv2-listenerrule-pathpatternconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStringConfig`

Information for a query string condition. Specify only when `Field` is
`query-string`.

_Required_: Conditional

_Type_: [QueryStringConfig](aws-properties-elasticloadbalancingv2-listenerrule-querystringconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegexValues`

The regular expressions to match against the condition field. The maximum length of each string is 128 characters.
Specify only when `Field` is `http-header`, `host-header`, or `path-pattern`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceIpConfig`

Information for a source IP condition. Specify only when `Field` is
`source-ip`.

_Required_: Conditional

_Type_: [SourceIpConfig](aws-properties-elasticloadbalancingv2-listenerrule-sourceipconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The condition value. Specify only when `Field` is `host-header` or
`path-pattern`. Alternatively, to specify multiple host names or multiple
path patterns, use `HostHeaderConfig` or `PathPatternConfig`.

If `Field` is `host-header` and you're not using
`HostHeaderConfig`, you can specify a single host name (for example,
my.example.com). A host name is case insensitive, can be up to 128 characters in length,
and can contain any of the following characters.

- A-Z, a-z, 0-9

- \- .

- \\* (matches 0 or more characters)

- ? (matches exactly 1 character)

If `Field` is `path-pattern` and you're not using
`PathPatternConfig`, you can specify a single path pattern (for example,
/img/\*). A path pattern is case-sensitive, can be up to 128 characters in length, and can
contain any of the following characters.

- A-Z, a-z, 0-9

- \_ - . $ / ~ " ' @ : +

- & (using &amp;)

- \\* (matches 0 or more characters)

- ? (matches exactly 1 character)

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RewriteConfigObject

SourceIpConfig

All content copied from https://docs.aws.amazon.com/.
