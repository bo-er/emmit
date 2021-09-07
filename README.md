## To Those Who Wish To Download All OCW's Resources

I am a big fan of The OCW, I wish to download all OCW's resources, this repository serves that purpose.


## How to use

find a school you feel like exploring,  I prefer EE school so I would do the following steps:
1. crawling all the links from EE school
```
go run . crawl -p "https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/" -f ee.info -s "electrical-engineering-and-computer-science"
```

2. select one course I would like to learn and merge its course materials(PDF version lecture notes)

```
go run . pdf ee.info 6_172

```

### Note

Since each course may have different lecture notes links. Thus you can use regex expression to filter specific link suffixes.
For example, use `go run . pdf ee.info 6_006 2011 "lec[0-9]+.pdf"` you can Download MIT course 6.006's lecture materials since they have `lec01.pdf` style suffixs.


## License Info
---
This software contains a package ([unipdf](https://github.com/unidoc/unipdf)) which is a commercial product and requires a license code to operate, 
Therefore the use of this software package is governed by the end-user license agreement (EULA) available at: https://unidoc.io/eula/
To obtain a Trial license code to evaluate the software, please visit https://unidoc.io/