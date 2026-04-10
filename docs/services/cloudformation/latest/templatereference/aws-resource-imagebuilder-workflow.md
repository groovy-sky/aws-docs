This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Workflow

Create a new workflow or a new version of an existing workflow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::Workflow",
  "Properties" : {
      "ChangeDescription" : String,
      "Data" : String,
      "Description" : String,
      "KmsKeyId" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...},
      "Type" : String,
      "Uri" : String,
      "Version" : String
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::Workflow
Properties:
  ChangeDescription: String
  Data: String
  Description: String
  KmsKeyId: String
  Name: String
  Tags:
    Key: Value
  Type: String
  Uri: String
  Version: String

```

## Properties

`ChangeDescription`

Describes what change has been made in this version of the workflow, or
what makes this version different from other versions of the workflow.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Data`

Contains the UTF-8 encoded YAML document content for the workflow.
Alternatively, you can specify the `uri` of a YAML document file stored in
Amazon S3. However, you cannot specify both properties.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `16000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

Describes the workflow.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The Amazon Resource Name (ARN) that uniquely identifies the KMS key used to encrypt this workflow resource.
This can be either the Key ARN or the Alias ARN. For more information, see [Key identifiers (KeyId)](../../../kms/latest/developerguide/concepts.md#key-id-key-ARN)
in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the workflow to create.

_Required_: Yes

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags that apply to the workflow resource.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The phase in the image build process for which the workflow resource
is responsible.

_Required_: Yes

_Type_: String

_Allowed values_: `BUILD | TEST | DISTRIBUTION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Uri`

The `uri` of a YAML component document file. This must be an S3 URL
( `s3://bucket/key`), and the requester must have permission to access the
S3 bucket it points to. If you use Amazon S3, you can specify component content up to your
service quota.

Alternatively, you can specify the YAML document inline, using the component
`data` property. You cannot specify both properties.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

The semantic version of this workflow resource. The semantic version syntax
adheres to the following rules.

###### Note

The semantic version has four nodes: <major>.<minor>.<patch>/<build>.
You can assign values for the first three, and can filter on all of them.

**Assignment:** For the first three nodes you can assign any positive integer value, including
zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder automatically assigns the
build number to the fourth node.

**Patterns:** You can use any numeric pattern that adheres to the assignment requirements for
the nodes that you can assign. For example, you might choose a software version pattern, such as 1.0.0, or
a date, such as 2021.01.01.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+\.[0-9]+\.[0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the workflow resource.

`LatestVersion.Arn`

The latest version Amazon Resource Name (ARN) of the Image Builder resource.

`LatestVersion.Major`

The latest version Amazon Resource Name (ARN) with the same `major` version of the Image Builder resource.

`LatestVersion.Minor`

The latest version Amazon Resource Name (ARN) with the same `minor` version of the Image Builder resource.

`LatestVersion.Patch`

The latest version Amazon Resource Name (ARN) with the same `patch` version of the Image Builder resource.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceSelection

LatestVersion

All content copied from https://docs.aws.amazon.com/.
