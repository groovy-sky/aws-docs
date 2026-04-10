This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup CookieMatchPattern

The filter to use to identify the subset of cookies to inspect in a web request.

You must specify exactly one setting: either `All`, `IncludedCookies`, or `ExcludedCookies`.

Example JSON: `"MatchPattern": { "IncludedCookies": [ "session-id-time", "session-id" ] }`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "All" : Json,
  "ExcludedCookies" : [ String, ... ],
  "IncludedCookies" : [ String, ... ]
}

```

### YAML

```yaml

  All: Json
  ExcludedCookies:
    - String
  IncludedCookies:
    - String

```

## Properties

`All`

Inspect all cookies.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExcludedCookies`

Inspect only the cookies whose keys don't match any of the strings specified here.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `60 | 199`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedCookies`

Inspect only the cookies that have a key that matches one of the strings specified here.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `60 | 199`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Set a cookie match pattern for all paths](#aws-properties-wafv2-rulegroup-cookiematchpattern--examples--Set_a_cookie_match_pattern_for_all_paths)

- [Set a cookie match pattern with included paths](#aws-properties-wafv2-rulegroup-cookiematchpattern--examples--Set_a_cookie_match_pattern_with_included_paths)

### Set a cookie match pattern for all paths

The following shows an example cookie match pattern specification for all cookies.

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

### Set a cookie match pattern with included paths

The following shows an example cookie match pattern specification with included
keys.

#### YAML

```yaml

MatchPattern:
  IncludedCookies:
      - "session-id"
      - "session-id-time"
```

#### JSON

```json

"MatchPattern": {
  "IncludedCookies": [
    "session-id",
    "session-id-time"
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ChallengeConfig

Cookies

All content copied from https://docs.aws.amazon.com/.
