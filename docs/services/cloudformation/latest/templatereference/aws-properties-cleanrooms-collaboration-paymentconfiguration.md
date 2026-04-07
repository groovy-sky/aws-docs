This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::Collaboration PaymentConfiguration

An object representing the collaboration member's payment responsibilities set by the
collaboration creator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JobCompute" : JobComputePaymentConfig,
  "MachineLearning" : MLPaymentConfig,
  "QueryCompute" : QueryComputePaymentConfig
}

```

### YAML

```yaml

  JobCompute:
    JobComputePaymentConfig
  MachineLearning:
    MLPaymentConfig
  QueryCompute:
    QueryComputePaymentConfig

```

## Properties

`JobCompute`

The compute configuration for the job.

_Required_: No

_Type_: [JobComputePaymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-collaboration-jobcomputepaymentconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MachineLearning`

An object representing the collaboration member's machine learning payment responsibilities set by the
collaboration creator.

_Required_: No

_Type_: [MLPaymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-collaboration-mlpaymentconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryCompute`

The collaboration member's payment responsibilities set by the collaboration creator for
query compute costs.

_Required_: Yes

_Type_: [QueryComputePaymentConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanrooms-collaboration-querycomputepaymentconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModelTrainingPaymentConfig

QueryComputePaymentConfig
