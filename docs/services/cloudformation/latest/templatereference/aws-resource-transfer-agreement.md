This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Agreement

Creates an agreement. An agreement is a bilateral trading partner agreement, or
partnership, between an AWS Transfer Family server and an AS2 process. The agreement defines the
file and message transfer relationship between the server and the AS2 process. To define
an agreement, Transfer Family combines a server, local profile, partner profile, certificate, and
other attributes.

The partner is identified with the `PartnerProfileId`, and the AS2 process
is identified with the `LocalProfileId`.

###### Note

Specify _either_ `BaseDirectory` or `CustomDirectories`, but not both.
Specifying both causes the command to fail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Transfer::Agreement",
  "Properties" : {
      "AccessRole" : String,
      "BaseDirectory" : String,
      "CustomDirectories" : CustomDirectories,
      "Description" : String,
      "EnforceMessageSigning" : String,
      "LocalProfileId" : String,
      "PartnerProfileId" : String,
      "PreserveFilename" : String,
      "ServerId" : String,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Transfer::Agreement
Properties:
  AccessRole: String
  BaseDirectory: String
  CustomDirectories:
    CustomDirectories
  Description: String
  EnforceMessageSigning: String
  LocalProfileId: String
  PartnerProfileId: String
  PreserveFilename: String
  ServerId: String
  Status: String
  Tags:
    - Tag

```

## Properties

`AccessRole`

Connectors are used to send files using either the AS2 or SFTP protocol. For the access role,
provide the Amazon Resource Name (ARN) of the AWS Identity and Access Management role to use.

**For AS2 connectors**

With AS2, you can send files by calling `StartFileTransfer` and specifying the
file paths in the request parameter, `SendFilePaths`. We use the file’s parent
directory (for example, for `--send-file-paths /bucket/dir/file.txt`, parent
directory is `/bucket/dir/`) to temporarily store a processed AS2 message file,
store the MDN when we receive them from the partner, and write a final JSON file containing
relevant metadata of the transmission. So, the `AccessRole` needs to provide read
and write access to the parent directory of the file location used in the
`StartFileTransfer` request. Additionally, you need to provide read and write
access to the parent directory of the files that you intend to send with
`StartFileTransfer`.

If you are using Basic authentication for your AS2 connector, the access role requires the
`secretsmanager:GetSecretValue` permission for the secret. If the secret is encrypted using
a customer-managed key instead of the AWS managed key in Secrets Manager, then the role also
needs the `kms:Decrypt` permission for that key.

**For SFTP connectors**

Make sure that the access role provides
read and write access to the parent directory of the file location
that's used in the `StartFileTransfer` request.
Additionally, make sure that the role provides
`secretsmanager:GetSecretValue` permission to AWS Secrets Manager.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*role/.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseDirectory`

The landing directory (folder) for files that are transferred by using the AS2
protocol.

_Required_: No

_Type_: String

_Pattern_: `^(|/.*)$`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDirectories`

A `CustomDirectoriesType` structure. This structure specifies custom directories for storing various AS2 message files. You can specify directories for the following types of files.

- Failed files

- MDN files

- Payload files

- Status files

- Temporary files

_Required_: No

_Type_: [CustomDirectories](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-agreement-customdirectories.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The name or short description that's used to identify the agreement.

_Required_: No

_Type_: String

_Pattern_: `^[\u0021-\u007E]+$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnforceMessageSigning`

Determines whether or not unsigned messages from your trading partners will be accepted.

- `ENABLED`: Transfer Family rejects unsigned messages from your trading partner.

- `DISABLED` (default value): Transfer Family accepts unsigned messages from your trading partner.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalProfileId`

A unique identifier for the AS2 local profile.

_Required_: Yes

_Type_: String

_Pattern_: `^p-([0-9a-f]{17})$`

_Minimum_: `19`

_Maximum_: `19`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartnerProfileId`

A unique identifier for the partner profile used in the agreement.

_Required_: Yes

_Type_: String

_Pattern_: `^p-([0-9a-f]{17})$`

_Minimum_: `19`

_Maximum_: `19`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreserveFilename`

Determines whether or not Transfer Family appends a unique string of characters to the end of the AS2 message payload
filename when saving it.

- `ENABLED`: the filename provided by your trading parter is preserved when the file is saved.

- `DISABLED` (default value): when Transfer Family saves the file, the filename is adjusted, as
described in [File names and locations](https://docs.aws.amazon.com/transfer/latest/userguide/send-as2-messages.html#file-names-as2).

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerId`

A system-assigned unique identifier for a server instance. This identifier indicates
the specific server that the agreement uses.

_Required_: Yes

_Type_: String

_Pattern_: `^s-([0-9a-f]{17})$`

_Minimum_: `19`

_Maximum_: `19`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

The current status of the agreement, either `ACTIVE` or
`INACTIVE`.

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | INACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs that can be used to group and search for agreements.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-agreement-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AgreementId`

The unique identifier for the AS2 agreement, returned after the API call
succeeds.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Transfer Family

CustomDirectories
