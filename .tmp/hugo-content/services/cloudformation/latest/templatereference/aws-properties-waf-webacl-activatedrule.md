This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAF::WebACL ActivatedRule

The `ActivatedRule` object in an `UpdateWebACL` request specifies a `Rule` that you want to insert or delete,
the priority of the `Rule` in the `WebACL`, and the action that you want AWS WAF to take when a web request matches the `Rule`
( `ALLOW`, `BLOCK`, or `COUNT`).

To specify whether to insert or delete a `Rule`, use the `Action` parameter in the `WebACLUpdate` data type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : WafAction,
  "Priority" : Integer,
  "RuleId" : String
}

```

### YAML

```yaml

  Action:
    WafAction
  Priority: Integer
  RuleId: String

```

## Properties

`Action`

Specifies the action that Amazon CloudFront or AWS WAF takes when a web request matches the conditions in the `Rule`.
Valid values for `Action` include the following:

- `ALLOW`: CloudFront responds with the requested object.

- `BLOCK`: CloudFront responds with an HTTP 403 (Forbidden) status code.

- `COUNT`: AWS WAF increments a counter of requests that match the conditions in the rule and then continues to
inspect the web request based on the remaining rules in the web ACL.

`ActivatedRule|OverrideAction` applies only when updating or adding a
`RuleGroup` to a `WebACL`. In this
case,
you do not use `ActivatedRule|Action`. For all other update requests,
`ActivatedRule|Action` is used instead of
`ActivatedRule|OverrideAction`.

_Required_: No

_Type_: [WafAction](aws-properties-waf-webacl-wafaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

Specifies the order in which the `Rules` in a `WebACL` are evaluated. Rules with a lower value for
`Priority` are evaluated before `Rules` with a higher value. The value must be a unique integer. If you add multiple
`Rules` to a `WebACL`, the values don't need to be consecutive.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleId`

The `RuleId` for a `Rule`. You use `RuleId` to get more information about a `Rule`,
update a `Rule`, insert a `Rule` into a `WebACL` or delete a
one from a `WebACL`, or delete a `Rule` from AWS WAF.

`RuleId` is returned by `CreateRule` and by `ListRules`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::WAF::WebACL

WafAction

All content copied from https://docs.aws.amazon.com/.
