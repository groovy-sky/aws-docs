This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessTrustProvider

A trust provider is a third-party entity that creates, maintains, and manages identity
information for users and devices. When an application request is made, the identity
information sent by the trust provider is evaluated by Verified Access before allowing or
denying the application request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VerifiedAccessTrustProvider",
  "Properties" : {
      "Description" : String,
      "DeviceOptions" : DeviceOptions,
      "DeviceTrustProviderType" : String,
      "NativeApplicationOidcOptions" : NativeApplicationOidcOptions,
      "OidcOptions" : OidcOptions,
      "PolicyReferenceName" : String,
      "SseSpecification" : SseSpecification,
      "Tags" : [ Tag, ... ],
      "TrustProviderType" : String,
      "UserTrustProviderType" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VerifiedAccessTrustProvider
Properties:
  Description: String
  DeviceOptions:
    DeviceOptions
  DeviceTrustProviderType: String
  NativeApplicationOidcOptions:
    NativeApplicationOidcOptions
  OidcOptions:
    OidcOptions
  PolicyReferenceName: String
  SseSpecification:
    SseSpecification
  Tags:
    - Tag
  TrustProviderType: String
  UserTrustProviderType: String

```

## Properties

`Description`

A description for the AWS Verified Access trust provider.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceOptions`

The options for device-identity trust provider.

_Required_: No

_Type_: [DeviceOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccesstrustprovider-deviceoptions.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeviceTrustProviderType`

The type of device-based trust provider.

_Required_: No

_Type_: String

_Allowed values_: `jamf | crowdstrike | jumpcloud`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NativeApplicationOidcOptions`

The OpenID Connect (OIDC) options.

_Required_: No

_Type_: [NativeApplicationOidcOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccesstrustprovider-nativeapplicationoidcoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OidcOptions`

The options for an OpenID Connect-compatible user-identity trust provider.

_Required_: No

_Type_: [OidcOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccesstrustprovider-oidcoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyReferenceName`

The identifier to be used when working with policy rules.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseSpecification`

The options for additional server side encryption.

_Required_: No

_Type_: [SseSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccesstrustprovider-ssespecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccesstrustprovider-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustProviderType`

The type of Verified Access trust provider.

_Required_: Yes

_Type_: String

_Allowed values_: `user | device`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserTrustProviderType`

The type of user-based trust provider.

_Required_: No

_Type_: String

_Allowed values_: `iam-identity-center | oidc`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Verified Access trust provider.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The creation time.

`LastUpdatedTime`

The last updated time.

`VerifiedAccessTrustProviderId`

The ID of the Verified Access trust provider.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VerifiedAccessTrustProvider

DeviceOptions
