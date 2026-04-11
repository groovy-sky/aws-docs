This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::MlflowTrackingServer

Creates an MLflow Tracking Server using a general purpose Amazon S3 bucket as the artifact
store. For more information, see [Create an MLflow Tracking\
Server](../../../sagemaker/latest/dg/mlflow-create-tracking-server.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::MlflowTrackingServer",
  "Properties" : {
      "ArtifactStoreUri" : String,
      "AutomaticModelRegistration" : Boolean,
      "MlflowVersion" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "TrackingServerName" : String,
      "TrackingServerSize" : String,
      "WeeklyMaintenanceWindowStart" : String
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::MlflowTrackingServer
Properties:
  ArtifactStoreUri: String
  AutomaticModelRegistration: Boolean
  MlflowVersion: String
  RoleArn: String
  Tags:
    - Tag
  TrackingServerName: String
  TrackingServerSize: String
  WeeklyMaintenanceWindowStart: String

```

## Properties

`ArtifactStoreUri`

The S3 URI for a general purpose bucket to use as the MLflow Tracking Server artifact
store.

_Required_: Yes

_Type_: String

_Pattern_: `^s3:\/\/([^\/]+)\/?(.*)$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutomaticModelRegistration`

Whether to enable or disable automatic registration of new MLflow models to the SageMaker Model Registry. To enable automatic model registration, set this value to `True`.
To disable automatic model registration, set this value to `False`. If not specified, `AutomaticModelRegistration` defaults to `False`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MlflowVersion`

The version of MLflow that the tracking server uses. To see which MLflow versions are
available to use, see [How it works](../../../sagemaker/latest/dg/mlflow.md#mlflow-create-tracking-server-how-it-works).

_Required_: No

_Type_: String

_Pattern_: `^\d+(\.\d+)+$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) for an IAM role in your account that the MLflow Tracking Server uses to
access the artifact store in Amazon S3. The role should have `AmazonS3FullAccess`
permissions. For more information on IAM permissions for tracking server creation, see
[Set up IAM permissions for MLflow](../../../sagemaker/latest/dg/mlflow-create-tracking-server-iam.md).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role\/?[a-zA-Z_0-9+=,.@\-_\/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

Tags consisting of key-value pairs used to manage metadata for the tracking server.

_Required_: No

_Type_: Array of [Tag](aws-properties-sagemaker-mlflowtrackingserver-tag.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackingServerName`

A unique string identifying the tracking server name. This string is part of the tracking server
ARN.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,255}$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TrackingServerSize`

The size of the tracking server you want to create. You can choose between
`"Small"`, `"Medium"`, and `"Large"`. The default MLflow
Tracking Server configuration size is `"Small"`. You can choose a size depending on
the projected use of the tracking server such as the volume of data logged, number of users,
and frequency of use.

We recommend using a small tracking server for teams of up to 25 users, a medium tracking
server for teams of up to 50 users, and a large tracking server for teams of up to 100 users.

_Required_: No

_Type_: String

_Allowed values_: `Small | Medium | Large`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WeeklyMaintenanceWindowStart`

The day and time of the week in Coordinated Universal Time (UTC) 24-hour standard time that weekly maintenance updates are scheduled. For example: TUE:03:30.

_Required_: No

_Type_: String

_Pattern_: `^(Mon|Tue|Wed|Thu|Fri|Sat|Sun):([01]\d|2[0-3]):([0-5]\d)$`

_Maximum_: `9`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`TrackingServerArn`

The ARN of the tracking server.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
