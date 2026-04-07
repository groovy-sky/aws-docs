This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate PresignedUrlConfig

Configuration for pre-signed S3 URLs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExpiresInSec" : Integer,
  "RoleArn" : String
}

```

### YAML

```yaml

  ExpiresInSec: Integer
  RoleArn: String

```

## Properties

`ExpiresInSec`

How long (in seconds) pre-signed URLs are valid. Valid values are 60 - 3600, the
default value is 3600 seconds. Pre-signed URLs are generated when Jobs receives an MQTT
request for the job document.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `3600`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of an IAM role that grants grants permission to download files from the S3
bucket where the job data/updates are stored. The role must also grant permission for IoT
to download the files.

###### Important

For information about addressing the confused deputy problem, see [cross-service confused deputy prevention](https://docs.aws.amazon.com/iot/latest/developerguide/cross-service-confused-deputy-prevention.html) in the _AWS IoT Core developer guide_.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MaintenanceWindow

RateIncreaseCriteria
