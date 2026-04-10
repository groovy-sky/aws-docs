This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Shield::Protection Action

Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the
protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature,
when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Block" : Json,
  "Count" : Json
}

```

### YAML

```yaml

  Block: Json
  Count: Json

```

## Properties

`Block`

Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Block` action.

You must specify exactly one action, either `Block` or `Count`.

Example JSON: `{ "Block": {} }`

Example YAML: `Block: {}`

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Count`

Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Count` action.

You must specify exactly one action, either `Block` or `Count`.

Example JSON: `{ "Count": {} }`

Example YAML: `Count: {}`

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Shield::Protection

ApplicationLayerAutomaticResponseConfiguration

All content copied from https://docs.aws.amazon.com/.
