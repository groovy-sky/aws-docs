# WorkSpaces Applications Windows Client Release Notes

The WorkSpaces Applications client is a native application that is designed for users who require additional functionality during their WorkSpaces Applications streaming sessions. The following table describes the latest updates that are available in released versions
of the WorkSpaces Applications client.

For more information about the client, see [Provide Access Through the WorkSpaces Applications Client](client-application.md).

Client versionRelease dateChanges

1.2.1778

03-31-2026

- Added support for automatically minimizing the application catalog window when using a streaming URL with a pre-configured application ID.

- Updated the Chromium Embedded Framework (CEF) to version 145.0.26.

- Updated the DCV SDK.

- Other bug fixes and enhancements.

1.2.1733

02-12-2026

- Added streaming support for Ubuntu-based Elastic fleets.

- Added support for the client to consume the centrally-managed time zone redirection setting from the enable-timezone-redirection parameter on the DCV server.

- Added a configurable indicator to provide real-time visual feedback on streaming data loss.

- Fixed an issue with incorrect monitor resolutions being applied when the display scaling factor was set to higher than 100% in a multi-display configuration.

- Fixed a timing issue that caused streaming session launches to result in a black screen in rare conditions.

- Fixed an issue to show the application display name during the application launch process.

- Fixed an issue that showed an incorrect session expiration error during user pool streaming sessions.

- Updated the DCV SDK.

- Updated the Chromium Embedded Framework (CEF) to version 140.1.140.

- Other bug fixes and enhancements.

1.2.1610

11-14-2025

- Fixed an issue introduced in client version 1.2.1581 that disabled smart card redirection functionality.

1.2.1581

11-06-2025

- Launching the new Amazon WorkSpaces Applications client

- Adds support for streaming over IPv6

- Additional bug fixes and improvements

1.1.1532

09-18-2025

- Includes bug fixes and improvements

1.1.1490

08-11-2025

- Includes bug fixes and improvements

1.1.1458

06-10-2025

- Upgrades the embedded Chromium browser to version
135.0.170

- Includes bug fixes and improvements

1.1.144006-06-2025

- Reverts the embedded Chromium browser to version
131.3.50

- Includes bug fixes and improvements

1.1.1437

05-21-2025

- Upgrades the embedded Chromium browser to version
135.0.170

- Includes bug fixes and improvements

1.1.1423

03-31-2025

- Upgrades the embedded Chromium browser to version
131.3.50

- Includes bug fixes and improvements

1.1.1414

01-16-2025

- Adds support for automatic time zone
redirection

- Includes bug fixes and improvements

1.1.1408

12-19-2024

- Includes bug fixes and improvements

1.1.1403

12-12-2024

- Adds support to save user preferences between
streaming sessions

- Includes bug fixes and improvements

1.1.1395

11-18-2024

- Upgrades the embedded Chromium browser to version
129.0.110

- Includes bug fixes and improvements

1.1.1360

08-01-2024

- Adds support for extending full-screen across selected
monitors

- Adds support to stream Red Hat Enterprise Linux
images

- Upgrades the embedded Chromium browser to version
125.0.210

- Includes bug fixes and improvements

1.1.1332

07-03-2024

- Includes bug fixes and improvements

1.1.1326

06-17-2024

- Improves the user experience for the IdP-initiated SSO
workflow by automatically opening the client after user
sign-in with the system browser

- Other bug fixes and improvements

1.1.1303

04-03-2024

- Includes bug fixes and improvements

1.1.1300

03-28-2024

- Added support for launching the WorkSpaces Applications client from
IdP-initiated streaming sessions

- Added support for new relay state regional
endpoints

- Upgrades the embedded Chromium browser to version
121.3.70

- Includes bug fixes and improvements

1.1.1259

02-08-2024

- Includes bug fixes and improvements

1.1.1246

01-18-2024

- Includes improved accessibility features

- Includes bug fixes and improvements

- Upgrades the embedded Chromium browser to version
119.4.30

1.1.1228

11-01-2023

- Includes bug fixes and improvements

- Upgrades the embedded Chromium browser to version
114.1.120

1.1.118306-22-2023

- Includes bug fixes and improvements

- Upgrades the embedded Chromium browser to version
111.2.20

1.1.115905-09-2023

- Includes bug fixes and improvements

1.1.113002-09-2023

- Upgrades the embedded Chromium browser to version
108.4.130

1.1.111811-07-2022

- Upgrades the embedded Chromium browser to version
106.0.26

1.1.1099 10-13-2022

- Includes bug fixes and improvements

1.1.1066 08-17-2022

- Upgrades the embedded Chromium browser to version 102.0.9.
Microsoft Visual C++ 2019 Redistributable must be installed
as a prerequisite.

1.1.1025 06-29-2022

