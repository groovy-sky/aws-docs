This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule QueryStringConfig

Information about a query string condition.

The query string component of a URI starts after the first '?' character and is terminated
by either a '#' character or the end of the URI. A typical query string contains key/value
pairs separated by '&' characters. The allowed characters are specified by RFC 3986. Any
character can be percentage encoded.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Values" : [ QueryStringKeyValue, ... ]
}

```

### YAML

```yaml

  Values:
    - QueryStringKeyValue

```

## Properties

`Values`

The key/value pairs or values to find in the query string. The maximum length of
each string is 128 characters. The comparison is case insensitive. The following wildcard
characters are supported: \* (matches 0 or more characters) and ? (matches exactly 1
character). To search for a literal '\*' or '?' character in a query string, you must escape
these characters in `Values` using a '\\' character.

If you specify multiple key/value pairs or values, the condition is satisfied if one of
them is found in the query string.

_Required_: No

_Type_: Array of [QueryStringKeyValue](aws-properties-elasticloadbalancingv2-listenerrule-querystringkeyvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PathPatternConfig

QueryStringKeyValue

All content copied from https://docs.aws.amazon.com/.
