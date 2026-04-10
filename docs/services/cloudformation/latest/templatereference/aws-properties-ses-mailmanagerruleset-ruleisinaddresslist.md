This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRuleSet RuleIsInAddressList

The structure type for a boolean condition that provides the address lists and address list attribute to evaluate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AddressLists" : [ String, ... ],
  "Attribute" : String
}

```

### YAML

```yaml

  AddressLists:
    - String
  Attribute: String

```

## Properties

`AddressLists`

The address lists that will be used for evaluation.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Attribute`

The email attribute that needs to be evaluated against the address list.

_Required_: Yes

_Type_: String

_Allowed values_: `RECIPIENT | MAIL_FROM | SENDER | FROM | TO | CC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleIpToEvaluate

RuleNumberExpression

All content copied from https://docs.aws.amazon.com/.
