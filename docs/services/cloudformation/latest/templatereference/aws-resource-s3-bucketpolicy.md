This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::BucketPolicy

Applies an Amazon S3 bucket policy to an Amazon S3 bucket. If you are using an identity
other than the root user of the AWS account that owns the bucket, the calling
identity must have the `PutBucketPolicy` permissions on the specified bucket and
belong to the bucket owner's account in order to use this operation.

If you don't have `PutBucketPolicy` permissions, Amazon S3 returns a `403
        Access Denied` error. If you have the correct permissions, but you're not using an
identity that belongs to the bucket owner's account, Amazon S3 returns a `405 Method Not
        Allowed` error.

###### Important

As a security precaution, the root user of the AWS account that owns a
bucket can always use this operation, even if the policy explicitly denies the root user the
ability to perform this action.

When using the `AWS::S3::BucketPolicy` resource,
you can create, update, and delete bucket policies for S3 buckets located in Regions that are different from the stack's Region. However, the CloudFormation stacks should be deployed in the US East (N. Virginia) or `us-east-1` Region. This cross-region bucket policy modification functionality is supported for backward compatibility with existing workflows.

###### Important

If the [DeletionPolicy attribute](../userguide/aws-attribute-deletionpolicy.md) is not specified or set to `Delete`,
the bucket policy will be removed when the stack is deleted. If set to `Retain`, the bucket policy will be preserved even after the stack is deleted.

For example, a CloudFormation stack in `us-east-1` can use the `AWS::S3::BucketPolicy` resource to manage the bucket policy for an S3 bucket in `us-west-2`.
The retention or removal of the bucket policy during the stack deletion is determined by the `DeletionPolicy` attribute specified in the stack template.

For more information, see [Bucket policy\
examples](../../../s3/latest/userguide/example-bucket-policies.md).

The following operations are related to `PutBucketPolicy`:

- [CreateBucket](../../../s3/latest/api/api-createbucket.md)

- [DeleteBucket](../../../s3/latest/api/api-deletebucket.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3::BucketPolicy",
  "Properties" : {
      "Bucket" : String,
      "PolicyDocument" : Json
    }
}

```

### YAML

```yaml

Type: AWS::S3::BucketPolicy
Properties:
  Bucket: String
  PolicyDocument: Json

```

## Properties

`Bucket`

The name of the Amazon S3 bucket to which the policy applies.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyDocument`

A policy document containing permissions to add to the specified bucket. In IAM, you must
provide policy documents in JSON format. However, in CloudFormation you can provide the policy
in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to
IAM. For more information, see the AWS::IAM::Policy [PolicyDocument](../userguide/aws-resource-iam-policy.md#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language\
Overview](../../../s3/latest/dev/access-policy-language-overview.md) in the _Amazon S3 User Guide_.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Bucket policy that allows GET requests from specific referers

The following sample is a bucket policy that is attached to the DOC-EXAMPLE-BUCKET
bucket and allows GET requests that originate from www.example.com and example.net:

###### Important

This key should be used carefully. It is dangerous to include a publicly known
referer header value. Unauthorized parties can use modified or custom browsers to
provide any `aws:referer` value that they choose. As a result,
`aws:referer` should not be used to prevent unauthorized parties from
making direct AWS requests. It is offered only to allow customers to
protect their digital content, such as content stored in Amazon S3, from being
referenced on unauthorized third-party sites. For more information, see [`aws:referer`](../../../iam/latest/userguide/reference-policies-condition-keys.md#condition-keys-referer) in the _IAM User Guide_.

#### JSON

```json

{
    "SampleBucketPolicy": {
        "Type": "AWS::S3::BucketPolicy",
        "Properties": {
            "Bucket": {
                "Ref": "DOC-EXAMPLE-BUCKET"
            },
            "PolicyDocument": {
                "Version": "2012-10-17",
                "Statement": [
                    {
                        "Action": [
                            "s3:GetObject"
                        ],
                        "Effect": "Allow",
                        "Resource": {
                            "Fn::Join": [
                                "",
                                [
                                    "arn:aws:s3:::",
                                    {
                                        "Ref": "DOC-EXAMPLE-BUCKET"
                                    },
                                    "/*"
                                ]
                            ]
                        },
                        "Principal": "*",
                        "Condition": {
                            "StringLike": {
                                "aws:Referer": [
                                    "http://www.example.com/*",
                                    "http://example.net/*"
                                ]
                            }
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

SampleBucketPolicy:
  Type: AWS::S3::BucketPolicy
  Properties:
    Bucket: !Ref DOC-EXAMPLE-BUCKET
    PolicyDocument:
      Version: 2012-10-17
      Statement:
        - Action:
            - 's3:GetObject'
          Effect: Allow
          Resource: !Join
            - ''
            - - 'arn:aws:s3:::'
              - !Ref DOC-EXAMPLE-BUCKET
              - /*
          Principal: '*'
          Condition:
            StringLike:
              'aws:Referer':
                - 'http://www.example.com/*'
                - 'http://example.net/*'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebsiteConfiguration

AWS::S3::MultiRegionAccessPoint

All content copied from https://docs.aws.amazon.com/.
