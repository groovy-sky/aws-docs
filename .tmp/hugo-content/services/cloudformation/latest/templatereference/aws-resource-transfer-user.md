This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::User

The `AWS::Transfer::User` resource creates a user and associates them with
an existing server. You can only create and associate users with servers that have the
`IdentityProviderType` set to `SERVICE_MANAGED` . Using
parameters for `CreateUser` , you can specify the user name, set the home
directory, store the user's public key, and assign the user's AWS
Identity and Access Management (IAM) role. You can also optionally add a session policy,
and assign metadata with tags that can be used to group and search for users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Transfer::User",
  "Properties" : {
      "HomeDirectory" : String,
      "HomeDirectoryMappings" : [ HomeDirectoryMapEntry, ... ],
      "HomeDirectoryType" : String,
      "Policy" : String,
      "PosixProfile" : PosixProfile,
      "Role" : String,
      "ServerId" : String,
      "SshPublicKeys" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "UserName" : String
    }
}

```

### YAML

```yaml

Type: AWS::Transfer::User
Properties:
  HomeDirectory: String
  HomeDirectoryMappings:
    - HomeDirectoryMapEntry
  HomeDirectoryType: String
  Policy: String
  PosixProfile:
    PosixProfile
  Role: String
  ServerId: String
  SshPublicKeys:
    - String
  Tags:
    - Tag
  UserName: String

```

## Properties

`HomeDirectory`

The landing directory (folder) for a user when they log in to the server using the client.

A `HomeDirectory` example is `/bucket_name/home/mydirectory`.

###### Note

You can use the `HomeDirectory` parameter for `HomeDirectoryType` when it is set to either `PATH` or `LOGICAL`.

_Required_: No

_Type_: String

_Pattern_: `^(|/.*)$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HomeDirectoryMappings`

Logical directory mappings that specify what Amazon S3 or Amazon EFS paths and keys should
be visible to your user and how you want to make them visible. You must specify the
`Entry` and `Target` pair, where `Entry` shows how the path
is made visible and `Target` is the actual Amazon S3 or Amazon EFS path. If you
only specify a target, it is displayed as is. You also must ensure that your AWS Identity and Access Management (IAM)
role provides access to paths in `Target`. This value
can be set only when `HomeDirectoryType` is set to
_LOGICAL_.

The following is an `Entry` and `Target` pair example.

`[ { "Entry": "/directory1", "Target": "/bucket_name/home/mydirectory" }
            ]`

In most cases, you can use this value instead of the session policy to lock your user
down to the designated home directory (" `chroot`"). To do this, you can set
`Entry` to `/` and set `Target` to the value the
user should see for their home directory when they log in.

The following is an `Entry` and `Target` pair example for `chroot`.

`[ { "Entry": "/", "Target": "/bucket_name/home/mydirectory" } ]`

_Required_: No

_Type_: Array of [HomeDirectoryMapEntry](aws-properties-transfer-user-homedirectorymapentry.md)

_Minimum_: `1`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HomeDirectoryType`

The type of landing directory (folder) that you want your users' home directory to be when they log in to the server.
If you set it to `PATH`, the user will see the absolute Amazon S3 bucket or Amazon EFS path as is in their file transfer
protocol clients. If you set it to `LOGICAL`, you need to provide mappings in the `HomeDirectoryMappings` for
how you want to make Amazon S3 or Amazon EFS paths visible to your users.

###### Note

If `HomeDirectoryType` is `LOGICAL`, you must provide mappings,
using the `HomeDirectoryMappings` parameter. If, on the other hand,
`HomeDirectoryType` is `PATH`, you provide an absolute path
using the `HomeDirectory` parameter. You cannot have both
`HomeDirectory` and `HomeDirectoryMappings` in your
template.

_Required_: No

_Type_: String

_Allowed values_: `PATH | LOGICAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policy`

A session policy for your user so you can use the same IAM role across multiple
users. This policy restricts user access to portions of their Amazon S3 bucket.
Variables that you can use inside this policy include `${Transfer:UserName}`
, `${Transfer:HomeDirectory}` , and `${Transfer:HomeBucket}` .

###### Note

For session policies, AWS Transfer Family stores the policy as a JSON blob,
instead of the Amazon Resource Name (ARN) of the policy. You save the policy as a
JSON blob and pass it in the `Policy` argument.

For an example of a session policy, see [Example session\
policy](../../../transfer/latest/userguide/session-policy.md) .

For more information, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the
_AWS Security Token Service API Reference_ .

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PosixProfile`

Specifies the full POSIX identity, including user ID ( `Uid`), group ID
( `Gid`), and any secondary groups IDs ( `SecondaryGids`), that
controls your users' access to your Amazon Elastic File System (Amazon EFS) file
systems. The POSIX permissions that are set on files and directories in your file system
determine the level of access your users get when transferring files into and out of
your Amazon EFS file systems.

_Required_: No

_Type_: [PosixProfile](aws-properties-transfer-user-posixprofile.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that controls your users' access to your Amazon S3
bucket or Amazon EFS file system. The policies attached to this role determine the level of access that you want to provide your users
when transferring files into and out of your Amazon S3 bucket or Amazon EFS file system. The IAM role should also contain a trust
relationship that allows the server to access your resources when servicing your users' transfer requests.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:.*role/\S+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerId`

A system-assigned unique identifier for a server instance. This is the specific server
that you added your user to.

_Required_: Yes

_Type_: String

_Pattern_: `^s-([0-9a-f]{17})$`

_Minimum_: `19`

_Maximum_: `19`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SshPublicKeys`

Specifies the public key portion of the Secure Shell (SSH) keys stored for the
described user.

###### Note

To delete the public key body, set its value to zero keys, as shown here:

`SshPublicKeys: []`

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs that can be used to group and search for users. Tags are metadata
attached to users for any purpose.

_Required_: No

_Type_: Array of [Tag](aws-properties-transfer-user-tag.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserName`

A unique string that identifies a user and is associated with a `ServerId`.
This user name must be a minimum of 3 and a maximum of 100 characters long. The
following are valid characters: a-z, A-Z, 0-9, underscore '\_', hyphen
'-', period '.', and at sign '@'. The user name can't
start with a hyphen, period, or at sign.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w][\w@.-]{2,99}$`

_Minimum_: `3`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the username, such as `transfer_user` .

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name associated with the user, in the form `
                arn:aws:transfer:region: account-id :user/
                    server-id / username` .

An example of a user ARN is:
`arn:aws:transfer:us-east-1:123456789012:user/user1` .

`ServerId`

The ID of the server to which the user is attached.

An example `ServerId` is `s-01234567890abcdef` .

`UserName`

A unique string that identifies a Transfer Family user account associated with a
server.

An example `UserName` is `transfer-user-1` .

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowDetails

HomeDirectoryMapEntry

All content copied from https://docs.aws.amazon.com/.
