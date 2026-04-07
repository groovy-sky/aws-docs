This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary

Creates or updates a canary. Canaries are scripts that monitor your endpoints and APIs from the
outside-in. Canaries help you check the availability and latency of your web services and
troubleshoot anomalies by investigating load time data, screenshots of the UI, logs, and
metrics. You can set up a canary to run continuously or just once.

Canaries are automated scripts that run at specified intervals against an endpoint. They include Python
or Node.js code to create a Lambda function. This code needs to be packaged in a certain way, depending
on the language. For more information, see [Writing a canary script](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-canaries-writingcanary.md).

To create canaries, you must have the `CloudWatchSyntheticsFullAccess` policy.
If you are creating a new IAM role for the canary, you also need the
the `iam:CreateRole`, `iam:CreatePolicy` and
`iam:AttachRolePolicy` permissions. For more information, see [Necessary\
Roles and Permissions](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-canaries-roles.md).

Do not include secrets or proprietary information in your canary names. The canary name
makes up part of the Amazon Resource Name (ARN) for the canary, and the ARN is included in
outbound calls over the internet. For more information, see [Security\
Considerations for Synthetics Canaries](../../../amazoncloudwatch/latest/monitoring/servicelens-canaries-security.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Synthetics::Canary",
  "Properties" : {
      "ArtifactConfig" : ArtifactConfig,
      "ArtifactS3Location" : String,
      "BrowserConfigs" : [ BrowserConfig, ... ],
      "Code" : Code,
      "DryRunAndUpdate" : Boolean,
      "ExecutionRoleArn" : String,
      "FailureRetentionPeriod" : Integer,
      "Name" : String,
      "ProvisionedResourceCleanup" : String,
      "ResourcesToReplicateTags" : [ String, ... ],
      "RunConfig" : RunConfig,
      "RuntimeVersion" : String,
      "Schedule" : Schedule,
      "StartCanaryAfterCreation" : Boolean,
      "SuccessRetentionPeriod" : Integer,
      "Tags" : [ Tag, ... ],
      "VisualReferences" : [ VisualReference, ... ],
      "VPCConfig" : VPCConfig
    }
}

```

### YAML

```yaml

Type: AWS::Synthetics::Canary
Properties:
  ArtifactConfig:
    ArtifactConfig
  ArtifactS3Location: String
  BrowserConfigs:
    - BrowserConfig
  Code:
    Code
  DryRunAndUpdate: Boolean
  ExecutionRoleArn: String
  FailureRetentionPeriod: Integer
  Name: String
  ProvisionedResourceCleanup: String
  ResourcesToReplicateTags:
    - String
  RunConfig:
    RunConfig
  RuntimeVersion: String
  Schedule:
    Schedule
  StartCanaryAfterCreation: Boolean
  SuccessRetentionPeriod: Integer
  Tags:
    - Tag
  VisualReferences:
    - VisualReference
  VPCConfig:
    VPCConfig

```

## Properties

`ArtifactConfig`

A structure that contains the configuration for canary artifacts, including
the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.

_Required_: No

_Type_: [ArtifactConfig](aws-properties-synthetics-canary-artifactconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ArtifactS3Location`

The location in Amazon S3 where Synthetics stores artifacts from the runs of this
canary. Artifacts include the log file, screenshots, and HAR files.
Specify the full location path, including `s3://` at the beginning
of the path.

_Required_: Yes

_Type_: String

_Pattern_: `^(s3|S3)://`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BrowserConfigs`

A structure that specifies the browser type to use for a canary run. CloudWatch Synthetics supports running canaries on both `CHROME` and `FIREFOX` browsers.

###### Note

If not specified, `browserConfigs` defaults to Chrome.

_Required_: No

_Type_: Array of [BrowserConfig](aws-properties-synthetics-canary-browserconfig.md)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Code`

Use this structure to input your script code for the canary. This structure contains the
Lambda handler with the location where the canary should start running the script. If the
script is stored in an S3 bucket, the bucket name, key, and version are also included. If
the script is passed into the canary directly, the script code is contained in the value
of `Script`.

_Required_: Yes

