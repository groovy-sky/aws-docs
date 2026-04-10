This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition FSxAuthorizationConfig

The authorization configuration details for Amazon FSx for Windows File Server file
system. See [FSxWindowsFileServerVolumeConfiguration](../../../../reference/amazonecs/latest/apireference/api-fsxwindowsfileservervolumeconfiguration.md) in the _Amazon ECS API_
_Reference_.

For more information and the input format, see [Amazon FSx for Windows File\
Server Volumes](../../../amazonecs/latest/developerguide/wfsx-volumes.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CredentialsParameter" : String,
  "Domain" : String
}

```

### YAML

```yaml

  CredentialsParameter: String
  Domain: String

```

## Properties

`CredentialsParameter`

The authorization credential option to use. The authorization credential options can
be provided using either the Amazon Resource Name (ARN) of an AWS Secrets Manager secret or SSM
Parameter Store parameter. The ARN refers to the stored credentials.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Domain`

A fully qualified domain name hosted by an [AWS Directory Service](../../../directoryservice/latest/admin-guide/directory-microsoft-ad.md) Managed Microsoft AD (Active Directory) or self-hosted
AD on Amazon EC2.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FirelensConfiguration

FSxWindowsFileServerVolumeConfiguration

All content copied from https://docs.aws.amazon.com/.
