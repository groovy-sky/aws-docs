This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Build

The `AWS::GameLift::Build` resource creates a game server build that is
installed and run on instances in an Amazon GameLift fleet. This resource points to an Amazon
S3 location that contains a zip file with all of the components of the game server
build.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLift::Build",
  "Properties" : {
      "Name" : String,
      "OperatingSystem" : String,
      "ServerSdkVersion" : String,
      "StorageLocation" : StorageLocation,
      "Tags" : [ Tag, ... ],
      "Version" : String
    }
}

```

### YAML

```yaml

Type: AWS::GameLift::Build
Properties:
  Name: String
  OperatingSystem: String
  ServerSdkVersion: String
  StorageLocation:
    StorageLocation
  Tags:
    - Tag
  Version: String

```

## Properties

`Name`

A descriptive label that is associated with a build. Build names do not need to be
unique.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperatingSystem`

The operating system that your game server binaries run on. This value determines the
type of fleet resources that you use for this build. If your game build contains
multiple executables, they all must run on the same operating system. You must specify a
valid operating system in this request. There is no default value. You can't change a
build's operating system later.

###### Note

Amazon Linux 2 (AL2) will reach end of support on 6/30/2025. See more details in
the [Amazon Linux 2 FAQs](https://aws.amazon.com/amazon-linux-2/faqs).
For game servers
that are hosted on AL2 and use server SDK version 4.x for Amazon GameLift Servers, first update the
game server build to server SDK 5.x, and then deploy to AL2023 instances. See
[Migrate to server SDK version 5.](https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-serversdk5-migration.html)

_Required_: No

_Type_: String

_Allowed values_: `AMAZON_LINUX | AMAZON_LINUX_2 | AMAZON_LINUX_2023 | WINDOWS_2012 | WINDOWS_2016 | WINDOWS_2022`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServerSdkVersion`

A server SDK version you used when integrating your game server build with Amazon GameLift Servers. For more information see [Integrate games\
with custom game servers](https://docs.aws.amazon.com/gamelift/latest/developerguide/integration-custom-intro.html). By default Amazon GameLift Servers sets this value to
`4.0.2`.

_Required_: No

_Type_: String

_Pattern_: `^\d+\.\d+\.\d+$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageLocation`

Information indicating where your game build files are stored. Use this parameter only
when creating a build with files stored in an Amazon S3 bucket that you own. The storage
location must specify an Amazon S3 bucket name and key. The location must also specify a role
ARN that you set up to allow Amazon GameLift Servers to access your Amazon S3 bucket. The S3 bucket and your
new build must be in the same Region.

If a `StorageLocation` is specified, the size of your file can be found in
your Amazon S3 bucket. Amazon GameLift Servers will report a `SizeOnDisk` of 0.

_Required_: No

_Type_: [StorageLocation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-build-storagelocation.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of labels to assign to the new build resource. Tags are developer defined
key-value pairs. Tagging AWS resources are useful for resource management, access
management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the
_AWS General Reference_. Once the resource is created, you can
use [TagResource](https://docs.aws.amazon.com/gamelift/latest/apireference/API_TagResource.html), [UntagResource](https://docs.aws.amazon.com/gamelift/latest/apireference/API_UntagResource.html), and
[ListTagsForResource](https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListTagsForResource.html) to add, remove, and view tags. The maximum tag limit
may be lower than stated. See the AWS General Reference for actual tagging
limits.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gamelift-build-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

Version information that is associated with this build. Version strings do not need to be unique.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the build ID, such as
`build-1111aaaa-22bb-33cc-44dd-5555eeee66ff`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`BuildArn`

The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that is assigned to a Amazon GameLift Servers build resource and uniquely identifies it. ARNs are unique across all Regions. Format is `arn:aws:gamelift:<region>::build/build-a1234567-b8c9-0d1e-2fa3-b45c6d7e8912`. In a GameLift build ARN, the resource ID matches the
_BuildId_ value.

`BuildId`

A unique identifier for the build.

## Examples

### Create a Custom Game Server Build

The following example creates a GameLift build named `MyGameServerBuild`. The
build package is located in an S3 bucket, specified by the `S3Bucket` and
`S3Key` input parameters. The example also creates the AWS Identity and Access Management role
that GameLift assumes so that it has permissions to download the build package files.

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
                "RoleName": "BuildIAMRole",
                "Policies": [
                    {
                        "PolicyName": "gamelift-s3-access-policy",
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
                                        "arn:aws:s3:::amzn-s3-demo-bucket/*"
                                    ]
                                }
                            ]
                        }
                    }
                ]
            }
        },
        "BuildResource": {
            "Type": "AWS::GameLift::Build",
            "Properties": {
                "Name": "MyGameServerBuild",
                "Version": "v1.0",
                "OperatingSystem": "WINDOWS_2012",
                "ServerSdkVersion: "5.0.0",
                "StorageLocation": {
                    "Bucket": "amzn-s3-demo-bucket",
                    "Key": "MyGameBuildFiles.zip",
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
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          - Effect: "Allow"
            Principal:
              Service: ["cloudformation.amazonaws.com", "gamelift.amazonaws.com"]
            Action: "sts:AssumeRole"
      RoleName: "BuildIAMRole"
      Policies:
        - PolicyName: gamelift-s3-access-policy
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
                  - "arn:aws:s3:::amzn-s3-demo-bucket/*"
  BuildResource:
    Type: AWS::GameLift::Build
    Properties:
      Name: MyGameServerBuild
      Version: v1.0
      OperatingSystem: WINDOWS_2012
      ServerSdkVersion: 5.0.0
      StorageLocation:
        Bucket: "amzn-s3-demo-bucket"
        Key: "MyGameBuildFiles.zip"
        RoleArn: !GetAtt IAMRole.Arn
```

## See also

- [Create GameLift resources using Amazon CloudFront](https://docs.aws.amazon.com/gamelift/latest/developerguide/resources-cloudformation.html) in the _Amazon_
_GameLift Developer Guide_

- [Create a build with files in Amazon S3](https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-cli-uploading.html#gamelift-build-cli-uploading-create-build) in the _Amazon GameLift_
_Developer Guide_

- [Upload script files in Amazon S3](https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-script-uploading.html#realtime-script-uploading-s3) in the _Amazon GameLift Developer_
_Guide_

- [CreateBuild](https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateBuild.html) in the _Amazon GameLift API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

StorageLocation
