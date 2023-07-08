# Learning about WaitGroups

Step1: Create a WG

Step2: Call the method `add` in order to indicate how many operations or goroutines you want to wait for

Step3: Run your goroutines and inside each goroutine, you want to call the method `done`

Step4: Once you called all of your goroutines, you have to call the method `wait` in order for those goroutines to execute and call their `done` method