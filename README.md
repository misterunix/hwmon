# hwmon

Collect hwmon values

Custom for my use. I dont expect anyone would be interested besides me.  
If you are, please contribute.

```
	err := Init("/sys/class/hwmon")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	GetValues()

	for _, s := range Sensors {
		fmt.Printf("Path:%s Name:%s Head:%s Tail:%s Value:%s\n", s.Path, s.Name, s.Head, s.Tail, s.Value)
	}
```



## Output of slice.

```
Path:/sys/class/hwmon/hwmon1 Name:acpitz Head:temp1 Tail:crit Value:107000
Path:/sys/class/hwmon/hwmon1 Name:acpitz Head:temp1 Tail:input Value:25000
Path:/sys/class/hwmon/hwmon2 Name:bat0 Head:curr1 Tail:input Value:1
Path:/sys/class/hwmon/hwmon2 Name:bat0 Head:in0 Tail:input Value:12809
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:fan1 Tail:input Value:0
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:fan1 Tail:label Value:processor fan
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:temp1 Tail:input Value:40000
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:temp1 Tail:label Value:cpu
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:temp2 Tail:input Value:33000
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:temp2 Tail:label Value:ambient
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:temp3 Tail:input Value:27000
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:temp3 Tail:label Value:sodimm
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:temp4 Tail:input Value:
Path:/sys/class/hwmon/hwmon3 Name:dell_smm Head:temp4 Tail:label Value:gpu
Path:/sys/class/hwmon/hwmon4 Name:radeon Head:freq1 Tail:input Value:
Path:/sys/class/hwmon/hwmon4 Name:radeon Head:temp1 Tail:crit Value:120000
Path:/sys/class/hwmon/hwmon4 Name:radeon Head:temp1 Tail:crit_hyst Value:90000
Path:/sys/class/hwmon/hwmon4 Name:radeon Head:temp1 Tail:input Value:
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp1 Tail:crit Value:100000
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp1 Tail:crit_alarm Value:0
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp1 Tail:input Value:49000
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp1 Tail:label Value:package id 0
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp1 Tail:max Value:84000
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp2 Tail:crit Value:100000
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp2 Tail:crit_alarm Value:0
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp2 Tail:input Value:49000
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp2 Tail:label Value:core 0
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp2 Tail:max Value:84000
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp3 Tail:crit Value:100000
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp3 Tail:crit_alarm Value:0
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp3 Tail:input Value:45000
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp3 Tail:label Value:core 1
Path:/sys/class/hwmon/hwmon5 Name:coretemp Head:temp3 Tail:max Value:84000
```

