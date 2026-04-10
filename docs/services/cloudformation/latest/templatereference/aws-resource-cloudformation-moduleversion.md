This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::ModuleVersion

The `AWS::CloudFormation::ModuleVersion` resource registers the specified
version of the module with the CloudFormation registry. Registering a module
makes it available for use in CloudFormation templates in your AWS account and Region.

For more information, see [Create reusable resource\
configurations that can be included across templates with CloudFormation\
modules](../userguide/modules.md) in the _CloudFormation User_
_Guide_.

For information about the CloudFormation registry, see [Managing\
extensions with the CloudFormation registry](../userguide/registry.md) in the
_CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::ModuleVersion",
  "Properties" : {
      "ModuleName" : String,
      "ModulePackage" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::ModuleVersion
Properties:
  ModuleName: String
  ModulePackage: String

```

## Properties

`ModuleName`

The name of the module being registered.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::MODULE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModulePackage`

A URL to the S3 bucket for the package that contains the template fragment and schema
files for the module version to register.

For more information, see [Module structure\
and requirements](../../../cloudformation-cli/latest/userguide/modules-structure.md) in the _CloudFormation Command Line Interface_
_(CLI) User Guide_.

###### Note

To register the module version, you must have `s3:GetObject`
permissions to access the S3 objects.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the module
version.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the module.

`Description`

The description of the module.

`DocumentationUrl`

The URL of a page providing detailed documentation for this module.

`IsDefaultVersion`

Whether the specified module version is set as the default version.

This applies only to private extensions you have registered in your account, and
extensions published by AWS. For public third-party extensions, whether
they are activated in your account, CloudFormation returns
`null`.

`Schema`

The schema that defines the module.

`TimeCreated`

When the specified private module version was registered or activated in your
account.

`VersionId`

The ID of this version of the module.

`Visibility`

The visibility level that determines who can see and use this module in CloudFormation operations:

- `PRIVATE`: The module is only visible and usable within the account
where it was registered. CloudFormation automatically marks any modules
you register as `PRIVATE`.

- `PUBLIC`: The module is publicly visible and usable within any
AWS account.

## Remarks

Considerations when managing module versions:

- The account in which you register the module version must have permission
to access the S3 bucket in which the module package resides.

- The first module version to be registered in an account and region remains
the default version CloudFormation uses, unless and until you
explicitly sets another version as the default. To specify a module version
as the default version, use the [AWS::CloudFormation::ModuleDefaultVersion](aws-resource-cloudformation-moduledefaultversion.md) resource.

- If your template contains multiple versions of the same module, we
strongly recommend using the `DependsOn` attribute to explicitly
set the order in which the versions are registered.

- If you delete an `AWS::CloudFormation::ModuleVersion` resource,
either by deleting it from a stack or deleting the entire stack, CloudFormation marks the corresponding module version as
`DEPRECATED`.

If you attempt to delete an
`AWS::CloudFormation::ModuleVersion` resource that represent
the default version, the operation will fail if there are other active
versions.

For more information on deprecating module versions, see [DeregisterType](../../../../reference/awscloudformation/latest/apireference/api-deregistertype.md) in the _CloudFormation API_
_Reference_.

- You can't edit a module version. Updating an
`AWS::CloudFormation::ModuleVersion` resource results in a
new module version being registered in the CloudFormation
registry.

## Examples

- [Registering a module version](#aws-resource-cloudformation-moduleversion--examples--Registering_a_module_version)

- [Registering multiple module versions](#aws-resource-cloudformation-moduleversion--examples--Registering_multiple_module_versions)

### Registering a module version

The following example registers a module version. If this is the only version
of the module registered in this account and Region, CloudFormation
sets this version as the default version.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "ModuleVersion": {
            "Type": "AWS::CloudFormation::ModuleVersion",
            "Properties": {
                "ModuleName": "My::Sample::Test::MODULE",
                "ModulePackage": "s3://my-sample-moduleversion-bucket/sample-module-package-v1.zip"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  ModuleVersion:
    Type: 'AWS::CloudFormation::ModuleVersion'
    Properties:
      ModuleName: 'My::Sample::Test::MODULE'
      ModulePackage: 's3://my-sample-moduleversion-bucket/sample-module-package-v1.zip'
```

### Registering multiple module versions

The following example registers two versions of a module. Note the
following:

- The example uses the `DependsOn` attribute to ensure that
CloudFormation provisions version one before version
two.

- CloudFormation sets version one of the module as the default
version, as it's registered first. (This assumes no other versions of
the module are currently registered in this account and Region.)

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "ModuleVersion1": {
            "Type": "AWS::CloudFormation::ModuleVersion",
            "Properties": {
                "ModuleName": "My::Sample::Test::MODULE",
                "ModulePackage": "s3://my-sample-moduleversion-bucket/sample-module-package-v1.zip"
            }
        },
        "ModuleVersion2": {
            "Type": "AWS::CloudFormation::ModuleVersion",
            "Properties": {
                "ModuleName": "My::Sample::Test::MODULE",
                "ModulePackage": "s3://my-sample-moduleversion-bucket/sample-module-package-v2.zip"
            },
            "DependsOn": "ModuleVersion1"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  ModuleVersion1:
    Type: 'AWS::CloudFormation::ModuleVersion'
    Properties:
      ModuleName: 'My::Sample::Test::MODULE'
      ModulePackage: 's3://my-sample-moduleversion-bucket/sample-module-package-v1.zip'
  ModuleVersion2:
    Type: 'AWS::CloudFormation::ModuleVersion'
    Properties:
      ModuleName: 'My::Sample::Test::MODULE'
      ModulePackage: 's3://my-sample-moduleversion-bucket/sample-module-package-v2.zip'
    DependsOn: ModuleVersion1
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFormation::ModuleDefaultVersion

AWS::CloudFormation::PublicTypeVersion

All content copied from https://docs.aws.amazon.com/.
