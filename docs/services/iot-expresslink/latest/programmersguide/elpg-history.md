# Document history

The following table describes important changes to the AWS IoT ExpressLink Programmer's Guide
starting with v1.0. We also update the documentation to address any errors found or feedback
received.

Change

Description

Date

version 1.3

The following sections were added:

- [13 Bluetooth Low Energy (BLE)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ble.html)
Section 13 was updated with:

- BLE Central Device Configuration was updated.

- BLE Peripheral Device Configuration was updated.

- BLE Characteristics Configuration was updated.

- BLE Descriptors Configuration was updated.

- [13.3.6 BLE AUTH\[#\] \[0/1\]   »Authorize access to a characteristic«](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ble.html#elpg-ble-auth-command) was introduced.

- [13.4.11 AT+BLE ALLOW\_CLEAR »Remove all devices from the Allow list«](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ble.html#elpg-ble-allow-clear-command) was introduced.

- A new section [14 GPIO control (new with v1.3)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-gpio-control.html) was added.

August 27, 2025

version 1.2

The following sections and tables were updated:

- [Table 1 - Error codes](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-table1) -
New error codes were introduced:

- 27 : BLE ERROR

- 28 : CONFIGURATION ERROR

- [Table 2 - Configuration Dictionary Persistent Keys](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-configuration-dictionary.html#elpg-table2) -
Added new non-persistent configuration parameters:

- BLECentral#

- BLEPeripheral

- BLEGATT#

- [Table 4 - ExpressLink event codes](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-event-handling.html#elpg-table4) -
New BLE Events introduced.

- [13 Bluetooth Low Energy (BLE)](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ble.html) -
New BLE commands introduced.

October 27, 2023

version 1.1.2

The following sections and tables were updated:

- [Table 1 - Error codes](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-table1)

- Introduction of ERROR 26 INVALID CERTIFICATE.

- [Table 2 - Configuration Dictionary Persistent Keys](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-configuration-dictionary.html#elpg-table2)

- 'Version' now allows for a suffix (for example,
'X.Y.Z beta2').

- OTACertificate is now write-only.

- [9 ExpressLink module Updates](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html)
(formerly Chapter 9)

- [9.5 Module OTA signature verification](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-ota-signature-verification)

- Underlined sentence requiring OTA Certificate
pre-provisioning.

- Reworded last sentence for clarity.

- [9.6 Module OTA certificate updates](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-ota-certificate-updates)

- OTACertificate is now write-only.

- ERR26 is produced when a corrupted or otherwise
invalid certificate is presented.

- [9.10 Host OTA Signature Verification](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-hota-sig-verification)

- [9.11 Host OTA certificate update](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-hota-cert-update)

- [9.11.2 CONF? {certificate} pem   »Special certificate output formatting option«](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-ota-confq-command)

- Special 'pem' qualifier is now case insensitive
(for example, 'pem' \| 'PEM').

- [9.11.3 CONF {certificate}=pem   »Special certificate input formatting option«](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-ota-updates.html#elpg-ota-conf-command)

- Special 'pem' qualifier is now case insensitive
(for example, 'pem' \| 'PEM').

July 25, 2023

version 1.1.1

The following sections were added:

- [10 AWS IoT Services](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-iot-services.html)

The following sections and tables were updated:

- [4.3 Delimiters and escaping](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-delimiters)

- [4.4 Maximum values](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-max-values)

- [Table 1 - Error codes](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-table1)

- [4.7.2 CONNECT   »Establish a connection to an AWS IoT Core Endpoint«](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-connect-command)

- [4.7.7 FACTORY\_RESET   »Request a factory reset of the ExpressLink module«](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-commands.html#elpg-factory-reset-command)

- [5.1.7 SUBSCRIBE\[#\]   »Subscribe to Topic#«](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-messaging.html#elpg-subscribeh-command)

- [5.1.8 UNSUBSCRIBE\[#\]   »Unsubscribe from Topic#«](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-messaging.html#elpg-unsubscribeh-command)

- [Table 2 - Configuration Dictionary Persistent Keys](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-configuration-dictionary.html#elpg-table2)

- [Table 3 - Configuration dictionary non-persistent keys](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-configuration-dictionary.html#elpg-table3)

- [Table 4 - ExpressLink event codes](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-event-handling.html#elpg-table4)

- [12 Provisioning and Onboarding](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-provisioning.html)

The following sections were removed:

- 4.1.3 SEND _{topic} message_   »Publish msg on the specified topic«

- 6 Status dictionary

    note that subsequent sections were renumbered

- 9 Additional services

November 17, 2022

version 1.0

Initial release.

June 20, 2022

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Appendix – Manufacturer Module Datasheet Requirements

Archive
