This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Invoicing::InvoiceUnit Rule

The `InvoiceUnitRule` object used to update invoice units.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LinkedAccounts" : [ String, ... ]
}

```

### YAML

```yaml

  LinkedAccounts:
    - String

```

## Properties

`LinkedAccounts`

The list of `LINKED_ACCOUNT` IDs where charges are included within the invoice unit.

_Required_: Yes

_Type_: Array of String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceTag

Next
