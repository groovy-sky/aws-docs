# Get application information in Amazon Connect Agent Workspace

Returns the application information for the given application instance ID in the
Amazon Connect agent workspace.

**Signature**

```typescript

getApp(instanceId: string): Promise<AppInfo>
```

**Usage**

```typescript

const applicationInfo: AppInfo = await appControllerClient.getApp(appInstanceId);
```

**Input**

**Parameter****Type****Description**appInstanceId _Required_stringThe instance ID of the application

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

focusApp()

getAppCatalog()

All content copied from https://docs.aws.amazon.com/.
