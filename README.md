## To Those Who Wish To Download All OCW's Resources

I am a big fan of The OCW, I wish to download all ocw's resources, this repository serves that purpose.


## How to use

find a school you feel like to exploer, for me it's EE school so I would do the following steps:
1. crawling all the links from EE school
```
go run . crawl -p "https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/" -f ee.info -s "electrical-engineering-and-computer-science"
```

2. select one course I would like to learn and merge it's course materials(PDF version lecture notes)

```
go run . pdf ee.info 6_172

```

## License Info
---
This softwae contains a package ([unipdf](https://github.com/unidoc/unipdf)) which is a commercial product and requires a license code to operate, 
Therefor the use of this software package is governed by the end-user license agreement (EULA) available at: https://unidoc.io/eula/
To obtain a Trial license code to evaluate the software, please visit https://unidoc.io/