# Congestion Tax Calculator
## Assumptions
### Single charge interval starting with a toll free pass
In this scenario a tool-free pass at 5:30 starts a 2h tool single charge block. In result the tax=8+13 instead of tax=max(8, 13).

Unit test covering this scenario:
```
{ // First pass free (1st interval start), second pass paid (within 1st interval), third paid (second interval).
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2022, 12, 27, 5, 30, 0, 0, time.Local), //0
			time.Date(2022, 12, 27, 6, 0, 0, 0, time.Local),  //8
			time.Date(2022, 12, 27, 8, 0, 0, 0, time.Local)}, //13
		tax: 21,
	},
```

### All dates for the same day
Assumption is that tax is calculated for a single day and input table will contain only dates for the same day (but different times).

## Fixes
- upgraded go version 1.15 -> 1.19
- moved vehicle related artifacts to vehicle package
- moved configuration related artifacts to config package
- created unit tests
- created http 
- added support for multiple configurations (cities)
- aligned naming motorbike -> motorcycle
- fixed logic around toll-free days. Added switch
- fixed logic around single fee calculation 
- added sorting for input dates

## Questions
- As described above, a toll-free pass will start the interval, which may resut in higher fee. Is this behavior intentional?
