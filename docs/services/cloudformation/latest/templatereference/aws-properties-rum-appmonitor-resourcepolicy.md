This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RUM::AppMonitor ResourcePolicy

Use this structure to assign a resource-based policy to a CloudWatch RUM app monitor to control access to it. Each app monitor can
have one resource-based policy. The maximum size of the policy is 4 KB. To learn more about using resource policies with RUM, see [Using resource-based policies with CloudWatch RUM](../../../amazoncloudwatch/latest/monitoring/cloudwatch-rum-resource-policies.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PolicyDocument" : String,
  "PolicyRevisionId" : String
}

```

### YAML

```yaml

  PolicyDocument: String
  PolicyRevisionId: String

```

## Properties

`PolicyDocument`

The JSON to use as the resource policy. The document can be up to 4 KB in size. For more information about the contents and syntax
for this policy, see [Using resource-based policies with CloudWatch RUM](../../../amazoncloudwatch/latest/monitoring/cloudwatch-rum-resource-policies.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyRevisionId`

A string value that you can use to conditionally update your policy. You can provide the revision ID of your existing policy to make mutating requests against that policy.

When you assign a policy revision ID, then later requests about that policy will be rejected with an `InvalidPolicyRevisionIdException` error if they don't provide the correct current revision ID.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetricDestination

Tag

All content copied from https://docs.aws.amazon.com/.
