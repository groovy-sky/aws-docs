# EphemeralStorage

The amount of ephemeral storage to allocate for the task. This parameter is used to
expand the total amount of ephemeral storage available, beyond the default amount, for
tasks hosted on AWS Fargate. For more information, see [Using data volumes in\
tasks](../../../../services/amazonecs/latest/developerguide/using-data-volumes.md) in the _Amazon ECS Developer Guide;_.

###### Note

For tasks using the Fargate launch type, the task requires the following
platforms:

- Linux platform version `1.4.0` or later.

- Windows platform version `1.0.0` or later.

## Contents

**sizeInGiB**

The total amount, in GiB, of ephemeral storage to set for the task. The minimum
supported value is `21` GiB and the maximum supported value is
`200` GiB.

Type: Integer

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/ephemeralstorage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/ephemeralstorage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/ephemeralstorage.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnvironmentFile

ExecuteCommandConfiguration
