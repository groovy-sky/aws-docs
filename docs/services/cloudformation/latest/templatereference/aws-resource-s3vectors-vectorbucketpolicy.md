This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Vectors::VectorBucketPolicy

The `AWS::S3Vectors::VectorBucketPolicy` resource defines an Amazon S3 vector bucket policy to control access to an Amazon S3 vector bucket.

Vector bucket policies are written in JSON and allow you to grant or deny permissions across all (or a subset of) objects within a vector bucket.

You must specify either `VectorBucketName` or `VectorBucketArn` to identify the target bucket.

To control how AWS CloudFormation handles the vector bucket policy when the stack is deleted, you can set a deletion policy for your policy. You can choose to _retain_ the policy or to _delete_ the policy. For more information, see [DeletionPolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).

Permissions

The required permissions for CloudFormation to use are based on the operations that are performed on the stack.

- Create

- s3vectors:GetVectorBucketPolicy

- s3vectors:PutVectorBucketPolicy

- Read

- s3vectors:GetVectorBucketPolicy

- Update

- s3vectors:GetVectorBucketPolicy

- s3vectors:PutVectorBucketPolicy

- Delete

- s3vectors:GetVectorBucketPolicy

- s3vectors:DeleteVectorBucketPolicy

- List

- s3vectors:GetVectorBucketPolicy

- s3vectors:ListVectorBuckets

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3Vectors::VectorBucketPolicy",
  "Properties" : {
      "Policy" : Json,
      "VectorBucketArn" : String,
      "VectorBucketName" : String
    }
}

```

### YAML

```yaml

Type: AWS::S3Vectors::VectorBucketPolicy
Properties:
  Policy: Json
  VectorBucketArn: String
  VectorBucketName: String

```

## Properties

`Policy`

A policy document containing permissions to add to the specified vector bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VectorBucketArn`

The Amazon Resource Name (ARN) of the S3 vector bucket to which the policy applies.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VectorBucketName`

The name of the S3 vector bucket to which the policy applies.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the vector bucket ARN.

Example: `arn:aws:s3vectors:us-east-1:123456789012:bucket/amzn-s3-demo-vector-bucket`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
