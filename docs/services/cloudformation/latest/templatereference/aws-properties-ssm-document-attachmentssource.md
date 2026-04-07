This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::Document AttachmentsSource

Identifying information about a document attachment, including the file name and a key-value
pair that identifies the location of an attachment to a document.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Name" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Name: String
  Values:
    - String

```

## Properties

`Key`

The key of a key-value pair that identifies the location of an attachment to a
document.

_Required_: No

_Type_: String

_Allowed values_: `SourceUrl | S3FileUrl | AttachmentReference`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Name`

The name of the document attachment file.

_Required_: No

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Values`

The value of a key-value pair that identifies the location of an attachment to a document.
The format for **Value** depends on the type of key you
specify.

- For the key _SourceUrl_, the value is an S3 bucket location. For
example:

`"Values": [ "s3://amzn-s3-demo-bucket/my-prefix" ]`

- For the key _S3FileUrl_, the value is a file in an S3 bucket. For
example:

`"Values": [ "s3://amzn-s3-demo-bucket/my-prefix/my-file.py" ]`

- For the key _AttachmentReference_, the value is constructed from the
name of another SSM document in your account, a version number of that document, and a file
attached to that document version that you want to reuse. For example:

`"Values": [ "MyOtherDocument/3/my-other-file.py" ]`

However, if the SSM document is shared with you from another account, the full SSM
document ARN must be specified instead of the document name only. For example:

`"Values": [
        "arn:aws:ssm:us-east-2:111122223333:document/OtherAccountDocument/3/their-file.py"
        ]`

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100000 | 1`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SSM::Document

DocumentRequires
