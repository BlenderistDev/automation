# Golang automation package
Package provides automation service with custom triggers, conditions and actions.


## Install
````
go get github.com/BlenderistDev/automation
````
## Automation
Automation struct includes three entities: triggers, condition and actions.

If there is trigger event name in automation triggers list and if condition returns true (or if there's no condition) automation actions execute.
## Trigger
Trigger is a name of event to execute automation. Automation can have multiple triggers.
## Condition
Automation can have at most one condition.
Automation actions will execute if it has no condition or condition check method will return true.

To implement condition interface you should implement check method
````
Check(trigger Trigger) (bool, error)
````

You can make complex conditions with built-in conditions:
* Or
* And
* Not

There are some other build-in useful conditions:
* Equal

## Action
Automation can have at least one action.

To implement action interface you should implement execute method
````
execute(Trigger trigger) error 
````

## Trigger event
Trigger event is a struct for event data. To implement trigger event you should implement two methods:

GetName method returns name of trigger. It is used to find automations for execute.
````
GetName() string
````
GetData method returns event data. Map keys is used as variable name and map value is used as variable value.
````
GetData() map[string]string
````

## Testing
Package is fully test covered.

To use interface mocks you need to install gomock dependency.

You can add mocks by import:

````
import("github.com/BlenderistDev/automation/testing/interfaces")
````
