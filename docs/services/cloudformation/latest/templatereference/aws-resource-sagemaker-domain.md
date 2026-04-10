This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain

Creates a `Domain`. A domain consists of an associated Amazon Elastic File System
volume, a list of authorized users, and a variety of security, application, policy, and
Amazon Virtual Private Cloud (VPC) configurations. Users within a domain can share notebook files
and other artifacts with each other.

**EFS storage**

When a domain is created, an EFS volume is created for use by all of the users within the
domain. Each user receives a private home directory within the EFS volume for notebooks, Git
repositories, and data files.

SageMaker AI uses the AWS Key Management Service (AWS
KMS) to encrypt the EFS volume attached to the domain with an AWS managed key
by default. For more control, you can specify a customer managed key. For more information,
see [Protect Data\
at Rest Using Encryption](../../../sagemaker/latest/dg/encryption-at-rest.md).

**VPC configuration**

All traffic between the domain and the Amazon EFS volume is through the specified
VPC and subnets. For other traffic, you can specify the `AppNetworkAccessType`
parameter. `AppNetworkAccessType` corresponds to the network access type that you
choose when you onboard to the domain. The following options are available:

- `PublicInternetOnly` \- Non-EFS traffic goes through a VPC managed by
Amazon SageMaker AI, which allows internet access. This is the default value.

- `VpcOnly` \- All traffic is through the specified VPC and subnets. Internet
access is disabled by default. To allow internet access, you must specify a NAT
gateway.

When internet access is disabled, you won't be able to run a Amazon SageMaker AI
Studio notebook or to train or host models unless your VPC has an interface endpoint to
the SageMaker AI API and runtime or a NAT gateway and your security groups allow
outbound connections.

###### Important

NFS traffic over TCP on port 2049 needs to be allowed in both inbound and outbound rules
in order to launch a Amazon SageMaker AI Studio app successfully.

For more information, see [Connect Amazon SageMaker AI Studio Notebooks to Resources in a VPC](../../../sagemaker/latest/dg/studio-notebooks-and-internet-access.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Domain",
  "Properties" : {
      "AppNetworkAccessType" : String,
      "AppSecurityGroupManagement" : String,
      "AuthMode" : String,
      "DefaultSpaceSettings" : DefaultSpaceSettings,
      "DefaultUserSettings" : UserSettings,
      "DomainName" : String,
      "DomainSettings" : DomainSettings,
      "KmsKeyId" : String,
      "SubnetIds" : [ String, ... ],
      "TagPropagation" : String,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Domain
Properties:
  AppNetworkAccessType: String
  AppSecurityGroupManagement: String
  AuthMode: String
  DefaultSpaceSettings:
    DefaultSpaceSettings
  DefaultUserSettings:
    UserSettings
  DomainName: String
  DomainSettings:
    DomainSettings
  KmsKeyId: String
  SubnetIds:
    - String
  TagPropagation: String
  Tags:
    - Tag
  VpcId: String

```

## Properties

`AppNetworkAccessType`

Specifies the VPC used for non-EFS traffic. The default value is `PublicInternetOnly`.

- `PublicInternetOnly` \- Non-EFS traffic is through a VPC managed by Amazon SageMaker AI, which
allows direct internet access

- `VpcOnly` \- All Studio traffic is through the specified VPC and subnets

_Valid Values_: `PublicInternetOnly | VpcOnly`

_Required_: No

_Type_: String

_Allowed values_: `PublicInternetOnly | VpcOnly`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppSecurityGroupManagement`

The entity that creates and manages the required security groups for inter-app communication in
`VpcOnly` mode. Required when `CreateDomain.AppNetworkAccessType` is `VpcOnly` and
`DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn` is provided. If setting up the
domain for use with RStudio, this value must be set to `Service`.

_Allowed Values_: `Service` \| `Customer`

_Required_: No

_Type_: String

_Allowed values_: `Service | Customer`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthMode`

The mode of authentication that members use to access the Domain.

_Valid Values_: `SSO | IAM`

_Required_: Yes

_Type_: String

_Allowed values_: `SSO | IAM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultSpaceSettings`

The default settings for shared spaces that users create in the domain.

SageMaker applies these settings only to shared spaces. It doesn't apply them to private
spaces.

_Required_: No

_Type_: [DefaultSpaceSettings](aws-properties-sagemaker-domain-defaultspacesettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultUserSettings`

The default user settings.

_Required_: Yes

_Type_: [UserSettings](aws-properties-sagemaker-domain-usersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The domain name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainSettings`

A collection of settings that apply to the `SageMaker Domain`. These settings are specified through
the `CreateDomain` API call.

_Required_: No

_Type_: [DomainSettings](aws-properties-sagemaker-domain-domainsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

SageMaker uses AWS KMS to encrypt the EFS volume attached to the Domain with an AWS managed customer master key (CMK) by default. For more control, specify a customer managed CMK.

_Length Constraints_: Maximum length of 2048.

_Pattern_: `.*`

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The VPC subnets that Studio uses for communication.

_Length Constraints_: Maximum length of 32.

_Array members_: Minimum number of 1 item. Maximum number of 16 items.

_Pattern_: `[-0-9a-zA-Z]+`

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `32 | 16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagPropagation`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to associated with the Domain. Each tag consists of a key and an optional value. Tag keys must be unique
per resource. Tags are searchable using the Search API.

Tags that you specify for the Domain are also added to all apps that are launched in the Domain.

_Array members_: Minimum number of 0 items. Maximum number of 50 items.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-domain-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcId`

The ID of the Amazon Virtual Private Cloud (Amazon VPC) that Studio uses for communication.

_Length Constraints_: Maximum length of 32.

_Pattern_: `[-0-9a-zA-Z]+`

_Required_: No

_Type_: String

_Pattern_: `[-0-9a-zA-Z]+`

_Maximum_: `32`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Domain ID, such as `d-xxxxxxxxxxxx`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainArn`

The Amazon Resource Name (ARN) of the Domain, such as
`arn:aws:sagemaker:us-west-2:account-id:domain/my-domain-name`.

`DomainId`

The Domain ID.

`HomeEfsFileSystemId`

The ID of the Amazon Elastic File System (EFS) managed by this Domain.

`SecurityGroupIdForDomainBoundary`

The ID of the security group that authorizes traffic between the `RSessionGateway` apps and the
`RStudioServerPro` app.

`SingleSignOnApplicationArn`

The ARN of the application managed by SageMaker in IAM Identity Center. This value is only returned for domains
created after October 1, 2023.

`SingleSignOnManagedApplicationInstanceId`

The IAM Identity Center managed application instance ID.

`Url`

The URL for the Domain.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AppLifecycleManagement

All content copied from https://docs.aws.amazon.com/.
