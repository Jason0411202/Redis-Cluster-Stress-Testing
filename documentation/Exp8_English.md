# The Impact of Enabling Redis Hook on Performance in Experiment 7
* This experiment will be conducted using the Standard_E4ds_v4 virtual machine provided by Azure.

## Read/Write Performance
* In this experiment, I will have the producer and consumer send and receive a total of 300,000 messages to test performance.
* Without Redis Hook: 52.033732367s, 51.992560416s, 52.214875285s, average 52.080389356s
* With Redis Hook: 52.610379859s, 52.387042074s, 53.172067901s, average 52.723163278s

## Summary
The Redis Hook may have a slight impact on performance; however, in this experiment, the impact is minimal.
