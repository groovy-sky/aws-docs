This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Signer::SigningProfile

Creates a signing profile. A signing profile is a code-signing template that can be used to
carry out a pre-defined signing job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Signer::SigningProfile",
  "Properties" : {
      "PlatformId" : String,
      "ProfileName" : String,
      "SignatureValidityPeriod" : SignatureValidityPeriod,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Signer::SigningProfile
Properties:
  PlatformId: String
  ProfileName: String
  SignatureValidityPeriod:
    SignatureValidityPeriod
  Tags:
    - Tag

```

## Properties

`PlatformId`

The ID of a platform that is available for use by a signing profile.

_Required_: Yes

_Type_: String

_Allowed values_: `AWSLambda-SHA384-ECDSA | Notation-OCI-SHA384-ECDSA`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProfileName`

The name of the signing profile.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-zA-Z_]{2,64}$`

_Minimum_: `2`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SignatureValidityPeriod`

The validity period override for any signature generated using this signing
profile. If unspecified, the default is 135 months.

_Required_: No

_Type_: [SignatureValidityPeriod](aws-properties-signer-signingprofile-signaturevalidityperiod.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags associated with the signing profile.

_Required_: No

_Type_: Array of [Tag](aws-properties-signer-signingprofile-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The signing profile ARN.

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the signing profile created.

`ProfileName`

The name of the signing profile created.

`ProfileVersion`

The version of the signing profile created.

`ProfileVersionArn`

The signing profile ARN, including the profile version.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Signer::ProfilePermission

SignatureValidityPeriod

All content copied from https://docs.aws.amazon.com/.
