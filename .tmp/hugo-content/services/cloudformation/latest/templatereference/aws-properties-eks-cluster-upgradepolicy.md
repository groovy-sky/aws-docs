This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster UpgradePolicy

The support policy to use for the cluster. Extended support allows you to remain on
specific Kubernetes versions for longer. Clusters in extended support have higher costs. The
default value is `EXTENDED`. Use `STANDARD` to disable extended
support.

[Learn more about EKS Extended Support in the _Amazon EKS User Guide_.](../../../eks/latest/userguide/extended-support-control.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SupportType" : String
}

```

### YAML

```yaml

  SupportType: String

```

## Properties

`SupportType`

If the cluster is set to `EXTENDED`, it will enter extended support at the
end of standard support. If the cluster is set to `STANDARD`, it will be
automatically upgraded at the end of standard support.

[Learn more about EKS Extended Support in the _Amazon EKS User Guide_.](../../../eks/latest/userguide/extended-support-control.md)

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | EXTENDED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ZonalShiftConfig

All content copied from https://docs.aws.amazon.com/.
