This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::ResourcePolicy

Use resource-based policies to grant permissions to other AWS accounts
or services to access your workspace.

Only Prometheus-compatible APIs can be used for workspace sharing. You can add
non-Prometheus-compatible APIs to the policy, but they will be ignored. For more
information, see [Prometheus-compatible APIs](../../../../reference/prometheus/latest/userguide/amp-apireference-prometheus-compatible-apis.md) in the _Amazon Managed Service for Prometheus User_
_Guide_.

If your workspace uses customer-managed AWS KMS keys for encryption, you
must grant the principals in your resource-based policy access to those AWS KMS keys. You can do this by creating AWS KMS grants. For more information,
see [CreateGrant](../../../../reference/kms/latest/apireference/api-creategrant.md) in the _AWS KMS API Reference_
and [Encryption at rest](../../../prometheus/latest/userguide/encryption-at-rest-amazon-service-prometheus.md) in the _Amazon Managed Service for Prometheus User_
_Guide_.

For more information about working with IAM, see [Using Amazon Managed Service for Prometheus with IAM](../../../prometheus/latest/userguide/security-iam-service-with-iam.md) in the _Amazon Managed Service for Prometheus User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::APS::ResourcePolicy",
  "Properties" : {
      "PolicyDocument" : String,
      "WorkspaceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::APS::ResourcePolicy
Properties:
  PolicyDocument: String
  WorkspaceArn: String

```

## Properties

`PolicyDocument`

The JSON to use as the Resource-based Policy.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkspaceArn`

An ARN identifying a Workspace.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):aps:[a-z0-9-]+:[0-9]+:workspace/[a-zA-Z0-9-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::APS::RuleGroupsNamespace

All content copied from https://docs.aws.amazon.com/.
