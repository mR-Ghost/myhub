/*
Copyright 2018 Sgoby.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package opt

import (
	"testing"
	"fmt"
)

func Test_OptSelect(t *testing.T){
	sql := "SELECT name FROM dealer_info_201609  GROUP BY id,ss";
	nSql,err := OptimizeSelectSql(sql)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println(sql)
	fmt.Println(nSql)
}
