This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::SecurityConfiguration LakeFormationConfiguration

AWS Lake Formation related configuration inputs for the security
configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizedSessionTagValue" : String,
  "QueryAccessControlEnabled" : Boolean,
  "QueryEngineRoleArn" : String,
  "SecureNamespaceInfo" : SecureNamespaceInfo
}

```

### YAML

```yaml

  AuthorizedSessionTagValue: String
  QueryAccessControlEnabled: Boolean
  QueryEngineRoleArn: String
  SecureNamespaceInfo:
    SecureNamespaceInfo

```

## Properties

`AuthorizedSessionTagValue`

The session tag to authorize Amazon EMR on EKS for API calls to AWS Lake Formation.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryAccessControlEnabled`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryEngineRoleArn`

The query engine IAM role ARN that is tied to the secure Spark job. The
`QueryEngine` role assumes the `JobExecutionRole` to execute all
the Lake Formation calls.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):iam::\d{12}:role/.+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecureNamespaceInfo`

The namespace input of the system job.

_Required_: No

_Type_: [SecureNamespaceInfo](aws-properties-emrcontainers-securityconfiguration-securenamespaceinfo.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InTransitEncryptionConfiguration

LocalDiskEncryptionConfiguration

All content copied from https://docs.aws.amazon.com/.
