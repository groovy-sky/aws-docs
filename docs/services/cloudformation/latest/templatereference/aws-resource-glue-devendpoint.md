This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::DevEndpoint

The `AWS::Glue::DevEndpoint` resource specifies a development endpoint
where a developer can remotely debug ETL scripts for AWS Glue. For more information, see
[DevEndpoint Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-dev-endpoint.html#aws-glue-api-jobs-dev-endpoint-DevEndpoint) in the AWS Glue Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::DevEndpoint",
  "Properties" : {
      "Arguments" : Json,
      "EndpointName" : String,
      "ExtraJarsS3Path" : String,
      "ExtraPythonLibsS3Path" : String,
      "GlueVersion" : String,
      "NumberOfNodes" : Integer,
      "NumberOfWorkers" : Integer,
      "PublicKey" : String,
      "PublicKeys" : [ String, ... ],
      "RoleArn" : String,
      "SecurityConfiguration" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SubnetId" : String,
      "Tags" : [ Tag, ... ],
      "WorkerType" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::DevEndpoint
Properties:
  Arguments: Json
  EndpointName: String
  ExtraJarsS3Path: String
  ExtraPythonLibsS3Path: String
  GlueVersion: String
  NumberOfNodes: Integer
  NumberOfWorkers: Integer
  PublicKey: String
  PublicKeys:
    - String
  RoleArn: String
  SecurityConfiguration: String
  SecurityGroupIds:
    - String
  SubnetId: String
  Tags:
    - Tag
  WorkerType: String

```

## Properties

`Arguments`

A map of arguments used to configure the `DevEndpoint`.

Valid arguments are:

- `"--enable-glue-datacatalog": ""`

- `"GLUE_PYTHON_VERSION": "3"`

- `"GLUE_PYTHON_VERSION": "2"`

You can specify a version of Python support for development endpoints by using the `Arguments` parameter in the `CreateDevEndpoint` or `UpdateDevEndpoint` APIs. If no arguments are provided, the version defaults to Python 2.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointName`

The name of the `DevEndpoint`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExtraJarsS3Path`

The path to one or more Java `.jar` files in an S3 bucket that should be
loaded in your `DevEndpoint`.

###### Note

You can only use pure Java/Scala libraries with a `DevEndpoint`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExtraPythonLibsS3Path`

The paths to one or more Python libraries in an Amazon S3 bucket that should be loaded
in your `DevEndpoint`. Multiple values must be complete paths separated by a
comma.

###### Note

You can only use pure Python libraries with a `DevEndpoint`. Libraries
that rely on C extensions, such as the [pandas](http://pandas.pydata.org/) Python data analysis library, are not currently
supported.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlueVersion`

The AWS Glue version determines the versions of Apache Spark and Python that AWS Glue supports. The Python version indicates the version supported for running your ETL scripts on development endpoints.

For more information about the available AWS Glue versions and corresponding Spark and Python versions, see [Glue version](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) in the developer guide.

Development endpoints that are created without specifying a Glue version default to Glue 0.9.

You can specify a version of Python support for development endpoints by using the `Arguments` parameter in the `CreateDevEndpoint` or `UpdateDevEndpoint` APIs. If no arguments are provided, the version defaults to Python 2.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfNodes`

The number of AWS Glue Data Processing Units (DPUs) allocated to this
`DevEndpoint`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfWorkers`

The number of workers of a defined `workerType` that are allocated to the development endpoint.

The maximum number of workers you can define are 299 for `G.1X`, and 149 for `G.2X`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicKey`

The public key to be used by this `DevEndpoint` for authentication. This
attribute is provided for backward compatibility because the recommended attribute to
use is public keys.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicKeys`

A list of public keys to be used by the `DevEndpoints` for authentication.
Using this attribute is preferred over a single public key because the public keys allow you
to have a different private key per client.

###### Note

If you previously created an endpoint with a public key, you must remove that key to be
able to set a list of public keys. Call the `UpdateDevEndpoint` API operation
with the public key content in the `deletePublicKeys` attribute, and the list of
new keys in the `addPublicKeys` attribute.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role used in this
`DevEndpoint`.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:iam::\d{12}:role/.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityConfiguration`

The name of the `SecurityConfiguration` structure to be used with this
`DevEndpoint`.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

A list of security group identifiers used in this `DevEndpoint`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The subnet ID for this `DevEndpoint`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to use with this DevEndpoint.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerType`

The type of predefined worker that is allocated to the development endpoint. Accepts a value of Standard, G.1X, or G.2X.

- For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.

- For the `G.1X` worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.

- For the `G.2X` worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.

Known issue: when a development endpoint is created with the `G.2X` `WorkerType` configuration, the Spark drivers for the development endpoint will run on 4 vCPU, 16 GB of memory, and a 64 GB disk.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the endpoint name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## See also

- [DevEndpoint Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-dev-endpoint.html#aws-glue-api-jobs-dev-endpoint-DevEndpoint) in the _AWS Glue Developer_
_Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataQualityTargetTable

AWS::Glue::IdentityCenterConfiguration
