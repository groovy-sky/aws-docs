This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment S3AccessPoint

Describes the S3 access point configuration of the S3 access point attachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Alias" : String,
  "Policy" : Json,
  "ResourceARN" : String,
  "VpcConfiguration" : S3AccessPointVpcConfiguration
}

```

### YAML

```yaml

  Alias: String
  Policy: Json
  ResourceARN: String
  VpcConfiguration:
    S3AccessPointVpcConfiguration

```

## Properties

`Alias`

The S3 access point's alias.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-z\\-]{1,63}`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The S3 access point's policy.

_Required_: No

_Type_: Json

_Minimum_: `1`

_Maximum_: `200000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceARN`

The S3 access point's ARN.

_Required_: No

_Type_: String

_Pattern_: `^arn:[^:]{1,63}:[^:]{0,63}:[^:]{0,63}:(?:|\d{12}):[^/].{0,1023}$`

_Minimum_: `8`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConfiguration`

The S3 access point's virtual private cloud (VPC) configuration.

_Required_: No

_Type_: [S3AccessPointVpcConfiguration](aws-properties-fsx-s3accesspointattachment-s3accesspointvpcconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenZFSPosixFileSystemUser

S3AccessPointOntapConfiguration

All content copied from https://docs.aws.amazon.com/.
