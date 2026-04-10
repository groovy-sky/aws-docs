This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptFilter IpFilter

A receipt IP address filter enables you to specify whether to accept or reject mail
originating from an IP address or range of IP addresses.

For information about setting up IP address filters, see the [Amazon SES\
Developer Guide](../../../ses/latest/dg/receiving-email-ip-filtering-console-walkthrough.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidr" : String,
  "Policy" : String
}

```

### YAML

```yaml

  Cidr: String
  Policy: String

```

## Properties

`Cidr`

A single IP address or a range of IP addresses to block or allow, specified in
Classless Inter-Domain Routing (CIDR) notation. An example of a single email address is
10.0.0.1. An example of a range of IP addresses is 10.0.0.1/24. For more information
about CIDR notation, see [RFC\
2317](https://tools.ietf.org/html/rfc2317).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

Indicates whether to block or allow incoming mail from the specified IP
addresses.

_Required_: Yes

_Type_: String

_Allowed values_: `Block | Allow`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

AWS::SES::ReceiptRule

All content copied from https://docs.aws.amazon.com/.
