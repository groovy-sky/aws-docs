# Launch an application in Amazon Connect Agent Workspace

Launch the application in the agent workspace for the given application ARN or name. It
supports optional launch options to fine tune the launch behavior.

**Signature**

```typescript

launchApp(arnOrName: string, options?: AppLaunchOptions): Promise<AppInfo>
```

**Usage**

```typescript

const applicationsConfig: AppConfig[] = await appControllerClient.getAppCatalog();
const appInfo: AppInfo = await appControllerClient.launchApp(applicationsConfig[0].arn);
```

**Input**

**Parameter****Type****Description**arnOrName _Required_stringThe ARN or name of the applicationoptions.parametersAppParametersKey value pair of parameters passed to the applicationoptions.launchKeystringA unique id to avoid duplicate application being launched with
multiple invocation of launchApp APIoptions.openInBackgroundbooleanIf set to true, the application won't be set to focus after its
launchedoptions.scopeContactScope \| IdleScopeIndicates if the application is launched with idle i.e when there
are no contacts or launched during an active contact

**Output - AppInfo**

**Parameter****Type****Description**instanceIdstringUnique ID of the application instanceconfigConfigThe configuration of the applicationstartTimeDateTime when the application is launchedstateAppStateCurrent state of the applicationappCreateOrdernumberSequentially incremented counter upon new application launchesaccessUrlstringAccess URL of the applicationparametersAppParameters \| undefinedKey value pair of parameters passed to the applicationlaunchKeystringA unique id to avoid duplicate application being launched with multiple invocation of launchApp APIscopeContactScope \| IdleScopeIndicates if the application is launched with idle i.e when there are no contacts or launched during an active contact

**Permissions required:**

```typescript

*
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

getApps()

Contact

All content copied from https://docs.aws.amazon.com/.
