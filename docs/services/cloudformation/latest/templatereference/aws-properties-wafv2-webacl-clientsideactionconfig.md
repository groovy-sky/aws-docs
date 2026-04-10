This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ClientSideActionConfig

This is part of the configuration for the managed rules `AWSManagedRulesAntiDDoSRuleSet`
in `ManagedRuleGroupConfig`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Challenge" : ClientSideAction
}

```

### YAML

```yaml

  Challenge:
    ClientSideAction

```

## Properties

`Challenge`

Configuration for the use of the `AWSManagedRulesAntiDDoSRuleSet` rules `ChallengeAllDuringEvent` and `ChallengeDDoSRequests`.

###### Note

This setting isn't related to the configuration of the `Challenge` action itself. It only
configures the use of the two anti-DDoS rules named here.

You can enable or disable the use of these rules, and you can configure how to use them when they are enabled.

_Required_: Yes

_Type_: [ClientSideAction](aws-properties-wafv2-webacl-clientsideaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientSideAction

CookieMatchPattern

All content copied from https://docs.aws.amazon.com/.
