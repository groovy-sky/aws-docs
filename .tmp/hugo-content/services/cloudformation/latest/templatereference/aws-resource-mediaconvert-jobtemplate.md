This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConvert::JobTemplate

The AWS::MediaConvert::JobTemplate resource is an AWS Elemental MediaConvert resource
type that you can use to generate transcoding jobs.

When you declare this entity in your CloudFormation template, you pass in your
transcoding job settings in JSON or YAML format. This settings specification must be
formed in a particular way that conforms to AWS Elemental MediaConvert job validation. For
more information about creating a job template model for the `SettingsJson`
property, see the Remarks section later in this topic.

For information about job templates,
see [Working with AWS Elemental MediaConvert Job Templates](../../../mediaconvert/latest/ug/working-with-job-templates.md) in the _AWS Elemental MediaConvert User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConvert::JobTemplate",
  "Properties" : {
      "AccelerationSettings" : AccelerationSettings,
      "Category" : String,
      "Description" : String,
      "HopDestinations" : [ HopDestination, ... ],
      "Name" : String,
      "Priority" : Integer,
      "Queue" : String,
      "SettingsJson" : Json,
      "StatusUpdateInterval" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaConvert::JobTemplate
Properties:
  AccelerationSettings:
    AccelerationSettings
  Category: String
  Description: String
  HopDestinations:
    - HopDestination
  Name: String
  Priority: Integer
  Queue: String
  SettingsJson:
    Json
  StatusUpdateInterval: String
  Tags:
    - Tag

```

## Properties

`AccelerationSettings`

Accelerated transcoding can significantly speed up jobs with long, visually complex
content. Outputs that use this feature incur pro-tier pricing. For information about
feature limitations, For more information, see [Job Limitations for Accelerated Transcoding in AWS Elemental MediaConvert](../../../mediaconvert/latest/ug/job-requirements.md) in the _AWS Elemental MediaConvert User Guide_.

_Required_: No

_Type_: [AccelerationSettings](aws-properties-mediaconvert-jobtemplate-accelerationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Category`

Optional. A category for the job template you are creating

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Optional. A description of the job template you are creating.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HopDestinations`

Optional. Configuration for a destination queue to which the job can hop once a
customer-defined minimum wait time has passed. For more information, see [Setting Up Queue Hopping to Avoid Long Waits](../../../mediaconvert/latest/ug/setting-up-queue-hopping-to-avoid-long-waits.md) in the _AWS Elemental MediaConvert User Guide_.

_Required_: No

_Type_: Array of [HopDestination](aws-properties-mediaconvert-jobtemplate-hopdestination.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the output group

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Priority`

Specify the relative priority for this job. In any given queue, the service begins
processing the job with the highest value first. When more than one job has the same
priority, the service begins processing the job that you submitted first. If you don't
specify a priority, the service uses the default value 0. Minimum: -50 Maximum:
50

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Queue`

Optional. The queue that jobs created from this template are assigned to. Specify the
Amazon Resource Name (ARN) of the queue. For example,
arn:aws:mediaconvert:us-west-2:505474453218:queues/Default. If you don't specify this,
jobs will go to the default queue.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SettingsJson`

Specify, in JSON format, the transcoding job settings for this job template. This
specification must conform to the AWS Elemental MediaConvert job validation. For
information about forming this specification, see the Remarks section later in this
topic.

For more information about MediaConvert job templates, see [Working with AWS Elemental MediaConvert Job Templates](../../../mediaconvert/latest/ug/working-with-job-templates.md) in the
_AWS Elemental MediaConvert User Guide_.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusUpdateInterval`

Specify how often MediaConvert sends STATUS\_UPDATE events to Amazon CloudWatch Events.
Set the interval, in seconds, between status updates. MediaConvert sends an update at
this interval from the time the service begins processing your job to the time it
completes the transcode or encounters an error.

Specify one of the following enums:

SECONDS\_10

SECONDS\_12

SECONDS\_15

SECONDS\_20

SECONDS\_30

SECONDS\_60

SECONDS\_120

SECONDS\_180

SECONDS\_240

SECONDS\_300

SECONDS\_360

SECONDS\_420

SECONDS\_480

SECONDS\_540

SECONDS\_600

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::MediaConvert::JobTemplate`
resource to the intrinsic `Ref` function, the function returns the name of
the job template, such as `Streaming stack DASH`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the job template, such as
`arn:aws:mediaconvert:us-west-2:123456789012`.

`Name`

The name of the job template, such as `Streaming stack DASH`.

## Remarks

_Creating a Job Template Model for the SettingsJson Property_

When you declare an AWS::MediaConvert::JobTemplate entity in your CloudFormation template, you pass in your transcoding job settings as the value
for the property `SettingsJson`. This settings specification must be in
in JSON or YAML format and must conform to AWS Elemental MediaConvert job
validation.

The following procedure is for generating the specification in JSON. If you need
it in YAML, you can create it in JSON and use a conversion utility.

**To create a JSON job template model for**
**`SettingsJson`**

1. Create a job template using the MediaConvert [https://console.aws.amazon.com/mediaconvert/](https://console.aws.amazon.com/mediaconvert). For more information, see [Working\
    with AWS Elemental MediaConvert Job Templates](../../../mediaconvert/latest/ug/working-with-job-templates.md).

2. Use the AWS CLI to get just the settings structure using the
    following command:

`aws mediaconvert
                               https://abcd1234.mediaconvert.region-name-1.amazonaws.com
                               get-job-template
                               DASH-stack-template --query
                               'JobTemplate.Settings'`

3. Copy the settings as the value for the property
    `SettingsJson`.

For an example job template model in JSON and YAML, see the Examples section of
this topic.

## Examples

### Job Template Model for SettingsJson

For more information about creating a job template model in JSON or YAML for
the `SettingsJson` property, see the Remarks section of this
topic.

#### JSON

```json

{
    "AdAvailOffset": 0,
    "OutputGroups": [
        {
            "Name": "File Group",
            "OutputGroupSettings": {
                "FileGroupSettings": {},
                "Type": "FILE_GROUP_SETTINGS"
            },
            "Outputs": [
                {
                    "Extension": "mp4",
                    "NameModifier": "_Generic_Uhd_Mp4_Hevc_Aac_16x9_3840x2160p_24Hz_8Mbps",
                    "Preset": "System-Generic_Uhd_Mp4_Hevc_Aac_16x9_3840x2160p_24Hz_8Mbps"
                },
                {
                    "Extension": "mp4",
                    "NameModifier": "_Generic_Hd_Mp4_Hevc_Aac_16x9_1920x1080p_24Hz_4.5Mbps",
                    "Preset": "System-Generic_Hd_Mp4_Hevc_Aac_16x9_1920x1080p_24Hz_4.5Mbps"
                },
                {
                    "Extension": "mp4",
                    "NameModifier": "_Generic_Hd_Mp4_Hevc_Aac_16x9_1280x720p_24Hz_3.0Mbps",
                    "Preset": "System-Generic_Hd_Mp4_Hevc_Aac_16x9_1280x720p_24Hz_3.0Mbps"
                },
                {
                    "Extension": "mp4",
                    "NameModifier": "_Generic_Hd_Mp4_Avc_Aac_16x9_1920x1080p_24Hz_6Mbps",
                    "Preset": "System-Generic_Hd_Mp4_Avc_Aac_16x9_1920x1080p_24Hz_6Mbps"
                },
                {
                    "Extension": "mp4",
                    "NameModifier": "_Generic_Hd_Mp4_Avc_Aac_16x9_1280x720p_24Hz_4.5Mbps",
                    "Preset": "System-Generic_Hd_Mp4_Avc_Aac_16x9_1280x720p_24Hz_4.5Mbps"
                },
                {
                    "Extension": "mp4",
                    "NameModifier": "_Generic_Sd_Mp4_Avc_Aac_4x3_640x480p_24Hz_1.5Mbps",
                    "Preset": "System-Generic_Sd_Mp4_Avc_Aac_4x3_640x480p_24Hz_1.5Mbps"
                }
            ]
        }
    ]
}

```

#### YAML

```yaml

---
AdAvailOffset: 0
OutputGroups:
- Name: File Group
  OutputGroupSettings:
    FileGroupSettings: {}
    Type: FILE_GROUP_SETTINGS
  Outputs:
  - Extension: mp4
    NameModifier: _Generic_Uhd_Mp4_Hevc_Aac_16x9_3840x2160p_24Hz_8Mbps
    Preset: System-Generic_Uhd_Mp4_Hevc_Aac_16x9_3840x2160p_24Hz_8Mbps
  - Extension: mp4
    NameModifier: _Generic_Hd_Mp4_Hevc_Aac_16x9_1920x1080p_24Hz_4.5Mbps
    Preset: System-Generic_Hd_Mp4_Hevc_Aac_16x9_1920x1080p_24Hz_4.5Mbps
  - Extension: mp4
    NameModifier: _Generic_Hd_Mp4_Hevc_Aac_16x9_1280x720p_24Hz_3.0Mbps
    Preset: System-Generic_Hd_Mp4_Hevc_Aac_16x9_1280x720p_24Hz_3.0Mbps
  - Extension: mp4
    NameModifier: _Generic_Hd_Mp4_Avc_Aac_16x9_1920x1080p_24Hz_6Mbps
    Preset: System-Generic_Hd_Mp4_Avc_Aac_16x9_1920x1080p_24Hz_6Mbps
  - Extension: mp4
    NameModifier: _Generic_Hd_Mp4_Avc_Aac_16x9_1280x720p_24Hz_4.5Mbps
    Preset: System-Generic_Hd_Mp4_Avc_Aac_16x9_1280x720p_24Hz_4.5Mbps
  - Extension: mp4
    NameModifier: _Generic_Sd_Mp4_Avc_Aac_4x3_640x480p_24Hz_1.5Mbps
    Preset: System-Generic_Sd_Mp4_Avc_Aac_4x3_640x480p_24Hz_1.5Mbps
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Elemental MediaConvert

AccelerationSettings

All content copied from https://docs.aws.amazon.com/.
