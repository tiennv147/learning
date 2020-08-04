package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/tiennv147/learning/proto/http"
	"github.com/tiennv147/learning/proto/security"
)


func main() {
	provider := &security.Provider{
		ProviderType: &security.Provider_TheBProvider{
			TheBProvider: &security.TheBProvider{
				Host: "abc.com",
				Path: "/verify_token",
				ServiceKey: "x123y456",
				ServiceToken: "sdd:4ab3ef2qw1",
			},
		},
	}

	securitySettings := &security.Settings{
		Providers: map[string]*security.Provider{
			"the_b_security" : provider,
		},
	}

	// First-way to create any

	//bytes, err := proto.Marshal(securitySettings)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//setting := &http.Settings{
	//	Filter:        "security",
	//	CrashOnErrors: false,
	//	Settings:      &any.Any{
	//		TypeUrl: "type.googleapis.com/ego.security.aa",
	//		Value: bytes,
	//	},
	//}

	// Second-way to create any
	marshalAny, err := ptypes.MarshalAny(securitySettings)
	if err != nil {
		fmt.Println(err)
		return
	}

	setting := &http.Settings{
		Filter:        "security",
		CrashOnErrors: false,
		Settings:      marshalAny,
	}


	fmt.Printf("setting: %v\n", setting)
	fmt.Printf("marshalAny: %v\n", marshalAny)

	// Try to modify value of bytes after Marshal...

}
