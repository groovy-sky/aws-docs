This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::ModuleDefaultVersion

Specifies the default version of a module. The default version of the module will be
used in CloudFormation operations for this account and Region.

For more information, see [Create reusable resource\
configurations that can be included across templates with CloudFormation\
modules](../userguide/modules.md) in the _CloudFormation User Guide_.

For information about the CloudFormation registry, see [Managing\
extensions with the CloudFormation registry](../userguide/registry.md) in the
_CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::ModuleDefaultVersion",
  "Properties" : {
      "Arn" : String,
      "ModuleName" : String,
      "VersionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::ModuleDefaultVersion
Properties:
  Arn: String
  ModuleName: String
  VersionId: String

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the module version to set as the default
version.

Conditional: You must specify either `Arn`, or `ModuleName` and
`VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/module/.+/[0-9]{8}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModuleName`

The name of the module.

Conditional: You must specify either `Arn`, or `ModuleName` and
`VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::MODULE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VersionId`

The ID for the specific version of the module.

Conditional: You must specify either `Arn`, or `ModuleName` and
`VersionId`.

_Required_: Conditional

_Type_: String

_Pattern_: `^[0-9]{8}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the module
version.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Remarks

Considerations when managing the default module version:

- The first module version to be registered in an account and Region remains
the default version CloudFormation uses, unless and until you
explicitly sets another version as the default.

- For ease of determining which module version is the default version, we
recommend that you only include a single
`AWS::CloudFormation::ModuleDefaultVersion` resource for a
given module in a template.

- If you delete an `AWS::CloudFormation::ModuleVersion` resource,
either by deleting it from a stack or deleting the entire stack, CloudFormation marks the corresponding module version as
`DEPRECATED`. For more information on deprecating module
versions, see [DeregisterType](../../../../reference/awscloudformation/latest/apireference/api-deregistertype.md) in the _CloudFormation API_
_Reference_.

- If you attempt to delete an
`AWS::CloudFormation::ModuleVersion` resource that represents
the default version, the operation will fail if there are other active
versions.

## Examples

### Specifying the default module version

The following example registers two versions of a module, and then sets the
second version as the default version for CloudFormation to use. Note
that the example uses the `DependsOn` attribute to ensure that
CloudFormation provisions version one before version two.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "ModuleVersion1": {
            "Type": "AWS::CloudFormation::ModuleVersion",
            "Properties": {
                "ModuleName": "My::Sample::Test::MODULE",
                "ModulePackage": "s3://amzn-s3-demo-bucket/sample-module-package-v1.zip"
            }
        },
        "ModuleVersion2": {
            "Type": "AWS::CloudFormation::ModuleVersion",
            "Properties": {
                "ModuleName": "My::Sample::Test::MODULE",
                "ModulePackage": "s3://amzn-s3-demo-bucket/sample-module-package-v2.zip"
            },
            "DependsOn": "ModuleVersion1"
        },
        "ModuleDefaultVersion": {
            "Type": "AWS::CloudFormation::ModuleDefaultVersion",
            "Properties": {
                "Arn": {
                    "Ref": "ModuleVersion2"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  ModuleVersion1:
    Type: AWS::CloudFormation::ModuleVersion
    Properties:
      ModuleName: 'My::Sample::Test::MODULE'
      ModulePackage: 's3://amzn-s3-demo-bucket/sample-module-package-v1.zip'
  ModuleVersion2:
    Type: AWS::CloudFormation::ModuleVersion
    Properties:
      ModuleName: 'My::Sample::Test::MODULE'
      ModulePackage: 's3://amzn-s3-demo-bucket/sample-module-package-v2.zip'
    DependsOn: ModuleVersion1
  ModuleDefaultVersion:
    Type: AWS::CloudFormation::ModuleDefaultVersion
    Properties:
      Arn: !Ref ModuleVersion2
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFormation::Macro

AWS::CloudFormation::ModuleVersion

All content copied from https://docs.aws.amazon.com/.
