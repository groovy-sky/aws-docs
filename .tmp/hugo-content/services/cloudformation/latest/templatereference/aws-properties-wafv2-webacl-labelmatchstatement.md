This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL LabelMatchStatement

A rule statement to match against labels that have been added to the web request by rules that have already run in the web ACL.

The label match statement provides the label or namespace string to search for. The label string can represent a part or all of the fully qualified label name that had been added to the web request. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. If you do not provide the fully qualified name in your label match string, AWS WAF performs the search for labels that were added in the same context as the label match statement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Scope" : String
}

```

### YAML

```yaml

  Key: String
  Scope: String

```

## Properties

`Key`

The string to match against. The setting you provide for this depends on the match
statement's `Scope` setting:

- If the `Scope` indicates `LABEL`, then this specification
must include the name and can include any number of preceding namespace
specifications and prefix up to providing the fully qualified label name.

- If the `Scope` indicates `NAMESPACE`, then this
specification can include any number of contiguous namespace strings, and can include
the entire label namespace prefix from the rule group or web ACL where the label
originates.

Labels are case sensitive and components of a label must be separated by colon, for
example `NS1:NS2:name`.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z_:-]{1,1024}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

Specify whether you want to match using the label name or just the namespace.

_Required_: Yes

_Type_: String

_Allowed values_: `LABEL | NAMESPACE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Label

ManagedRuleGroupConfig

All content copied from https://docs.aws.amazon.com/.
