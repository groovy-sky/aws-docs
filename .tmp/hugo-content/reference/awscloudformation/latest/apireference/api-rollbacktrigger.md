# RollbackTrigger

A rollback trigger CloudFormation monitors during creation and updating of stacks. If any of the
alarms you specify goes to ALARM state during the stack operation or within the specified
monitoring period afterwards, CloudFormation rolls back the entire stack operation.

## Contents

**Arn**

The Amazon Resource Name (ARN) of the rollback trigger.

If a specified trigger is missing, the entire stack operation fails and is rolled
back.

Type: String

Required: Yes

**Type**

The resource type of the rollback trigger. Specify either [AWS::CloudWatch::Alarm](../../../../services/cloudformation/latest/templatereference/aws-resource-cloudwatch-alarm.md) or [AWS::CloudWatch::CompositeAlarm](../../../../services/cloudformation/latest/templatereference/aws-resource-cloudwatch-compositealarm.md) resource types.

Type: String

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/rollbacktrigger.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/rollbacktrigger.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/rollbacktrigger.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RollbackConfiguration

ScanFilter

All content copied from https://docs.aws.amazon.com/.
