This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL OrStatement

A logical rule statement used to combine other rule statements with OR logic. You provide more than one [Statement](../userguide/aws-properties-wafv2-webacl-notstatement.md#cfn-wafv2-webacl-notstatement-statement) within the `OrStatement`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Statements" : [ Statement, ... ]
}

```

### YAML

```yaml

  Statements:
    - Statement

```

## Properties

`Statements`

The statements to combine with OR logic. You can use any statements that can be
nested.

_Required_: Yes

_Type_: Array of [Statement](aws-properties-wafv2-webacl-statement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnSourceDDoSProtectionConfig

OverrideAction

All content copied from https://docs.aws.amazon.com/.
