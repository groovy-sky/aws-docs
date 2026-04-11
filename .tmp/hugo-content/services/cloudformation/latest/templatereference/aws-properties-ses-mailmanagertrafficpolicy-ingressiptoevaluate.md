This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerTrafficPolicy IngressIpToEvaluate

The structure for an IP based condition matching on the incoming mail.

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

An enum type representing the allowed attribute types for an IP condition.

_Required_: Yes

_Type_: String

_Allowed values_: `SENDER_IP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IngressBooleanToEvaluate

IngressIpv4Expression

All content copied from https://docs.aws.amazon.com/.
