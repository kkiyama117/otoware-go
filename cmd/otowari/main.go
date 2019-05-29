/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import "log"

import (
	"github.com/gordonklaus/portaudio"
	"math/rand"
	"time"
)

func main() {
	log.Printf("hello, world!")
	portaudio.Initialize()
	defer portaudio.Terminate()
	h, err := portaudio.DefaultHostApi()
	chk(err)
	stream, err := portaudio.OpenStream(portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice), func(out []int32) {
		for i := range out {
			out[i] = int32(rand.Uint32())
		}
	})
	chk(err)
	defer stream.Close()
	chk(stream.Start())
	time.Sleep(time.Second)
	chk(stream.Stop())
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
