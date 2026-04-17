---
title: "BlueGreenDeploymentTask"
---

# BlueGreenDeploymentTask

Details about a task for a blue/green deployment.

For more information, see [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/userguide/blue-green-deployments.md) in the _Amazon RDS User_
_Guide_ and [Using Amazon RDS\
Blue/Green Deployments for database updates](../../../../services/amazonrds/latest/aurorauserguide/blue-green-deployments.md) in the _Amazon Aurora_
_User Guide_.

## Contents

###### Note

In the following list, the required parameters are described first.

**Name**

The name of the blue/green deployment task.

Type: String

Required: No

**Status**

The status of the blue/green deployment task.

Valid Values:

- `PENDING` \- The resource is being prepared for deployment.

- `IN_PROGRESS` \- The resource is being deployed.

- `COMPLETED` \- The resource has been deployed.

- `FAILED` \- Deployment of the resource failed.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/rds-2014-10-31/BlueGreenDeploymentTask)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/rds-2014-10-31/BlueGreenDeploymentTask)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/rds-2014-10-31/BlueGreenDeploymentTask)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlueGreenDeployment

Certificate

All content copied from https://docs.aws.amazon.com/.
