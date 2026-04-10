This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain MatchingRule

Specifies how the rule-based matching process should match profiles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rule" : [ String, ... ]
}

```

### YAML

```yaml

  Rule:
    - String

```

## Properties

`Rule`

A single rule level of the `MatchRules`. Configures how the rule-based
matching process should match profiles.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `255 | 15`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Matching

Readiness

All content copied from https://docs.aws.amazon.com/.
