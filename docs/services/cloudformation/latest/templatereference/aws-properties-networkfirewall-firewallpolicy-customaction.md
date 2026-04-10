This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::FirewallPolicy CustomAction

An optional, non-standard action to use for stateless packet handling. You can define
this in addition to the standard action that you must specify.

You define and name the custom actions that you want to be able to use, and then you
reference them by name in your actions settings.

You can use custom actions in the following places:

- In an `StatelessRulesAndCustomActions`.
The custom actions are available for use by name inside the
`StatelessRulesAndCustomActions` where you define them. You can use
them for your stateless rule actions to specify what to do with a packet that matches
the rule's match attributes.

- In an firewall policy specification, in
`StatelessCustomActions`. The custom actions are available for use
inside the policy where you define them. You can use them for the policy's default
stateless actions settings to specify what to do with packets that don't match any of
the policy's stateless rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionDefinition" : ActionDefinition,
  "ActionName" : String
}

```

### YAML

```yaml

  ActionDefinition:
    ActionDefinition
  ActionName: String

```

## Properties

`ActionDefinition`

The custom action associated with the action name.

_Required_: Yes

_Type_: [ActionDefinition](aws-properties-networkfirewall-firewallpolicy-actiondefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ActionName`

The descriptive name of the custom action. You can't change the name of a custom action after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionDefinition

Dimension

All content copied from https://docs.aws.amazon.com/.
