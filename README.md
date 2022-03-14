# Golang automation package
Package provides automation service with custom triggers, conditions and actions.


## Install
````
go get github.com/BlenderistDev/automation
````
## Automation
Automation struct includes three entities: triggers, condition and actions.
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
Automation can have any count of actions.
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