_Type_: [Code](aws-properties-synthetics-canary-code.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DryRunAndUpdate`

Specifies whether to perform a dry run before updating the canary. If set to `true`, CloudFormation will execute a dry run to validate the changes before applying them to the canary.
If the dry run succeeds, the canary will be updated with the changes. If the dry run fails, the CloudFormation deployment will fail with the dry run’s failure reason.

If set to `false` or omitted, the canary will be updated directly without first performing a dry run. The default value is `false`.

For more information, see [Performing safe canary updates](../../../amazoncloudwatch/latest/monitoring/performing-safe-canary-upgrades.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The ARN of the IAM role to be used to run the canary. This role must already exist,
and must include `lambda.amazonaws.com` as a principal in the trust
policy. The role must also have the following permissions:

- `s3:PutObject`

- `s3:GetBucketLocation`

- `s3:ListAllMyBuckets`

- `cloudwatch:PutMetricData`

- `logs:CreateLogGroup`

- `logs:CreateLogStream`

- `logs:PutLogEvents`

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws[a-zA-Z-]*)?:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailureRetentionPeriod`

The number of days to retain data about failed runs of this canary. If you omit
this field, the default of 31 days is used. The valid range is 1 to 455 days.

This setting affects the range of information returned by [GetCanaryRuns](../../../../reference/amazonsynthetics/latest/apireference/api-getcanaryruns.md), as well as
the range of information displayed in the Synthetics console.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for this canary. Be sure to give it a descriptive name
that distinguishes it from other canaries in your account.

Do not include secrets or proprietary information in your canary names. The canary name
makes up part of the canary ARN, and the ARN is included in outbound calls over the
internet. For more information, see [Security\
Considerations for Synthetics Canaries](../../../amazoncloudwatch/latest/monitoring/servicelens-canaries-security.md).

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-z_\-]{1,255}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisionedResourceCleanup`

Specifies whether to also delete the Lambda functions and layers used by this canary
when the canary is deleted. If it is `AUTOMATIC`, the Lambda functions and layers will be deleted
when the canary is deleted.

If the value of this parameter is `OFF`, then the value of the `DeleteLambda` parameter
of the [DeleteCanary](../../../../reference/amazonsynthetics/latest/apireference/api-deletecanary.md) operation
determines whether the Lambda functions and layers will be deleted.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | OFF`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcesToReplicateTags`

To have the tags that you apply to this canary also be applied to the Lambda function that the canary uses, specify this property with the value `lambda-function`. If
you do this, CloudWatch Synthetics will keep the tags of the canary and the Lambda function synchronized. Any future changes you make to the canary's tags will also be applied to the function.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunConfig`

A structure that contains input information for a canary run. If you omit
this structure, the
frequency of the canary is used as canary's timeout value, up to a maximum of 900 seconds.

_Required_: No

