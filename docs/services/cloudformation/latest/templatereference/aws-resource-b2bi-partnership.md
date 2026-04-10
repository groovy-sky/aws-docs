This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership

Creates a partnership between a customer and a trading partner, based on the supplied
parameters. A partnership represents the connection between you and your trading partner. It ties
together a profile and one or more trading capabilities.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::B2BI::Partnership",
  "Properties" : {
      "Capabilities" : [ String, ... ],
      "CapabilityOptions" : CapabilityOptions,
      "Email" : String,
      "Name" : String,
      "Phone" : String,
      "ProfileId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::B2BI::Partnership
Properties:
  Capabilities:
    - String
  CapabilityOptions:
    CapabilityOptions
  Email: String
  Name: String
  Phone: String
  ProfileId: String
  Tags:
    - Tag

```

## Properties

`Capabilities`

Returns one or more capabilities associated with this partnership.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapabilityOptions`

Contains the details for an Outbound EDI capability.

_Required_: No

_Type_: [CapabilityOptions](aws-properties-b2bi-partnership-capabilityoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Email`

Specifies the email address associated with this trading partner.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\.\-]+@[\w\.\-]+$`

_Minimum_: `5`

_Maximum_: `254`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Returns the name of the partnership.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `254`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phone`

Specifies the phone number associated with the partnership.

_Required_: No

_Type_: String

_Pattern_: `^\+?([0-9 \t\-()\/]{7,})(?:\s*(?:#|x\.?|ext\.?|extension) \t*(\d+))?$`

_Minimum_: `7`

_Maximum_: `22`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProfileId`

Returns the unique, system-generated identifier for the profile connected to this partnership.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A key-value pair for a specific partnership. Tags are metadata that you can use to
search for and group capabilities for various purposes.

_Required_: No

_Type_: Array of [Tag](aws-properties-b2bi-partnership-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

Returns a timestamp for creation date and time of the partnership.

`ModifiedAt`

Returns a timestamp that identifies the most recent date and time that the partnership
was modified.

`PartnershipArn`

Returns an Amazon Resource Name (ARN) for a specific AWS resource, such as a capability, partnership, profile, or transformer.

`PartnershipId`

Returns the unique, system-generated identifier for a partnership.

`TradingPartnerId`

Returns the unique, system-generated identifier for a trading partner.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12Details

CapabilityOptions

All content copied from https://docs.aws.amazon.com/.
