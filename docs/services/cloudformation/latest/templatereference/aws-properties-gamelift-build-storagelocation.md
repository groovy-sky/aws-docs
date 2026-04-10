This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Build StorageLocation

The location in Amazon S3 where build or script files are stored for access by
Amazon GameLift.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "Key" : String,
  "ObjectVersion" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  Bucket: String
  Key: String
  ObjectVersion: String
  RoleArn: String

```

## Properties

`Bucket`

An Amazon S3 bucket identifier. The name of the S3 bucket.

###### Note

Amazon GameLift doesn't support uploading from Amazon S3 buckets with names that contain a dot
(.).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Key`

The name of the zip file that contains the build files or script files.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectVersion`

A version of a stored file to retrieve, if the object versioning feature is turned on for the S3 bucket.
Use this parameter to specify a specific version. If this parameter isn't set, Amazon GameLift Servers retrieves
the latest version of the file.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARNfor an IAM role that allows Amazon GameLift to access the S3 bucket.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Create GameLift resources using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Create a build with files in Amazon S3](../../../gamelift/latest/developerguide/gamelift-build-cli-uploading.md#gamelift-build-cli-uploading-create-build) in the _Amazon GameLift_
_Developer Guide_

- [Upload script files in Amazon S3](../../../gamelift/latest/developerguide/realtime-script-uploading.md#realtime-script-uploading-s3) in the _Amazon GameLift Developer_
_Guide_

- [S3Location](../../../../reference/gamelift/latest/apireference/api-s3location.md) in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GameLift::Build

Tag

All content copied from https://docs.aws.amazon.com/.
