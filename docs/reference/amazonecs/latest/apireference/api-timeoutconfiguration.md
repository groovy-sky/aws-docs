# TimeoutConfiguration

An object that represents the timeout configurations for Service Connect.

###### Note

If `idleTimeout` is set to a time that is less than
`perRequestTimeout`, the connection will close when the
`idleTimeout` is reached and not the
`perRequestTimeout`.

## Contents

**idleTimeoutSeconds**

The amount of time in seconds a connection will stay active while idle. A value of
`0` can be set to disable `idleTimeout`.

The `idleTimeout` default for
`HTTP`/ `HTTP2`/ `GRPC` is 5 minutes.

The `idleTimeout` default for `TCP` is 1 hour.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2147483647.

Required: No

**perRequestTimeoutSeconds**

The amount of time waiting for the upstream to respond with a complete response per
request. A value of `0` can be set to disable `perRequestTimeout`.
`perRequestTimeout` can only be set if Service Connect
`appProtocol` isn't `TCP`. Only `idleTimeout` is
allowed for `TCP` `appProtocol`.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2147483647.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecs-2014-11-13/timeoutconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecs-2014-11-13/timeoutconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecs-2014-11-13/timeoutconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TaskVolumeConfiguration

Tmpfs
