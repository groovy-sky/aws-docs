This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ReceiptFilter Filter

Specifies an IP address filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IpFilter" : IpFilter,
  "Name" : String
}

```

### YAML

```yaml

  IpFilter:
    IpFilter
  Name: String

```

## Properties

`IpFilter`

A structure that provides the IP addresses to block or allow, and whether to block or
allow incoming mail from them.

_Required_: Yes

_Type_: [IpFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-receiptfilter-ipfilter.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the IP address filter. The name must meet the following
requirements:

- Contain only ASCII letters (a-z, A-Z), numbers (0-9), underscores (\_), or
dashes (-).

- Start and end with a letter or number.

- Contain 64 characters or fewer.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SES::ReceiptFilter

IpFilter
