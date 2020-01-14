// Copyright (c) 2012-2020 Grabtaxi Holdings PTE LTD (GRAB), All Rights Reserved. NOTICE: All information contained herein
// is, and remains the property of GRAB. The intellectual and technical concepts contained herein are confidential, proprietary
// and controlled by GRAB and may be covered by patents, patents in process, and are protected by trade secret or copyright law.
//
// You are strictly forbidden to copy, download, store (in any medium), transmit, disseminate, adapt or change this material
// in any way unless prior written permission is obtained from GRAB. Access to the source code contained herein is hereby
// forbidden to anyone except current GRAB employees or contractors with binding Confidentiality and Non-disclosure agreements
// explicitly covering such access.
//
// The copyright notice above does not evidence any actual or intended publication or disclosure of this source code,
// which includes information that is confidential and/or proprietary, and is a trade secret, of GRAB.
//
// ANY REPRODUCTION, MODIFICATION, DISTRIBUTION, PUBLIC PERFORMANCE, OR PUBLIC DISPLAY OF OR THROUGH USE OF THIS SOURCE
// CODE WITHOUT THE EXPRESS WRITTEN CONSENT OF GRAB IS STRICTLY PROHIBITED, AND IN VIOLATION OF APPLICABLE LAWS AND
// INTERNATIONAL TREATIES. THE RECEIPT OR POSSESSION OF THIS SOURCE CODE AND/OR RELATED INFORMATION DOES NOT CONVEY
// OR IMPLY ANY RIGHTS TO REPRODUCE, DISCLOSE OR DISTRIBUTE ITS CONTENTS, OR TO MANUFACTURE, USE, OR SELL ANYTHING
// THAT IT MAY DESCRIBE, IN WHOLE OR IN PART.

package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Field1 string `json:"field1,omitempty"`
	Field2 string `json:"field2,omitempty"`
	Field3 int64  `json:"field3,omitempty"`
	RefB   *B     `json:"ref,omitempty"`
}

type B struct {
	Field4 string `json:"field4,omitempty"`
	Field5 string `json:"field5,omitempty"`
	Field6 int64  `json:"field6,omitempty"`
}

func main() {
	b := B{
		Field4: "hello",
		Field5: "man",
		Field6: 6,
	}

	a := A{
		Field1: "hello",
		Field2: "man",
		Field3: 4,
		RefB:   &b,
	}

	bytes, err := json.Marshal(a)
	if err == nil {
		// {"field1":"hello","field2":"man","field3":4,"ref":{"field4":"hello","field5":"man","field6":6}}
		// How to marshal a but for reference just only show field4
		fmt.Println(string(bytes))
	}
}
