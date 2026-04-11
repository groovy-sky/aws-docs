This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::Association InstanceAssociationOutputLocation

`InstanceAssociationOutputLocation` is a property of the [AWS::SSM::Association](../userguide/aws-resource-ssm-association.md) resource that specifies an Amazon S3 bucket
where you want to store the results of this association request.

For the minimal permissions required to enable Amazon S3 output for an
association, see [Creating\
associations](../../../systems-manager/latest/userguide/sysman-state-assoc.md) in the _Systems Manager User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Location" : S3OutputLocation
}

```

### YAML

```yaml

  S3Location:
    S3OutputLocation

```

## Properties

`S3Location`

`S3OutputLocation` is a property of the [InstanceAssociationOutputLocation](../userguide/aws-properties-ssm-association-instanceassociationoutputlocation.md) property that specifies an Amazon S3 bucket where you want to store the results of this request.

_Required_: No

_Type_: [S3OutputLocation](aws-properties-ssm-association-s3outputlocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSM::Association

S3OutputLocation

All content copied from https://docs.aws.amazon.com/.
