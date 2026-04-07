This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Membership MembershipPaymentConfiguration

An object representing the payment responsibilities accepted by the collaboration
member.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JobCompute" : MembershipJobComputePaymentConfig,
  "MachineLearning" : MembershipMLPaymentConfig,
  "QueryCompute" : MembershipQueryComputePaymentConfig
}

```

### YAML

```yaml

  JobCompute:
    MembershipJobComputePaymentConfig
  MachineLearning:
    MembershipMLPaymentConfig
  QueryCompute:
    MembershipQueryComputePaymentConfig

```

## Properties

`JobCompute`

The payment responsibilities accepted by the collaboration member for job compute
costs.

_Required_: No

_Type_: [MembershipJobComputePaymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-membership-membershipjobcomputepaymentconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MachineLearning`

The payment responsibilities accepted by the collaboration member for machine learning
costs.

_Required_: No

_Type_: [MembershipMLPaymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-membership-membershipmlpaymentconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryCompute`

The payment responsibilities accepted by the collaboration member for query compute
costs.

_Required_: Yes

_Type_: [MembershipQueryComputePaymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-membership-membershipquerycomputepaymentconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MembershipModelTrainingPaymentConfig

MembershipProtectedJobOutputConfiguration
