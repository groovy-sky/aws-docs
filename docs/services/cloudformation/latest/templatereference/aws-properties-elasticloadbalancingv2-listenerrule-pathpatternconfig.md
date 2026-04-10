This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule PathPatternConfig

Information about a path pattern condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RegexValues" : [ String, ... ],
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  RegexValues:
    - String
  Values:
    - String

```

## Properties

`RegexValues`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The path patterns to compare against the request URL. The maximum size of each
string is 128 characters. The comparison is case sensitive. The following wildcard
characters are supported: \* (matches 0 or more characters) and ? (matches exactly 1
character).

If you specify multiple strings, the condition is satisfied if one of them matches the
request URL. The path pattern is compared only to the path of the URL, not to its query
string.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JwtValidationConfig

QueryStringConfig

All content copied from https://docs.aws.amazon.com/.
