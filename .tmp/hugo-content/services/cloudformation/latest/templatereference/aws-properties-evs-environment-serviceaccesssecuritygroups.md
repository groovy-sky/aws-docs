This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EVS::Environment ServiceAccessSecurityGroups

The security groups that allow traffic between the Amazon EVS control plane and your VPC for Amazon EVS service access.
If a security group is not specified, Amazon EVS uses the default security group in your account for service access.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroups" : [ String, ... ]
}

```

### YAML

```yaml

  SecurityGroups:
    - String

```

## Properties

`SecurityGroups`

The security groups that allow service access.

_Required_: No

_Type_: Array of String

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Secret

Tag

All content copied from https://docs.aws.amazon.com/.
