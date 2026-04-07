This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessInstance VerifiedAccessTrustProvider

A trust provider is a third-party entity that creates, maintains, and manages identity
information for users and devices. When an application request is made, the identity
information sent by the trust provider is evaluated by Verified Access before allowing or
denying the application request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "DeviceTrustProviderType" : String,
  "TrustProviderType" : String,
  "UserTrustProviderType" : String,
  "VerifiedAccessTrustProviderId" : String
}

```

### YAML

```yaml

  Description: String
  DeviceTrustProviderType: String
  TrustProviderType: String
  UserTrustProviderType: String
  VerifiedAccessTrustProviderId: String

```

## Properties

`Description`

A description for the AWS Verified Access trust provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceTrustProviderType`

The type of device-based trust provider.

_Required_: No

_Type_: String

_Allowed values_: `jamf | crowdstrike | jumpcloud`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustProviderType`

The type of Verified Access trust provider.

_Required_: No

_Type_: String

_Allowed values_: `user | device`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserTrustProviderType`

The type of user-based trust provider.

_Required_: No

_Type_: String

_Allowed values_: `iam-identity-center | oidc`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerifiedAccessTrustProviderId`

The ID of the AWS Verified Access trust provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VerifiedAccessLogs

AWS::EC2::VerifiedAccessTrustProvider
