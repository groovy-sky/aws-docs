This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3ObjectLambda::AccessPointPolicy

The `AWS::S3ObjectLambda::AccessPointPolicy` resource specifies the Object
Lambda Access Point resource policy document.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::S3ObjectLambda::AccessPointPolicy",
  "Properties" : {
      "ObjectLambdaAccessPoint" : String,
      "PolicyDocument" : Json
    }
}

```

### YAML

```yaml

Type: AWS::S3ObjectLambda::AccessPointPolicy
Properties:
  ObjectLambdaAccessPoint: String
  PolicyDocument: Json

```

## Properties

`ObjectLambdaAccessPoint`

An access point with an attached AWS Lambda function used to access transformed data from an Amazon S3
bucket.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

_Minimum_: `3`

_Maximum_: `45`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyDocument`

Object Lambda Access Point resource policy document.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransformationConfiguration

Next
