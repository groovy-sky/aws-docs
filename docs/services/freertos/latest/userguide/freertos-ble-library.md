# Bluetooth Low Energy library

###### Important

This library is hosted on the Amazon-FreeRTOS repository which is deprecated. We recommend
that you [start here](freertos-getting-started-modular.md) when you
create a new project. If you already have an existing FreeRTOS project based on the now
deprecated Amazon-FreeRTOS repository, see the [Amazon-FreeRTOS Github Repository Migration Guide](github-repo-migration.md).

## Overview

FreeRTOS supports publishing and subscribing to Message Queuing Telemetry Transport
(MQTT) topics over Bluetooth Low Energy through a proxy device, such as a mobile
phone. With the FreeRTOS
[Bluetooth \
Low Energy](https://docs.aws.amazon.com/freertos/latest/lib-ref/html2/ble/index.html) (BLE) library, your microcontroller can securely communicate with
the AWS IoT MQTT broker.

![BLE devices connecting to AWS IoT Core via MQTT/HTTP/Websocket through AWS Cognito.](https://docs.aws.amazon.com/images/freertos/latest/userguide/images/blediagram.jpg)

Using the Mobile SDKs for FreeRTOS Bluetooth Devices, you can write native mobile
applications that communicate with the embedded applications on your microcontroller
over BLE. For more information about the mobile SDKs, see
[Mobile SDKs for FreeRTOS Bluetooth devices](https://docs.aws.amazon.com/freertos/latest/userguide/freertos-ble-mobile.html).

The FreeRTOS BLE library includes services for configuring Wi-Fi networks, transferring
large amounts of data, and providing network abstractions over BLE. The FreeRTOS BLE
library also includes middleware and lower-level APIs for more direct control over
your BLE stack.

## Architecture

Three layers make up the FreeRTOS
BLE
library: services, middleware, and low-level wrappers.

![Cloud architecture layers: User Application, Services, Middleware, Low-level Wrappers, Manufacturer BLE Stack.](https://docs.aws.amazon.com/images/freertos/latest/userguide/images/ble-architecture.png)

### Services

The FreeRTOS BLE services layer consists of four Generic Attribute (GATT) services that leverage
the middleware APIs:

- Device information

- Wi-Fi provisioning

- Network abstraction

- Large object transfer

#### Device information

The Device information service gathers details about your microcontroller, including:

- The version of FreeRTOS that your device is using.

- The AWS IoT endpoint of the account for which the device is registered.

- Bluetooth Low Energy Maximum Transmission Unit (MTU).

#### Wi-Fi provisioning

The Wi-Fi provisioning service enables microcontrollers with Wi-Fi
capabilities to do the following:

- List networks in range.

- Save networks and network credentials to flash memory.

- Set network priority.

- Delete networks and network credentials from flash memory.

#### Network abstraction

The network abstraction service abstracts the network connection type for applications.
A common API interacts with your device's Wi-Fi, Ethernet, and Bluetooth Low Energy
hardware stack, enabling an application to be compatible with multiple connection types.

#### Large Object Transfer

The Large Object Transfer service sends data to, and receives data from, a client.
Other services, like Wi-Fi provisioning and Network abstraction, use the Large Object
Transfer service to send and receive data. You can also use the Large Object Transfer
API to directly interact with the service.

#### MQTT over BLE

MQTT over BLE contains the GATT profile for creating an MQTT proxy service
over BLE.
The
MQTT proxy service allows an MQTT client to communicate with the AWS MQTT
broker through a gateway device. For example, you can use the proxy service to
connect a device running FreeRTOS to AWS MQTT through a smartphone
app. The BLE device is the GATT server and exposes services
and characteristics for the gateway device.
The
GATT server uses these exposed services and characteristics to
perform MQTT operations with the cloud
for
that device. For more details, refer to [Appendix A: MQTT over BLE GATT profile](#freertos-ble-gatt-profile).

### Middleware

FreeRTOS Bluetooth Low Energy middleware is an abstraction from the lower-level APIs. The middleware APIs
make up a more user-friendly interface to the Bluetooth Low Energy stack.

Using middleware APIs, you can register several callbacks, across multiple layers, to a single event.
Initializing the Bluetooth Low Energy middleware also initializes services and starts advertising.

#### Flexible callback subscription

Suppose your Bluetooth Low Energy hardware disconnects, and the MQTT over Bluetooth Low Energy service needs to detect the
disconnection. An application that you wrote might also need to detect
the same disconnection event. The Bluetooth Low Energy middleware can route the event to
different parts of the code where you have registered callbacks, without
making the higher layers compete for lower-level resources.

### Low-level wrappers

The low-level FreeRTOS Bluetooth Low Energy wrappers are an abstraction from the manufacturer's Bluetooth Low Energy stack.
Low-level wrappers offer a common set of APIs for direct control over
the hardware. The low-level APIs optimize RAM usage, but are limited in
functionality.

Use the Bluetooth Low Energy service APIs to interact with the Bluetooth Low Energy services.
The service APIs demand more resources than the low-level APIs.

## Dependencies and requirements

The Bluetooth Low Energy library has the following direct dependencies:

- [Linear \
Containers](https://docs.aws.amazon.com/freertos/latest/lib-ref/c-sdk/linear_containers/index.html) library

- A platform layer that interfaces with the operating system for thread management, timers, clock functions, and network access.

![Architecture diagram showing components: BLE, List/Queue, Network, and Clock, with directional arrows indicating interactions.](https://docs.aws.amazon.com/images/freertos/latest/userguide/images/ble-dependencies.png)

Only the Wi-Fi Provisioning service has FreeRTOS library dependencies:

GATT ServiceDependencyWi-Fi Provisioning[Wi-Fi library](freertos-wifi.md)

To communicate with the AWS IoT MQTT broker, you must have an AWS account and you must
register your devices as AWS IoT things. For more information about setting up,
see the [AWS IoT Developer Guide](https://docs.aws.amazon.com/iot/latest/developerguide).

FreeRTOS Bluetooth Low Energy uses Amazon Cognito for user authentication on your mobile device. To use MQTT proxy
services, you must create an Amazon Cognito identity and user pools. Each Amazon Cognito Identity must
have the appropriate policy attached to it. For more information, see the [Amazon Cognito Developer Guide](https://docs.aws.amazon.com/cognito/latest/developerguide).

## Library configuration file

Applications that use the FreeRTOS MQTT over Bluetooth Low Energy service must provide an `iot_ble_config.h` header file,
in which configuration parameters are defined. Undefined configuration parameters take the default values specified in
`iot_ble_config_defaults.h`.

Some important configuration parameters include:

**`IOT_BLE_ADD_CUSTOM_SERVICES`**

Allows users to create their own services.

**`IOT_BLE_SET_CUSTOM_ADVERTISEMENT_MSG`**

Allows users to customize the advertisement and scan response messages.

For more information, see
[Bluetooth Low Energy API Reference](https://docs.aws.amazon.com/freertos/latest/lib-ref/html2/ble/index.html).

## Optimization

When optimizing your board's performance, consider the following:

- Low-level APIs use less RAM, but offer limited functionality.

- You can set the `bleconfigMAX_NETWORK` parameter in the
`iot_ble_config.h` header file to a lower value to reduce the amount of stack consumed.

- You can increase the MTU size to its maximum value to limit message buffering, and make code run faster and consume less RAM.

## Usage restrictions

By default, the FreeRTOS Bluetooth Low Energy library sets the `eBTpropertySecureConnectionOnly` property to TRUE, which places the device in a Secure Connections Only mode. As specified by
the [Bluetooth Core Specification](https://www.bluetooth.com/specifications/bluetooth-core-specification) v5.0, Vol 3, Part C, 10.2.4,
when a device is in a Secure Connections Only mode, the highest LE security mode 1 level, level 4, is required for access to any attribute that has permissions higher than the
lowest LE security mode 1 level, level 1. At the LE security mode 1 level 4, a device must have input and output capabilities for numeric comparison.

Here are the supported modes, and their associated properties:

**Mode 1, Level 1 (No security)**

```c

/* Disable numeric comparison */
#define IOT_BLE_ENABLE_NUMERIC_COMPARISON        ( 0 )
#define IOT_BLE_ENABLE_SECURE_CONNECTION         ( 0 )
#define IOT_BLE_INPUT_OUTPUT                     ( eBTIONone )
#define IOT_BLE_ENCRYPTION_REQUIRED               ( 0 )
```

**Mode 1, Level 2 (Unauthenticated pairing with encryption)**

```c

#define IOT_BLE_ENABLE_NUMERIC_COMPARISON        ( 0 )
#define IOT_BLE_ENABLE_SECURE_CONNECTION         ( 0 )
#define IOT_BLE_INPUT_OUTPUT                     ( eBTIONone )
```

**Mode 1, Level 3 (Authenticated pairing with encryption)**

This mode is not supported.

**Mode 1, Level 4 (Authenticated LE Secure Connections pairing with encryption)**

This mode is supported by default.

For information about LE security modes, see the
[Bluetooth Core Specification](https://www.bluetooth.com/specifications/bluetooth-core-specification) v5.0, Vol 3, Part C, 10.2.1.

## Initialization

If your application interacts with the Bluetooth Low Energy stack through middleware, you only need to initialize the middleware. Middleware takes care of initializing the lower layers of the stack.

### Middleware

**To initialize the middleware**

1. Initialize any Bluetooth Low Energy hardware drivers before you call the
    Bluetooth Low Energy middleware API.

2. Enable Bluetooth Low Energy.

3. Initialize the middleware with `IotBLE_Init()`.

###### Note

This initialization step is not required if you are running the AWS demos. Demo
initialization is handled by the Network Manager, located at
`freertos/demos/network_manager`.

### Low-level APIs

If you don't want to use the FreeRTOS Bluetooth Low Energy GATT services, you
can bypass the middleware and interact directly with the low-level APIs to save resources.

**To initialize the low-level APIs**

1. Initialize any Bluetooth Low Energy hardware drivers before you call the APIs. Driver
    initialization is not part of the Bluetooth Low Energy low-level APIs.

2. The Bluetooth Low Energy low-level API provides an enable/disable call to the Bluetooth Low Energy stack for optimizing
    power and resources. Before calling the APIs, you must enable
    Bluetooth Low Energy.

```C

const BTInterface_t * pxIface = BTGetBluetoothInterface();
xStatus = pxIface->pxEnable( 0 );
```

3. The Bluetooth manager contains APIs that are common to both Bluetooth Low Energy and Bluetooth classic.
    The callbacks for the common manager must be initialized
    second.

```C

xStatus = xBTInterface.pxBTInterface->pxBtManagerInit( &xBTManagerCb );
```

4. The Bluetooth Low Energy adapter fits on top of the common API. You must initialize its callbacks like
    you initialized the common API.

```C

xBTInterface.pxBTLeAdapterInterface = ( BTBleAdapter_t * ) xBTInterface.pxBTInterface->pxGetLeAdapter();
xStatus = xBTInterface.pxBTLeAdapterInterface->pxBleAdapterInit( &xBTBleAdapterCb );
```

5. Register your new user application.

```C

xBTInterface.pxBTLeAdapterInterface->pxRegisterBleApp( pxAppUuid );
```

6. Initialize the callbacks to the GATT servers.

```C

xBTInterface.pxGattServerInterface = ( BTGattServerInterface_t * ) xBTInterface.pxBTLeAdapterInterface->ppvGetGattServerInterface();
xBTInterface.pxGattServerInterface->pxGattServerInit( &xBTGattServerCb );
```

After you initialize the Bluetooth Low Energy adapter, you can add a GATT server. You can register only
    one GATT server at a time.

```C

xStatus = xBTInterface.pxGattServerInterface->pxRegisterServer( pxAppUuid );
```

7. Set application properties like secure connection only and MTU size.

```C

xStatus = xBTInterface.pxBTInterface->pxSetDeviceProperty( &pxProperty[ usIndex ] );
```

## API reference

For a full API reference, see
[Bluetooth Low Energy API Reference](https://docs.aws.amazon.com/freertos/latest/lib-ref/html2/ble/index.html).

## Example usage

The examples below demonstrate how to use the Bluetooth Low Energy library for advertising and creating new services.
For full FreeRTOS Bluetooth Low Energy demo applications, see [Bluetooth Low Energy Demo Applications](https://docs.aws.amazon.com/freertos/latest/userguide/ble-demo.html).

### Advertising

1. In your application, set the advertising UUID:

```C

static const BTUuid_t _advUUID =
{
       .uu.uu128 = IOT_BLE_ADVERTISING_UUID,
       .ucType   = eBTuuidType128
};
```

2. Then define the `IotBle_SetCustomAdvCb` callback function:

```c

void IotBle_SetCustomAdvCb( IotBleAdvertisementParams_t * pAdvParams,  IotBleAdvertisementParams_t * pScanParams)
{
       memset(pAdvParams, 0, sizeof(IotBleAdvertisementParams_t));
       memset(pScanParams, 0, sizeof(IotBleAdvertisementParams_t));

       /* Set advertisement message */
       pAdvParams->pUUID1 = &_advUUID;
       pAdvParams->nameType = BTGattAdvNameNone;

       /* This is the scan response, set it back to true. */
       pScanParams->setScanRsp = true;
       pScanParams->nameType = BTGattAdvNameComplete;
}
```

This callback sends the UUID in the advertisement message and the full name in the scan response.

3. Open `vendors/vendor/boards/board/aws_demos/config_files/iot_ble_config.h`, and
    set `IOT_BLE_SET_CUSTOM_ADVERTISEMENT_MSG` to `1`. This triggers the `IotBle_SetCustomAdvCb` callback.

### Adding a new service

For full examples of services, see `freertos/.../ble/services`.

1. Create UUIDs for the service's characteristic and descriptors:

```C

#define xServiceUUID_TYPE \
{\
       .uu.uu128 = gattDemoSVC_UUID, \
       .ucType   = eBTuuidType128 \
}
#define xCharCounterUUID_TYPE \
{\
       .uu.uu128 = gattDemoCHAR_COUNTER_UUID,\
       .ucType   = eBTuuidType128\
}
#define xCharControlUUID_TYPE \
{\
       .uu.uu128 = gattDemoCHAR_CONTROL_UUID,\
       .ucType   = eBTuuidType128\
}
#define xClientCharCfgUUID_TYPE \
{\
       .uu.uu16 = gattDemoCLIENT_CHAR_CFG_UUID,\
       .ucType  = eBTuuidType16\
}
```

2. Create a buffer to register the handles of the characteristic and descriptors:

```C

static uint16_t usHandlesBuffer[egattDemoNbAttributes];
```

3. Create the attribute table. To save some RAM, define the table as a `const`.

###### Important

Always create the attributes in order, with the service as the first attribute.

```C

static const BTAttribute_t pxAttributeTable[] = {
        {
            .xServiceUUID =  xServiceUUID_TYPE
        },
       {
            .xAttributeType = eBTDbCharacteristic,
            .xCharacteristic =
            {
                 .xUuid = xCharCounterUUID_TYPE,
                 .xPermissions = ( IOT_BLE_CHAR_READ_PERM ),
                 .xProperties = ( eBTPropRead | eBTPropNotify )
             }
        },
        {
            .xAttributeType = eBTDbDescriptor,
            .xCharacteristicDescr =
            {
                .xUuid = xClientCharCfgUUID_TYPE,
                .xPermissions = ( IOT_BLE_CHAR_READ_PERM | IOT_BLE_CHAR_WRITE_PERM )
             }
        },
       {
            .xAttributeType = eBTDbCharacteristic,
            .xCharacteristic =
            {
                 .xUuid = xCharControlUUID_TYPE,
                 .xPermissions = ( IOT_BLE_CHAR_READ_PERM | IOT_BLE_CHAR_WRITE_PERM  ),
                 .xProperties = ( eBTPropRead | eBTPropWrite )
             }
        }
};
```

4. Create an array of callbacks. This array of callbacks must follow the same order as the table array defined above.

For example, if `vReadCounter` gets triggered when `xCharCounterUUID_TYPE` is accessed, and
    `vWriteCommand` gets triggered when `xCharControlUUID_TYPE` is accessed, define the array as follows:

```C

static const IotBleAttributeEventCallback_t pxCallBackArray[egattDemoNbAttributes] =
       {
     NULL,
     vReadCounter,
     vEnableNotification,
     vWriteCommand
};
```

5. Create the service:

```C

static const BTService_t xGattDemoService =
{
     .xNumberOfAttributes = egattDemoNbAttributes,
     .ucInstId = 0,
     .xType = eBTServiceTypePrimary,
     .pusHandlesBuffer = usHandlesBuffer,
     .pxBLEAttributes = (BTAttribute_t *)pxAttributeTable
};
```

6. Call the API `IotBle_CreateService` with the structure that you created in the previous step.
    The middleware synchronizes the creation of all services, so any new services need to already be defined when the `IotBle_AddCustomServicesCb` callback is triggered.
1. Set `IOT_BLE_ADD_CUSTOM_SERVICES` to `1` in `vendors/vendor/boards/board/aws_demos/config_files/iot_ble_config.h`.

2. Create IotBle\_AddCustomServicesCb in your application:

      ```c

      void IotBle_AddCustomServicesCb(void)
      {
          BTStatus_t xStatus;
          /* Select the handle buffer. */
          xStatus = IotBle_CreateService( (BTService_t *)&xGattDemoService, (IotBleAttributeEventCallback_t *)pxCallBackArray );
      }
      ```

## Porting

### User input and output peripheral

A secure connection requires both input and output for numeric comparison. The `eBLENumericComparisonCallback` event
can be registered using the event manager:

```C

xEventCb.pxNumericComparisonCb = &prvNumericComparisonCb;
xStatus = BLE_RegisterEventCb( eBLENumericComparisonCallback, xEventCb );

```

The peripheral must display the numeric passkey and take the result of the comparison as
an input.

### Porting API implementations

To port FreeRTOS to a new target, you must implement some APIs for the Wi-Fi Provisioning
service and Bluetooth Low Energy functionality.

#### Bluetooth Low Energy APIs

To use the FreeRTOS Bluetooth Low Energy middleware, you must implement some APIs.

##### APIs common between GAP for Bluetooth Classic and GAP for Bluetooth Low Energy

- `pxBtManagerInit`

- `pxEnable`

- `pxDisable`

- `pxGetDeviceProperty`

- `pxSetDeviceProperty` (All options are mandatory expect `eBTpropertyRemoteRssi` and `eBTpropertyRemoteVersionInfo`)

- `pxPair`

- `pxRemoveBond`

- `pxGetConnectionState`

- `pxPinReply`

- `pxSspReply`

- `pxGetTxpower`

- `pxGetLeAdapter`

- `pxDeviceStateChangedCb`

- `pxAdapterPropertiesCb`

- `pxSspRequestCb`

- `pxPairingStateChangedCb`

- `pxTxPowerCb`

##### APIs specific to GAP for Bluetooth Low Energy

- `pxRegisterBleApp`

- `pxUnregisterBleApp`

- `pxBleAdapterInit`

- `pxStartAdv`

- `pxStopAdv`

- `pxSetAdvData`

- `pxConnParameterUpdateRequest`

- `pxRegisterBleAdapterCb`

- `pxAdvStartCb`

- `pxSetAdvDataCb`

- `pxConnParameterUpdateRequestCb`

- `pxCongestionCb`

##### GATT server

- `pxRegisterServer`

- `pxUnregisterServer`

- `pxGattServerInit`

- `pxAddService`

- `pxAddIncludedService`

- `pxAddCharacteristic`

- `pxSetVal`

- `pxAddDescriptor`

- `pxStartService`

- `pxStopService`

- `pxDeleteService`

- `pxSendIndication`

- `pxSendResponse`

- `pxMtuChangedCb`

- `pxCongestionCb`

- `pxIndicationSentCb`

- `pxRequestExecWriteCb`

- `pxRequestWriteCb`

- `pxRequestReadCb`

- `pxServiceDeletedCb`

- `pxServiceStoppedCb`

- `pxServiceStartedCb`

- `pxDescriptorAddedCb`

- `pxSetValCallbackCb`

- `pxCharacteristicAddedCb`

- `pxIncludedServiceAddedCb`

- `pxServiceAddedCb`

- `pxConnectionCb`

- `pxUnregisterServerCb`

- `pxRegisterServerCb`

For more information about porting the FreeRTOS Bluetooth Low Energy library to your platform, see [Porting the Bluetooth Low Energy Library](https://docs.aws.amazon.com/freertos/latest/portingguide/afr-porting-ble.html) in the FreeRTOS Porting Guide.

## Appendix A: MQTT over BLE GATT profile

### GATT Service Details

MQTT over BLE uses an instance of the data transfer GATT service to send MQTT
Concise Binary Object Representation (CBOR) messages between the FreeRTOS device
and the proxy device. The data transfer service exposes certain characteristics
that help send and receive raw data over the BLE GATT protocol. It also handles
the fragmentation and assembly of payloads greater than the BLE maximum transfer
unit (MTU) size.

**Service UUID**

`A9D7-166A-D72E-40A9-A002-4804-4CC3-FF00`

**Service Instances**

One instance of the GATT service is created for each MQTT session
with the broker. Each service has a unique UUID (two bytes) that
identifies its type. Each individual instance is differentiated by
the instance ID.

Each service is instantiated as a primary service on each BLE server
device. You can create multiple instances of the service on a given
device. The MQTT proxy service type has a unique UUID.

**Characteristics**

Characteristic content format: **CBOR**

Max characteristic value size : 512 bytes

CharacteristicRequirementMandatory PropertiesOptional PropertiesSecurity PermissionsBrief DescriptionUUIDControlMWriteNoneWrite Needs Encryption Used to start and stop the MQTT proxy.`A9D7-166A-D72E-40A9-A002-4804-4CC3-FF01`TXMessage M Read, Notification None Read Needs Encryption Used to send a notification containing a message to a broker
via a proxy.`A9D7-166A-D72E-40A9-A002-4804-4CC3-FF02`RXMessage M Read, Write Without Response None Read, Write Needs Encryption Used to receive a message from a broker via a proxy.`A9D7-166A-D72E-40A9-A002-4804-4CC3-FF03`TXLargeMessage M Read, Notification None Read Needs Encryption Used to send a large message (Message
\> BLE MTU Size) to a broker via a proxy.`A9D7-166A-D72E-40A9-A002-4804-4CC3-FF04`RXLargeMessage M Read,
Write Without Response None Read, Write Needs Encryption Used to receive large message (Message
\> BLE MTU Size) from a broker via a proxy.`A9D7-166A-D72E-40A9-A002-4804-4CC3-FF05`

**GATT Procedure Requirements**

Read Characteristic Values  Mandatory  Read Long Characteristic Values  Mandatory  Write Characteristic Values  Mandatory  Write Long Characteristic Values  Mandatory  Read Characteristic descriptors  Mandatory  Write Characteristic descriptors  Mandatory  Notifications  Mandatory  Indications  Mandatory

**Message Types**

The following message types are exchanged.

Message Type  Message  Map with these key / value pairs 0x01  CONNECT

- Key = "w", value = Type 0 Integer, Message type (1)

- Key = "d", value = Type 3, Text String, Client Identifier for the session

- Key = "a", value = Type 3, Text String, Broker endpoint for the session

- Key = "c", Value = Simple Value Type True/False

0x02  CONNACK

- Key = "w, value = Type 0 Integer, Message type (2)

- Key = "s", Value = Type 0 Integer, Status code

0x03  PUBLISH

- Key = "w", value = Type 0 Integer, Message type (3)

- Key = "u", value = Type 3, Text String, Topic for publish

- Key = "n", value = Type 0, Integer, QoS for publish

- Key = "i", value = Type 0, Integer, Message Identifier, Only for QoS 1 Publishes

- Key = "k", Value = Type 2, Byte String, Payload for publish

0x04  PUBACK

- Sent Only for QoS 1 messages.

- Key = "w", value = Type 0 Integer, Message type (4)

- Key = "i", value = Type 0, Integer, Message Identifier

0x08  SUBSCRIBE

- Key = "w", value = Type 0 Integer, Message type (8)

- Key = "v", value = Type 4, Array of text strings, topics for subscription

- Key = "o", value = Type 4, Array of Integers, QoS for subscription

- Key = "i", value = Type 0, Integer, Message Identifier

0x09  SUBACK

- Key = "w", value = Type 0 Integer, Message type (9)

- Key = "i", value = Type 0, Integer, Message Identifier

- Key = "s", value = Type 0, Integer, Status code for Subscription

0X0A  UNSUBSCRIBE

- Key = "w", value = Type 0 Integer, Message type (10)

- Key = "v", value = Type 4, Array of text strings, topics for unsubscription

- Key = "i", value = Type 0, Integer, Message Identifier

0x0B  UNSUBACK

- Key = "w", value = Type 0 Integer, Message type (11)

- Key = "i", value = Type 0, Integer, Message Identifier

- Key = "s", value = Type 0, Integer, Status code for UnSubscription

0X0C  PINGREQ

- Key = "w", value = Type 0 Integer, Message type (12)

0x0D  PINGRESP

- Key = "w", value = Type 0 Integer, Message type (13)

0x0E  DISCONNNECT

- Key = "w", value = Type 0 Integer, Message type (14)

**Large Payload Transfer Characteristics**

**TXLargeMessage**

TXLargeMessage is used by the device to send a large payload that
is greater than the MTU size negotiated for the BLE connection.

- The device sends the first MTU bytes of the payload
as a notification through the characteristic.

- The proxy sends a read request on this characteristic
for the remaining bytes.

- The device sends up to the MTU size or the remaining bytes
of the payload, whichever is less. Each time, it increases the
offset read by the size of the payload sent.

- The proxy will continue to read the characteristic until it
gets a zero length payload or a payload less than the MTU size.

- If the device doesn't get a read request within a specified timeout,
the transfer fails and the proxy and gateway release the buffer.

- If the proxy
doesn't
get a read response within a specified timeout, the
transfer fails and the proxy releases the buffer.

**RXLargeMessage**

RXLargeMessage is used by the device to receive a large payload that is
greater than the MTU size negotiated for the BLE connection.

- The proxy writes messages, up to the MTU size, one by one, using
write with response on this characteristic.

- The device buffers the message until it receives a write
request with zero length or a length less than the MTU size.

- If the device
doesn't
get a write request within a specified timeout, the
transfer fails and the device releases the buffer.

- If the proxy
doesn't
get a write response within a specified timeout, the
transfer fails and the proxy releases the buffer.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

backoffAlgorithm

Mobile SDKs for FreeRTOS Bluetooth devices
