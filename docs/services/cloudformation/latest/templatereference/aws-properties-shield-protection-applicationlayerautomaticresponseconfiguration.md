This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Shield::Protection ApplicationLayerAutomaticResponseConfiguration

The automatic application layer DDoS mitigation settings for a [AWS::Shield::Protection](aws-resource-shield-protection.md).
This configuration determines whether Shield Advanced automatically
manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.

If you use CloudFormation to manage the web ACLs that you use with Shield Advanced automatic mitigation, see the guidance
for the `AWS::WAFv2::WebACL` resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : Action,
  "Status" : String
}

```

### YAML

```yaml

  Action:
    Action
  Status: String

```

## Properties

`Action`

Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the
protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature,
when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.

_Required_: Yes

_Type_: [Action](aws-properties-shield-protection-action.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Indicates whether automatic application layer DDoS mitigation is enabled for the protection.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

Tag

All content copied from https://docs.aws.amazon.com/.
