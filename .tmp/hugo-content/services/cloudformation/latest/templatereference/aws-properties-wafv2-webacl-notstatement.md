This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL NotStatement

A logical rule statement used to negate the results of another rule statement. You provide one [Statement](../userguide/aws-properties-wafv2-webacl-notstatement.md#cfn-wafv2-webacl-notstatement-statement) within the `NotStatement`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Statement" : Statement
}

```

### YAML

```yaml

  Statement:
    Statement

```

## Properties

`Statement`

The statement to negate. You can use any statement that can be nested.

_Required_: Yes

_Type_: [Statement](aws-properties-wafv2-webacl-statement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedRuleGroupStatement

OnSourceDDoSProtectionConfig

All content copied from https://docs.aws.amazon.com/.
