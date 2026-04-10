This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job

Specifies a new DataBrew job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataBrew::Job",
  "Properties" : {
      "DatabaseOutputs" : [ DatabaseOutput, ... ],
      "DataCatalogOutputs" : [ DataCatalogOutput, ... ],
      "DatasetName" : String,
      "EncryptionKeyArn" : String,
      "EncryptionMode" : String,
      "JobSample" : JobSample,
      "LogSubscription" : String,
      "MaxCapacity" : Integer,
      "MaxRetries" : Integer,
      "Name" : String,
      "OutputLocation" : OutputLocation,
      "Outputs" : [ Output, ... ],
      "ProfileConfiguration" : ProfileConfiguration,
      "ProjectName" : String,
      "Recipe" : Recipe,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "Timeout" : Integer,
      "Type" : String,
      "ValidationConfigurations" : [ ValidationConfiguration, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataBrew::Job
Properties:
  DatabaseOutputs:
    - DatabaseOutput
  DataCatalogOutputs:
    - DataCatalogOutput
  DatasetName: String
  EncryptionKeyArn: String
  EncryptionMode: String
  JobSample:
    JobSample
  LogSubscription: String
  MaxCapacity: Integer
  MaxRetries: Integer
  Name: String
  OutputLocation:
    OutputLocation
  Outputs:
    - Output
  ProfileConfiguration:
    ProfileConfiguration
  ProjectName: String
  Recipe:
    Recipe
  RoleArn: String
  Tags:
    - Tag
  Timeout: Integer
  Type: String
  ValidationConfigurations:
    - ValidationConfiguration

```

## Properties

`DatabaseOutputs`

Represents a list of JDBC database output objects which defines the output
destination for a DataBrew recipe job to write into.

_Required_: No

_Type_: Array of [DatabaseOutput](aws-properties-databrew-job-databaseoutput.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataCatalogOutputs`

One or more artifacts that represent the AWS Glue Data Catalog output
from running the job.

_Required_: No

_Type_: Array of [DataCatalogOutput](aws-properties-databrew-job-datacatalogoutput.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetName`

A dataset that the job is to process.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyArn`

The Amazon Resource Name (ARN) of an encryption key that is used to protect the job
output. For more information, see [Encrypting data\
written by DataBrew jobs](../../../databrew/latest/dg/encryption-security-configuration.md)

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionMode`

The encryption mode for the job, which can be one of the following:

- `SSE-KMS` \- Server-side encryption with keys managed by AWS KMS.

- `SSE-S3` \- Server-side encryption with keys managed by Amazon S3.

_Required_: No

_Type_: String

_Allowed values_: `SSE-KMS | SSE-S3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobSample`

A sample configuration for profile jobs only, which determines the number of rows on which the
profile job is run. If a `JobSample` value isn't provided, the default value
is used. The default value is CUSTOM\_ROWS for the mode parameter and 20,000 for the
size parameter.

_Required_: No

_Type_: [JobSample](aws-properties-databrew-job-jobsample.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogSubscription`

The current status of Amazon CloudWatch logging for the job.

_Required_: No

_Type_: String

_Allowed values_: `ENABLE | DISABLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCapacity`

The maximum number of nodes that can be consumed when the job processes data.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRetries`

The maximum number of times to retry the job after a job run fails.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique name of the job.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputLocation`

The location in Amazon S3 where the job writes its output.

_Required_: No

_Type_: [OutputLocation](aws-properties-databrew-job-outputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Outputs`

One or more artifacts that represent output from running the job.

_Required_: No

_Type_: Array of [Output](aws-properties-databrew-job-output.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProfileConfiguration`

Configuration for profile jobs. Configuration can be used to select columns, do evaluations, and override default
parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all
supported columns.

_Required_: No

_Type_: [ProfileConfiguration](aws-properties-databrew-job-profileconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProjectName`

The name of the project that the job is associated with.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Recipe`

A series of data transformation steps that the job runs.

_Required_: No

_Type_: [Recipe](aws-properties-databrew-job-recipe.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the role to be assumed for this job.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata tags that have been applied to the job.

_Required_: No

_Type_: Array of [Tag](aws-properties-databrew-job-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The job's timeout in minutes. A job that attempts to run longer than this timeout
period ends with a status of `TIMEOUT`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The job type of the job, which must be one of the following:

- `PROFILE` \- A job to analyze a dataset, to determine its size, data
types, data distribution, and more.

- `RECIPE` \- A job to apply one or more transformations to a
dataset.

_Required_: Yes

_Type_: String

_Allowed values_: `PROFILE | RECIPE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidationConfigurations`

List of validation configurations that are applied to the profile job.

_Required_: No

_Type_: Array of [ValidationConfiguration](aws-properties-databrew-job-validationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the resource name. For example:

`{ "Ref": "myJob" }`

For an AWS Glue DataBrew job named `myJob`, `Ref`
returns the name of the job.

## Examples

### Creating jobs

The following examples create new DataBrew profile jobs.

#### YAML

```yaml

Resources:
  TestDataBrewJob:
    Type: AWS::DataBrew::Job
    Properties:
      Type: PROFILE
      Name: job-name
      DatasetName: dataset-name
      RoleArn: arn:aws:iam::12345678910:role/PassRoleAdmin
      JobSample:
        Mode: 'CUSTOM_ROWS'
        Size: 50000
      OutputLocation:
        Bucket: !Join [ '', ['databrew-cfn-integration-tests-', !Ref 'AWS::Region', '-', !Ref 'AWS::AccountId' ] ]
      Tags: [{Key: key00AtCreate, Value: value001AtCreate}]

```

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "This CloudFormation template specifies a DataBrew Profile Job",
    "Resources": {
        "MyDataBrewProfileJob": {
            "Type": "AWS::DataBrew::Job",
            "Properties": {
                "Type": "PROFILE",
                "Name": "job-test",
                "DatasetName": "dataset-test",
                "RoleArn": "arn:aws:iam::1234567891011:role/PassRoleAdmin",
                "JobSample": {
                    "Mode": "FULL_DATASET"
                },
                "OutputLocation": {
                    "Bucket": "test-output",
                    "Key": "job-output.json"
                },
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

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AllowedStatistics

All content copied from https://docs.aws.amazon.com/.
