This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL JsonMatchPattern

The patterns to look for in the JSON body. AWS WAF inspects the results of these
pattern matches against the rule inspection criteria. This is used with the [FieldToMatch](../userguide/aws-properties-wafv2-rulegroup-regexpatternsetreferencestatement.md#cfn-wafv2-rulegroup-regexpatternsetreferencestatement-fieldtomatch) option `JsonBody`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "All" : Json,
  "IncludedPaths" : [ String, ... ]
}

```

### YAML

```yaml

  All: Json
  IncludedPaths:
    - String

```

## Properties

`All`

Match all of the elements. See also `MatchScope` in the `JsonBody` `FieldToMatch` specification.

You must specify either this setting or the `IncludedPaths` setting, but not
both.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedPaths`

Match only the specified include paths. See also `MatchScope` in the
`JsonBody` `FieldToMatch` specification.

Provide the include paths using JSON Pointer syntax. For example, `"IncludedPaths":
            ["/dogs/0/name", "/dogs/1/name"]`. For information about this syntax, see the
Internet Engineering Task Force (IETF) documentation [JavaScript Object Notation (JSON)\
Pointer](https://tools.ietf.org/html/rfc6901).

You must specify either this setting or the `All` setting, but not
both.

###### Note

Don't use this option to include all paths. Instead, use the `All`
setting.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Set a JSON match pattern for all paths](#aws-properties-wafv2-webacl-jsonmatchpattern--examples--Set_a_JSON_match_pattern_for_all_paths)

- [Set a JSON match pattern with included paths](#aws-properties-wafv2-webacl-jsonmatchpattern--examples--Set_a_JSON_match_pattern_with_included_paths)

### Set a JSON match pattern for all paths

The following shows an example JSON match pattern specification for all paths.

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

### Set a JSON match pattern with included paths

The following shows an example JSON match pattern specification with included
paths.

#### YAML

```yaml

MatchPattern:
  IncludedPaths:
    - "/dogs/0/name"
    - "/cats/0/name"
```

#### JSON

```json

"MatchPattern": {
  "IncludedPaths": [
    "/dogs/0/name",
    "/cats/0/name"
  ]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JsonBody

Label

All content copied from https://docs.aws.amazon.com/.
