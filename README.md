# sprout
Elastic Beanstalk CLI wrapper for application creation management and maintenance
![example](http://bleacher-report.d.pr/19L8V.png)

Example Json Skeleton:
```json
{
  "ebeans":[{
    "name":"prod-carburetor-app-s2",
    "tier": "webserver",
    "instance_type": "m3.2xLarge",
    "platform": "docker",
    "db": "",
    "db_engine": "",
    "db_instance": "",
    "db_size":"",
    "db_user": "",
    "db_pass": "",
    "profile": "EbApp",
    "key": "default",
    "size": "2",
    "tags": "cluster=carburetor,role=prod-carburetor"
  },
  {
    "name": "stag-carburetor-app-s2",
    "tier": "webserver",
    "instance_type": "m3.medium",
    "platform": "docker",
    "db": "",
    "db_engine": "",
    "db_instance": "",
    "db_size":"",
    "db_user": "",
    "db_pass": "",
    "profile": "EbApp",
    "key": "default",
    "size": "1",
    "tags": "cluster=carburetor,role=stag-carburetor"
  }]
}
```
