This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Express::AccessPoint VpcConfiguration

The `VpcConfiguration` property type specifies Property description not available. for an [AWS::S3Express::AccessPoint](aws-resource-s3express-accesspoint.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "VpcId" : String
}

```

### YAML

```yaml

  VpcId: String

```

## Properties

`VpcId`

If this field is specified, this access point will only allow connections from the specified VPC ID.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::S3Express::BucketPolicy
