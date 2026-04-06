# AWS IoT ExpressLink programmer's guide v1.3

This document defines the Application Programming Interface (API) that all AWS IoT ExpressLink compliant
connectivity modules are required to implement to connect any host processor to the AWS cloud.

If you have questions or issues that are not answered here, please visit the
[AWS re:Post \
for AWS IoT ExpressLink](https://repost.aws/tags/TADqOo0ODORl2pC69DWwUFug/aws-io-t-express-link) page.

See the [Document history](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-history.html) for changes in this version.

###### Topics

- [1 Overview](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-overview.html)

- [2 Hardware](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-hardware.html)

- [3 Run states](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-run-states.html)

- [4 ExpressLink commands](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html)

- [5 Messaging](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-messaging.html)

- [6 Configuration Dictionary](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-configuration-dictionary.html)

- [7 Status dictionary](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-status-dictionary.html)

- [8 Event handling](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-event-handling.html)

- [9 ExpressLink module Updates](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html)

- [10 AWS IoT Services](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-iot-services.html)

- [11 Additional services](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-additional-services.html)

- [12 Provisioning and Onboarding](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-provisioning.html)

- [13 Bluetooth Low Energy (BLE)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ble.html)

- [14 GPIO control (new with v1.3)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-gpio-control.html)

- [Appendix – Manufacturer Module Datasheet Requirements](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-appendix-b.html)

- [Document history](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-history.html)

- [Archive](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-archive.html)

**AWS IoT ExpressLink commands**

See these sections for descriptions of AWS IoT ExpressLink commands in the following general
categories:

- [Power and connection control (CONNECT, ...)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-power-control.title)

- [Messaging topic model (SEND, GET, ... )](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-messaging.html#elpg-messaging-topic.title)

- [Dictionary data access (CONF, ...)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-configuration-dictionary.html#elpg-data-accessed-conf.title)

- [Event handling commands (EVENT, ...)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-event-handling.html#elpg-event-handling-commands.title)

- [Diagnostic commands (DIAG, ...)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-event-handling.html#elpg-diagnostic-commands.title)

- [OTA commands (OTA)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-ota-commands.title)

- [Host OTA certificate update](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-hota-cert-update.title)

- [Over the Wire (OTW) module firmware update command](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-otw-firmware-update.title)

- [AWS IoT Device Defender (using CONF)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-iot-services.html#elpg-device-defender.title)

- [AWS IoT Device Shadow (SHADOW)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-iot-services.html#elpg-device-shadow.title)

- [Additional services (TIME?, WHERE?)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-additional-services.html#elpg-additional-services.title)

- [Bluetooth Low Energy (BLE)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ble.html#elpg-ble.title)

**Tables**

- [Table 1 - Error codes](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-table1)

- [Table 2 - Configuration Dictionary Persistent Keys](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-configuration-dictionary.html#elpg-table2)

- [Table 3 - Configuration dictionary non-persistent keys](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-configuration-dictionary.html#elpg-table3)

- [Table 4 - ExpressLink event codes](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-event-handling.html#elpg-table4)

- [Table 5 - OTA codes and descriptions](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-table5)

- [Table 6 - Reserved OTA file type codes (0-255)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-table6)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

1 Overview
