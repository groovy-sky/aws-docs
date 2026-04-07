This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DirectoryService::MicrosoftAD

The `AWS::DirectoryService::MicrosoftAD` resource specifies a Microsoft Active
Directory in AWS so that your directory users and groups can access the
AWS Management Console and AWS applications using their existing
credentials. For more information, see [AWS Managed Microsoft AD](https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_microsoft_ad.html) in the _Directory Service Admin Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DirectoryService::MicrosoftAD",
  "Properties" : {
      "CreateAlias" : Boolean,
      "Edition" : String,
      "EnableSso" : Boolean,
      "Name" : String,
      "Password" : String,
      "ShortName" : String,
      "VpcSettings" : VpcSettings
    }
}

```

### YAML

```yaml

Type: AWS::DirectoryService::MicrosoftAD
Properties:
  CreateAlias: Boolean
  Edition: String
  EnableSso: Boolean
  Name: String
  Password: String
  ShortName: String
  VpcSettings:
    VpcSettings

```

## Properties

`CreateAlias`

Specifies an alias for a directory and assigns the alias to the directory. The alias is
used to construct the access URL for the directory, such as
`http://<alias>.awsapps.com`. By default, CloudFormation does not
create an alias.

###### Important

After an alias has been created, it cannot be deleted or reused, so this operation
should only be used when absolutely necessary.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Edition`

AWS Managed Microsoft AD is available in two editions: `Standard` and
`Enterprise`. `Enterprise` is the default.

_Required_: No

_Type_: String

_Allowed values_: `Enterprise | Standard | Hybrid`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableSso`

Whether to enable single sign-on for a Microsoft Active Directory in AWS.
Single sign-on allows users in your directory to access certain AWS services
from a computer joined to the directory without having to enter their credentials separately.
If you don't specify a value, CloudFormation disables single sign-on by default.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The fully qualified domain name for the AWS Managed Microsoft AD directory, such as
`corp.example.com`. This name will resolve inside your VPC only. It does not need
to be publicly resolvable.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9]+[\\.-])+([a-zA-Z0-9])+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Password`

The password for the default administrative user named `Admin`.

If you need to change the password for the administrator account, see the [ResetUserPassword](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_ResetUserPassword.html) API call in the _Directory Service API Reference_.

_Required_: Yes

_Type_: String

_Pattern_: `(?=^.{8,64}$)((?=.*\d)(?=.*[A-Z])(?=.*[a-z])|(?=.*\d)(?=.*[^A-Za-z0-9\s])(?=.*[a-z])|(?=.*[^A-Za-z0-9\s])(?=.*[A-Z])(?=.*[a-z])|(?=.*\d)(?=.*[A-Z])(?=.*[^A-Za-z0-9\s]))^.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ShortName`

The NetBIOS name for your domain, such as `CORP`. If you don't specify a
NetBIOS name, it will default to the first part of your directory DNS. For example,
`CORP` for the directory DNS `corp.example.com`.

_Required_: No

_Type_: String

_Pattern_: `^[^\\/:*?"<>|.]+[^\\/:*?"<>|]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcSettings`

Specifies the VPC settings of the Microsoft AD directory server in AWS.

_Required_: Yes

_Type_: [VpcSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-directoryservice-microsoftad-vpcsettings.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic
function, `Ref` returns the resource ID.

In the following sample, the `Ref` function returns the ID of the
`myDirectory` directory, such as `d-12345ab592`.

`{ "Ref": "myDirectory" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Alias`

The alias for a directory. For example: `alias4-mydirectory-12345abcgmzsk` (if
you have the `CreateAlias` property set to true).

`DnsIpAddresses`

The IP addresses of the DNS servers for the directory, such as `[ "192.0.2.1",
        "192.0.2.2" ]`.

`Id`

The directory ID. For example: `d-12373a053a`.

## Examples

The following example creates a Microsoft Active Directory in AWS,
where the directory DNS name is `corp.example.com`:

### Create an AWS Managed Microsoft AD

#### JSON

```json

"myDirectory" : {
  "Type" : "AWS::DirectoryService::MicrosoftAD",
  "Properties" : {
    "Name" : "corp.example.com",
    "Password" : { "Ref" : "MicrosoftADPW" },
    "ShortName" : { "Ref" : "MicrosoftADShortName" },
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
  Type: AWS::DirectoryService::MicrosoftAD
  Properties:
    Name: "corp.example.com"
    Password:
      Ref: MicrosoftADPW
    ShortName:
      Ref: MicrosoftADShortName
    VpcSettings:
      SubnetIds:
        - Ref: subnetID1
        - Ref: subnetID2
      VpcId:
        Ref: vpcID
```

## See also

- [Getting Started\
with AWS Managed Microsoft AD](https://docs.aws.amazon.com/directoryservice/latest/admin-guide/ms_ad_getting_started.html) in the _Directory Service Admin Guide_..

- [CreateMicrosoftAD](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_CreateMicrosoftAD.html) in the _Directory Service API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Directory Service

VpcSettings
