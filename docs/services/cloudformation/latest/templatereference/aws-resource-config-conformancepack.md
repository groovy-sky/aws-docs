This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConformancePack

A conformance pack is a collection of AWS Config rules and remediation actions
that can be easily deployed in an account and a region.
ConformancePack creates a service linked role in your account.
The service linked role is created only when the role does not exist in your account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Config::ConformancePack",
  "Properties" : {
      "ConformancePackInputParameters" : [ ConformancePackInputParameter, ... ],
      "ConformancePackName" : String,
      "DeliveryS3Bucket" : String,
      "DeliveryS3KeyPrefix" : String,
      "TemplateBody" : String,
      "TemplateS3Uri" : String,
      "TemplateSSMDocumentDetails" : TemplateSSMDocumentDetails
    }
}

```

### YAML

```yaml

Type: AWS::Config::ConformancePack
Properties:
  ConformancePackInputParameters:
    - ConformancePackInputParameter
  ConformancePackName: String
  DeliveryS3Bucket: String
  DeliveryS3KeyPrefix: String
  TemplateBody: String
  TemplateS3Uri: String
  TemplateSSMDocumentDetails:
    TemplateSSMDocumentDetails

```

## Properties

`ConformancePackInputParameters`

A list of ConformancePackInputParameter objects.

_Required_: No

_Type_: Array of [ConformancePackInputParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-config-conformancepack-conformancepackinputparameter.html)

_Minimum_: `0`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConformancePackName`

Name of the conformance pack you want to create.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z][-a-zA-Z0-9]*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeliveryS3Bucket`

The name of the Amazon S3 bucket where AWS Config stores conformance pack templates.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryS3KeyPrefix`

The prefix for the Amazon S3 bucket.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateBody`

A string containing full conformance pack template body. Structure containing the template body with a
minimum length of 1 byte and a maximum length of 51,200 bytes.

###### Note

You can only use a YAML template with two resource types: config rule ( `AWS::Config::ConfigRule`) and a remediation action ( `AWS::Config::RemediationConfiguration`).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateS3Uri`

Location of file containing the template body (s3://bucketname/prefix). The uri must point to the conformance pack template (max size: 300 KB)
that is located in an Amazon S3 bucket.

###### Note

You must have access to read Amazon S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `s3://.*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateSSMDocumentDetails`

An object that contains the name or Amazon Resource Name (ARN) of the AWS Systems Manager document (SSM document) and the version of the SSM document that is used to create a conformance pack.

_Required_: No

_Type_: [TemplateSSMDocumentDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-config-conformancepack-templatessmdocumentdetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the conformance pack.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Conformance Pack

The following example creates a conformance pack.

#### JSON

```json

{
    "Resources": {
        "ConformancePack": {
            "Type": "AWS::Config::ConformancePack",
            "Properties": {
                "ConformancePackName": "ConformancePackName",
                "DeliveryS3Bucket": "DeliveryS3Bucket",
                "TemplateS3Uri": "s3://bucketname/prefix"
            }
        }
    }
}

```

#### YAML

```yaml

---
AWSTemplateFormatVersion: 2010-09-09
Resources:
    CloudFormationCanaryPack:
        Type: AWS::Config::ConformancePack
        Properties:
          ConformancePackName: ConformancePackName
          DeliveryS3Bucket: DeliveryS3Bucket
          TemplateS3Uri: s3://bucketname/prefix

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RecordingStrategy

ConformancePackInputParameter