_Type_: [RunConfig](aws-properties-synthetics-canary-runconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuntimeVersion`

Specifies the runtime version to use for the canary. For more information about
runtime versions, see [Canary Runtime Versions](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-canaries-library.md).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

A structure that contains information about how often the canary is to run, and when
these runs are to stop.

_Required_: Yes

_Type_: [Schedule](aws-properties-synthetics-canary-schedule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartCanaryAfterCreation`

Specify TRUE to have the canary start making runs immediately after it is created.

A canary that you create using CloudFormation can't be used to monitor the
CloudFormation stack that creates the canary or to roll back that stack if there is a failure.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuccessRetentionPeriod`

The number of days to retain data about successful runs of this canary. If you omit
this field, the default of 31 days is used. The valid range is 1 to 455 days.

This setting affects the range of information returned by [GetCanaryRuns](../../../../reference/amazonsynthetics/latest/apireference/api-getcanaryruns.md), as well as
the range of information displayed in the Synthetics console.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of key-value pairs that are associated with the canary.

_Required_: No

_Type_: Array of [Tag](aws-properties-synthetics-canary-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualReferences`

A list of visual reference configurations for the canary, one for each browser type that the canary is configured to run on. Visual references are used for visual monitoring comparisons.

`syn-nodejs-puppeteer-11.0` and above, and `syn-nodejs-playwright-3.0` and above, only supports `visualReferences`. `visualReference` field is not supported.

Versions older than `syn-nodejs-puppeteer-11.0` supports both `visualReference` and `visualReferences` for backward compatibility. It is recommended to use `visualReferences`
for consistency and future compatibility.

_Required_: No

_Type_: Array of [VisualReference](aws-properties-synthetics-canary-visualreference.md)

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VPCConfig`

If this canary is to test an endpoint in a VPC, this structure contains
information about the subnet and security groups of the VPC endpoint.
For more information, see [Running a Canary in a VPC](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-canaries-vpc.md).

_Required_: No

_Type_: [VPCConfig](aws-properties-synthetics-canary-vpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the canary, such as `MyCanary`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Code.SourceLocationArn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the Lambda layer where Synthetics stores the canary script code.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`Id`

The ID of the canary.

`State`

The state of the canary. For example, `RUNNING`.

## Examples

- [Canary with script stored in an Amazon S3 bucket](#aws-resource-synthetics-canary--examples--Canary_with_script_stored_in_an_Amazon_S3_bucket)

- [Canary with script passed through CloudFormation](#aws-resource-synthetics-canary--examples--Canary_with_script_passed_through_CloudFormation)

### Canary with script stored in an Amazon S3 bucket

This example creates a canary that uses an existing script stored in
an S3 bucket. The canary is started as soon as it is created.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "CloudFormation Sample Template for CloudWatch Synthetics: Create a Canary using this template",
    "Resources": {
        "SyntheticsCanary": {
            "Type": "AWS::Synthetics::Canary",
            "Properties": {
                "Name": {
                    "Ref": "samplecanary"
                },
                "ExecutionRoleArn": {
                    "Ref": "arn:aws:iam::123456789012:role/my-lambda-execution-role-to-run-canary"
                },
                "Code": {
                    "Handler": "pageLoadBlueprint.handler",
                    "S3Bucket": "aws-synthetics-code-myaccount-canary1",
                    "S3Key": "my-script-location"
                },
                "ArtifactS3Location": "s3://my-results-bucket",
                "RuntimeVersion": "syn-nodejs-puppeteer-6.2",
                "Schedule": {
                    "Expression": "rate(1 minute)",
                    "DurationInSeconds": 3600
                },
                "RunConfig": {
                    "TimeoutInSeconds": 60
                },
                "FailureRetentionPeriod": 30,
                "SuccessRetentionPeriod": 30,
                "StartCanaryAfterCreation": true,
                "Tags": [
                    {
                        "Key": "key00AtCreate",
                        "Value": "value001AtCreate"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
    SyntheticsCanary:
        Type: 'AWS::Synthetics::Canary'
        Properties:
            Name: samplecanary
            ExecutionRoleArn: 'arn:aws:iam::123456789012:role/my-lambda-execution-role-to-run-canary'
            Code: {Handler: pageLoadBlueprint.handler, S3Bucket: aws-synthetics-code-myaccount-canary1, S3Key: my-script-location}
            ArtifactS3Location: s3://my-results-bucket
            RuntimeVersion: syn-nodejs-puppeteer-6.2
            Schedule: {Expression: 'rate(1 minute)', DurationInSeconds: 3600}
            RunConfig: {TimeoutInSeconds: 60}
            FailureRetentionPeriod: 30
            SuccessRetentionPeriod: 30
            Tags: [{Key: key00AtCreate, Value: value001AtCreate}]
            StartCanaryAfterCreation: true
```

### Canary with script passed through CloudFormation

This example creates a canary and passes the script code directly
into the canary.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "CloudFormation Sample Template for CloudWatch Synthetics: Create a Canary using this template",
    "Resources": {
        "SyntheticsCanary": {
            "Type": "AWS::Synthetics::Canary",
            "Properties": {
                "Name": {
                    "Ref": "samplecanary"
                },
                "ExecutionRoleArn": {
                    "Ref": "arn:aws:iam::123456789012:role/my-lambda-execution-role-to-run-canary"
                },
                "Code": {
                    "Handler": "pageLoadBlueprint.handler",
                    "Script": "var synthetics = require('Synthetics');\nconst log = require('SyntheticsLogger');\n\nconst pageLoadBlueprint = async function () {\n\n    // INSERT URL here\n    const URL = \"https://amazon.com\";\n\n    let page = await synthetics.getPage();\n    const response = await page.goto(URL, {waitUntil: 'domcontentloaded', timeout: 30000});\n    //Wait for page to render.\n    //Increase or decrease wait time based on endpoint being monitored.\n    await page.waitFor(15000);\n    await synthetics.takeScreenshot('loaded', 'loaded');\n    let pageTitle = await page.title();\n    log.info('Page title: ' + pageTitle);\n    if (response.status() !== 200) {\n        throw \"Failed to load page!\";\n    }\n};\n\nexports.handler = async () => {\n    return await pageLoadBlueprint();\n};\n"
                },
                "ArtifactS3Location": "s3://my-results-bucket",
                "RuntimeVersion": "syn-nodejs-puppeteer-6.2",
                "Schedule": {
                    "Expression": "rate(1 minute)",
                    "DurationInSeconds": 3600
                },
                "RunConfig": {
                    "TimeoutInSeconds": 60
                },
                "FailureRetentionPeriod": 30,
                "SuccessRetentionPeriod": 30,
                "StartCanaryAfterCreation": false,
                "Tags": [
                    {
                        "Id": "key00AtCreate",
                        "Value": "value001AtCreate"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
    SyntheticsCanary:
        Type: 'AWS::Synthetics::Canary'
        Properties:
            Name: samplecanary
            ExecutionRoleArn: 'arn:aws:iam::123456789012:role/my-lambda-execution-role-to-run-canary'
            Code: {Handler: pageLoadBlueprint.handler, Script: "var synthetics = require('Synthetics');\nconst log = require('SyntheticsLogger');\nconst pageLoadBlueprint = async function () {\n// INSERT URL here\nconst URL = \"https://amazon.com\";\n\nlet page = await synthetics.getPage();\nconst response = await page.goto(URL, {waitUntil: 'domcontentloaded', timeout: 30000});\n//Wait for page to render.\n//Increase or decrease wait time based on endpoint being monitored.\nawait page.waitFor(15000);\nawait synthetics.takeScreenshot('loaded', 'loaded');\nlet pageTitle = await page.title();\nlog.info('Page title: ' + pageTitle);\nif (response.status() !== 200) {\n     throw \"Failed to load page!\";\n}\n};\n\nexports.handler = async () => {\nreturn await pageLoadBlueprint();\n};\n"}
            ArtifactS3Location: s3://my-results-bucket
            RuntimeVersion: syn-nodejs-puppeteer-6.2
            Schedule: {Expression: 'rate(1 minute)', DurationInSeconds: 3600}
            RunConfig: {TimeoutInSeconds: 60}
            FailureRetentionPeriod: 30
            SuccessRetentionPeriod: 30
            Tags: [{Key: key00AtCreate, Value: value001AtCreate}]
            StartCanaryAfterCreation: false
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatch Synthetics

ArtifactConfig
