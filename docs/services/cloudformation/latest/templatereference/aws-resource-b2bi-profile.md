This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Profile

Creates a customer profile. You can have up to five customer profiles, each representing
a distinct private network. A profile is the mechanism used to create the concept of
a private network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::B2BI::Profile",
  "Properties" : {
      "BusinessName" : String,
      "Email" : String,
      "Logging" : String,
      "Name" : String,
      "Phone" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::B2BI::Profile
Properties:
  BusinessName: String
  Email: String
  Logging: String
  Name: String
  Phone: String
  Tags:
    - Tag

```

## Properties

`BusinessName`

Returns the name for the business associated with this profile.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `254`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Email`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[\w\.\-]+@[\w\.\-]+$`

_Minimum_: `5`

_Maximum_: `254`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logging`

Specifies whether or not logging is enabled for this profile.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Returns the display name for profile.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `254`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Phone`

Specifies the phone number associated with the profile.

_Required_: Yes

_Type_: String

_Pattern_: `^\+?([0-9 \t\-()\/]{7,})(?:\s*(?:#|x\.?|ext\.?|extension) \t*(\d+))?$`

_Minimum_: `7`

_Maximum_: `22`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A key-value pair for a specific profile. Tags are metadata that you can use to search
for and group capabilities for various purposes.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-profile-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

Returns the timestamp for creation date and time of the profile.

`LogGroupName`

Returns the name of the logging group.

`ModifiedAt`

Returns the timestamp that identifies the most recent date and time that the profile was
modified.

`ProfileArn`

Returns an Amazon Resource Name (ARN) for the profile.

`ProfileId`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

X12OutboundEdiHeaders

Tag
