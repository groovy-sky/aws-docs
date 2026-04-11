This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::WebACL Rule

A combination of `ByteMatchSet`, `IPSet`, and/or `SqlInjectionMatchSet` objects that identify the web requests that you
want to allow, block, or count. For example, you might create a `Rule` that includes the following predicates:

- An `IPSet` that causes AWS WAF to search for web requests that originate from the IP address `192.0.2.44`

- A `ByteMatchSet` that causes AWS WAF to search for web requests for which the value of the `User-Agent`
header is `BadBot`.

To match the settings in this `Rule`, a request must originate from `192.0.2.44` AND include a `User-Agent`
header for which the value is `BadBot`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : Action,
  "Priority" : Integer,
  "RuleId" : String
}

```

### YAML

```yaml

  Action:
    Action
  Priority: Integer
  RuleId: String

```

## Properties

`Action`

The action that AWS WAF takes when a web request matches all conditions in the rule, such as allow, block, or count the request.

_Required_: Yes

_Type_: [Action](aws-properties-wafregional-webacl-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Priority`

The order in which AWS WAF evaluates the rules in a web ACL. AWS WAF evaluates rules with a lower value before rules with a higher value. The value must be a unique integer. If you have multiple rules in a web ACL, the priority numbers do not need to be consecutive.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleId`

The ID of an AWS WAF Regional rule to associate with a web ACL.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

AWS::WAFRegional::WebACLAssociation

All content copied from https://docs.aws.amazon.com/.
