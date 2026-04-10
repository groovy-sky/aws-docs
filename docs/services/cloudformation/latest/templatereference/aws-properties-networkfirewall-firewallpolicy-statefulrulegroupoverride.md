This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy StatefulRuleGroupOverride

The setting that allows the policy owner to change the behavior of the rule group within a policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String
}

```

### YAML

```yaml

  Action: String

```

## Properties

`Action`

The action that changes the rule group from `DROP` to `ALERT`. This only applies to
managed rule groups.

_Required_: No

_Type_: String

_Allowed values_: `DROP_TO_ALERT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StatefulEngineOptions

StatefulRuleGroupReference

All content copied from https://docs.aws.amazon.com/.
