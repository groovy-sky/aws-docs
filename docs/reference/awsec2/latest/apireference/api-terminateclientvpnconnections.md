# TerminateClientVpnConnections

Terminates active Client VPN endpoint connections. This action can be used to terminate a specific client connection, or up to five connections established by a specific user.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientVpnEndpointId**

The ID of the Client VPN endpoint to which the client is connected.

Type: String

Required: Yes

**ConnectionId**

The ID of the client connection to be terminated.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Username**

The name of the user who initiated the connection. Use this option to terminate all active connections for
the specified user. This option can only be used if the user has established up to five connections.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**clientVpnEndpointId**

The ID of the Client VPN endpoint.

Type: String

**connectionStatuses**

The current state of the client connections.

Type: Array of [TerminateConnectionStatus](api-terminateconnectionstatus.md) objects

**requestId**

The ID of the request.

Type: String

**username**

The user who established the terminated client connections.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example terminates a Client VPN endpoint connection.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=TerminateClientVpnConnections
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&ConnectionId=cvpn-connection-010b1282b7EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<TerminateClientVpnConnectionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
	<clientVpnEndpointId>cvpn-endpoint-00c5d11fc4EXAMPLE</clientVpnEndpointId>
	<connectionStatuses>
		<Item>
			<connectionId>cvpn-connection-010b1282b7EXAMPLE</connectionId>
		       <currentStatus>
				<code>terminating</code>
		       </currentStatus>
		       <previousStatus>
			       <code>active</code>
		       </previousStatus>
	       </Item>
    </connectionStatuses>
	<requestId>00d80748-708d-40f7-8635-f34acEXAMPLE</requestId>
</TerminateClientVpnConnectionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/terminateclientvpnconnections.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/terminateclientvpnconnections.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StopInstances

TerminateInstances
