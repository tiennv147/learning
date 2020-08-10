package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/types/known/anypb"
	gproto "google.golang.org/protobuf/proto"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/tiennv147/learning/proto/http"
	"github.com/tiennv147/learning/proto/security"
)

func main() {
	//tryFirstWayToCreateAny()
	//tryMergeFunction()
	//tryMergeAny()
	tryUnmarshalNew()
}

func tryFirstWayToCreateAny() {
	fmt.Println("-----begin tryFirstWayToCreateAny")
	provider := &security.Provider{
		ProviderType: &security.Provider_TheBProvider{
			TheBProvider: &security.TheBProvider{
				Host:         "abc.com",
				Path:         "/verify_token",
				ServiceKey:   "x123y456",
				ServiceToken: "sdd:4ab3ef2qw1",
			},
		},
	}

	securitySettings := &security.Settings{
		Providers: map[string]*security.Provider{
			"the_b_security": provider,
		},
	}


	bytes, err := proto.Marshal(securitySettings)
	if err != nil {
		fmt.Println(err)
		return
	}

	setting := &http.Settings{
		Filter:        "security",
		CrashOnErrors: false,
		Settings:      &any.Any{
			TypeUrl: "type.googleapis.com/ego.security.aa",
			Value: bytes,
		},
	}
	fmt.Printf("setting: %v\n", setting)
	fmt.Println("-----end tryFirstWayToCreateAny")
}

func tryMergeFunction() {
	fmt.Println("-----begin tryMergeFunction")

	provider := &security.Provider{
		ProviderType: &security.Provider_TheBProvider{
			TheBProvider: &security.TheBProvider{
				Host:         "abc.com",
				Path:         "/verify_token",
				ServiceKey:   "x123y456",
				ServiceToken: "sdd:4ab3ef2qw1",
			},
		},
	}

	securitySettings := &security.Settings{
		Providers: map[string]*security.Provider{
			"the_b_security": provider,
		},
	}

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

	setting2 := &http.Settings{
		Filter:        "security2",
		CrashOnErrors: true,
	}

	fmt.Printf("setting: %v\n", setting)
	fmt.Printf("marshalAny: %v\n", marshalAny)

	// Try merge function from proto
	proto.Merge(setting, setting2)

	fmt.Printf("after merge setting: %v\n", setting)
	fmt.Printf("after merge setting2: %v\n", setting2)

	fmt.Println("-----end tryMergeFunction")

}

func tryMergeAny() {
	fmt.Println("-----begin tryMergeAny")

	provider1 := &security.Provider{
		ProviderType: &security.Provider_TheBProvider{
			TheBProvider: &security.TheBProvider{
				Host:         "abc.com",
				Path:         "/verify_token",
				ServiceKey:   "x123y456",
				ServiceToken: "sdd:4ab3ef2qw1",
			},
		},
	}

	securitySettings1 := &security.Settings{
		Providers: map[string]*security.Provider{
			"the_b_security": provider1,
		},
	}

	marshalAny1, err := ptypes.MarshalAny(securitySettings1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("marshalAny1: %v\n", marshalAny1)

	provider2 := &security.Provider{
		ProviderType: &security.Provider_TheBProvider{
			TheBProvider: &security.TheBProvider{
				ServiceKey:   "--will place at the last--",
			},
		},
	}

	securitySettings2 := &security.Settings{
		Providers: map[string]*security.Provider{
			"the_b_security": provider2,
		},
	}

	marshalAny2, err := ptypes.MarshalAny(securitySettings2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("marshalAny2: %v\n", marshalAny2)

	// Try to put value of marshalAny2 -> marshalAny1
	proto.Merge(marshalAny1, marshalAny2)

	fmt.Println("-----after merged")

	fmt.Printf("marshalAny1: %v\n", marshalAny1)

	fmt.Println("-----end tryMergeAny")

}

func tryUnmarshalNew() {
	fmt.Println("-----begin tryUnmarshalNew")

	provider1 := &security.Provider{
		ProviderType: &security.Provider_TheBProvider{
			TheBProvider: &security.TheBProvider{
				Host:         "abc.com",
				Path:         "/verify_token",
				ServiceKey:   "x123y456",
				ServiceToken: "sdd:4ab3ef2qw1",
			},
		},
	}

	securitySettings1 := &security.Settings{
		Providers: map[string]*security.Provider{
			"the_b_security": provider1,
		},
	}

	marshalAny1, err := ptypes.MarshalAny(securitySettings1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("marshalAny1: %v\n\n", marshalAny1)

	options := gproto.UnmarshalOptions{
	}

	unmarshalNew, err := anypb.UnmarshalNew(marshalAny1, options)
	if err != nil {
		fmt.Println(err)
		return
	}

	fields := unmarshalNew.ProtoReflect().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fmt.Printf("%v ", fields.Get(i).MapValue())
	}


	//fmt.Printf("unmarshalNew: %v\n", unmarshalNew)


	fmt.Println("-----end tryUnmarshalNew")

}