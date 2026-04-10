This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::MLTransform

The AWS::Glue::MLTransform is an AWS Glue resource type that manages machine learning transforms.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::MLTransform",
  "Properties" : {
      "Description" : String,
      "GlueVersion" : String,
      "InputRecordTables" : InputRecordTables,
      "MaxCapacity" : Number,
      "MaxRetries" : Integer,
      "Name" : String,
      "NumberOfWorkers" : Integer,
      "Role" : String,
      "Tags" : [ Tag, ... ],
      "Timeout" : Integer,
      "TransformEncryption" : TransformEncryption,
      "TransformParameters" : TransformParameters,
      "WorkerType" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::MLTransform
Properties:
  Description: String
  GlueVersion: String
  InputRecordTables:
    InputRecordTables
  MaxCapacity: Number
  MaxRetries: Integer
  Name: String
  NumberOfWorkers: Integer
  Role: String
  Tags:
    - Tag
  Timeout: Integer
  TransformEncryption:
    TransformEncryption
  TransformParameters:
    TransformParameters
  WorkerType: String

```

## Properties

`Description`

A user-defined, long-form description text for the machine learning transform.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlueVersion`

This value determines which version of AWS Glue this machine learning transform is compatible with. Glue 1.0 is recommended for most customers. If the value is not set, the Glue compatibility defaults to Glue 0.9. For more information, see [AWS Glue Versions](../../../glue/latest/dg/release-notes.md#release-notes-versions) in the developer guide.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputRecordTables`

A list of AWS Glue table definitions used by the transform.

_Required_: Yes

_Type_: [InputRecordTables](aws-properties-glue-mltransform-inputrecordtables.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxCapacity`

The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from 2 to 100 DPUs; the default is 10. A DPU is a relative measure of
processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory. For more
information, see the [AWS Glue pricing\
page](https://aws.amazon.com/glue/pricing).

`MaxCapacity` is a mutually exclusive option with `NumberOfWorkers` and `WorkerType`.

- If either `NumberOfWorkers` or `WorkerType` is set, then `MaxCapacity` cannot be set.

- If `MaxCapacity` is set then neither `NumberOfWorkers` or `WorkerType` can be set.

- If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).

- `MaxCapacity` and `NumberOfWorkers` must both be at least 1.

When the `WorkerType` field is set to a value other than `Standard`, the `MaxCapacity` field is set automatically and becomes read-only.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRetries`

The maximum number of times to retry after an `MLTaskRun` of the machine
learning transform fails.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A user-defined name for the machine learning transform. Names are required to be unique. `Name` is optional:

- If you supply `Name`, the stack cannot be repeatedly created.

- If `Name` is not provided, a randomly generated name will be used instead.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfWorkers`

The number of workers of a defined `workerType` that are allocated when a task of the transform runs.

If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The name or Amazon Resource Name (ARN) of the IAM role with the required permissions. The required permissions include both AWS Glue service role permissions to AWS Glue resources, and Amazon S3 permissions required by the transform.

- This role needs AWS Glue service role permissions to allow access to resources in AWS Glue. See [Attach a Policy to IAM Users That Access AWS Glue](../../../glue/latest/dg/attach-policy-iam-user.md).

- This role needs permission to your Amazon Simple Storage Service (Amazon S3) sources, targets, temporary directory, scripts, and any libraries used by the task run for this transform.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to use with this machine learning transform. You may use tags to limit access to the machine learning transform. For more information about tags in AWS Glue, see [AWS Tags in AWS Glue](../../../glue/latest/dg/monitor-tags.md) in the developer guide.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timeout`

The timeout in minutes of the machine learning transform.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformEncryption`

The encryption-at-rest settings of the transform that apply to accessing user data. Machine learning
transforms can access user data encrypted in Amazon S3 using KMS.

Additionally, imported labels and trained transforms can now be encrypted using a customer provided
KMS key.

_Required_: No

_Type_: [TransformEncryption](aws-properties-glue-mltransform-transformencryption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformParameters`

The algorithm-specific parameters that are associated with the machine learning transform.

_Required_: Yes

_Type_: [TransformParameters](aws-properties-glue-mltransform-transformparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerType`

The type of predefined worker that is allocated when a task of this transform runs. Accepts a value of Standard, G.1X, or G.2X.

- For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.

- For the `G.1X` worker type, each worker provides 4 vCPU, 16 GB of memory and a 64GB disk, and 1 executor per worker.

- For the `G.2X` worker type, each worker provides 8 vCPU, 32 GB of memory and a 128GB disk, and 1 executor per worker.

`MaxCapacity` is a mutually exclusive option with `NumberOfWorkers` and `WorkerType`.

- If either `NumberOfWorkers` or `WorkerType` is set, then `MaxCapacity` cannot be set.

- If `MaxCapacity` is set then neither `NumberOfWorkers` or `WorkerType` can be set.

- If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).

- `MaxCapacity` and `NumberOfWorkers` must both be at least 1.

_Required_: No

_Type_: String

_Allowed values_: `Standard | G.1X | G.2X | G.025X | G.4X | G.8X | Z.2X`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the transform ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationProperty

FindMatchesParameters

All content copied from https://docs.aws.amazon.com/.
