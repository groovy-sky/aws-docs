This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VoiceID::Domain

###### Important

End of support notice: On May 20, 2026, AWS will end support for
Amazon Connect Voice ID. After May 20, 2026, you will no longer be able to access Voice ID on the Amazon Connect
console, access Voice ID features on the Amazon Connect admin website or Contact Control Panel, or access
Voice ID resources. For more information, visit
[Amazon Connect Voice ID end of support](../../../connect/latest/adminguide/amazonconnect-voiceid-end-of-support.md).

Creates a domain that contains all Amazon Connect Voice ID data, such as speakers, fraudsters,
customer audio, and voiceprints. Every domain is created with a default watchlist that
fraudsters can be a part of.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VoiceID::Domain",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "ServerSideEncryptionConfiguration" : ServerSideEncryptionConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VoiceID::Domain
Properties:
  Description: String
  Name: String
  ServerSideEncryptionConfiguration:
    ServerSideEncryptionConfiguration
  Tags:
    - Tag

```

## Properties

`Description`

The description of the domain.

_Required_: No

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-%@]*)$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideEncryptionConfiguration`

The server-side encryption configuration containing the KMS key
identifier you want Voice ID to use to encrypt your data.

_Required_: Yes

_Type_: [ServerSideEncryptionConfiguration](aws-properties-voiceid-domain-serversideencryptionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-voiceid-domain-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `DomainId` of the domain.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainId`

The identifier of the domain.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Connect Voice ID

ServerSideEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
