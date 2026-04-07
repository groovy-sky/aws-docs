This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL VisibilityConfig

Defines and enables Amazon CloudWatch metrics and web request sample collection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchMetricsEnabled" : Boolean,
  "MetricName" : String,
  "SampledRequestsEnabled" : Boolean
}

```

### YAML

```yaml

  CloudWatchMetricsEnabled: Boolean
  MetricName: String
  SampledRequestsEnabled: Boolean

```

## Properties

`CloudWatchMetricsEnabled`

Indicates whether the associated resource sends metrics to Amazon CloudWatch. For the
list of available metrics, see [AWS WAF\
Metrics](https://docs.aws.amazon.com/waf/latest/developerguide/monitoring-cloudwatch.html#waf-metrics) in the _AWS WAF Developer Guide_.

For web ACLs, the metrics are for web requests that have the web ACL default action applied.
AWS WAF applies the default action to web requests that pass the inspection of all rules
in the web ACL without being either allowed or blocked. For more information,
see [The web ACL default action](https://docs.aws.amazon.com/waf/latest/developerguide/web-acl-default-action.html) in the _AWS WAF Developer Guide_.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

A name of the Amazon CloudWatch metric dimension. The name can contain only the characters: A-Z, a-z, 0-9,
\- (hyphen), and \_ (underscore). The name can be from one to 128 characters long. It can't
contain whitespace or metric names that are reserved for AWS WAF, for example `All` and
`Default_Action`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampledRequestsEnabled`

Indicates whether AWS WAF should store a sampling of the web requests that
match the rules. You can view the sampled requests through the AWS WAF console.

If you configure data protection for the web ACL, the protection applies to the web ACL's sampled web request data.

###### Note

Request sampling doesn't provide a field redaction option, and any field redaction that you specify in your logging configuration doesn't affect sampling.
You can only exclude fields from request sampling by disabling sampling in the web ACL visibility configuration
or by configuring data protection for the web ACL.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UriFragment

XssMatchStatement
