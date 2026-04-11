This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CodeSigningConfig CodeSigningPolicies

Code signing configuration [policies](../../../lambda/latest/dg/configuration-codesigning.md#config-codesigning-policies) specify the validation failure action for signature mismatch or
expiry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UntrustedArtifactOnDeployment" : String
}

```

### YAML

```yaml

  UntrustedArtifactOnDeployment: String

```

## Properties

`UntrustedArtifactOnDeployment`

Code signing configuration policy for deployment validation failure. If you set the policy to
`Enforce`, Lambda blocks the deployment request if signature validation checks
fail. If you set the policy to `Warn`, Lambda allows the deployment and issues a
new Amazon CloudWatch metric ( `SignatureValidationErrors`) and also stores the
warning in the CloudTrail log.

Default value: `Warn`

_Required_: Yes

_Type_: String

_Allowed values_: `Warn | Enforce`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AllowedPublishers

Tag

All content copied from https://docs.aws.amazon.com/.
