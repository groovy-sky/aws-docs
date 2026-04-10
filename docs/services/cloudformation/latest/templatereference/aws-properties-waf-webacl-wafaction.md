This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAF::WebACL WafAction

###### Note

AWS WAF Classic support will end on September 30, 2025.

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

For the action that is associated with a rule in a `WebACL`, specifies the action that you want AWS WAF to perform when a
web request matches all of the conditions in a rule. For the default action in a `WebACL`, specifies the action that you want
AWS WAF to take when a web request doesn't match all of the conditions in any of the rules in a `WebACL`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

Specifies how you want AWS WAF to respond to requests that match the settings in a `Rule`. Valid settings include the following:

- `ALLOW`: AWS WAF allows requests

- `BLOCK`: AWS WAF blocks requests

- `COUNT`: AWS WAF increments a counter of the requests that match all of the conditions in the rule.
AWS WAF then continues to inspect the web request based on the remaining rules in the web ACL. You can't specify `COUNT`
for the default action for a `WebACL`.

_Required_: Yes

_Type_: String

_Allowed values_: `BLOCK | ALLOW | COUNT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActivatedRule

AWS::WAF::XssMatchSet

All content copied from https://docs.aws.amazon.com/.
