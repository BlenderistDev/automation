# Golang automation package
Package provides automation service with custom triggers, conditions and actions.


## Install
````
go get github.com/BlenderistDev/automation
````
## Triggers
Trigger is a name of event to execute automation. Automation can have multiple triggers.
## Condition
Automation can have only one condition.

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

## Actions
Automation can have any count of actions.
To implement action interface you should implement execute method
````
execute(Trigger trigger) error 
````
