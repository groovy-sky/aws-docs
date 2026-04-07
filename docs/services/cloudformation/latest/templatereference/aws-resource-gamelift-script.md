This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Script

The `AWS::GameLift::Script` resource creates a new script record for your
Realtime Servers script. Realtime scripts are JavaScript that provide configuration settings
and optional custom game logic for your game. The script is deployed when you create a
Realtime Servers fleet to host your game sessions. Script logic is executed during an active
game session.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::Script",
  "Properties" : {
      "Name" : String,
      "NodeJsVersion" : String,
      "StorageLocation" : S3Location,
      "Tags" : [ Tag, ... ],
      "Version" : String
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::Script
Properties:
  Name: String
  NodeJsVersion: String
  StorageLocation:
    S3Location
  Tags:
    - Tag
  Version: String

```

## Properties

`Name`

A descriptive label that is associated with a script. Script names do not need to be
unique.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeJsVersion`

The Node.js version used for execution of your Realtime script. The valid values are
`10.x | 24.x`. By default, `NodeJsVersion` is `10.x`.
This value cannot be updated later.

###### Note

Node.js 10 will reach end of support on September 30, 2026. See more details
in the [Node.js 10 FAQs](https://aws.amazon.com/gamelift/faq/nodejs10). For migration guidance,
see [Migrating from Node.js 10 to 24](https://docs.aws.amazon.com/gamelift/latest/realtimeguide/realtime-script.html#realtime-script-nodejs-migration).

_Required_: No

_Type_: String

_Pattern_: `^\d+\.[x0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageLocation`

The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is
stored. The storage location must specify the Amazon S3 bucket name, the zip file name (the
"key"), and a role ARN that allows Amazon GameLift Servers to access the Amazon S3 storage location. The S3
bucket must be in the same Region where you want to create a new script. By default,
Amazon GameLift Servers uploads the latest version of the zip file; if you have S3 object versioning
turned on, you can use the `ObjectVersion` parameter to specify an earlier
version.

_Required_: Yes

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-script-s3location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of labels to assign to the new script resource. Tags are developer-defined
key-value pairs. Tagging
AWS resources are useful for resource management, access management and cost allocation.
For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the
_AWS General Reference_. Once the resource is created, you can
use TagResource, UntagResource, and
ListTagsForResource to add, remove, and view tags. The
maximum tag limit may be lower than stated. See the AWS General Reference for actual
tagging limits.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-script-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version that is associated with a build or script. Version strings do not need to
be unique.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ScriptId`, such as
`script-1111aaaa-22bb-33cc-44dd-5555eeee66ff`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The unique Amazon Resource Name (ARN) for the script.

`CreationTime`

A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example `"1469498468.057"`).

`Id`

A unique identifier for a Realtime script.

`SizeOnDisk`

The file size of the uploaded Realtime script, expressed in bytes. When files are
uploaded from an S3 location, this value remains at "0".

## Examples

### Create a Realtime Servers Script

The following example creates a GameLift script named `MyRealtimeScript`.
The zipped script files are located in an S3 bucket, specified by the
`S3Bucket` and `S3Key` input parameters. The example also creates
the AWS Identity and Access Management role that GameLift assumes so that it has
permissions to download the script files.

#### JSON

```json

{
    "Resources": {
        "IAMRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": [
                        {
                            "Effect": "Allow",
                            "Principal": {
                                "Service": [
                                    "cloudformation.amazonaws.com",
                                    "gamelift.amazonaws.com"
                                ]
                            },
                            "Action": "sts:AssumeRole"
                        }
                    ]
                },
                "RoleName": "ScriptIAMRole",
                "Policies": [
                    {
                        "PolicyName": "ScriptResourceIAMPolicy",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "s3:GetObject",
                                        "s3:GetObjectVersion",
                                        "s3:GetObjectMetadata",
                                        "s3:*Object*"
                                    ],
                                    "Resource": [
                                        "*"
                                    ]
                                }
                            ]
                        }
                    }
                ]
            }
        },
        "ScriptResource": {
            "Type": "AWS::GameLift::Script",
            "Properties": {
                "Name": "MyRealtimeScript",
                "Version": "v1.0",
                "NodeJsVersion": "24.x",
                "StorageLocation": {
                    "Bucket": "amzn-s3-demo-bucket",
                    "Key": "MyScriptFiles.zip",
                    "RoleArn": {
                        "Fn::GetAtt": [
                            "IAMRole",
                            "Arn"
                        ]
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  IAMRole:
    Type: AWS::IAM::Role
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: "Allow"
            Principal:
              Service: ["cloudformation.amazonaws.com", "gamelift.amazonaws.com"]
            Action: "sts:AssumeRole"
      RoleName: "ScriptIAMRole"
      Policies:
        - PolicyName: ScriptResourceIAMPolicy
          PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Effect: Allow
                Action:
                  - "s3:GetObject"
                  - "s3:GetObjectVersion"
                  - "s3:GetObjectMetadata"
                  - "s3:*Object*"
                Resource:
                  - "*"
  ScriptResource:
    Type: AWS::GameLift::Script
    Properties:
      Name: MyRealtimeScript
      Version: v1.0
      NodeJsVersion: 24.x
      StorageLocation:
        Bucket: "amzn-s3-demo-bucket"
        Key: "MyScriptFiles.zip"
        RoleArn: !GetAtt IAMRole.Arn
```

## See also

- [Create GameLift resources using Amazon CloudFront](https://docs.aws.amazon.com/gamelift/latest/developerguide/resources-cloudformation.html) in the _Amazon_
_GameLift Developer Guide_

- [Upload script files in Amazon S3](https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-script-uploading.html#realtime-script-uploading-s3) in the _Amazon GameLift Developer_
_Guide_

- [CreateScript](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateScript.html) in the _Amazon GameLift API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

S3Location
