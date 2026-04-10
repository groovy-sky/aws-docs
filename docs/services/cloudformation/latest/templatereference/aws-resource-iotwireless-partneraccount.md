This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::PartnerAccount

A partner account. If `PartnerAccountId` and `PartnerType` are
`null`, returns all partner accounts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTWireless::PartnerAccount",
  "Properties" : {
      "AccountLinked" : Boolean,
      "PartnerAccountId" : String,
      "PartnerType" : String,
      "Sidewalk" : SidewalkAccountInfo,
      "SidewalkResponse" : SidewalkAccountInfoWithFingerprint,
      "SidewalkUpdate" : SidewalkUpdateAccount,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTWireless::PartnerAccount
Properties:
  AccountLinked: Boolean
  PartnerAccountId: String
  PartnerType: String
  Sidewalk:
    SidewalkAccountInfo
  SidewalkResponse:
    SidewalkAccountInfoWithFingerprint
  SidewalkUpdate:
    SidewalkUpdateAccount
  Tags:
    - Tag

```

## Properties

`AccountLinked`

Whether the partner account is linked to the AWS account.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartnerAccountId`

The ID of the partner account to update.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PartnerType`

The partner type.

_Required_: No

_Type_: String

_Allowed values_: `Sidewalk`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sidewalk`

The Sidewalk account credentials.

_Required_: No

_Type_: [SidewalkAccountInfo](aws-properties-iotwireless-partneraccount-sidewalkaccountinfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SidewalkResponse`

Information about a Sidewalk account.

_Required_: No

_Type_: [SidewalkAccountInfoWithFingerprint](aws-properties-iotwireless-partneraccount-sidewalkaccountinfowithfingerprint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SidewalkUpdate`

Sidewalk update.

_Required_: No

_Type_: [SidewalkUpdateAccount](aws-properties-iotwireless-partneraccount-sidewalkupdateaccount.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags are an array of key-value pairs to attach to the specified resource. Tags can
have a minimum of 0 and a maximum of 50 items.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotwireless-partneraccount-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the partner account.

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the resource.

`Fingerprint`

The fingerprint of the Sidewalk application server private key.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TraceContent

SidewalkAccountInfo

All content copied from https://docs.aws.amazon.com/.
