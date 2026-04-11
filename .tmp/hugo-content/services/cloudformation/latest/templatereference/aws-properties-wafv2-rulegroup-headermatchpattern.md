This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup HeaderMatchPattern

The filter to use to identify the subset of headers to inspect in a web request.

You must specify exactly one setting: either `All`, `IncludedHeaders`, or `ExcludedHeaders`.

Example JSON: `"MatchPattern": { "ExcludedHeaders": [ "KeyToExclude1", "KeyToExclude2" ] }`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "All" : Json,
  "ExcludedHeaders" : [ String, ... ],
  "IncludedHeaders" : [ String, ... ]
}

```

### YAML

```yaml

  All: Json
  ExcludedHeaders:
    - String
  IncludedHeaders:
    - String

```

## Properties

`All`

Inspect all headers.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludedHeaders`

Inspect only the headers whose keys don't match any of the strings specified here.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `64 | 199`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedHeaders`

Inspect only the headers that have a key that matches one of the strings specified here.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `64 | 199`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Set a header match pattern for all paths](#aws-properties-wafv2-rulegroup-headermatchpattern--examples--Set_a_header_match_pattern_for_all_paths)

- [Set a header match pattern with included paths](#aws-properties-wafv2-rulegroup-headermatchpattern--examples--Set_a_header_match_pattern_with_included_paths)

### Set a header match pattern for all paths

The following shows an example header match pattern specification for all headers.

#### YAML

```yaml

MatchPattern:
  All: {}
```

#### JSON

```json

"MatchPattern": {
  "All": {}
}
```

### Set a header match pattern with included paths

The following shows an example header match pattern specification with included
keys.

#### YAML

```yaml

MatchPattern:
  IncludedHeaders:
      - "User-Agent"
      - "Referer"
```

#### JSON

```json

"MatchPattern": {
  "IncludedHeaders": [
    "User-Agent",
    "Referer"
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GeoMatchStatement

HeaderOrder

All content copied from https://docs.aws.amazon.com/.
