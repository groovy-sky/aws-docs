This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster JwtAuth

The JWT authentication configuration for Slurm REST API access.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JwtKey" : JwtKey
}

```

### YAML

```yaml

  JwtKey:
    JwtKey

```

## Properties

`JwtKey`

The JWT key for Slurm REST API authentication.

_Required_: No

_Type_: [JwtKey](aws-properties-pcs-cluster-jwtkey.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ErrorInfo

JwtKey

All content copied from https://docs.aws.amazon.com/.
