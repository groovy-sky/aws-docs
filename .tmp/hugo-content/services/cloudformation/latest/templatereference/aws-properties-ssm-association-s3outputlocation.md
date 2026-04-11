This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::Association S3OutputLocation

`S3OutputLocation` is a property of the [AWS::SSM::Association](../userguide/aws-resource-ssm-association.md) resource that specifies an Amazon S3 bucket
where you want to store the results of this association request.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OutputS3BucketName" : String,
  "OutputS3KeyPrefix" : String,
  "OutputS3Region" : String
}

```

### YAML

```yaml

  OutputS3BucketName: String
  OutputS3KeyPrefix: String
  OutputS3Region: String

```

## Properties

`OutputS3BucketName`

The name of the S3 bucket.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputS3KeyPrefix`

The S3 bucket subfolder.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputS3Region`

The AWS Region of the S3 bucket.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceAssociationOutputLocation

Target

All content copied from https://docs.aws.amazon.com/.
