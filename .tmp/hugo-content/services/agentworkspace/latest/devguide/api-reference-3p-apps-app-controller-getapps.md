# Get all active application information in Amazon Connect Agent Workspace

Returns the application information for all active application instances in the
Amazon Connect agent workspace.

**Signature**

```typescript

getApps(): Promise<AppInfo[]>
```

**Usage**

```typescript

const applicationInfo: AppInfo[] = await appControllerClient.getApps();
```

**Output - AppInfo**

**Parameter****Type****Description**instanceIdstringUnique ID of the application instanceconfigConfigThe configuration of the applicationstartTimeDateTime when the application is launchedstateAppStateCurrent state of the applicationappCreateOrdernumberSequentially incremented counter upon new application
launchesaccessUrlstringAccess URL of the applicationparametersAppParameters \| undefinedKey value pair of parameters passed to the applicationlaunchKeystringA unique id to avoid duplicate application being launched with
multiple invocation of launchApp APIscopeContactScope \| IdleScopeIndicates if the application is launched with idle i.e when there
are no contacts or launched during an active contact

**Permissions required:**

```typescript

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getAppConfig()

launchApp()

All content copied from https://docs.aws.amazon.com/.
