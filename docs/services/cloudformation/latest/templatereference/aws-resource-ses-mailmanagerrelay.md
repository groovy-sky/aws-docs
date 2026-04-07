This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MailManagerRelay

Resource to create an SMTP relay, which can be used within a Mail Manager rule set to
forward incoming emails to defined relay destinations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::MailManagerRelay",
  "Properties" : {
      "Authentication" : RelayAuthentication,
      "RelayName" : String,
      "ServerName" : String,
      "ServerPort" : Number,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::MailManagerRelay
Properties:
  Authentication:
    RelayAuthentication
  RelayName: String
  ServerName: String
  ServerPort: Number
  Tags:
    - Tag

```

## Properties

`Authentication`

Authentication for the relay destination server—specify the secretARN where
the SMTP credentials are stored.

_Required_: Yes

_Type_: [RelayAuthentication](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-mailmanagerrelay-relayauthentication.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelayName`

The unique relay name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerName`

The destination relay server address.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-\.]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerPort`

The destination relay server port.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for the resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-mailmanagerrelay-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`RelayArn`

The Amazon Resource Name (ARN) of the relay.

`RelayId`

The unique relay identifier.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

RelayAuthentication
