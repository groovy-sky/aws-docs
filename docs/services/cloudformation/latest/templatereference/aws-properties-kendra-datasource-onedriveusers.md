This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource OneDriveUsers

User accounts whose documents should be indexed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OneDriveUserList" : [ String, ... ],
  "OneDriveUserS3Path" : S3Path
}

```

### YAML

```yaml

  OneDriveUserList:
    - String
  OneDriveUserS3Path:
    S3Path

```

## Properties

`OneDriveUserList`

A list of users whose documents should be indexed. Specify the user names in email
format, for example, `username@tenantdomain`. If you need to index the
documents of more than 10 users, use the `OneDriveUserS3Path` field to
specify the location of a file containing a list of users.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OneDriveUserS3Path`

The S3 bucket location of a file containing a list of users whose documents should be
indexed.

_Required_: No

_Type_: [S3Path](aws-properties-kendra-datasource-s3path.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OneDriveConfiguration

ProxyConfiguration

All content copied from https://docs.aws.amazon.com/.
