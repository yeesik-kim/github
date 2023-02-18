# Pull Spinnaker Branch
* Pull the spinnaker branch into my repository.

## 사용법
``` go
package main

import (
	"fmt"

	"github.com/yeesik-kim/github/spinnaker"
)

func main() {
	spinRepo := &spinnaker.RepositInfo{[spinnaker module name], [my github username], [my github token]}
	myRepo := &spinnaker.MyInfo{[my github url], [spinnaker module url], [my github username], [my github token]}
	branches := spinRepo.GetBranchNames([target release name])
	for _, b := range branches {
		fmt.Println(b)
		myRepo.PushRelease(b)
	}
}
```

* spinnaker module name : spinnaker의 micro service 중 하나를 타겟
* my github username, my github token : spinnaker branch를 가져오기 위한 계정 정보
* my github url : spinnaker branch를 당겨올 나의 repository url
* spinnaker module url : spinnaker target url
* my github username, my github token : 나의 respoistory 계정 정보
