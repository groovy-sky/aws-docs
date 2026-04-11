This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL Headers

Inspect all headers in the web request. You can specify the parts of the headers to
inspect and you can narrow the set of headers to inspect by including or excluding specific
keys.

This is used to indicate the web request component to inspect, in the
`FieldToMatch` specification.

If you want to inspect just the value of a single header, use the
`SingleHeader` `FieldToMatch` setting instead.

Example JSON: `"Headers": { "MatchPattern": { "All": {} }, "MatchScope": "KEY",
            "OversizeHandling": "MATCH" }`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MatchPattern" : HeaderMatchPattern,
  "MatchScope" : String,
  "OversizeHandling" : String
}

```

### YAML

```yaml

  MatchPattern:
    HeaderMatchPattern
  MatchScope: String
  OversizeHandling: String

```

## Properties

`MatchPattern`

The filter to use to identify the subset of headers to inspect in a web request.

You must specify exactly one setting: either `All`, `IncludedHeaders`, or `ExcludedHeaders`.

Example JSON: `"MatchPattern": { "ExcludedHeaders": [ "KeyToExclude1", "KeyToExclude2" ] }`

_Required_: Yes

_Type_: [HeaderMatchPattern](aws-properties-wafv2-webacl-headermatchpattern.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchScope`

The parts of the headers to match with the rule inspection criteria. If you specify
`ALL`, AWS WAF inspects both keys and values.

`All` does not require a match to be found in the keys
and a match to be found in the values. It requires a match to be found in the keys
or the values or both. To require a match in the keys and in the values, use a logical `AND` statement
to combine two match rules, one that inspects the keys and another that inspects the values.

_Required_: Yes

_Type_: String

_Allowed values_: `ALL | KEY | VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OversizeHandling`

What AWS WAF should do if the headers determined by your match scope are more numerous or larger than AWS WAF can inspect.
AWS WAF does not support inspecting the entire contents of request headers
when they exceed 8 KB (8192 bytes) or 200 total headers. The underlying host service forwards a maximum of 200 headers
and at most 8 KB of header contents to AWS WAF.

The options for oversize handling are the following:

- `CONTINUE` \- Inspect the available headers normally, according to the rule inspection criteria.

- `MATCH` \- Treat the web request as matching the rule statement. AWS WAF
applies the rule action to the request.

- `NO_MATCH` \- Treat the web request as not matching the rule
statement.

_Required_: Yes

_Type_: String

_Allowed values_: `CONTINUE | MATCH | NO_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Set the Headers specification

The following shows an example Headers field to match specification.

#### YAML

```yaml

FieldToMatch:
  Headers:
    MatchPattern:
      IncludedHeaders:
      - "User-Agent"
      - "Referer"
    MatchScope: ALL
    OversizeHandling: MATCH
```

#### JSON

```json

"FieldToMatch": {
  "Headers": {
    "MatchPattern": {
      "IncludedHeaders": [
        "User-Agent",
        "Referer"
      ]
    },
    "MatchScope": "ALL",
    "OversizeHandling": "MATCH"
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeaderOrder

ImmunityTimeProperty

All content copied from https://docs.aws.amazon.com/.
