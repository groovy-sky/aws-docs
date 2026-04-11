This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleIpToEvaluate

The IP address to evaluate for this condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attribute" : String
}

```

### YAML

```yaml

  Attribute: String

```

## Properties

`Attribute`

The attribute of the email to evaluate.

_Required_: Yes

_Type_: String

_Allowed values_: `SOURCE_IP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleIpExpression

RuleIsInAddressList

All content copied from https://docs.aws.amazon.com/.
