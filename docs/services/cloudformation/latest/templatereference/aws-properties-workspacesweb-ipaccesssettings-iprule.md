This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::IpAccessSettings IpRule

The IP rules of the IP access settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "IpRange" : String
}

```

### YAML

```yaml

  Description: String
  IpRange: String

```

## Properties

`Description`

The description of the IP rule.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpRange`

The IP range of the IP rule. This can either be a single IP address or a range using CIDR notation.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?:/([0-9]|[12][0-9]|3[0-2])|)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::WorkSpacesWeb::IpAccessSettings

Tag
