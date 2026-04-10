This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket OwnershipControls

Specifies the container element for Object Ownership rules.

S3 Object Ownership is an Amazon S3 bucket-level setting that you can use to disable
access control lists (ACLs) and take ownership of every object in your bucket, simplifying
access management for data stored in Amazon S3. For more information, see [Controlling\
ownership of objects and disabling ACLs](../../../s3/latest/userguide/about-object-ownership.md) in the _Amazon S3 User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ OwnershipControlsRule, ... ]
}

```

### YAML

```yaml

  Rules:
    - OwnershipControlsRule

```

## Properties

`Rules`

Specifies the container element for Object Ownership rules.

_Required_: Yes

_Type_: Array of [OwnershipControlsRule](aws-properties-s3-bucket-ownershipcontrolsrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Object ownership - BucketOwnerEnforced](#aws-properties-s3-bucket-ownershipcontrols--examples--Object_ownership_-_BucketOwnerEnforced)

- [Object Ownership - BucketOwnerPreferred](#aws-properties-s3-bucket-ownershipcontrols--examples--Object_Ownership_-_BucketOwnerPreferred)

### Object ownership - BucketOwnerEnforced

The following examples show Object Ownership set to
`BucketOwnerEnforced`.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "OwnershipControls": {
                    "Rules": [
                        {
                            "ObjectOwnership": "BucketOwnerEnforced"
                        }
                    ]
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
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      OwnershipControls:
        Rules:
          - ObjectOwnership: BucketOwnerEnforced
```

### Object Ownership - BucketOwnerPreferred

The following examples show Object Ownership set to
`BucketOwnerPreferred`.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "OwnershipControls": {
                    "Rules": [
                        {
                            "ObjectOwnership": "BucketOwnerPreferred"
                        }
                    ]
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
  S3Bucket:
    Type: 'AWS::S3::Bucket'
    Properties:
      OwnershipControls:
        Rules:
          - ObjectOwnership: BucketOwnerPreferred
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectLockRule

OwnershipControlsRule

All content copied from https://docs.aws.amazon.com/.
