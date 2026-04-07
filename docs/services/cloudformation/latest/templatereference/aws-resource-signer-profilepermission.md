This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Signer::ProfilePermission

Adds cross-account permissions to a signing profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Signer::ProfilePermission",
  "Properties" : {
      "Action" : String,
      "Principal" : String,
      "ProfileName" : String,
      "ProfileVersion" : String,
      "StatementId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Signer::ProfilePermission
Properties:
  Action: String
  Principal: String
  ProfileName: String
  ProfileVersion: String
  StatementId: String

```

## Properties

`Action`

The AWS Signer action permitted as part of cross-account
permissions.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Principal`

The AWS principal receiving cross-account permissions. This
may be an IAM role or another AWS account ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProfileName`

The human-readable name of the signing profile.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z_]{2,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProfileVersion`

The version of the signing profile.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-zA-Z]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StatementId`

A unique identifier for the cross-account permission statement.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

The StatementId and ProfileName in the form StatementId\|ProfileName

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Signer

AWS::Signer::SigningProfile
