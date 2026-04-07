# AWS IoT Over the air (OTA) library

###### Note

The content on this page may not be up-to-date. Please refer to the
[FreeRTOS.org library page](https://www.freertos.org/Documentation/03-Libraries/01-Library-overview/01-All-libraries)
for the latest update.

## Introduction

The [AWS IoT Over-the-air (OTA) update library](https://docs.aws.amazon.com/embedded-csdk/latest/lib-ref/libraries/aws/ota-for-aws-iot-embedded-sdk/docs/doxygen/output/html/index.html) enables you to manage the notification, download, and verification
of firmware updates for FreeRTOS devices using HTTP or MQTT as the protocol. By using the OTA Agent library, you can
logically separate firmware updates and the application running on your devices. The OTA Agent can share a network
connection with the application. By sharing a network connection, you can potentially save a significant amount of
RAM. In addition, the OTA Agent library lets you define application-specific logic for testing, committing, or
rolling back a firmware update.

The Internet of Things (IoT) extends internet connectivity to embedded devices that were traditionally not
connected. These devices can be programmed to communicate usable data over the internet, and can be remotely
monitored and controlled. With advances in technology, these traditional embedded devices are getting internet
capabilities in consumer, industrial, and enterprise spaces at a fast pace.

IoT devices are typically deployed in large quantities and often in places that are difficult or impractical
for a human operator to access. Imagine a scenario where a security vulnerability that can expose data is
discovered. In such scenarios, it is important to update the affected devices with security fixes quickly
and reliably. Without the ability to perform OTA updates, it can also be difficult to update devices that are
geographically dispersed. Having a technician update these devices will be costly, time consuming, and often
times impractical. The time required to update these devices leaves them exposed to security vulnerabilities
for a longer period. Recalling these devices for updating will also be costly and may cause significant disruption
to consumers due to downtime.

Over the Air (OTA) Updates make it possible to update device firmware without an expensive recall or technician
visit. This method adds the following benefits:

- **Security** \- The ability to quickly respond to security vulnerabilities
and software bugs that are discovered after the devices are deployed in the field.

- **Innovation** \- Products can be updated frequently as new features are
developed, driving the innovation cycle. The updates can take effect quickly with minimum downtime compared
to traditional update methods.

- **Cost** \- OTA updates can reduce maintenance costs significantly compared
to methods traditionally used to update these devices.

Providing the OTA functionality requires the following design considerations:

- **Secure Communication** \- Updates must use encrypted communication
channels to prevent the downloads from being tampered with during transit.

- **Recovery** \- Updates can fail due to things like intermittent network
connectivity or receiving an invalid update. In these scenarios, the device needs to be able to return
to a stable state and avoid becoming bricked.

- **Author Verification** \- Updates must be verified to be from a trusted
source, along with other validations like checking the version and compatibility.

For more information about setting up OTA updates with FreeRTOS, see [FreeRTOS Over-the-Air Updates](freertos-ota-dev.md).

### AWS IoT Over the air (OTA) library

The AWS IoT OTA library enables you to manage notifications of a newly available updates, download
them, and perform cryptographic verification of firmware updates. Using the over-the-air (OTA) client
library, you can logically separate the firmware update mechanisms from the application running on your
device. The over-the-air (OTA) client library can share a network connection with the application, saving
memory in resource-constrained devices. In addition, the over-the-air (OTA) client library lets you define
application-specific logic for testing, committing, or rolling back a firmware update. The library supports
different application protocols like Message Queuing Telemetry Transport (MQTT) and Hypertext Transfer
Protocol (HTTP) and provides various configuration options you can fine tune for your network type and
conditions.

This library's APIs provide these major functions:

- Register for notifications or poll for new update requests that are available.

- Receive, parse and validate the update request.

- Download and verify the file according to the information in the update request.

- Run a self-test before activating the received update to ensure the functional validity of
the update.

- Update the status of the device.

This library uses AWS services to manage various cloud related functions such as sending firmware updates,
monitoring large numbers of devices across multiple regions, reducing the blast radius of faulty deployments,
and verifying the security of updates. This library can be used with any MQTT or HTTP library.

The demos for this library demonstrate complete over-the-air updates using the coreMQTT Library and AWS
Services on a FreeRTOS device.

## Features

Here is the complete OTA Agent interface:

**`**
**OTA_Init`**

Initializes the OTA engine by starting OTA Agent ("OTA Task") in the system. Only one OTA Agent may
exist.

**`**
**OTA_Shutdown`**

Signal to the OTA Agent to shut down. The OTA agent will optionally unsubscribe from all MQTT job
notification topics, stop in progress OTA jobs, if any, and clear all resources.

**`**
**OTA_GetState`**

Gets the current state of the OTA Agent.

**`**
**OTA_ActivateNewImage`**

Activates the newest microcontroller firmware image received through OTA. (The detailed job status
should now be self-test.)

**`**
**OTA_SetImageState`**

Sets the validation state of the currently running microcontroller firmware image (testing, accepted
or rejected).

**`**
**OTA_GetImageState`**

Gets the state of the currently running microcontroller firmware image (testing, accepted or
rejected).

**`**
**OTA_CheckForUpdate`**

Requests the next available OTA update from the OTA Update service.

**`**
**OTA_Suspend`**

Suspend all OTA Agent operations.

**`**
**OTA_Resume`**

Resume OTA Agent operations.

**`**
**OTA_SignalEvent`**

Signal an event to the OTA Agent task.

**`**
**OTA_EventProcessingTask`**

OTA agent event processing loop.

**`**
**OTA_GetStatistics`**

Get the statistics of OTA message packets which includes the number of packets received, queued,
processed and dropped.

**`**
**OTA_Err_strerror`**

Error code to string conversion for OTA errors.

**`**
**OTA_JobParse_strerror**
**`**

Convert an OTA Job Parsing error code to a string.

**`**
**OTA_PalStatus_strerror`**

Status code to string conversion for OTA PAL status.

**`**
**OTA_OsStatus_strerror`**

Status code to string conversion for OTA OS status.

## API reference

For more information, see the [AWS IoT Over-the-air Update: Functions](https://aws.github.io/ota-for-aws-iot-embedded-sdk/v3.4.0/ota_functions.html).

## Example usage

A typical OTA-capable device application using the MQTT protocol drives the OTA Agent by using the
following sequence of API calls.

1. Connect to the AWS IoT coreMQTT Agent. For more information, see [coreMQTT Agent library](https://docs.aws.amazon.com/freertos/latest/userguide/coremqtt-agent.html).

2. Initialize the OTA Agent by calling `OTA_Init`, including the buffers, the required ota
    interfaces, the thing name and the application callback. The callback implements application-specific logic that executes after completing an OTA update job.

3. When the OTA update is complete, FreeRTOS calls the job completion callback with one of the following events:
    `accepted`, `rejected`, or `self test`.

4. If the new firmware image has been rejected (for example, due to a validation error), the application can
    typically ignore the notification and wait for the next update.

5. If the update is valid and has been marked as accepted, call `OTA_ActivateNewImage` to reset
    the device and boot the new firmware image.

## Porting

For information about porting OTA functionality to your platform, see [Porting the OTA Library](https://docs.aws.amazon.com/freertos/latest/portingguide/afr-porting-ota.html) in the FreeRTOS Porting Guide.

## Memory use

Code Size of AWS IoT OTA (example generated with GCC for ARM Cortex-M)FileWith -O1 OptimizationWith -Os Optimizationota.c8.3K7.5Kota\_interface.c0.1K0.1Kota\_base64.c0.6K0.6Kota\_mqtt.c2.4K2.2Kota\_cbor.c0.8K0.6Kota\_http.c0.3K0.3K**Total estimates****12.5K****11.3K**

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

coreMQTT Agent

corePKCS11