- Adds support for UDP streaming. For more information, see
[Amazon WorkSpaces Applications enables UDP streaming for Windows native\
client](https://aws.amazon.com/about-aws/whats-new/2022/06/amazon-appstream-2-0-udp-streaming-windows-native-client).

1.1.42105-19-2022

- Includes bug fixes

1.1.41404-26-2022

- Includes bug fixes and UI improvements

1.1.39802-23-2022

- Includes bug fixes

1.1.39402-08-2022

- Upgrades the embedded Chromium browser to version
97

1.1.38612-20-2021

- Upgrades the embedded Chromium browser to version
94.4

- Includes bug fixes

1.1.36011-15-2021

- Adds support for Linux application streaming

- Adds support for Elastic fleets. For more information, see
[Amazon AppStream 2.0 launches Elastic\
fleets](https://aws.amazon.com/about-aws/whats-new/2021/11/amazon-appstream-2-0-launches-elastic-fleets-serverless-fleet-type).

- Fixes a bug with the Japanese keyboard

1.1.33309-08-2021

- Bug fixes for the embedded Chromium browser

1.1.319 08-16-2021

- Resolves an issue with the caps lock, number lock, and
scroll lock keys

- Resolves an issue for the domain join sign-in experience

1.1.30408-02-2021

- Upgrades the embedded Chromium browser to version
91

- Updated USB driver to include important fixes

1.1.29404-26-2021

- Resolves an issue with SAML 2.0 authentication

- Resolves a client stability issue with Windows 7

- Resolves an issue with folder sharing on client
reconnection

1.1.28503-08-2021

- Includes fixes that improve compatibility with antivirus
software

1.1.25712-28-2020

- Adds support for real-time audio-video (AV)

- Adds support for using a smart card for Windows sign in to
Active Directory-joined streaming instances and in-session
authentication for streaming applications

- Resolves an issue that causes Microsoft Excel sheets to
lose focus during streaming sessions

1.1.19508-18-2020

- Improves the experience of sharing local drives and
folders that belong to cloud-based persistent storage
solutions such as OneDrive when file redirection is used
during streaming sessions

- Upgrades the embedded Chromium browser to version
81

- Resolves `AS2TrustedDomain` DNS TXT record
lookup failures for domains specified in the
`AS2TrustedDomains` list. These failures may
occur with some URI schemes. For more information, see [Create the AS2TrustedDomains DNS TXT Record to Enable Your Domain for the WorkSpaces Applications Client Without Registry Changes](install-client-configure-settings.md#create-AS2TrustedDomains-DNS-TXT-record-client)

- Resolves an intermittent issue that causes the client to
stop functioning when audio is enabled

1.1.17907-08-2020

- Adds support for local printer redirection

- Resolves an issue with concurrent HTTP connections that
prevents streaming with some proxy settings

- Resolves an issue that causes file downloads for files
greater than a few gigabytes to stop, and then fail

- Resolves an issue that causes subsequent connection
attempts to WorkSpaces Applications to fail if users sign in and connect to
WorkSpaces Applications over SAML, disconnect from the session without
closing the WorkSpaces Applications client, and then try to start a new
WorkSpaces Applications streaming session

1.1.16004-28-2020

- Resolves an issue that prevents the application catalog
page from opening on a Windows PC that has .NET Framework
version 4.7.1 or earlier installed

- Resolves an intermittent issue that causes the client to
stop responding when users close the client
application

1.1.15604-22-2020

- Adds support for defining trusted subdomains for user
connections in a DNS TXT record

- Adds support for on-demand diagnostic log and minidump
uploads

- Adds support for displaying custom branding for users who
stream in native application mode

###### Note

Users who have this version of the WorkSpaces Applications client installed
must have .NET Framework version 4.7.2 or later installed on the
same PC. For a list of the .NET Framework versions available for
download, see [Download .NET Framework](https://dotnet.microsoft.com/download/dotnet-framework).

1.1.13703-08-2020

- Reverts the updates in version 1.1.136

1.1.13603-05-2020

- Adds support for defining trusted subdomains for user
connections in a DNS TXT record

1.1.12902-28-2020

- Adds support for native application mode

- Improves the user interface for the DCV Printer
experience

- Resolves an issue with using Surface Pro Pen with
streaming applications

- Resolves an issue with downloading files with file names
that have international characters

1.0.52512-12-2019

- Resolves a DPI issue that causes the mouse cursor to point
to the wrong location when a user clicks on an application
during a streaming session

1.0.51110-16-2019

- Adds support for up to 4 monitors with a maximum display resolution of 4096x2160 pixels per monitor

- Adds support for up to 2 monitors with a maximum display resolution of 4096x2160 pixels per monitor on the Graphics Design and Graphics Pro instance types

- Adds support for seamless user connections to streaming
sessions that were started using custom uniform resource
identifier (URI) redirects

- Adds support for adding trusted domains for start
URLs

1.0.49909-26-2019

- Resolves an issue with client-side hardware
rendering

- Resolves an issue with the client not working correctly
when Bluetooth headsets are connected to the local
computer

1.0.48008-20-2019

- Adds support for WorkSpaces Applications file system redirection

1.0.46707-29-2019

- Includes fixes and enhancements to ensure compatibility
with updates made to WorkSpaces Applications portal endpoints

1.0.40705-16-2019

- Adds support for configuring the amount of time that users
can be idle (inactive) before they are disconnected from
their streaming session. For more information, see "Create a
Fleet" in [Create an Amazon WorkSpaces Applications Fleet and Stack](set-up-stacks-fleets.md).

- Resolves an issue with the “session alert” window
appearing when a SAML 2.0 session has expired

- Includes bug fixes for printing a document to a print
server

1.0.37503-07-2019

- Adds touch screen support on Windows PCs

- Adds support for automatically connecting USB devices when
a new streaming session starts

- Adds support for running session scripts

- Adds support for delivering virtualized applications using
the WorkSpaces Applications dynamic application framework APIs

1.0.32001-19-2019

- Adds multi monitor support for Graphics Design
instances

- Adds support for client display scaling factors greater
than 100 percent

- Adds support for WorkSpaces Applications regional settings

- Adds support for the WorkSpaces Applications user pool

- Adds support for honoring client-side proxy
settings

1.0.24711-20-2018 Initial release

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting

Client for macOS

All content copied from https://docs.aws.amazon.com/.
