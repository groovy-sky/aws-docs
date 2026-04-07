This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow WorkflowStep

The basic building block of a workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CopyStepDetails" : CopyStepDetails,
  "CustomStepDetails" : CustomStepDetails,
  "DecryptStepDetails" : DecryptStepDetails,
  "DeleteStepDetails" : DeleteStepDetails,
  "TagStepDetails" : TagStepDetails,
  "Type" : String
}

```

### YAML

```yaml

  CopyStepDetails:
    CopyStepDetails
  CustomStepDetails:
    CustomStepDetails
  DecryptStepDetails:
    DecryptStepDetails
  DeleteStepDetails:
    DeleteStepDetails
  TagStepDetails:
    TagStepDetails
  Type: String

```

## Properties

`CopyStepDetails`

Details for a step that performs a file copy.

Consists of the following values:

- A description

- An Amazon S3 location for the destination of the file copy.

- A flag that indicates whether to overwrite an existing file of the same name.
The default is `FALSE`.

_Required_: No

_Type_: [CopyStepDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-workflow-copystepdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CustomStepDetails`

Details for a step that invokes an AWS Lambda function.

Consists of the Lambda function's name, target, and timeout (in seconds).

_Required_: No

_Type_: [CustomStepDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-workflow-customstepdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DecryptStepDetails`

Details for a step that decrypts an encrypted file.

Consists of the following values:

- A descriptive name

- An Amazon S3 or Amazon Elastic File System (Amazon EFS) location for the
source file to decrypt.

- An S3 or Amazon EFS location for the destination of the file
decryption.

- A flag that indicates whether to overwrite an existing file of the same name.
The default is `FALSE`.

- The type of encryption that's used. Currently, only PGP encryption is
supported.

_Required_: No

_Type_: [DecryptStepDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-workflow-decryptstepdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeleteStepDetails`

Details for a step that deletes the file.

_Required_: No

_Type_: [DeleteStepDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-workflow-deletestepdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagStepDetails`

Details for a step that creates one or more tags.

You specify one or more tags. Each tag contains a key-value pair.

_Required_: No

_Type_: [TagStepDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-workflow-tagstepdetails.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

Currently, the following step types are supported.

- **`COPY`** \- Copy the file to another location.

- **`CUSTOM`** \- Perform a custom step with an AWS Lambda function target.

- **`DECRYPT`** \- Decrypt a file that was encrypted before it was uploaded.

- **`DELETE`** \- Delete the file.

- **`TAG`** \- Add a tag to the file.

_Required_: No

_Type_: String

_Allowed values_: `COPY | CUSTOM | DECRYPT | DELETE | TAG`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagStepDetails

Next
