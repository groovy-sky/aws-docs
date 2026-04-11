This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Profile

Creates the local or partner profile to use for AS2 transfers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Transfer::Profile",
  "Properties" : {
      "As2Id" : String,
      "CertificateIds" : [ String, ... ],
      "ProfileType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Transfer::Profile
Properties:
  As2Id: String
  CertificateIds:
    - String
  ProfileType: String
  Tags:
    - Tag

```

## Properties

`As2Id`

The `As2Id` is the _AS2-name_, as defined in the
[RFC 4130](https://datatracker.ietf.org/doc/html/rfc4130). For inbound transfers, this is the `AS2-From` header for the AS2 messages
sent from the partner. For outbound connectors, this is the `AS2-To` header for the
AS2 messages sent to the partner using the `StartFileTransfer` API operation. This ID cannot include spaces.

_Required_: Yes

_Type_: String

_Pattern_: `^[\u0020-\u007E\s]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateIds`

An array of identifiers for the imported certificates. You use this identifier for working with profiles and partner profiles.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProfileType`

Indicates whether to list only `LOCAL` type profiles or only `PARTNER` type profiles.
If not supplied in the request, the command lists all types of profiles.

_Required_: Yes

_Type_: String

_Allowed values_: `LOCAL | PARTNER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key-value pairs that can be used to group and search for profiles.

_Required_: No

_Type_: Array of [Tag](aws-properties-transfer-profile-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name associated with the profile, in the form
`arn:aws:transfer:region:account-id:profile/profile-id/` .

`ProfileId`

The unique identifier for the AS2 profile, returned after the API call
succeeds.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
