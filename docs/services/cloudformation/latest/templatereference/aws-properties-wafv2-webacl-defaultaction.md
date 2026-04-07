This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL DefaultAction

In a [AWS::WAFv2::WebACL](aws-resource-wafv2-webacl.md), this is the action that you want AWS WAF to perform
when a web request doesn't match any of the rules in the `WebACL`. The default
action must be a terminating action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Allow" : AllowAction,
  "Block" : BlockAction
}

```

### YAML

```yaml

  Allow:
    AllowAction
  Block:
    BlockAction

```

## Properties

`Allow`

Specifies that AWS WAF should allow requests by default.

_Required_: No

_Type_: [AllowAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-allowaction.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Block`

Specifies that AWS WAF should block requests by default.

_Required_: No

_Type_: [BlockAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-blockaction.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Set a web ACL default action](#aws-properties-wafv2-webacl-defaultaction--examples--Set_a_web_ACL_default_action)

- [Set a customized web ACL default action](#aws-properties-wafv2-webacl-defaultaction--examples--Set_a_customized_web_ACL_default_action)

### Set a web ACL default action

The following shows an example web ACL default action specification that sets the
default action to "Block".

#### YAML

```yaml

DefaultAction:
    Block: {}
```

#### JSON

```json

"DefaultAction": {
    "Block": {}
}
```

### Set a customized web ACL default action

The following shows an example web ACL default action specification with
customization.

#### YAML

```yaml

DefaultAction:
  Allow:
    CustomRequestHandling:
      InsertHeaders:
        - Name: AllowActionHeader1Name
          Value: AllowActionHeader1Value
        - Name: AllowActionHeader2Name
          Value: AllowActionHeader2Value
```

#### JSON

```json

"DefaultAction": {
  "Allow": {
    "CustomRequestHandling": {
      "InsertHeaders": [
        {
          "Name": "AllowActionHeader1Name",
          "Value": "AllowActionHeader1Value"
        },
        {
          "Name": "AllowActionHeader2Name",
          "Value": "AllowActionHeader2Value"
        }
      ]
    }
  }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataProtectionConfig

ExcludedRule
