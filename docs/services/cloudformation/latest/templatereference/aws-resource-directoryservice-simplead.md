This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectoryService::SimpleAD

The `AWS::DirectoryService::SimpleAD` resource specifies an Directory Service Simple Active Directory (Simple AD) in AWS so that your
directory users and groups can access the AWS Management Console and AWS
applications using their existing credentials. Simple AD is a Microsoft
Active Directory–compatible directory. For more information, see [Simple Active\
Directory](https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_simple_ad.html) in the _Directory Service Admin Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DirectoryService::SimpleAD",
  "Properties" : {
      "CreateAlias" : Boolean,
      "Description" : String,
      "EnableSso" : Boolean,
      "Name" : String,
      "Password" : String,
      "ShortName" : String,
      "Size" : String,
      "Tags" : [ Tag, ... ],
      "VpcSettings" : VpcSettings
    }
}

```

### YAML

```yaml

Type: AWS::DirectoryService::SimpleAD
Properties:
  CreateAlias: Boolean
  Description: String
  EnableSso: Boolean
  Name: String
  Password: String
  ShortName: String
  Size: String
  Tags:
    - Tag
  VpcSettings:
    VpcSettings

```

## Properties

`CreateAlias`

If set to `true`, specifies an alias for a directory and assigns the alias to
the directory. The alias is used to construct the access URL for the directory, such as
`http://<alias>.awsapps.com`. By default, this property is set to
`false`.

###### Important

After an alias has been created, it cannot be deleted or reused, so this operation
should only be used when absolutely necessary.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description for the directory.

_Required_: No

_Type_: String

_Pattern_: `^([a-zA-Z0-9_])[\\a-zA-Z0-9_@#%*+=:?./!\s-]*$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableSso`

Whether to enable single sign-on for a directory. If you don't specify a value, CloudFormation disables single sign-on by default.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The fully qualified name for the directory, such as `corp.example.com`.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Password`

The password for the directory administrator. The directory creation process creates a
directory administrator account with the user name `Administrator` and this
password.

If you need to change the password for the administrator account, see the [ResetUserPassword](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_ResetUserPassword.html) API call in the _Directory Service API Reference_.

_Required_: No

_Type_: String

_Pattern_: `(?=^.{8,64}$)((?=.*\d)(?=.*[A-Z])(?=.*[a-z])|(?=.*\d)(?=.*[^A-Za-z0-9\s])(?=.*[a-z])|(?=.*[^A-Za-z0-9\s])(?=.*[A-Z])(?=.*[a-z])|(?=.*\d)(?=.*[A-Z])(?=.*[^A-Za-z0-9\s]))^.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ShortName`

The NetBIOS name of the directory, such as `CORP`.

_Required_: No

_Type_: String

_Pattern_: `^[^\\/:*?"<>|.]+[^\\/:*?"<>|]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Size`

The size of the directory. For valid values, see [CreateDirectory](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_CreateDirectory.html) in
the _Directory Service API Reference_.

_Required_: Yes

_Type_: String

_Allowed values_: `Small | Large`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-directoryservice-simplead-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcSettings`

A [DirectoryVpcSettings](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_DirectoryVpcSettings.html) object that contains additional information for the
operation.

_Required_: Yes

_Type_: [VpcSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-directoryservice-simplead-vpcsettings.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic
function, `Ref` returns the resource ID.

In the following sample, the `Ref` function returns the ID of the
`myDirectory` directory, such as `d-1a2b3c4d5e`.

`{ "Ref": "myDirectory" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Alias`

The alias for a directory. For example: `alias4-mydirectory-12345abcgmzsk` (if
you have the `CreateAlias` property set to true).

`DirectoryId`

The directory ID. For example: `d-12373a053a`.

`DnsIpAddresses`

The IP addresses of the DNS servers for the directory, such as `[ "172.31.3.154",
        "172.31.63.203" ]`.

## Examples

The following example creates a Simple AD directory, where the
directory DNS name is `corp.example.com`:

### Create a Simple AD Directory

#### JSON

```json

"myDirectory" : {
  "Type" : "AWS::DirectoryService::SimpleAD",
  "Properties" : {
    "Name" : "corp.example.com",
    "Password" : { "Ref" : "SimpleADPW" },
    "Size" : "Small",
    "VpcSettings" : {
      "SubnetIds" : [ { "Ref" : "subnetID1" }, { "Ref" : "subnetID2" } ],
      "VpcId" : { "Ref" : "vpcID" }
    }
  }
}
```

#### YAML

```yaml

myDirectory:
  Type: AWS::DirectoryService::SimpleAD
  Properties:
    Name: "corp.example.com"
    Password:
      Ref: SimpleADPW
    Size: "Small"
    VpcSettings:
      SubnetIds:
        - Ref: subnetID1
        - Ref: subnetID2
      VpcId:
        Ref: vpcID
```

## See also

- [Getting\
Started with Simple AD](https://docs.aws.amazon.com/directoryservice/latest/admin-guide/simple_ad_getting_started.html) in the _Directory Service Admin Guide_..

- [CreateDirectory](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_CreateDirectory.html) in the _Directory Service API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcSettings

Tag
