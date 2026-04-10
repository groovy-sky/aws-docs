This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate

Represents a job template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::JobTemplate",
  "Properties" : {
      "AbortConfig" : AbortConfig,
      "Description" : String,
      "DestinationPackageVersions" : [ String, ... ],
      "Document" : String,
      "DocumentSource" : String,
      "JobArn" : String,
      "JobExecutionsRetryConfig" : JobExecutionsRetryConfig,
      "JobExecutionsRolloutConfig" : JobExecutionsRolloutConfig,
      "JobTemplateId" : String,
      "MaintenanceWindows" : [ MaintenanceWindow, ... ],
      "PresignedUrlConfig" : PresignedUrlConfig,
      "Tags" : [ Tag, ... ],
      "TimeoutConfig" : TimeoutConfig
    }
}

```

### YAML

```yaml

Type: AWS::IoT::JobTemplate
Properties:
  AbortConfig:
    AbortConfig
  Description: String
  DestinationPackageVersions:
    - String
  Document: String
  DocumentSource: String
  JobArn: String
  JobExecutionsRetryConfig:
    JobExecutionsRetryConfig
  JobExecutionsRolloutConfig:
    JobExecutionsRolloutConfig
  JobTemplateId: String
  MaintenanceWindows:
    - MaintenanceWindow
  PresignedUrlConfig:
    PresignedUrlConfig
  Tags:
    - Tag
  TimeoutConfig:
    TimeoutConfig

```

## Properties

`AbortConfig`

The criteria that determine when and how a job abort takes place.

_Required_: No

_Type_: [AbortConfig](aws-properties-iot-jobtemplate-abortconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of the job template.

_Required_: Yes

_Type_: String

_Pattern_: `[^\p{C}]+`

_Maximum_: `2028`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPackageVersions`

The package version Amazon Resource Names (ARNs) that are installed on the device’s
reserved named shadow ( `$package`) when the job successfully completes.

**Note:** Up to 25 package version ARNS are allowed.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Document`

The job document.

Required if you don't specify a value for `documentSource`.

_Required_: No

_Type_: String

_Maximum_: `32768`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DocumentSource`

An S3 link, or S3 object URL, to the job document. The link is an Amazon S3 object URL
and is required if you don't specify a value for `document`.

For example, `--document-source
                https://s3.region-code.amazonaws.com/example-firmware/device-firmware.1.0`

For more information, see [Methods for accessing a\
bucket](../../../s3/latest/userguide/access-bucket-intro.md).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1350`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobArn`

The ARN of the job to use as the basis for the job template.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobExecutionsRetryConfig`

Allows you to create the criteria to retry a job.

_Required_: No

_Type_: [JobExecutionsRetryConfig](aws-properties-iot-jobtemplate-jobexecutionsretryconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobExecutionsRolloutConfig`

Allows you to create a staged rollout of a job.

_Required_: No

_Type_: [JobExecutionsRolloutConfig](aws-properties-iot-jobtemplate-jobexecutionsrolloutconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`JobTemplateId`

A unique identifier for the job template. We recommend using a UUID. Alpha-numeric
characters, "-", and "\_" are valid for use here.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaintenanceWindows`

An optional configuration within the SchedulingConfig to setup a recurring maintenance
window with a predetermined start time and duration for the rollout of a job document to
all devices in a target group for a job.

_Required_: No

_Type_: Array of [MaintenanceWindow](aws-properties-iot-jobtemplate-maintenancewindow.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PresignedUrlConfig`

Configuration for pre-signed S3 URLs.

_Required_: No

_Type_: [PresignedUrlConfig](aws-properties-iot-jobtemplate-presignedurlconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that can be used to manage the job template.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-jobtemplate-tag.md)

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeoutConfig`

Specifies the amount of time each device has to finish its execution of the job. A timer
is started when the job execution status is set to `IN_PROGRESS`. If the job
execution status is not set to another terminal state before the timer expires, it will be
automatically set to `TIMED_OUT`.

_Required_: No

_Type_: [TimeoutConfig](aws-properties-iot-jobtemplate-timeoutconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the job template Id. For example:

`{ "Ref": "MyJobTemplate-12345" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the job to use as the basis for the job template.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AbortConfig

All content copied from https://docs.aws.amazon.com/.
